//实现与ssdb通信。
package ssdb

import (
	"net"
	"fmt"
	"bufio"
	"strings"
	"strconv"
	"errors"
	"sync"
)
//锁。保证命令已串行的方式执行。确保多线程情况下，运行正常
var lock sync.Mutex

type Client struct {
	Host   string
	Port   string
	conn   net.Conn
	reader *bufio.Reader
}

func NewSSDBClient(host, port string) Client {
	return Client{
		Host:host,
		Port:port,
		conn:nil,
		reader:nil,
	}
}
func (c *Client)Connect() (error) {
	if c.conn != nil {
		return errors.New("aready connected")
	}
	address := fmt.Sprintf("%s:%s", c.Host, c.Port)
	addr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		return err
	}
	conn, err := net.Dial("tcp", addr.String())
	if err != nil {
		return err
	}
	c.conn = conn
	c.reader = bufio.NewReader(conn)
	return nil

}


func handleResponse(response Response) error {
	var err error
	if (!response.responseStatus) {
		err = errors.New(response.responseText)
	}
	return err
}
func (c *Client)send(cmd...interface{}) {
	lock.Lock()
	str := make([]string, 0)
	for _, commond := range cmd {
		str = append(str, strconv.Itoa(len(commond)))
		str = append(str, commond)
	}
	str = append(str, "\n")
	_, err := c.conn.Write([]byte(strings.Join(str, "\n")))
	if err != nil {
		fmt.Println(err)
	}
}
func (c *Client)read() (Response) {
	defer lock.Unlock()
	//使用buff reader。因为ssdb的server 不会发送EOF.作为结束。所以不能用ReadAll之类的。已EOF或exception为返回条件的函数
	//边界判断 ssdb的协议为连续两个换行"\n" 为结束
	endCount := 0
	response := NewSSDBResponse()

	for {
		//开始一行行读取
		line, _ := c.reader.ReadBytes('\n')
		//判断是否读到了一个空行
		if (len(line) == 1&&string(line) == "\n") {
			endCount++
			//如果是连续2个则说明读完了。可以结束并返回结果了
			if endCount >= 2 {
				return response
			}
		} else {
			//如果不是单独空行则重置边界值
			endCount = 0
			//去掉每行最后面的一个换行符
			line = line[:len(line) - 1]

			//先取长度值
			payloadSize, _ := strconv.ParseInt(string(line), 10, 64)
			next := make([]byte, payloadSize)
			//然后按照长度。读出payload
			c.reader.Read(next)

			//如果是第一个包那么是返回头标识返回状态
			if response.packages == nil {
				if string(next) != "ok" {
					response.responseText = string(next)
					response.responseStatus = false
				}
				response.packages = make([]string, 0)
			} else {
				//压入response
				response.packages = append(response.packages, string(next))
			}
		}
	}

}


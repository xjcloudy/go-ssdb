package ssdb

import (
	//"fmt"
	"testing"
	//"sync"
	"fmt"
)

func setApi(client Client){
	client.Set("t1","nil")
	client.Get("t1")
	client.Setx("sx","5秒后消失",5)
	client.Setnx("t1","should not set")
	client.Expire("t1",5)
	ttl,_:=client.Ttl("t1")
	fmt.Println("t1 ttl=",ttl)
	r,_:=client.Get("t1")
	fmt.Println(r)
	old,_:=client.Getset("t1","new")
	fmt.Println(old)
        client.Del("t1")
	client.Incr("t1",-2)
	t1Exists,_:=client.Exists("t2")
	fmt.Println("t2 is exists?",t1Exists)
	bit,_:=client.Getbit("t1",-1)
	fmt.Println("the offset 0 is ",bit)
	oldbit,_:=client.Setbit("t1",1,1)
	fmt.Println("the 1 offset bit of t1 is",oldbit)

	client.Set("str","abcdefg")
	sub,_:=client.Substr("str",1,20)
	fmt.Println("substr is ",sub)

	len,_:=client.Strlen("str")
	fmt.Println("str length is ",len)

	keys,_:=client.Keys("a","z",10)
	fmt.Println("get keys array ",keys)

	rkeys,_:=client.RKeys("","",10)
	fmt.Println("get rkeys array ",rkeys)

	maps,_:=client.Scan("","",1000)
	fmt.Println("get maps ",maps)

	rmaps,_:=client.Rscan("","",1000)
	fmt.Println("get rmaps ",rmaps)

	client.MultiSet(map[string]string{"red":"#da1337","orange":"#e95a22"})

	mget,_:=client.MultiGet([]string{"red","t1","t3"})
	fmt.Println("multiget result ",mget)

	client.MultiDel([]string{"t1","red"})
}

func TestSSDBClient_Hscan(t *testing.T) {
	ssdb := NewSSDBClient("localhost", "8888")
	err := ssdb.Connect()
	if err != nil {
		t.Error(err)
	} else {
		setApi(ssdb)
		//var wait sync.WaitGroup
		//for i := 0; i < 100; i++ {
		//	wait.Add(1)
		//	go func() {
		//		//result, err := ssdb.Hscan("ha", "", "", 10)
		//		//result3, err := ssdb.Hget("ha", "a")
		//		result4, err := ssdb.Set("a", "中文中文")
		//		result2, err := ssdb.Get("a")
		//
		//		//result5, err := ssdb.Hset("hb", "hello", "world")
		//		//result6, err := ssdb.Hget("hb", "hello")
		//
		//		if err != nil {
		//			t.Error(err)
		//		} else {
		//			fmt.Println( result4,result2)
		//		}
		//		wait.Done()
		//	}()
		//
		//}
		//wait.Wait()
	}
}

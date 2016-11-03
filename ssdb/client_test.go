package ssdb

import (
	//"fmt"
	"testing"
	//"sync"
	"fmt"
)

func setApi(client Client) {
	client.Set("t1", "nil")
	client.Get("t1")
	client.Setx("sx", "5秒后消失", 5)
	client.Setnx("t1", "should not set")
	client.Expire("t1", 5)
	ttl, _ := client.Ttl("t1")
	fmt.Println("t1 ttl=", ttl)
	r, _ := client.Get("t1")
	fmt.Println(r)
	old, _ := client.Getset("t1", "new")
	fmt.Println(old)
	client.Del("t1")
	client.Incr("t1", -2)
	t1Exists, _ := client.Exists("t2")
	fmt.Println("t2 is exists?", t1Exists)
	bit, _ := client.Getbit("t1", -1)
	fmt.Println("the offset 0 is ", bit)
	oldbit, _ := client.Setbit("t1", 1, 1)
	fmt.Println("the 1 offset bit of t1 is", oldbit)

	client.Set("str", "abcdefg")
	sub, _ := client.Substr("str", 1, 20)
	fmt.Println("substr is ", sub)

	len, _ := client.Strlen("str")
	fmt.Println("str length is ", len)

	keys, _ := client.Keys("a", "z", 10)
	fmt.Println("get keys array ", keys)

	rkeys, _ := client.RKeys("", "", 10)
	fmt.Println("get rkeys array ", rkeys)

	maps, _ := client.Scan("", "", 1000)
	fmt.Println("get maps ", maps)

	rmaps, _ := client.Rscan("", "", 1000)
	fmt.Println("get rmaps ", rmaps)

	client.MultiSet(map[string]string{"blue":"#da1337", "black":"#e95a22"})

	mget, _ := client.MultiGet([]string{"red", "t1", "t3"})
	fmt.Println("multiget result ", mget)

	client.MultiDel([]string{"t1", "red"})
}
func hsetApi(client Client) {
	client.Hset("map1", "m1", "v1");
	mval, _ := client.Hget("map1", "m1")
	fmt.Println("value of m1 from map1 is", mval)
	client.Hdel("map1", "m1")

	oval, _ := client.Hincr("map1", "m1", 10)
	fmt.Println("after hincr value is ", oval)

	isExists, _ := client.Hexists("map1", "m1")
	fmt.Println("is m1 in map1 ", isExists)

	mapSize, _ := client.Hsize("map1")
	fmt.Println("size of map1 is ", mapSize)

	mapNames, _ := client.Hlist("", "", 10)
	fmt.Println("get map names ", mapNames)

	keys, _ := client.Hkeys("map1", "", "", 10)
	fmt.Println("get keys from map1", keys)

	allKvs, _ := client.Hgetall("map1")
	fmt.Println("all kvs of map1", allKvs)

	map1kvs, _ := client.Hscan("map1", "", "", 10)
	fmt.Println("values of map1", map1kvs)

	count, _ := client.Hclear("map1")
	fmt.Println("clear maps ", count)

	client.MultiHset("map1", map[string]string{"b1":"1", "b2":"2", "b3":"3"})

	mkvs, _ := client.MultiHget("map1", []string{"b3", "b1"})
	fmt.Println("multihget ", mkvs)
	client.MultiHdel("map1", []string{"b1", "a1"})
}
func zsetApi(client Client){
	client.Zset("z1","k1",100)
	k1val,_:=client.Zget("z1","k1")
	fmt.Println("k1 =",k1val)
	client.Zdel("z1","k1")
	nval,_:=client.Zincr("z1","k1",1000)
	fmt.Println("new score=",nval)
	size,_:=client.Zsize("z1")
	fmt.Println("size =",size)

}
func TestSSDBClient_Hscan(t *testing.T) {
	ssdb := NewSSDBClient("localhost", "8888")
	err := ssdb.Connect()
	if err != nil {
		t.Error(err)
	} else {
		setApi(ssdb)
		hsetApi(ssdb)
		zsetApi(ssdb)
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

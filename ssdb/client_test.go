package ssdb

import (
	"testing"
	//"sync"
	"fmt"
)

func setApi(client *Client) {
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
func hsetApi(client *Client) {
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
func zsetApi(client *Client) {
	client.Zset("z1", "k1", 100)
	k1val, _ := client.Zget("z1", "k1")
	fmt.Println("k1 =", k1val)
	client.Zdel("z1", "k1")
	nval, _ := client.Zincr("z1", "k1", 1000)
	fmt.Println("new score=", nval)
	zkeys, _ := client.Zlist("", "", 1)
	fmt.Println("get key list", zkeys)
	zrkeys, _ := client.Zrlist("", "", 1)
	fmt.Println("get r key list", zrkeys)
	size, _ := client.Zsize("z1")
	fmt.Println("size =", size)

	ks, _ := client.Zkeys("z1", "", "", "", 10)
	fmt.Println("kyes", ks)

	count, _ := client.Zcount("z1", "", "")
	fmt.Println("count =", count)

	sum, _ := client.Zsum("z1", "", "")
	fmt.Println("sum=", sum)

	avg, _ := client.Zavg("z1", "", "")
	fmt.Println("avg=", avg)

	rm1, _ := client.Zremrangebyrank("z1", "k1", "k1")
	fmt.Println("rm key ", rm1)

	top, _ := client.ZpopFront("z1", 1)
	fmt.Println("pop from zset", top)

	end, _ := client.ZpopBack("z1", 1)
	fmt.Println("pop from zset", end)
	client.MultiZset("z2", map[string]string{"k1":"1", "k2":"2", "k3":"3"})
	kvs, _ := client.MultiZget("z2", []string{"k1", "k2"})
	fmt.Println("multi get ", kvs)
	client.MultiZdel("z2", []string{"k1", "k2"})
	client.Zclear("z1")

}
func qsetApi(client *Client){

        qsize,_:=client.QpushFront("q1","a","b","c","d")
	fmt.Println("qsize ",qsize)
	setr,_:=client.Qset("q1",1,"v2")
	fmt.Println("qset ",setr)
	getr,_:=client.Qget("q1",100)
	fmt.Println("qget ",getr)
	list,_:=client.Qlist("","",100)
	fmt.Println("queue values ",list)

	q1size,_:=client.Qsize("q1")
	fmt.Println("q1 size ",q1size)

	fr,_:=client.Qfront("q1")
	fmt.Println("front of q1 ",fr)

	bc,_:=client.Qback("q1")
	fmt.Println("back of q1 ",bc)


	rlist,_:=client.Qrange("q1",0,10)
	fmt.Println("range of q1 ",rlist)

	slice,_:=client.Qslice("q1",1,12)
	fmt.Println("slice of q1 ",slice)

	client.QpushBack("q1",[]string{"b1","b2"}...)

	popitem,_:=client.QpopFront("q1",10)
	fmt.Println("item pop from q1 is ",popitem)
	backitem,_:=client.QpopBack("q1",2)
	fmt.Println("items back from q1 is ",backitem)

	psize,_:=client.QtrimFront("q1",2)
	fmt.Printf("delete %d items from q1 at top \n",psize)

	bsize,_:=client.QtrimBack("q1",12)
	fmt.Printf("delete %d items from q1 at bottom",bsize)

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
		qsetApi(ssdb)
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



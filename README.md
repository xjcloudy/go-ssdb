#go-ssdb
 go语言版本的ssdb驱动。
 
## usage

```
   client := ssdb.NewSSDBClient("localhost", "8888")
   err:=client.Connect()
   resp,err := client.Set("hello","world")
   
```
## api list

 - get
 - set
 - hget
 - hscan
 - hset
 
## todo

 - 实现更多的api
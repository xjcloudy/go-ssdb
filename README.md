#go-ssdb
 go语言版本的ssdb驱动。
 
## usage

```
   client:=ssdb.NewSSDBClient("localhost","8888")
   resp,err := client.Set("hello","world")
```

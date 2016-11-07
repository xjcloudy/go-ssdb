#go-ssdb
 go语言版本的ssdb驱动。
 
## usage

```
   client := ssdb.NewSSDBClient("localhost", "8888")
   client.Connect()
   resp,err := client.Set("hello","world")
   
```

## type

api的参数和返回类型限定于以下几种。
- `string`
  存取的值得类型都为`string`
- `int64`
- `unit64`
  用于大部分的limit参数。和统计类api的返回类型
- `float64`
  只有zavg的需要返回浮点类型
- `bool`
  ssdb的api用false表示操作失败，其他值表示成功。golang里面统一处理成true和false.
- `[]string`
- `map[string]string`
 
 

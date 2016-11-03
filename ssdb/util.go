package ssdb

import "strconv"

func sInt(val int64)string{
	return strconv.FormatInt(val,10)
}
func sUint(val uint64)string {
	return strconv.FormatUint(val,10)
}
func sMap(val map[string]string)[]string{
	var strArray []string
	for k,v:=range(val){
		strArray=append(strArray,k,v)
	}
	return strArray
}


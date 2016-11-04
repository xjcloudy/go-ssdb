package ssdb

import (
	"strings"
	"strconv"
)

type Response struct {
	responseStatus bool
	responseText   string
	packages       []string
}

func (r *Response) toString() (string) {
	return strings.Join(r.packages, "")
}
func (r *Response) toBool() (bool) {
	if r.responseStatus {
		b := true
		rsize := len(r.packages)
		if rsize > 0&&r.packages[0] == "false" {
			b = false
		}
		return b
	} else {
		return false
	}

}
func (r *Response) toInt() (int64) {
	if i, err := strconv.ParseInt(r.packages[0], 10, 64); err == nil {
		return i
	} else {
		return 0
	}
}
func (r *Response)toFloat() (float64) {
	if i, err := strconv.ParseFloat(r.packages[0], 64); err == nil {
		return i
	} else {
		return 0
	}
}
func (r *Response) toArray() ([]string) {
	//todo should return new array fill with r.packages
	return r.packages
}
func (r *Response) toMap() (map[string]string) {

	mapData := make(map[string]string)
	size := len(r.packages)
	i := 0
	if r.responseStatus&&size % 2 == 0 {
		for i < size {
			key := r.packages[i]
			value := r.packages[i + 1]
			mapData[key] = value
			i += 2
		}
	}

	return mapData
}
func NewSSDBResponse() (Response) {
	return Response{
		responseStatus:true,
		responseText:"ok",
		packages:nil,
	}
}

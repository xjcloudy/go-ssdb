package ssdb

import "strings"

type Response struct {
	responseStatus bool
	responseText   string
	packages       []string
}

func (r *Response) String() (string) {
	return strings.Join(r.packages, "")
}
func (r *Response) Map() (map[string]string) {
	mapData := make(map[string]string)
	size := len(r.packages)
	i := 0
	for i < size {
		key := r.packages[i]
		value := r.packages[i + 1]
		mapData[key] = value
		i += 2
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

package ssdb

import "strings"

type SSDBResponse struct {
	responseStatus bool
	responseText   string
	packages       []string
}

func (r *SSDBResponse) String() (string) {
	return strings.Join(r.packages, "")
}
func (r *SSDBResponse) Map() (map[string]string) {
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
func NewSSDBResponse() (SSDBResponse) {
	return SSDBResponse{true, "ok", nil}
}

package ssdb

import "strconv"

func (c *Client)Hget(key, value string) (bool, error) {
	response := c.do("set", key, value)
	return response.toBool(), handleResponse(response)
}
func (c *Client)Hset(mapName,key, value string) (bool, error) {
	response := c.do("hset", mapName,key, value)
	return response.toBool(), handleResponse(response)
}
func (c *Client)Hscan(mapName, startKey,endKey string,limit uint64) (map[string]string, error) {
	response := c.do("hscan", mapName, startKey,endKey,strconv.FormatUint(limit,10))
	return response.toMap(), handleResponse(response)
}
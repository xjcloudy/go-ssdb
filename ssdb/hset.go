package ssdb

func (c *Client)Hset(mapName, key, value string) (bool, error) {
	response := c.do("hset", mapName, key, value)
	return response.toBool(), handleResponse(response)
}
func (c *Client)Hget(mapName, key string) (string, error) {
	response := c.do("hget", mapName, key)
	return response.toString(), handleResponse(response)
}
func (c *Client)Hdel(mapName, key string) (bool, error) {
	response := c.do("hdel", mapName, key)
	return response.toBool(), handleResponse(response)
}
func (c *Client)Hincr(mapName, key string, num int64) (int64, error) {
	response := c.do("hincr", mapName, key, sInt(num))
	return response.toInt(), handleResponse(response)
}
func (c *Client)Hexists(mapName, key string) (bool, error) {
	response := c.do("hexists", mapName, key)
	return response.toBool(), handleResponse(response)
}
func (c *Client)Hsize(mapName string) (int64, error) {
	response := c.do("hsize", mapName)
	return response.toInt(), handleResponse(response)
}
func (c *Client)Hlist(nameStart, nameEnd string, limit int64) ([]string, error) {
	response := c.do("hlist", nameStart, nameEnd, sInt(limit))
	return response.toArray(), handleResponse(response)
}
func (c *Client)Hrlist(nameStart, nameEnd string, limit int64) ([]string, error) {
	response := c.do("hrlist", nameStart, nameEnd, sInt(limit))
	return response.toArray(), handleResponse(response)
}
func (c *Client)Hkeys(mapName, nameStart, nameEnd string, limit int64) ([]string, error) {
	response := c.do("hkeys", mapName, nameStart, nameEnd, sInt(limit))
	return response.toArray(), handleResponse(response)
}
func (c *Client)Hgetall(mapName string) (map[string]string, error) {
	response := c.do("hgetall", mapName)
	return response.toMap(), handleResponse(response)
}
func (c *Client)Hscan(mapName, nameStart, nameEnd string, limit int64) (map[string]string, error) {
	response := c.do("hscan",mapName, nameStart, nameEnd, sInt(limit))
	return response.toMap(), handleResponse(response)
}
func (c *Client)Hclear(mapName string) (int64, error) {
	response := c.do("hclear", mapName)
	return response.toInt(), handleResponse(response)
}
func (c *Client)MultiHset(mapName string, kvs map[string]string) (bool, error) {
	response := c.do("multi_hset", append([]string{mapName}, sMap(kvs)...)...)
	return response.toBool(), handleResponse(response)
}
func (c *Client)MultiHget(mapName string, keys []string) (map[string]string, error) {
	response := c.do("multi_hget", append([]string{mapName}, keys...)...)
	return response.toMap(), handleResponse(response)
}
func (c *Client)MultiHdel(mapName string, keys []string) (bool, error) {
	response := c.do("multi_hdel", append([]string{mapName}, keys...)...)
	return response.toBool(), handleResponse(response)
}
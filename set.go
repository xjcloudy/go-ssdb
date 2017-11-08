package ssdb

func (c *Client)Get(key string) (string, error) {
	response := c.do("get", key)
	return response.toString(), handleResponse(response)
}
func (c *Client)Set(key, value string) (bool, error) {
	response := c.do("set", key, value)
	return response.toBool(), handleResponse(response)
}
func (c *Client)Setx(key, value string, ttl int64) (bool, error) {
	response := c.do("setx", key, value, sInt(ttl))
	return response.toBool(), handleResponse(response)
}
func (c *Client)Setnx(key, value string) (bool, error) {
	response := c.do("setnx", key, value)
	return response.toBool(), handleResponse(response)
}
func (c *Client)Expire(key string, ttl int64) (bool, error) {
	response := c.do("expire", key, sInt(ttl))
	return response.toBool(), handleResponse(response)
}
func (c *Client)Ttl(key string) (int64, error) {
	response := c.do("ttl", key)
	return response.toInt(), handleResponse(response)
}
func (c *Client)Getset(key, value string) (string, error) {
	response := c.do("getset", key, value)
	return response.toString(), handleResponse(response)
}
func (c *Client)Del(key string) (bool, error) {
	response := c.do("del", key)
	return response.toBool(), handleResponse(response)
}
func (c *Client)Incr(key string, num int64) (bool, error) {
	response := c.do("incr", key, sInt(num))
	return response.toBool(), handleResponse(response)
}
func (c *Client)Exists(key string) (bool, error) {
	response := c.do("exists", key)
	return response.toBool(), handleResponse(response)
}
func (c *Client)Getbit(key string, offset int64) (int64, error) {
	response := c.do("getbit", key, sInt(offset))
	return response.toInt(), handleResponse(response)
}
func (c *Client)Setbit(key string, offset int64, value uint64) (int64, error) {
	response := c.do("setbit", key, sInt(offset), sUint(value))
	return response.toInt(), handleResponse(response)
}
func (c *Client)Bitcount(key string, between ...int) (int64, error) {
	response := c.do("bitcount", key)
	return response.toInt(), handleResponse(response)
}
func (c *Client)Substr(key string, between ...int64) (string, error) {
	var pos []string
	for _, i := range (between) {
		pos = append(pos, sInt(i))
	}
	response := c.do("substr", pos...)
	return response.toString(), handleResponse(response)
}
func (c *Client)Strlen(key string) (int64, error) {
	response := c.do("strlen", key)
	return response.toInt(), handleResponse(response)
}
func (c *Client)Keys(keyStart, keyEnd string, limit uint64) ([]string, error) {
	response := c.do("keys", keyStart, keyEnd, sUint(limit))
	return response.toArray(), handleResponse(response)
}
func (c *Client)RKeys(keyStart, keyEnd string, limit uint64) ([]string, error) {
	response := c.do("rkeys", keyStart, keyEnd, sUint(limit))
	return response.toArray(), handleResponse(response)
}
func (c *Client)Scan(keyStart, keyEnd string, limit uint64) (map[string]string, error) {
	response := c.do("scan", keyStart, keyEnd, sUint(limit))
	return response.toMap(), handleResponse(response)
}
func (c *Client)Rscan(keyStart, keyEnd string, limit uint64) (map[string]string, error) {
	response := c.do("rscan", keyStart, keyEnd, sUint(limit))
	return response.toMap(), handleResponse(response)
}
func (c *Client)MultiSet(kvs map[string]string) (bool, error) {
	response := c.do("multi_set", sMap(kvs)...)
	return response.toBool(), handleResponse(response)
}
func (c *Client)MultiGet(keys []string) (map[string]string, error) {
	response := c.do("multi_get", keys...)
	return response.toMap(), handleResponse(response)
}
func (c *Client)MultiDel(keys []string) (bool, error) {
	response := c.do("multi_del", keys...)
	return response.toBool(), handleResponse(response)
}
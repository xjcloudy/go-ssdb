package ssdb

func (c *Client)Zset(name, key string, score int64) (bool, error) {
	response := c.do("zset", name, key, sInt(score))
	return response.toBool(), handleResponse(response)
}
func (c *Client)Zget(name, key string) (int64, error) {
	response := c.do("zget", name, key)
	return response.toInt(), handleResponse(response)
}
func (c *Client)Zdel(name, key string) (bool, error) {
	response := c.do("zdel", name, key)
	return response.toBool(), handleResponse(response)
}
func (c *Client)Zincr(name, key string, num int64) (int64, error) {
	response := c.do("zincr", name, key, sInt(num))
	return response.toInt(), handleResponse(response)
}
func (c *Client)Zexists(name, key string) (bool, error) {
	response := c.do("zexists", name, key)
	return response.toBool(), handleResponse(response)
}
func (c *Client)Zsize(name string) (int64, error) {
	response := c.do("zsize", name)
	return response.toInt(), handleResponse(response)
}
func (c *Client)Zlist(nameStart, nameEnd string, limit int64) ([]string, error) {
	response := c.do("zlist", nameStart, nameEnd, sInt(limit))
	return response.toArray(), handleResponse(response)
}
func (c *Client)Zrlist(nameStart, nameEnd string, limit int64) ([]string, error) {
	response := c.do("zrlist", nameStart, nameEnd, sInt(limit))
	return response.toArray(), handleResponse(response)
}
func (c *Client)Zkeys(name, nameStart, scoreStart,scoreEnd string, limit int64) ([]string, error) {
	response := c.do("zkeys", name, nameStart, scoreStart,scoreEnd ,sInt(limit))
	return response.toArray(), handleResponse(response)
}

func (c *Client)Zscan(name, nameStart, scoreStart,scoreEnd string, limit int64) (map[string]string, error) {
	response := c.do("zscan",name, nameStart, scoreStart,scoreEnd ,sInt(limit))
	return response.toMap(), handleResponse(response)
}
func (c *Client)Zrscan(name, nameStart, scoreStart,scoreEnd string, limit int64) (map[string]string, error){
	response := c.do("zrscan",name, nameStart, scoreStart,scoreEnd ,sInt(limit))
	return response.toMap(), handleResponse(response)
}
func (c *Client)Zcount(name,scoreStart,scoreEnd string)(int64,error){
	response := c.do("zcount",name, scoreStart, scoreEnd)
	return response.toInt(), handleResponse(response)
}
func (c *Client)Zsum(name,scoreStart,scoreEnd string) (int64, error) {
	response := c.do("zsum", name,scoreStart,scoreEnd)
	return response.toInt(), handleResponse(response)
}
func (c *Client)Zavg(name,scoreStart,scoreEnd string) (int64, error) {
	response := c.do("zavg", name,scoreStart,scoreEnd)
	return response.toInt(), handleResponse(response)
}
func (c *Client)Zremrangebyrank(name,start,end string) (int64, error) {
	response := c.do("zclear", name)
	return response.toInt(), handleResponse(response)
}
func (c *Client)Zremrangebyscore(name,startScore,endScore   string) (int64, error) {
	response := c.do("hclear", name)
	return response.toInt(), handleResponse(response)
}
func (c *Client)ZpopFront(name string,limit int64) (map[string]string, error) {
	response := c.do("zpop_front", name,sInt(limit))
	return response.toMap(), handleResponse(response)
}
func (c *Client)ZpopBack(name string,limit int64) (map[string]string, error) {
	response := c.do("zpop_back", name,sInt(limit))
	return response.toMap(), handleResponse(response)
}
func (c *Client)Zclear(mapName string) (int64, error) {
	response := c.do("hclear", mapName)
	return response.toInt(), handleResponse(response)
}
func (c *Client)MultiZset(name string, kvs map[string]string) (bool, error) {
	response := c.do("multi_zset", append([]string{name}, sMap(kvs)...)...)
	return response.toBool(), handleResponse(response)
}
func (c *Client)MultiZget(name string, keys []string) (map[string]string, error) {
	response := c.do("multi_zget", append([]string{name}, keys...)...)
	return response.toMap(), handleResponse(response)
}
func (c *Client)MultiZdel(name string, keys []string) (bool, error) {
	response := c.do("multi_zdel", append([]string{name}, keys...)...)
	return response.toBool(), handleResponse(response)
}
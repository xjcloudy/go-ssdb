package ssdb

func (c *Client) Qsize(qname string) (int64, error) {
	response := c.do("qsize", qname)
	return response.toInt(), handleResponse(response)
}
func (c *Client) Qlist(nameStart, nameEnd string, limit uint64) ([]string, error) {
	response := c.do("qlist", nameStart, nameEnd,sUint(limit))
	return response.toArray(), handleResponse(response)
}
func (c *Client) Qrlist(qname string) ([]string, error) {
	response := c.do("qrlist", qname)
	return response.toArray(), handleResponse(response)
}
func (c *Client) Qclear(qname string) (bool, error) {
	response := c.do("qclear", qname)
	return response.toBool(), handleResponse(response)
}
func (c *Client) Qfront(qname string) (string, error) {
	response := c.do("qfront", qname)
	return response.toString(), handleResponse(response)
}
func (c *Client) Qback(qname string) (string, error) {
	response := c.do("qback", qname)
	return response.toString(), handleResponse(response)
}
func (c *Client) Qget(qname string, index int64) (string, error) {
	response := c.do("qget", qname, sInt(index))
	return response.toString(), handleResponse(response)
}
func (c *Client) Qset(qname string, index int64, value string) (bool, error) {
	response := c.do("qset", qname,sInt(index),value)
	return response.toBool(), handleResponse(response)
}
func (c *Client) Qrange(qname string, offset int64, limit uint64) ([]string, error) {
	response := c.do("qrange", qname, sInt(offset), sUint(limit))
	return response.toArray(), handleResponse(response)
}
func (c *Client) Qslice(qname string, begin, end int64) ([]string, error) {
	response := c.do("qslice", qname, sInt(begin), sInt(end))
	return response.toArray(), handleResponse(response)
}
func (c *Client) QpushFront(qname string, item... string) (int64, error) {
	response := c.do("qpush_front", append([]string{qname},item...)...)
	return response.toInt(), handleResponse(response)
}
func (c *Client) QpushBack(qname string, item... string) (int64, error) {
	response := c.do("qpush_back", append([]string{qname}, item ...)...)
	return response.toInt(), handleResponse(response)
}
func (c *Client) QpopFront(qname string, size uint64) ([]string, error) {
	response := c.do("qpop_front", qname, sUint(size))
	return response.toArray(), handleResponse(response)
}
func (c *Client) QpopBack(qname string, size uint64) ([]string, error) {
	response := c.do("qpop_back", qname, sUint(size))
	return response.toArray(), handleResponse(response)
}
func (c *Client) QtrimFront(qname string, size uint64) (int64, error) {
	response := c.do("qtrim_front", qname, sUint(size))
	return response.toInt(), handleResponse(response)
}
func (c *Client) QtrimBack(qname string, size uint64) (int64, error) {
	response := c.do("qtrim_front", qname, sUint(size))
	return response.toInt(), handleResponse(response)
}











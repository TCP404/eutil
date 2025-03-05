package cerr

type ApiErrMsg = string

type Msg struct {
	Code  int
	enStr ApiErrMsg
	cnStr ApiErrMsg
}

func (c *Msg) String() string {
	return string(c.enStr)
}

func (c *Msg) EN() string {
	return string(c.enStr)
}

func (c *Msg) CN() string {
	return string(c.cnStr)
}

func (c *Msg) SetCode(code int) *Msg {
	c.Code = code
	return c
}

func (c *Msg) SetEN(enMsg ApiErrMsg) *Msg {
	c.enStr = enMsg
	return c
}

func (c *Msg) SetCN(cnMsg ApiErrMsg) *Msg {
	c.cnStr = cnMsg
	return c
}

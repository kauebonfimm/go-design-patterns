package chainresponsabilty

type BaseHandler struct {
	next IProcessorUserSignUp
}

func (c *BaseHandler) LinkChain(next IProcessorUserSignUp) {
	c.next = next
}

func (c *BaseHandler) NextOrSuccess(user *User) Status {
	if c.next != nil {
		return c.next.Execute(user)
	}

	return Status{
		Success, "",
	}
}

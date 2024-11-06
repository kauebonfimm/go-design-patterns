package chainresponsabilty

import "fmt"

type AggregaterHandler struct {
	BaseHandler
}

func (h *AggregaterHandler) Execute(user *User) Status {
	user.ID = fmt.Sprint(1)
	user.Active = true
	user.Consumer = true
	return h.NextOrSuccess(user)
}

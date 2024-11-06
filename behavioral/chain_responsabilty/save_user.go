package chainresponsabilty

import "fmt"

type SaveUserHandler struct {
	BaseHandler
}

func (h *SaveUserHandler) Execute(user *User) Status {
	fmt.Println("Saved User")
	return h.NextOrSuccess(user)
}

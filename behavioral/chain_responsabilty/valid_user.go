package chainresponsabilty

import "regexp"

type ValidUserHandler struct {
	BaseHandler
}

func (h *ValidUserHandler) Execute(user *User) Status {
	if !hasAgeValid(user.Age) || !hasNameValid(user.Name) {
		return Status{Error, "invalid user"}
	}

	return h.NextOrSuccess(user)
}

func hasAgeValid(age int) bool {
	return age > 0 && age < 150
}

func hasNameValid(name string) bool {
	rx, _ := regexp.Compile("[a-zA-Z]*")
	return rx.Match([]byte(name))
}

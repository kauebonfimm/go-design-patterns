package chainresponsabilty

type TypeStatus int

const (
	Success TypeStatus = iota + 1
	Pending
	Error
)

type Status struct {
	Type   TypeStatus
	Reason string
}

func (s *Status) ToString() string {
	switch s.Type {
	case Success:
		return "success"
	case Pending:
		return "pending"
	case Error:
		return "error"
	}
	return ""
}

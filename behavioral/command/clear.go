package command

type ClearCommand struct {
	d DispatchEvent
}

func NewClearCommand() Command {
	return &ClearCommand{d: DispatchEvent{EventName: "clear"}}
}

func (c *ClearCommand) Execute() error {
	c.d.Send()
	return nil
}

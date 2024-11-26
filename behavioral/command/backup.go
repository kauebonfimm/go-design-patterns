package command

type BackupCommand struct {
	d DispatchEvent
}

func NewBackupCommand() Command {
	return &BackupCommand{d: DispatchEvent{EventName: "backup"}}
}

func (b *BackupCommand) Execute() error {
	b.d.Send()
	return nil
}

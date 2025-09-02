package cmd

type Command struct {
	initFunc  func() error
	startFunc func() error
	stopFunc  func() error
}

func NewCommand(initFunc func() error, startFunc func() error, stopFunc func() error) *Command {
	return &Command{
		initFunc:  initFunc,
		startFunc: startFunc,
		stopFunc:  stopFunc,
	}
}

func (c *Command) Main() error {
	
}

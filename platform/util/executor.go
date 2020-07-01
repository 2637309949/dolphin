package util

// Executor defined
type Executor []func() error

// Add defined
func (act *Executor) Add(i ...func() error) *Executor {
	*act = append(*act, i...)
	return act
}

// Execute defined
func (act *Executor) Execute() error {
	for _, v := range *act {
		if err := v(); err != nil {
			return err
		}
	}
	return nil
}

// NewExecutor defined
func NewExecutor(i ...func() error) *Executor {
	e := &Executor{}
	e.Add(i...)
	return e
}

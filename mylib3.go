package mylib3

type Counter struct {
    Value int
}

func (c *Counter) Inc() { c.Value++ }

func NewCounter() *Counter { return &Counter{ 5 } }

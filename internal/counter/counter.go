package counter

type Counter struct {
	Count int
}

func (c Counter) Get() int {
	return c.Count
}

func (c *Counter) Incr() {
	c.Count++
}

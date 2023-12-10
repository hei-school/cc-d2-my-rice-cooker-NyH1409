package main

type Cooker struct {
	IsOn         bool
	IsPlugIn     bool
	containWater bool
}

func (c *Cooker) turnOnOrOff() bool {
	c.IsOn = !c.IsOn
	return c.IsOn
}

func (c *Cooker) plugInOrOff() bool {
	c.IsPlugIn = !c.IsPlugIn
	return c.IsPlugIn
}

func main() {

}

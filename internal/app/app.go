package app

type Application interface {
	Run()
}

type Core struct {
}

func NewApplication() Application {
	return &Core{}
}

func (c *Core) Run() {

}

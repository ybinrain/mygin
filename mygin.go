package mygin

type Engine struct {
	Name string
}

func (e *Engine) Ping() string {
	return e.Name
}

func Default() *Engine {
	engine := &Engine{
		Name: "ybin",
	}
	return engine
}
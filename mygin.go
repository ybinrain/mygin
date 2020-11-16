package mygin

type Engine struct {
	Name string
}

func (e *Engine) Ping() string {
	return "pong"
}

func Default() *Engine {
	engine := &Engine{
		Name: "ybin",
	}
	return engine
}
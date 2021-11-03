package core

// Engine defined TODO
type Engine interface {
	Context
	Handler
}

func SetEngine(e Engine) {
	defaultHandler = e
	defaultContext = e
}

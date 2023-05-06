package types

type Endpoint struct {
	Uri         string
	Handler     func(req Request, res Response)
	Middlewares []Middleware
}

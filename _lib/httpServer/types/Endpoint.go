package types

import "main/_lib/httpServer/request"

type Endpoint struct {
	Uri         string
	Handler     func(req request.Request, res Response)
	Middlewares []Middleware
}

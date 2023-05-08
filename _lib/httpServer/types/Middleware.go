package types

import "main/_lib/httpServer/request"

type Middleware func(req request.Request, res Response) (request.Request, Response, error)

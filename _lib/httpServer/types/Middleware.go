package types

type Middleware func(req Request, res Response) (Request, Response, error)

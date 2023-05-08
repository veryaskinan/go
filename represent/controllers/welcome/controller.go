package welcome

import (
	"main/_lib/httpServer/request"
	httpServerTypes "main/_lib/httpServer/types"
	"main/domen/modules/welcome"
)

func handler(req request.Request, res httpServerTypes.Response) {
	welcomeResult := welcome.Welcome()
	res.SendJson(welcomeResult)
}

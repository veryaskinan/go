package welcome

import (
	httpServerTypes "main/_lib/httpServer/types"
	"main/domen/modules/welcome"
)

func handler(req httpServerTypes.Request, res httpServerTypes.Response) {
	welcomeResult := welcome.Welcome()
	res.SendJson(welcomeResult)
}

package request

import (
	"fmt"
	httpServerTypes "main/_lib/httpServer/types"
	"main/domen/auth/signIn/request"
)

func Handler(req httpServerTypes.Request, res httpServerTypes.Response) {
	signUpInfo, err := validate(req)

	if err != nil {
		res.SetStatus(400)
		res.SendJson(fmt.Sprint(err))
		return
	} else {
		result, err := request.Request(signUpInfo)
		if err != nil {
			res.SendJson(err)
		}
		res.SendJson(result)
	}
}

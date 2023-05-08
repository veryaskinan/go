package request

import (
	"fmt"
	request2 "main/_lib/httpServer/request"
	httpServerTypes "main/_lib/httpServer/types"
	"main/domen/modules/auth/signIn/request"
)

func Handler(req request2.Request, res httpServerTypes.Response) {
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

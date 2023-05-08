package confirm

import (
	"fmt"
	"main/_lib/httpServer/request"
	httpServerTypes "main/_lib/httpServer/types"
	"main/domen/modules/auth/signIn/confirm"
)

func Handler(req request.Request, res httpServerTypes.Response) {
	signUpInfo, err := validate(req)

	if err != nil {
		res.SetStatus(400)
		res.SendJson(fmt.Sprint(err))
		return
	} else {
		result, err := confirm.Confirm(signUpInfo)
		if err != nil {
			res.SendError(err.Error())
		} else {
			res.SendJson(result)
		}
	}
}

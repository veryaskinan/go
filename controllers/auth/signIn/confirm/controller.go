package confirm

import (
	"fmt"
	httpServerTypes "main/_lib/httpServer/types"
	"main/domen/auth/signIn/confirm"
)

func Handler(req httpServerTypes.Request, res httpServerTypes.Response) {
	signUpInfo, err := validate(req)

	if err != nil {
		res.SetStatus(400)
		res.SendJson(fmt.Sprint(err))
		return
	} else {
		result, err := confirm.Confirm(signUpInfo)
		if err != nil {
			res.SendError(err)
		} else {
			res.SendJson(result)
		}
	}
}

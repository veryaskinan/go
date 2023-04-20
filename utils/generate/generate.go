package generate

import (
	"errors"
	"fmt"
	"github.com/zach-klippenstein/goregen"
)

func validate(length uint8) (bool, error) {
	if length < 1 {
		return false, errors.New("Wrong length")
	}
	return true, nil
}

func Number(length uint8) (string, error) {
	_, err := validate(length)
	if err != nil {
		return "", errors.New("Length validation error")
	}
	number, err := regen.Generate(fmt.Sprintf("[0-9]{%d}", length))

	return number, nil
}

func String(length uint8) string {
	token, _ := regen.Generate(fmt.Sprintf("[0-9a-zA-Z]{%d}", length))
	return token
}

package generate

import (
	"errors"
	"fmt"
	"github.com/zach-klippenstein/goregen"
	"main/_lib/logger"
	"math/rand"
	"time"
)

func validate(length uint8) (bool, error) {
	if length < 1 {
		return false, errors.New("Wrong length")
	}
	return true, nil
}

func Number(length uint8) string {
	_, err := validate(length)
	if err != nil {
		logger.Info("ERROR! Length validation error")
	}
	return generate(fmt.Sprintf("[0-9]{%d}", length))
}

func String(length uint8) string {
	_, err := validate(length)
	if err != nil {
		logger.Info("ERROR! Length validation error")
	}
	return generate(fmt.Sprintf("[0-9a-zA-Z]{%d}", length))
}

func generate(pattern string) string {
	rand.Seed(time.Now().Unix())
	seed := rand.Int63()
	generator, err := regen.NewGenerator(pattern, &regen.GeneratorArgs{RngSource: rand.NewSource(seed)})
	if err != nil {
		logger.Info("ERROR! New generator error error")
	}
	return generator.Generate()
}

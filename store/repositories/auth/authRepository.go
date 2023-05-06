package auth

import (
	"errors"
	"main/_lib/logger"
	"main/_lib/redis"
)

var authCodePrefix string = "auth:code:"
var authTokenPrefix string = "auth:token:"

func getAuthCodeRedisKey(phone string) string {
	return authCodePrefix + phone
}

func getAuthTokenRedisKey(phone string) string {
	return authTokenPrefix + phone
}

func SaveAuthCode(phone string, code string, durationMinutes uint) {
	key := getAuthCodeRedisKey(phone)
	redis.Client.SetEx(key, code, durationMinutes*60)
}

func GetAuthCode(phone string) string {
	key := getAuthCodeRedisKey(phone)
	code, err := redis.Client.Get(key)
	if err != nil {
		logger.Info("ERROR! GetAuthCode() Code not found")
	}
	return code
}

func GetByToken(token string) (string, error) {
	key := getAuthTokenRedisKey(token)
	userInfo, err := redis.Client.Get(key)
	if err != nil {
		logger.Info("ERROR! GetByToken() userInfo not found")
		return "", errors.New("tokenNotFound")
	}
	return userInfo, nil
}
func SaveToken(phone string, token string, durationHours uint) {
	key := getAuthTokenRedisKey(token)
	redis.Client.SetEx(key, phone, durationHours*60*60)
}

package auth

import (
	"main/_lib/redis"
	"main/store/repositories/auth/errors"
)

var authCodePrefix string = "auth:code:"
var authTokenPrefix string = "auth:token:"

func getAuthCodeRedisKey(phone string) string {
	return authCodePrefix + phone
}

func getAuthTokenRedisKey(phone string) string {
	return authTokenPrefix + phone
}

func SaveAuthCode(phone string, code string, durationMinutes uint) error {
	key := getAuthCodeRedisKey(phone)
	err := redis.Client.SetEx(key, code, durationMinutes*60)
	if err != nil {
		return errors.NewSaveAuthCodeError(code, phone, durationMinutes, err)
	}
	return nil
}

func GetAuthCode(phone string) (string, error) {
	key := getAuthCodeRedisKey(phone)
	code, err := redis.Client.Get(key)
	if err != nil {
		return "", errors.NewGetAuthCodeError(phone, err)
	}
	return code, nil
}

func GetByToken(token string) (string, error) {
	key := getAuthTokenRedisKey(token)
	userInfo, err := redis.Client.Get(key)
	if err != nil {
		return "", errors.NewGetUserInfoError(token, err)
	}
	return userInfo, nil
}
func SaveToken(phone string, token string, durationHours uint) error {
	key := getAuthTokenRedisKey(token)
	err := redis.Client.SetEx(key, phone, durationHours*60*60)
	if err != nil {
		return errors.NewSaveTokenError(phone, token, durationHours, err)
	}
	return nil
}

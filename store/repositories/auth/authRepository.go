package auth

import "main/_lib/redis"

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
	redis.NewClient().SetEx(key, code, durationMinutes*60)
}

func GetAuthCode(phone string) string {
	key := getAuthCodeRedisKey(phone)
	value := redis.NewClient().Get(key)
	return value
}

func SaveToken(phone string, token string, durationHours uint) {
	key := getAuthTokenRedisKey(phone)
	redis.NewClient().SetEx(key, token, durationHours*60*60)
}

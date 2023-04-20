package logger

import (
	"encoding/json"
	"fmt"
	"time"
)

type AnyMap = map[string]interface{}

type LogMessage struct {
	Time    string      `json:"time"`
	Level   string      `json:"level"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload"`
}

func InfoPayload(message string, payload AnyMap) {
	logMessage := LogMessage{
		Time:    time.Now().Format(time.RFC3339),
		Level:   "info",
		Message: message,
	}
	if len(payload) > 0 {
		logMessage.Payload = payload
	}
	logMessageByte, err := json.Marshal(logMessage)
	if err == nil {
		fmt.Println(string(logMessageByte))
	}
}

func Info(message string) {
	InfoPayload(message, AnyMap{})
}

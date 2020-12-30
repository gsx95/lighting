package util

import (
	"encoding/json"
	"fmt"
	"time"
)

type logMessage struct {
	Timestamp string `json:"timestamp"`
	Type      string `json:"type"`
	Sender    string `json:"sender"`
	RequestID string `json:"request_id"`
	Message   string `json:"message"`
	Body      string `json:"body"`
}


func Log(actionType string, msg string, requestId string, body string) {

	message := logMessage{
		Timestamp: time.Now().Format(time.RFC3339),
		Type: actionType,
		Sender: "lighting",
		RequestID: requestId,
		Message: msg,
		Body: body,
	}

	jsonB, err := json.Marshal(message)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(jsonB))
}

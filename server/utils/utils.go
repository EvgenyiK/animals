package utils

import (
	"encoding/json"
	"net/http"
)

func Message(status bool, message string) (map[string]interface{}) {
	return map[string]interface{}{"status":status, "message":message}
}

func Respond(w http.Responsewriter, data map[string]interface{})  {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

//utils.go содержат удобные функции utils для создания  json сообщений и возврата  json ответа. 
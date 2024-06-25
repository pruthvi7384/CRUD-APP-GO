package utils

import (
	"crudApp/dto"
	"encoding/json"
	"fmt"
	"log"
)

/*
*
Response Prepare Func
*/
func ResponsePrepare(message string, status string, data interface{}) dto.Response {
	return dto.Response{ErrorDescription: message, Data: data, ErrorCode: status}
}

/*
*
Marshal Json Data
*/
func JsonFormater(data interface{}) string {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("[Helper] Error While Json Formate - %v", fmt.Sprint(err))
	}
	return string(jsonData)
}

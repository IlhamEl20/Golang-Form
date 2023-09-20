package services

import (
	"Si-KP/config"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func Auth_pusri_validation(username, password string) int {

	url := config.Env("URL_AUTH_PUSRI")

	iam_user := username
	iam_pass := password

	values := map[string]string{"username": iam_user, "password": iam_pass}

	jsonStr, _ := json.Marshal(values)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	// defer resp.Body.Close()

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	return resp.StatusCode
}

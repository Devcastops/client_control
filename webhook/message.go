package webhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Body struct {
	Content string `json:"content"`
}

func SendMessage(url, msg string) error {
	body, err := json.Marshal(Body{Content: msg})
	if err != nil {
		return err
	}
	resp, err := http.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		return err
	}
	fmt.Println(resp)
	return nil
}

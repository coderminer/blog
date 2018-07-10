package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

func (resp *Response) WriteJson(w http.ResponseWriter) {
	data, err := json.Marshal(resp)
	if err != nil {
		w.Write([]byte(`{Status:-1}`))
	} else {
		w.Write(data)
	}
}

package http

import (
	"bytes"

	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/udonetsm/client/models"
)

func CallUpdateNumber(target, newnumber string) {
	respObj := &models.RequestJSON{
		Target:  target,
		Upgrade: newnumber,
	}
	data, err := json.Marshal(respObj)
	if err != nil {
		log.Fatal(err)
	}
	resp := DoReq("http://127.0.0.1:8080/update/number", http.MethodPost, data)
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data), resp.Header.Get("ANSWER"))
}

func DoReq(url, method string, data []byte) *http.Response {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{
		Timeout: 3 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}

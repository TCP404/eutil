package eutil

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"
)

func TestPrintReq(t *testing.T) {
	var body = []byte(`{"key1": "value1", "key2": "value2"}`)
	req, _ := http.NewRequest(http.MethodPost, "http://httpbin.org/post", bytes.NewBuffer(body))
	req.AddCookie(&http.Cookie{
		Name:  "root",
		Value: "yes",
	})
	req.Header.Add("Host", "httpbin.org")
	req.Header.Set("Content-Type", "application/json")
	d := DumpReq(req)
	fmt.Println(string(d))

	res, _ := http.DefaultClient.Do(req)
	dr := DumpResp(res)
	fmt.Println(string(dr))
}

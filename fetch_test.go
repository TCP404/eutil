package eutil

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestFetch(t *testing.T) {
	var (
		url  = "http://httpbin.org/post"
		body = []byte(`{"key1": "value1", "key2": "value2"}`)
	)
	p := Fetch(url, body, nil).
		Then(func(res *http.Response) {
			fmt.Println(res.StatusCode)
		}).
		Then(func(res *http.Response) {
			b, _ := io.ReadAll(res.Body)
			fmt.Println(string(b))
			_ = res.Body.Close()
		}).
		Catch(func(err error) {
			fmt.Println(err.Error())
		})

	p.Await()
}

func BenchmarkFetch(b *testing.B) {
	var (
		url  = "http://httpbin.org/post"
		body = []byte(`{"key1": "value1", "key2": "value2"}`)
	)
	p := Fetch(url, body, nil).
		Then(func(res *http.Response) {
			// fmt.Println(res.StatusCode)
		}).
		Then(func(res *http.Response) {
			// b, _ := io.ReadAll(res.Body)
			// fmt.Println(string(b))
			// res.Body.Close()
		}).
		Catch(func(err error) {
			// fmt.Println(err.Error())
		})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p.Await()
	}
}

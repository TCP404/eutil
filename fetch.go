package eutil

import (
	"bytes"
	"net/http"
	"sync"
	"time"
)

type promise struct {
	res      *http.Response
	execFn   func()
	err      error
	pending  bool
	thenFns  []func(*http.Response)
	catchFns []func(error)
	wg       *sync.WaitGroup
}

func Fetch(url string, body []byte, headers map[string]string) *promise {
	pro := &promise{
		pending: true,
		wg:      &sync.WaitGroup{},
	}

	pro.execFn = func() {
		defer pro.wg.Done()
		req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
		if err != nil {
			pro.err = err
			pro.pending = false
		}
		for k, v := range headers {
			req.Header.Add(k, v)
		}
		client := &http.Client{
			Timeout: 3 * time.Second,
		}
		res, err := client.Do(req)
		if err != nil {
			pro.err = err
			pro.pending = false
		}
		pro.res = res
		pro.pending = false
	}
	return pro
}

func (p *promise) Then(fn func(*http.Response)) *promise {
	p.thenFns = append(p.thenFns, fn)
	return p
}

func (p *promise) Catch(fn func(err error)) *promise {
	p.catchFns = append(p.catchFns, fn)
	return p
}

func (p *promise) Await() {
	p.wg.Add(1)
	go p.execFn()
	p.wg.Wait()

	if p.err != nil {
		for _, catchFn := range p.catchFns {
			catchFn(p.err)
		}
	} else {
		for _, thenFn := range p.thenFns {
			thenFn(p.res)
		}
	}
}

package eutil

import (
	"bytes"
	"net/http"
	"sync"
	"time"

	"github.com/pkg/errors"
)

type (
	catchFunc = func(res *http.Response, err error)
	thenFunc  = func(*http.Response)
	promise   struct {
		res      *http.Response
		execFn   func()
		err      error
		pending  bool
		thenFns  []thenFunc
		catchFns []catchFunc
		wg       *sync.WaitGroup
	}
)

func Fetch(method, url string, body []byte, headers map[string]string, timeout time.Duration) *promise {
	if timeout == 0 {
		timeout = 3 * time.Second
	}
	pro := &promise{
		pending: true,
		wg:      &sync.WaitGroup{},
	}

	pro.execFn = func() {
		defer pro.wg.Done()
		pro.pending = false

		req, err := http.NewRequest(method, url, bytes.NewReader(body))
		if err != nil {
			pro.err = err
			return
		}

		for k, v := range headers {
			req.Header.Add(k, v)
		}
		client := &http.Client{Timeout: timeout}
		res, err := client.Do(req)
		if err != nil {
			pro.err = err
		} else if res.StatusCode != http.StatusOK {
			pro.err = errors.Errorf("status code is not ok: [%v]", res.Status)

		}
		pro.res = res
	}
	return pro
}

func (p *promise) Then(fn thenFunc) *promise {
	p.thenFns = append(p.thenFns, fn)
	return p
}

func (p *promise) Catch(fn catchFunc) *promise {
	p.catchFns = append(p.catchFns, fn)
	return p
}

func (p *promise) Await() (success bool) {
	p.wg.Add(1)
	go p.execFn()
	p.wg.Wait()

	if p.err != nil {
		success = false
		for _, catchFn := range p.catchFns {
			catchFn(p.res, p.err)
		}
	} else {
		success = true
		for _, thenFn := range p.thenFns {
			thenFn(p.res)
		}
	}
	return
}

package fetch

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"reflect"
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

		timeout time.Duration
		body    io.Reader
		headers map[string]string
		params  url.Values
	}
)

type FetchOption func(*promise)

func newPromise(opts ...FetchOption) *promise {
	pro := &promise{
		pending: true,
		wg:      &sync.WaitGroup{},
		headers: make(map[string]string),
		params:  make(url.Values),
		timeout: 3 * time.Second,
	}
	for _, opt := range opts {
		opt(pro)
	}
	return pro
}

func WithTimeout(timeout time.Duration) FetchOption {
	return func(p *promise) {
		p.timeout = timeout
	}
}

func WithByteBody(body []byte) FetchOption {
	return func(p *promise) {
		p.body = bytes.NewReader(body)
	}
}

func WithReaderBody(body io.Reader) FetchOption {
	return func(p *promise) {
		p.body = body
	}
}

func WithHeader(key, value string) FetchOption {
	return func(p *promise) {
		p.headers[key] = value
	}
}

func WithHeaders(headers map[string]string) FetchOption {
	return func(p *promise) {
		for k, v := range headers {
			p.headers[k] = v
		}
	}
}

func WithParam(key string, value string) FetchOption {
	return func(p *promise) {
		p.params[key] = []string{value}
	}
}

func WithArrParam(key string, value []string) FetchOption {
	return func(p *promise) {
		p.params[key] = value
	}
}

func WithParams(params map[string][]string) FetchOption {
	return func(p *promise) {
		for k, element := range params {
			p.params[k] = append(p.params[k], element...)
		}
	}
}

func Fetch(method, u string, opts ...FetchOption) *promise {
	pro := newPromise(opts...)
	pro.execFn = func() {
		defer pro.wg.Done()
		pro.pending = false

		uu, err := url.Parse(u)
		if err != nil {
			pro.err = err
			return
		}
		uu.RawQuery = pro.params.Encode()
		req, err := http.NewRequest(method, uu.String(), pro.body)
		if err != nil {
			pro.err = err
			return
		}

		for k, v := range pro.headers {
			req.Header.Add(k, v)
		}

		client := &http.Client{Timeout: pro.timeout}
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

func (p *promise) ThenWriteTo(responseSchema any) *promise {
	k := reflect.TypeOf(responseSchema)
	if k.Kind() != reflect.Ptr {
		p.err = errors.New("responseSchema must be a pointer")
		return p
	}
	if k.Elem().Kind() != reflect.Struct {
		p.err = errors.New("responseSchema must be a struct pointer")
		return p
	}

	p.thenFns = append(p.thenFns, func(res *http.Response) {
		buf := &bytes.Buffer{}
		_, err := io.Copy(buf, p.res.Body)
		if err != nil {
			p.err = err
			return
		}
		err = json.Unmarshal(buf.Bytes(), responseSchema)
		if err != nil {
			p.err = err
			return
		}
	})
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

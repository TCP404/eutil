package eutil

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
)

const (
	red_     = "\033[97;41m"
	green_   = "\033[97;42m"
	yellow_  = "\033[90;43m"
	blue_    = "\033[97;44m"
	magenta_ = "\033[97;45m"
	cyan_    = "\033[97;46m"
	white_   = "\033[90;47m"

	reset = "\033[0m"

	_red     = "\x1b[1;31;40m"
	_green   = "\x1b[1;32;40m"
	_yellow  = "\x1b[1;33;40m"
	_blue    = "\x1b[1;34;40m"
	_magenta = "\x1b[1;35;40m"
	_cyan    = "\x1b[1;36;40m"
	_white   = "\x1b[1;37;40m"
)

type dumpRequest struct {
	dump
	req *http.Request
}

type dump struct {
	buf bytes.Buffer
}

type dumpResponse struct {
	dump
	res *http.Response
}

func (d *dumpRequest) Method() *dumpRequest {
	switch d.req.Method {
	case http.MethodGet:
		d.buf.WriteString(blue_)
	case http.MethodPost:
		d.buf.WriteString(cyan_)
	case http.MethodPut:
		d.buf.WriteString(yellow_)
	case http.MethodDelete:
		d.buf.WriteString(red_)
	case http.MethodPatch:
		d.buf.WriteString(green_)
	case http.MethodHead:
		d.buf.WriteString(magenta_)
	case http.MethodOptions:
		d.buf.WriteString(white_)
	default:
		panic("Method not allow")
	}
	d.buf.WriteByte(' ')
	d.buf.WriteString(d.req.Method)
	d.buf.WriteByte(' ')
	d.buf.WriteString(reset)
	return d
}

func (d *dumpResponse) StatusCode() *dumpResponse {
	switch {
	case d.res.StatusCode >= 200 && d.res.StatusCode < 300:
		d.buf.WriteString(green_)
	case d.res.StatusCode >= 300 && d.res.StatusCode < 400:
		d.buf.WriteString(white_)
	case d.res.StatusCode >= 400 && d.res.StatusCode < 500:
		d.buf.WriteString(yellow_)
	default:
		d.buf.WriteString(red_)
	}
	d.buf.WriteByte(' ')
	d.buf.WriteString(d.res.Status)
	d.buf.WriteByte(' ')
	d.buf.WriteString(reset)
	return d
}

func (d *dumpRequest) Url() *dumpRequest {
	d.buf.WriteString(_cyan + d.req.URL.String() + reset)
	return d
}

func (d *dump) KV(k, v string) *dump {
	d.buf.WriteString(_blue + k + reset)
	d.buf.WriteString(": ")
	d.buf.WriteString(_green + v + reset)
	return d
}

func (d *dump) Body(dest *bytes.Buffer) *dump {
	var buf bytes.Buffer
	json.Indent(&buf, dest.Bytes(), "", "    ")
	d.buf.WriteString(_yellow)
	d.buf.Write(buf.Bytes())
	d.buf.WriteString(reset)
	return d
}

func drainBody(b io.ReadCloser) (r1, r2 io.ReadCloser, err error) {
	if b == nil || b == http.NoBody {
		// No copying needed. Preserve the magic sentinel meaning of NoBody.
		return http.NoBody, http.NoBody, nil
	}
	var buf bytes.Buffer
	if _, err = buf.ReadFrom(b); err != nil {
		return nil, b, err
	}
	if err = b.Close(); err != nil {
		return nil, b, err
	}
	return io.NopCloser(&buf), io.NopCloser(bytes.NewReader(buf.Bytes())), nil
}

func (d *dump) HTTPVersion(major, minor string) *dump {
	d.buf.WriteString(_white + "HTTP/")
	d.buf.WriteString(major)
	d.buf.WriteByte('.')
	d.buf.WriteString(minor)
	d.buf.WriteString(reset)
	return d
}

func (d *dump) space() *dump {
	d.buf.WriteByte(' ')
	return d
}

func (d *dump) newLine() *dump {
	d.buf.WriteString("\r\n")
	return d
}
func DumpReq(req *http.Request) []byte {
	if req == nil {
		return nil
	}
	d := &dumpRequest{req: req}
	d.Method().space()
	d.Url().space()
	d.HTTPVersion(strconv.Itoa(d.req.ProtoMajor), strconv.Itoa(d.req.ProtoMinor)).newLine()
	for k, v := range req.Header {
		d.KV(k, strings.Join(v, ";")).newLine()
	}
	if req.Body != nil {
		d.newLine()
		var (
			save io.ReadCloser
			err  error
			b    []byte
			buf  = bytes.NewBuffer(b)
		)
		save, req.Body, err = drainBody(req.Body)
		if err != nil {
			return nil
		}
		io.Copy(buf, save)
		d.Body(buf).newLine()
	}
	return d.buf.Bytes()
}

func DumpResp(resp *http.Response) []byte {
	if resp == nil {
		return nil
	}
	d := &dumpResponse{res: resp}
	d.HTTPVersion(strconv.Itoa(d.res.ProtoMajor), strconv.Itoa(d.res.ProtoMinor)).space()
	d.StatusCode().newLine()

	for k, v := range d.res.Header {
		d.KV(k, strings.Join(v, ";")).newLine()
	}
	if d.res.Body != nil {
		d.newLine()
		var (
			b   []byte
			buf = bytes.NewBuffer(b)
		)
		reqBody, _ := io.ReadAll(d.res.Body)
		json.Indent(buf, reqBody, "", "    ")
		d.Body(buf).newLine()
	}
	return d.buf.Bytes()
}

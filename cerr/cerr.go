package cerr

import (
	"fmt"
	"runtime"
	"strings"
)

type Err struct {
	StackPCS
	Msg
	text string
}

func Wrap(err error, enStr ApiErrMsg, skip ...int) *Err {
	return &Err{
		StackPCS: Callers(skip...),
		Msg: Msg{
			enStr: enStr,
		},
		text: fmt.Sprintf("%v: %v", err.Error(), enStr),
	}
}

func Wrapf(err error, enStr ApiErrMsg, args ...any) *Err {
	return &Err{
		StackPCS: Callers(),
		Msg: Msg{
			enStr: enStr,
		},
		text: fmt.Sprintf("%v: %v", fmt.Sprintf(err.Error(), args...), enStr),
	}
}

func New(enStr ApiErrMsg, skip ...int) *Err {
	return &Err{
		StackPCS: Callers(skip...),
		Msg: Msg{
			enStr: enStr,
		},
		text: string(enStr),
	}
}

func Newf(enStr ApiErrMsg, args ...any) *Err {
	return &Err{
		StackPCS: Callers(),
		Msg: Msg{
			enStr: enStr,
		},
		text: fmt.Sprintf(string(enStr), args...),
	}
}

func (e *Err) Error() string {
	return e.text
}

func (e *Err) Stack() StackPCS {
	return e.StackPCS
}

func (e *Err) StackString() string {
	var build strings.Builder
	for i, pc := range e.StackPCS {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		fnName := fn.Name()
		build.WriteString(fmt.Sprintf("\n%d. %s:%d", i, file, line))
		build.WriteString(fmt.Sprintf("\n\t %s", fnName))
	}
	return build.String()
}

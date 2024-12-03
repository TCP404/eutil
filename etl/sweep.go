package etl

import (
	"errors"
	"fmt"
)

type SweepInfo struct {
	err error
}

type defaultSweeper struct {
	sweepC chan SweepInfo
}

var _ Runner = (*defaultSweeper)(nil)

func NewDefaultSweeper(sweepC chan SweepInfo) Runner {
	return &defaultSweeper{sweepC: sweepC}
}

// Run of defaultSweeper sweep the error. It will collect all error after the channel is closed
// and panic with the error list. If you want panic when error occurs, you can throw the panic
// in ExtractFunc, TransformFunc, LoadFunc, or Reporter.Run.
func (s defaultSweeper) Run() {
	// Collect all the sweeps
	var sweeps []error
	for sweepInfo := range s.sweepC {
		sweeps = append(sweeps, sweepInfo.err)
	}
	if len(sweeps) == 0 {
		return
	}

	fmt.Println("xxxxxx sweeps xxxxxx ")
	for _, sweepInfo := range sweeps {
		fmt.Println(sweepInfo)
	}
	fmt.Println("xxxxxx sweeps xxxxxx ")
	panic(errors.Join(sweeps...))
}

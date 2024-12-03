package etl

import (
	"fmt"
	"runtime"

	"github.com/sourcegraph/conc/pool"
)

func (e *ETL[E, L]) Transform() {
	defer func() {
		if err := recover(); err != nil {
			panic(fmt.Sprintf("Transform panic: %v", err))
		}
	}()

	var (
		total     int
		batch     = make([]E, 0, e.TransformBatchSize)
		p         = pool.New().WithMaxGoroutines(max(runtime.NumCPU()*2, e.transformerNum))
		batchFunc = func(data []E) func() {
			return func() {
				target, err := e.transformFunc(data)
				if err != nil {
					e.sweepC <- SweepInfo{err: err}
					return
				}
				for _, tar := range target {
					e.t2l <- tar
				}
			}
		}
	)
	for datum := range e.e2t {
		datum := datum
		if len(batch) < e.TransformBatchSize {
			batch = append(batch, datum)
			total++
			continue
		}
		p.Go(batchFunc(batch))
		e.reportC <- ReportInfo{Name: "Transform", Value: total, Status: StatusProcessing}
		batch = batch[:0]
	}
	if len(batch) > 0 {
		p.Go(batchFunc(batch))
		e.reportC <- ReportInfo{Name: "Transform", Value: total, Status: StatusProcessing}
	}
	p.Wait()
	e.reportC <- ReportInfo{Name: "Transform", Value: total, Status: StatusComplete}
}

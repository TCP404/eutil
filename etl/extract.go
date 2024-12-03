package etl

import (
	"context"
	"fmt"
)

func (etl *ETL[E, L]) Extract(ctx context.Context) {
	defer func() {
		if err := recover(); err != nil {
			panic(fmt.Sprintf("Extract panic: %v", err))
		}
	}()

	iter, err := etl.extractFunc(ctx)
	if err != nil {
		panic(err)
	}

	var total int
	for iter.Next() {
		select {
		case <-ctx.Done():
			etl.reportC <- ReportInfo{Name: "Extract", Value: total, Status: StatusComplete}
			return
		default:
			record := iter.Value()
			if err := iter.Err(); err != nil {
				etl.sweepC <- SweepInfo{err: err}
				continue
			}

			etl.e2t <- record
			total++
			if total%etl.extractBatchSize == 0 {
				etl.reportC <- ReportInfo{Name: "Extract", Value: total, Status: StatusProcessing}
			}
		}
	}
	etl.reportC <- ReportInfo{Name: "Extract", Value: total, Status: StatusComplete}
}

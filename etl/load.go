package etl

import "fmt"

func (etl *ETL[E, L]) Load() {
	defer func() {
		if err := recover(); err != nil {
			panic(fmt.Sprintf("Load panic: %v", err))
		}
	}()

	var (
		batch     = make([]L, 0, etl.LoadBatchSize)
		unique    = make(map[string]struct{})
		uniqueKey = etl.loadUniqFunc
		total     int
	)

	for v := range etl.t2l {
		v := v
		if uniqueKey != nil {
			if _, exist := unique[uniqueKey(v)]; exist {
				continue
			}
			unique[uniqueKey(v)] = struct{}{}
		}

		batch = append(batch, v)
		if len(batch) < etl.LoadBatchSize {
			continue
		}

		cnt, err := etl.loadFunc(batch)
		if err != nil {
			etl.sweepC <- SweepInfo{err: err}
		} else {
			total += cnt
			etl.reportC <- ReportInfo{Name: "Load", Value: total, Status: StatusProcessing}
		}
		batch = batch[:0]
	}

	// handle rest
	cnt, err := etl.loadFunc(batch)
	if err != nil {
		etl.sweepC <- SweepInfo{err: err}
	} else {
		total += cnt
		etl.reportC <- ReportInfo{Name: "Load", Value: total, Status: StatusProcessing}
	}

	etl.reportC <- ReportInfo{Name: "Load", Value: total, Status: StatusComplete}
}

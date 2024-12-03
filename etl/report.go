package etl

import (
	"fmt"
	"sync"

	"github.com/TCP404/eutil"
)

const (
	StatusProcessing = "processing"
	StatusComplete   = "complete"
)

type Reporter interface {
	Runner
	Progress() uint64
}

type ReportInfo struct {
	Name   string
	Value  int
	Status string
}

type progressReporter struct {
	reportC         chan ReportInfo
	total           uint64
	mux             sync.RWMutex
	progressPercent uint64
}

var _ Reporter = (*progressReporter)(nil)

func NewProgressReporter(reportC chan ReportInfo, total uint64) Reporter {
	return &progressReporter{
		reportC: reportC,
		total:   total,
		mux:     sync.RWMutex{},
	}
}

func ProgressReporterFactory(total uint64) ReporterFactory {
	return func(c *chan ReportInfo) Reporter {
		return NewProgressReporter(*c, total)
	}
}

func (d *progressReporter) Run() {
	var finishedReport []ReportInfo

	for r := range d.reportC {
		if r.Status == StatusComplete {
			finishedReport = append(finishedReport, r)
		}

		// slog.Info(fmt.Sprintf("%s task count: %v\n", r.Name, r.Value))
		if r.Name == "Extract" {
			d.setProgress(uint64(r.Value))
		}
	}

	fmt.Println("======= ETL report =======")
	for _, v := range finishedReport {
		fmt.Printf("Count of %v: %v\n", v.Name, v.Value)
	}
	fmt.Println("======= ETL report =======")
}

func (d *progressReporter) setProgress(currTotal uint64) {
	d.mux.Lock()
	defer d.mux.Unlock()
	d.progressPercent = min(100, eutil.If(d.total <= 0, 0, currTotal*100/d.total))
}

func (d *progressReporter) Progress() uint64 {
	d.mux.RLock()
	defer d.mux.RUnlock()
	return d.progressPercent
}

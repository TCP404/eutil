package etl

import (
	"context"
	"errors"
	"fmt"

	"github.com/sourcegraph/conc"
)

type (
	Runner          interface{ Run() }
	Iterator[T any] interface {
		Next() bool
		Value() T
		Err() error
	}

	ExtractFunc[E any]      func(context.Context) (Iterator[E], error)
	TransformFunc[E, L any] func(datum []E) ([]L, error)
	LoadFunc[L any]         func(batch []L) (int, error)
	LoadUniqFunc[L any]     func(L) string

	ReporterFactory  func(*chan ReportInfo) Reporter
	SweeperFactory   func(*chan SweepInfo) Runner
	Option[E, L any] func(*ETL[E, L])
)

type ETL[E, L any] struct {
	extractorNum       int
	extractBatchSize   int
	extractFunc        ExtractFunc[E]
	e2t                chan E
	transformerNum     int
	TransformBatchSize int
	transformFunc      TransformFunc[E, L]
	t2l                chan L
	loaderNum          int
	LoadBatchSize      int
	loadFunc           LoadFunc[L]
	loadUniqFunc       func(L) string

	sweepC  chan SweepInfo
	sweeper Runner

	reportC  chan ReportInfo
	reporter Reporter
}

func WithE2TSize[E, L any](size int) Option[E, L] {
	return func(etl *ETL[E, L]) { etl.e2t = make(chan E, size) }
}
func WithT2LSize[E, L any](size int) Option[E, L] {
	return func(etl *ETL[E, L]) { etl.t2l = make(chan L, size) }
}
func WithSweepCSize[E, L any](size int) Option[E, L] {
	return func(etl *ETL[E, L]) { etl.sweepC = make(chan SweepInfo, size) }
}
func WithReportCSize[E, L any](size int) Option[E, L] {
	return func(etl *ETL[E, L]) { etl.reportC = make(chan ReportInfo, size) }
}
func WithSweeper[E, L any](factory SweeperFactory) Option[E, L] {
	return func(e *ETL[E, L]) {
		if e.sweepC == nil {
			e.sweepC = make(chan SweepInfo, 100)
		}
		e.sweeper = factory(&e.sweepC)
	}
}
func WithReporter[E, L any](factory ReporterFactory) Option[E, L] {
	return func(e *ETL[E, L]) {
		if e.reportC == nil {
			e.reportC = make(chan ReportInfo, 100)
		}
		e.reporter = factory(&e.reportC)
	}
}
func WithExtractorNum[E, L any](num int) Option[E, L] {
	return func(etl *ETL[E, L]) { etl.extractorNum = num }
}
func WithTransformerNum[E, L any](num int) Option[E, L] {
	return func(etl *ETL[E, L]) { etl.transformerNum = num }
}
func WithLoaderNum[E, L any](num int) Option[E, L] {
	return func(etl *ETL[E, L]) { etl.loaderNum = num }
}
func WithExtractBatchSize[E, L any](size int) Option[E, L] {
	return func(etl *ETL[E, L]) { etl.extractBatchSize = size }
}
func WithTransformBatchSize[E, L any](size int) Option[E, L] {
	return func(etl *ETL[E, L]) { etl.TransformBatchSize = size }
}
func WithLoadBatchSize[E, L any](size int) Option[E, L] {
	return func(etl *ETL[E, L]) { etl.LoadBatchSize = size }
}

func New[E, L any](extractFunc ExtractFunc[E], transformFunc TransformFunc[E, L], loadFunc LoadFunc[L], opts ...Option[E, L]) *ETL[E, L] {
	obj := ETL[E, L]{
		extractorNum:   1,
		transformerNum: 1,
		loaderNum:      1,

		extractFunc:   extractFunc,
		transformFunc: transformFunc,
		loadFunc:      loadFunc,

		extractBatchSize:   1,
		TransformBatchSize: 1,
		LoadBatchSize:      1,
	}
	for _, opt := range opts {
		opt(&obj)
	}

	if obj.e2t == nil {
		obj.e2t = make(chan E, 100)
	}
	if obj.t2l == nil {
		obj.t2l = make(chan L, 100)
	}

	if obj.sweeper == nil {
		if obj.sweepC == nil {
			obj.sweepC = make(chan SweepInfo, 100)
		}
		obj.sweeper = NewDefaultSweeper(obj.sweepC)
	}
	if obj.reporter == nil {
		if obj.reportC == nil {
			obj.reportC = make(chan ReportInfo, 100)
		}
		obj.reporter = NewProgressReporter(obj.reportC, 0)
	}

	return &obj
}

func (e *ETL[E, L]) Run(ctx context.Context) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Run panic: %v", err)
		}
	}()

	if err := e.check(); err != nil {
		return err
	}

	var wg = conc.NewWaitGroup()

	wg.Go(func() {
		e.sweeper.Run()
	})

	wg.Go(func() {
		e.reporter.Run()
	})

	wg.Go(func() {
		defer func() { close(e.e2t) }()
		e.Extract(ctx)
	})

	wg.Go(func() {
		defer func() { close(e.t2l) }()
		e.Transform()
	})

	wg.Go(func() {
		defer func() { close(e.sweepC); close(e.reportC) }()
		e.Load()
	})

	wg.Wait()

	return nil
}

func (etl *ETL[E, L]) check() error {
	if etl.extractFunc == nil ||
		etl.transformFunc == nil ||
		etl.loadFunc == nil {
		return errors.New("ExtractFunc, TransformFunc, LoadFunc can not be nil")
	}
	return nil
}

func (e *ETL[E, L]) Progress() uint64 {
	return e.reporter.Progress()
}

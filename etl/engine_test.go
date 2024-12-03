package etl

import (
	"context"
	"strconv"
	"testing"
)

type iterator[E any] struct {
	data  []E
	index int
}

func (i *iterator[E]) Next() bool {
	return i.index < len(i.data)
}

func (i *iterator[E]) Value() E {
	datum := i.data[i.index]
	i.index++
	return datum
}

func (i *iterator[E]) Err() error {
	return nil
}

func TestETL(t *testing.T) {

	extractFunc := func(ctx context.Context) (Iterator[int], error) {
		dataSize := 15
		data := make([]int, 0, dataSize)
		for i := 1; i <= dataSize; i++ {
			data = append(data, i)
		}

		return &iterator[int]{data: data, index: 0}, nil
	}
	transformFunc := func(data []int) ([]string, error) {
		ret := make([]string, 0, len(data))
		for _, datum := range data {
			ret = append(ret, strconv.Itoa(datum))
		}
		return ret, nil
	}
	loadFunc := func(data []string) (int, error) {
		for _, datum := range data {
			t.Logf("load: %v", datum)
		}
		return len(data), nil
	}
	etl := New(
		extractFunc, transformFunc, loadFunc,
		WithExtractorNum[int, string](2),
		WithTransformerNum[int, string](5),
		WithLoaderNum[int, string](1),
		WithExtractBatchSize[int, string](2),
		WithTransformBatchSize[int, string](3),
		WithLoadBatchSize[int, string](4),
		WithReporter[int, string](ProgressReporterFactory(100)),
	)
	err := etl.Run(context.TODO())
	if err != nil {
		t.Errorf("etl.Run() error = %v", err)
	}
	t.Log("etl.Run() success")
}

package packer

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type TestStruct struct {
	AnoStruct
	F1            string    `json:"f1"`
	F2            int       `json:"f2"`
	F3            SubStruct `json:"f3"`
	AnoJsonStruct `json:"ano-json"`
}

type SubStruct struct {
	SF1 string        `json:"sf1"`
	SF2 float64       `json:"sf2"`
	SF3 time.Time     `json:"sf3"`
	SF4 *time.Time    `json:"sf4"`
	SF5 time.Duration `json:"sf5"`
	SF6 time.Month    `json:"sf6"`
}

type AnoStruct struct {
	SF1 string `json:"af1"`
	SF2 bool   `json:"af2"`
}

type AnoJsonStruct struct {
	SF1 string  `json:"af1"`
	SF2 float64 `json:"af2"`
}

func TestPacker(t *testing.T) {
	tdate := time.Date(2022, time.January, 1, 0, 0, 0, 0, time.Local)
	tdate2 := time.Date(1997, time.July, 1, 0, 0, 0, 0, time.Local)
	tt := TestStruct{
		F1: "A",
		F2: 1_000_000_000_000_000,
		F3: SubStruct{
			SF1: "SA",
			SF2: 0.001,
			SF3: tdate,
			SF4: &tdate,
		},
	}

	ta := Parse(&tt)

	sf4 := ta.Get("f3.sf3")
	t.Logf("%v", sf4.ToString())
	sf4.Set(&tdate2)
	t.Logf("%v", sf4.ToString())
	t.Log(sf4.LastError())
}

func TestAgent_Set(t *testing.T) {
	newYear := time.Date(2022, time.January, 1, 0, 0, 0, 0, time.Local)
	nationalDay := time.Date(2022, time.October, 1, 0, 0, 0, 0, time.Local)
	ts := TestStruct{
		F1: "A",
		F2: 1_000_000_000_000_000,
		F3: SubStruct{
			SF1: "SA",
			SF2: 0.001,
			SF3: newYear,
			SF4: &newYear,
		},
	}

	tp := Parse(&ts)
	require.Nil(t, tp.LastError())

	type args struct {
		path  string
		value any
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			name: "time.Time",
			args: args{
				path:  "f3.sf3",
				value: nationalDay,
			},
			want: "2022-10-01 00:00:00 +0800 CST",
		},
		{
			name: "*time.Time",
			args: args{
				path:  "f3.sf4",
				value: &nationalDay,
			},
			want: "2022-10-01 00:00:00 +0800 CST",
		},
		{
			name: "string",
			args: args{
				path:  "f1",
				value: "ABC",
			},
			want: "ABC",
		},
		{
			name: "int",
			args: args{
				path:  "f2",
				value: 18,
			},
			want: 18,
		},
		{
			name: "float",
			args: args{
				path:  "f3.sf2",
				value: 0.112,
			},
			want: 0.112,
		},
		{
			name: "bool",
			args: args{
				path:  "af2",
				value: true,
			},
			want: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			f := tp.Get(tc.args.path).Set(tc.args.value)
			require.Nil(t, f.LastError())
			switch f.ValueType() {
			case StringValue, ObjectValue:
				got := f.ToString()
				if !reflect.DeepEqual(got, tc.want) {
					t.Errorf("got %v, want %v", got, tc.want)
				}
			case NumberValue:
				switch {
				case f.Val().CanInt():
					got := f.ToInt()
					if !reflect.DeepEqual(got, tc.want) {
						t.Errorf("got %v, want %v", got, tc.want)
					}
				case f.Val().CanUint():
					got := f.ToUint()
					if !reflect.DeepEqual(got, tc.want) {
						t.Errorf("got %v, want %v", got, tc.want)
					}
				case f.Val().CanFloat():
					got := f.ToFloat64()
					if !reflect.DeepEqual(got, tc.want) {
						t.Errorf("got %v, want %v", got, tc.want)
					}
				default:
					t.Fatal("no match type.")
				}
			case BoolValue:
				got := f.ToBool()
				if !reflect.DeepEqual(got, tc.want) {
					t.Errorf("got %v, want %v", got, tc.want)
				}
			default:
				t.Fatal("no match type.")
			}
		})
	}
}

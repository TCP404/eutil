package packer

import (
	"fmt"
)

type ExampleStruct struct {
	AnoSonStruct
	F1               string    `json:"f1"`
	F2               int       `json:"f2"`
	F3               SonStruct `json:"f3"`
	AnoJsonSonStruct `json:"ano-json"`
}

type SonStruct struct {
	SF1 string  `json:"sf1"`
	SF2 float64 `json:"sf2"`
}

type AnoSonStruct struct {
	SF1 string `json:"af1"`
	SF2 bool   `json:"af2"`
}

type AnoJsonSonStruct struct {
	SF1 string  `json:"af1"`
	SF2 float64 `json:"af2"`
}

func ExampleParse() {
	ts := TestStruct{
		F1: "A",
		F2: 1_000_000_000_000_000,
		F3: SubStruct{
			SF1: "SA",
			SF2: 0.001,
		},
	}

	parse := Parse(&ts)
	if err := parse.LastError(); err != nil {
		fmt.Println(err.Error())
	}
	f1 := parse.Get("f1").ToString()
	f2 := parse.Get("f2").ToInt()
	sf1 := parse.Get("f3.sf1").ToString()
	sf2 := parse.Get("f3.sf2").ToFloat64()
	fmt.Println(f1)
	fmt.Println(f2)
	fmt.Println(sf1)
	fmt.Println(sf2)
	// Output:
	// A
	// 1000000000000000
	// SA
	// 0.001
}

func ExampleParseWithTD() {
	ts := ExampleStruct{
		F1: "A",
		F2: 1_000_000_000_000_000,
		F3: SonStruct{
			SF1: "SA",
			SF2: 0.001,
		},
	}
	td, err := DescribeType(&ts)
	if err != nil {
		fmt.Println(err.Error())
	}

	parse := ParseWithTD(&ts, td)
	if err := parse.LastError(); err != nil {
		fmt.Println(err.Error())
	}
	f1 := parse.Get("f1").ToString()
	f2 := parse.Get("f2").ToInt()
	sf1 := parse.Get("f3.sf1").ToString()
	sf2 := parse.Get("f3.sf2").ToFloat64()
	fmt.Println(f1)
	fmt.Println(f2)
	fmt.Println(sf1)
	fmt.Println(sf2)
	// Output:
	// A
	// 1000000000000000
	// SA
	// 0.001
}

func Example() {
	es := ExampleStruct{
		F1: "A",
		F2: 1_000_000_000_000_000,
		F3: SonStruct{
			SF1: "SA",
			SF2: 0.001,
		},
	}

	// ParseWithTD need a type descriptor. You can get a type descriptor
	// by call DescribeType(). You also use Parse() and it will call the
	// DescribeType() first.
	parse := Parse(&es)
	if err := parse.LastError(); err != nil {
		fmt.Println(err.Error())
	}

	// Now you have an object of type any. You can call the Get() method,
	// it will return an agent of the field. Then you can take the information
	// of field type and call the corresponding conversion method . You can
	// use the parse.Set("f1", "B") to set the field's value, or use
	// f1.Set("B") conveniently.
	// Don't forget to call the LastError() method to check the error.
	f1 := parse.Get("f1")
	switch f1.ValueType() {
	case StringValue:
		fmt.Println(f1.ToString())
		f1.Set("B")
		if err := f1.LastError(); err != nil {
			fmt.Println(err)
		}
		fmt.Println(f1.ToString())

	case NumberValue:
		// if the field is not type int, that will return zero value.
		fmt.Println(f1.ToInt())
	case BoolValue:
		fmt.Println(f1.ToBool())
	}

	// You can use Marshal() method in direct.
	b := parse.Marshal()

	es2 := &ExampleStruct{}
	Parse(es2).Unmarshal(b)

	fmt.Println(string(b))
	fmt.Println(es2.F3.SF2)
	// Output:
	// A
	// B
	// {"af1":"","af2":false,"f1":"B","f2":1000000000000000,"f3":{"sf1":"SA","sf2":0.001},"ano-json":{"af1":"","af2":0}}
	// 0.001
}

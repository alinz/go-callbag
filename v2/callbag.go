package callbag

type Value = interface{}

type Callbag interface {
	Register(Callbag)
	Push(Value)
	Done(error)
}

//
// FromSlice
//

type fromSlice struct {
	i      int
	values []interface{}
	sink   Callbag
}

func (f *fromSlice) Register(sink Callbag) {
	f.sink = sink
	sink.Register(f)
}

func (f *fromSlice) Push(Value) {
	if f.i < len(f.values) {
		f.sink.Push(f.values[f.i])
		f.i++
	}
}

func (f *fromSlice) Done(error) {}

func FromValues(values ...interface{}) Callbag {
	return &fromSlice{
		values: values,
	}
}

//
// ForEach
//

type forEach struct {
	op     func(Value)
	source Callbag
}

func (f *forEach) Register(source Callbag) {
	f.source = source
}

func (f *forEach) Push(Value) {}
func (f *forEach) Done(error) {}

func ForEach(op func(Value)) Callbag {
	return &forEach{op: op}
}

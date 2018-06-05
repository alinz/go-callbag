package callbag

func FromIter(arr ...interface{}) Source {
	return func(p Payload) {
		switch v := p.(type) {
		case Greets:
			i := 0
			sink := v.Source()

			sink(NewGreets(func(p Payload) {
				switch p.(type) {
				case Data:
					if i < len(arr) {
						val := arr[i]
						i++
						sink(NewData(val))
					} else {
						sink(NewTerminate(nil))
					}
				}
			}))

		default:
			return
		}
	}
}
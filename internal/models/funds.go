package models

type Funds struct {
	value int
}

func (f *Funds) SetValue(v int) {
	f.value = v
}

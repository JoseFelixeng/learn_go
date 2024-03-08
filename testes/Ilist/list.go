package list

type Ilist interface {
	Add(value int)
	AddOnIndex(value int, index int) error
	RemoveOnIndex(index int) error
	Get(index int) (int, error)
	Set(Value int, index int) error
	Size() int
}

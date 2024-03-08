package list

import "errors"

type ArrayList struct {
	values   []int
	inserted int
}

func (list *ArrayList) Init(size int) error {
	if size > 0 {
		list.values = make([]int, size)
		return nil
	} else {
		return errors.New("Can't init ArrayList with size <= 0.")

	}
}

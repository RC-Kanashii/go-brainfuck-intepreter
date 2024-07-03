package intepreter

import "errors"

type Stack []int

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Size() int {
	return len(*s)
}

func (s *Stack) Top() int {
	if s.IsEmpty() {
		panic(errors.New("stack is empty"))
	}
	return (*s)[len(*s)-1]
}

func (s *Stack) Push(i int) {
	*s = append(*s, i)
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		panic(errors.New("stack is empty"))
	}
	top := s.Top()
	*s = (*s)[:len(*s)-1]
	return top
}

func (s *Stack) Has(i int) bool {
	for _, v := range *s {
		if v == i {
			return true
		}
	}
	return false
}

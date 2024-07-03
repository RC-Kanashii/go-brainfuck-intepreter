package intepreter

import "fmt"

type tape struct {
	data []int
	ptr  int
}

func (t *tape) Get() int {
	return t.data[t.ptr]
}

func (t *tape) Incr() {
	t.data[t.ptr]++
}

func (t *tape) Decr() {
	t.data[t.ptr]--
}

func (t *tape) MoveRight() {
	t.ptr++
	if t.ptr >= len(t.data) {
		t.data = append(t.data, 0)
	}
}

func (t *tape) MoveLeft() {
	t.ptr--
	if t.ptr < 0 {
		t.resize()
	}
}

func (t *tape) Input() {
	var i int
	_, err := fmt.Scanf("Input an integer: %d", &i)
	if err != nil {
		panic(err)
	}
	t.data[t.ptr] = i
}

func (t *tape) Output() {
	fmt.Println(string(rune(t.data[t.ptr])))
}

// resize resizes the tape if the pointer goes out of the left bound.
func (t *tape) resize() {
	t.ptr += cap(t.data)
	newData := make([]int, cap(t.data)*2)
	copy(newData[cap(t.data)-1:], t.data)
	t.data = newData
}

func createTape(length int) *tape {
	return &tape{
		data: make([]int, length),
		ptr:  length / 2,
	}
}

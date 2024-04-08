package main

type Stack struct {
	data     []int
	capacity int
}

func NewStack() *Stack {
	return &Stack{
		data:     make([]int, 0),
		capacity: 10, // Start with initial capacity of 2
	}
}

func (s *Stack) Push(item int) {
	if len(s.data) == s.capacity {
		s.capacity *= 2
		newData := make([]int, len(s.data), s.capacity)
		copy(newData, s.data)
		s.data = newData
	}
	s.data = append(s.data, item)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.data) == 0 {
		return 0, false
	}
	item := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return item, true
}

func (s *Stack) Peek() (int, bool) {
	if len(s.data) == 0 {
		return 0, false
	}
	return s.data[len(s.data)-1], true
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

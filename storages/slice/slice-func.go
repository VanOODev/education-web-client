package slice

import "fmt"

func NewSliceStorage() *SliceStorage {
	return &SliceStorage{
		st: nil,
	}
}

func (s *SliceStorage) Add(data int64) {
	s.st = append(s.st, data)
}

func (s *SliceStorage) Delete(index int64) {
	s.st = append(s.st[:index], s.st[index+1:]...)
}

func (s *SliceStorage) Get(index int64) (data int64) {
	data = s.st[index]
	return data
}

func (s *SliceStorage) Print() {
	fmt.Println(s.st)
}

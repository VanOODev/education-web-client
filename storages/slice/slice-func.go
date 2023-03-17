package slice

import "fmt"

func NewSliceStorage() *SliceStorage {
	return &SliceStorage{
		st: nil,
	}
}

func (s *SliceStorage) SlAdd(data int64) {
	s.st = append(s.st, data)
}

func (s *SliceStorage) SlDelete(index int64) {
	s.st = append(s.st[:index], s.st[index+1:]...)
}

func (s *SliceStorage) SlGet(index int64) (data int64) {
	data = s.st[index]
	return data
}

func (s *SliceStorage) SlPrint() {
	fmt.Println(s.st)
}

package main

import (
	"errors"
	"sync"
)

type Employee struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Sex    string `json:"sex"`
	Age    int    `json:"age"`
	Salary int    `json:"salary"`
}

type Storage interface {
	Insert(e *Employee)
	Get(id int) (Employee, error)
	Update(id int, e Employee)
	Delete(id int)
}

type MemoryStorage struct {
	counter int
	data    map[int]Employee
	sync.Mutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data:    make(map[int]Employee),
		counter: 1,
	}
}

func (s *MemoryStorage) Insert(e *Employee) {
	s.Lock()
	defer s.Unlock()
	e.ID = s.counter
	s.data[e.ID] = *e
	s.counter++
}

func (s *MemoryStorage) Get(id int) (Employee, error) {
	s.Lock()
	defer s.Unlock()
	val, ok := s.data[id]
	if !ok {
		return val, errors.New("Invalid ID")
	}
	return val, nil
}

func (s *MemoryStorage) Update(id int, e Employee) {
	s.Lock()
	defer s.Unlock()
	s.data[id] = e
}

func (s *MemoryStorage) Delete(id int) {
	s.Lock()
	defer s.Unlock()
	delete(s.data, id)
}

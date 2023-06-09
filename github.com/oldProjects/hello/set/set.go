package set

import (
	"errors"

	"github.com/google/uuid"
)

type Set struct {
	Storage map[string]string
	Count   int
}

func (s *Set) initSet() {
	s.Storage = make(map[string]string)
}

func (s *Set) AddToSet(value string) string {
	if s.Storage == nil {
		s.initSet()
	}
	uuid := uuid.New().String()
	s.Count++
	s.Storage[uuid] = value

	return uuid
}

func (s *Set) GetValueByID(uuid string) (string, error) {
	value, ok := s.Storage[uuid]
	if !ok {
		return "", errors.New("value is not found")
	}

	return value, nil
}

func (s *Set) GetAll() ([]string, error) {
	if s.Storage == nil {
		return []string{}, errors.New("values are not found")
	}
	values := []string{}
	for _, v := range s.Storage {
		values = append(values, v)
	}

	return values, nil
}

func (s *Set) GetCount() int {
	return s.Count
}

func (s *Set) Update(uuid string, newValue string) error {
	_, ok := s.Storage[uuid]
	if !ok {
		return errors.New("uuid is not found")
	}

	s.Storage[uuid] = newValue

	return nil
}

func (s *Set) Delete(uuid string) error {
	_, ok := s.Storage[uuid]
	if !ok {
		return errors.New("uuid is not found")
	}

	delete(s.Storage, uuid)
	s.Count--

	return nil
}

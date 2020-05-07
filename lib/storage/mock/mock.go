package mock

import "github.com/goldenCRM.git/lib/models"

type Mock struct {
	flats []models.Flat
}

func New() *Mock {
	return &Mock{}
}

func (m *Mock) Add(flat models.Flat) error {
	m.flats = append(m.flats, flat)
	return nil
}

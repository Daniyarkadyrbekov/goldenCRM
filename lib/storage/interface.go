package storage

import "github.com/goldenCRM.git/lib/models"

type Storage interface {
	Add(flat models.Flat) error
	List() ([]models.Flat, error)
}

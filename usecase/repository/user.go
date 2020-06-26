package repository

import "github.com/TechLoCo/sample-go-test/model"

type User interface {
	Get(id int) (*model.User, error)
}

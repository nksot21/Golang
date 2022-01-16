package repository

import (
	models "DBdemo/src/model"
)

type ClientRepo interface {
	Select() ([]models.Client, error)
	Insert(u models.Client) error
}

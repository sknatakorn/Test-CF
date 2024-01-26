package deleteCustomerByIdDb

import (
	"Test-CF/infrastructure"
	"gorm.io/gorm"
	"net/http"
)

type adapterDb struct {
	db *gorm.DB
}

func NewAdapterDb(db *gorm.DB) Port {
	return &adapterDb{db}
}

func (a adapterDb) Execute(req *Request) error {
	result := a.db.Delete(&req)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return infrastructure.NewError(http.StatusNotFound, "this id don't stored")
	}
	return nil
}

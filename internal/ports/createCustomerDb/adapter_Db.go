package createCustomerDb

import "gorm.io/gorm"

type adapterDb struct {
	db *gorm.DB
}

func NewAdapterDb(db *gorm.DB) Port {
	return &adapterDb{db}
}

func (a adapterDb) Execute(req *Request) error {
	result := a.db.Create(req)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

package getCustomerByIdDb

import "gorm.io/gorm"

type adapterDb struct {
	db *gorm.DB
}

func NewAdapterDb(db *gorm.DB) Port {
	return &adapterDb{db}
}

func (a adapterDb) Execute(id int) (*Response, error) {
	var user Response
	result := a.db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

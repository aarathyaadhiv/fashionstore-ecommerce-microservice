package repository

import (
	"github.com/aarathyaadhiv/fashionstore-admin-service/pkg/models"
	interfaces "github.com/aarathyaadhiv/fashionstore-admin-service/pkg/repository/interface"
	"gorm.io/gorm"
	"errors"
)



type AdminRepository struct{
	DB *gorm.DB
}

func NewAdminRepository(db *gorm.DB)interfaces.AdminRepository{
	return &AdminRepository{db}
}

func (c *AdminRepository) CheckAdminAvailability(email string) bool {
	var count int
	if err := c.DB.Raw(`select count(*) from users where email=? and role='admin'`, email).Scan(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func (c *AdminRepository) FindByEmail(email string) (models.LoginCheck, error) {
	var adminDetails models.LoginCheck
	if err := c.DB.Raw(`select id,name,email,ph_no,password from users where email=? and role='admin'`, email).Scan(&adminDetails).Error; err != nil {
		return models.LoginCheck{}, errors.New("error in fetching admin details")
	}
	return adminDetails, nil
}
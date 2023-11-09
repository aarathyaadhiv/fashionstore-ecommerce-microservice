package interfaces

import "github.com/aarathyaadhiv/fashionstore-admin-service/pkg/models"

type AdminRepository interface{
	CheckAdminAvailability(email string) bool
	FindByEmail(email string) (models.LoginCheck, error)
}
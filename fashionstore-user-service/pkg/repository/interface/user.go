package interfaces

import (
	

	"github.com/aarathyaadhiv/fashionstore-user-service/pkg/pb"
	
	"github.com/aarathyaadhiv/fashionstore-user-service/pkg/models"
)

type UserRepository interface{
	CheckUserAvailability(email string) bool
	FindByEmail(email string) (models.UserLoginCheck, error)
	Save(user *pb.SignUpRequest) (*pb.UserDetails, error)
	IsBlocked(email string) (bool,error)
}
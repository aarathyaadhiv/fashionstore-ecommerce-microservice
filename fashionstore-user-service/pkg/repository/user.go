package repository

import (
	"errors"

	"github.com/aarathyaadhiv/fashionstore-user-service/pkg/models"
	"github.com/aarathyaadhiv/fashionstore-user-service/pkg/pb"
	interfaces "github.com/aarathyaadhiv/fashionstore-user-service/pkg/repository/interface"
	"gorm.io/gorm"
)


type UserRepository struct{
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB)interfaces.UserRepository{
	return &UserRepository{db}
}

func (c *UserRepository) CheckUserAvailability(email string) bool {
	var count int

	if err := c.DB.Raw(`select count(*) from users where email=? and role='user'`, email).Scan(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func (c *UserRepository) FindByEmail(email string) (models.UserLoginCheck, error) {
	var user models.UserLoginCheck

	if err := c.DB.Raw(`select id,name,email,ph_no as phno,password from users where email=? and role='user'`, email).Scan(&user).Error; err != nil {
		return models.UserLoginCheck{}, errors.New("error in checking userdetails")
	}
	return user, nil
}



func (c *UserRepository) Save(user *pb.SignUpRequest) (*pb.UserDetails, error) {
	var userdetails *pb.UserDetails
	if err := c.DB.Raw(`insert into users(name,email,ph_no,password,role) values($1,$2,$3,$4,$5) returning id,name,email,ph_no`, user.Name, user.Email, user.Phno, user.Password, "user").Scan(&userdetails).Error; err != nil {
		return nil, errors.New("error saving in database")
	}

	return userdetails, nil
}

func (c *UserRepository) IsBlocked(email string) (bool,error) {
	var block bool
	if err := c.DB.Raw(`select block from users where email=?`, email).Scan(&block).Error; err != nil {
		return false,errors.New("error in fetching block detail")
	}
	return block,nil
}
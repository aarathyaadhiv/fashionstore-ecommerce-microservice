package service

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"github.com/aarathyaadhiv/fashionstore-user-service/pkg/pb"
	"github.com/aarathyaadhiv/fashionstore-user-service/pkg/helper"
	interfaces "github.com/aarathyaadhiv/fashionstore-user-service/pkg/repository/interface"
	
)


type UserService struct{
	Repo interfaces.UserRepository
	pb.UnimplementedUserServiceServer
}

func NewUserService(repo interfaces.UserRepository)*UserService{
	return &UserService{Repo: repo}
}

func (c *UserService) SignUp(ctx context.Context,user *pb.SignUpRequest)(*pb.SignUpResponse,error){
	isValidEmail := helper.IsValidEmail(user.Email)
	if !isValidEmail {
		return nil, errors.New("please enter a valid email")
	}
	isValidNumber := helper.IsValidPhoneNumber(user.Phno)
	if !isValidNumber {
		return nil, errors.New("please enter a valid phone number")
	}
	if ok := c.Repo.CheckUserAvailability(user.Email); ok {
		return nil, errors.New("already existing email")
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return nil, errors.New("error in password hashing")
	}
	user.Password = string(hashPassword)

	userDetails, err := c.Repo.Save(user)
	if err != nil {
		return nil, errors.New("error in saving user data")
	}

	tokenString, err := helper.GenerateUserToken(userDetails)
	if err != nil {
		return nil, err
	}
	return &pb.SignUpResponse{Userdetails: userDetails, Token: tokenString}, nil
}

func (c *UserService) Login(ctx context.Context,user *pb.LoginRequest) (*pb.LoginResponse, error) {
	if ok := c.Repo.CheckUserAvailability(user.Email); !ok {
		return nil, errors.New("no such user exist")
	}
	if ok,_ := c.Repo.IsBlocked(user.Email); ok {
		return nil, errors.New("user is blocked")
	}
	fmt.Println("user",user)
	userCompare, err := c.Repo.FindByEmail(user.Email)
	fmt.Println("usercompare",userCompare)
	if err != nil {
		return nil, errors.New("error in fetching userdata")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(userCompare.Password), []byte(user.Password)); err != nil {
		return nil, errors.New("password is incorrect")
	}

	userDetails :=&pb.UserDetails{Id: userCompare.ID,Name: userCompare.Name,Email: userCompare.Email,Phno: userCompare.PhNo}
	fmt.Println("userdetails:",userDetails)
	tokenString, err := helper.GenerateUserToken(userDetails)
	if err != nil {
		return nil, err
	}
	fmt.Println("token",tokenString)
	return &pb.LoginResponse{Userdetails: userDetails, Token: tokenString}, nil

}

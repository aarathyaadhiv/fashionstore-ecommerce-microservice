package service

import (
	"context"
	"errors"

	
	"github.com/aarathyaadhiv/fashionstore-admin-service/pkg/helper"
	"github.com/aarathyaadhiv/fashionstore-admin-service/pkg/pb"
	interfaces "github.com/aarathyaadhiv/fashionstore-admin-service/pkg/repository/interface"
	"golang.org/x/crypto/bcrypt"
)


type AdminService struct{
	Repo interfaces.AdminRepository
	pb.UnimplementedAdminServiceServer
}

func NewAdminService(repo interfaces.AdminRepository)*AdminService{
	return &AdminService{Repo: repo}
}

func (a *AdminService) Login(ctx context.Context,admin *pb.AdminLoginRequest)(*pb.AdminLoginResponse,error){
	if ok := a.Repo.CheckAdminAvailability(admin.Email); !ok {
		return nil, errors.New("no such user exist")
	}

	adminCompare, err := a.Repo.FindByEmail(admin.Email)
	if err != nil {
		return nil, errors.New("error in fetching userdata")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(adminCompare.Password), []byte(admin.Password)); err != nil {
		return nil, errors.New("password is incorrect")
	}

	adminDetails:=&pb.AdminDetails{Id: adminCompare.ID,Name: adminCompare.Name,Email: adminCompare.Email,Phno: adminCompare.PhNo}
	
	tokenString, err := helper.GenerateAdminToken(adminDetails)
	if err != nil {
		return nil, err
	}

	return &pb.AdminLoginResponse{Admindetails: adminDetails, Token: tokenString}, nil
}
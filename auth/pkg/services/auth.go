package services

import (
	"auth_svc/pkg/db"
	"auth_svc/pkg/models"
	"auth_svc/pkg/pb"
	"context"
	"net/http"
)

type Server struct {
	H db.Handler
}

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	var user models.Users

	if result := s.H.DB.Where(&models.Users{Email: req.Email}).First(&user); result.Error == nil {
		return &pb.RegisterResponse{
			Status: http.StatusConflict,
			Error:  "Email is exist",
		}, nil
	}

	user.Email = req.Email
	user.Password = req.Password

	s.H.DB.Create(&user)

	return &pb.RegisterResponse{
		Status: http.StatusOK,
	}, nil
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var userLogin models.Users

	if result := s.H.DB.Where(&models.Users{Email: req.Email}).First(&userLogin); result.Error != nil {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	passwordValid := req.Password == userLogin.Password

	if !passwordValid {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	accessToken, _ := db.GenerateAccessToken(&userLogin)

	return &pb.LoginResponse{
		Status: http.StatusOK,
		Token:  accessToken,
	}, nil
}

package client

import (
	"github.com/husanmusa/bpro_api_gateway/config"
	pba "github.com/husanmusa/bpro_api_gateway/genproto/auth_service"
	pbb "github.com/husanmusa/bpro_api_gateway/genproto/book_pro_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceManagerI interface {
	UserService() pba.UserServiceClient
	BookService() pbb.BookProServiceClient
	BookCategoryService() pbb.BookCategoryServiceClient
}

type grpcClients struct {
	userService  pba.UserServiceClient
	bookService  pbb.BookProServiceClient
	bookCategory pbb.BookCategoryServiceClient
}

func NewGrpcClients(cfg config.Config) (ServiceManagerI, error) {
	connUserService, err := grpc.Dial(
		cfg.AuthServiceHost+cfg.AuthServicePort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	connBookService, err := grpc.Dial(
		cfg.BookServiceHost+cfg.BookServicePort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		userService:  pba.NewUserServiceClient(connUserService),
		bookService:  pbb.NewBookProServiceClient(connBookService),
		bookCategory: pbb.NewBookCategoryServiceClient(connBookService),
	}, nil
}
func (g *grpcClients) UserService() pba.UserServiceClient {
	return g.userService
}
func (g *grpcClients) BookService() pbb.BookProServiceClient {
	return g.bookService
}
func (g *grpcClients) BookCategoryService() pbb.BookCategoryServiceClient {
	return g.bookCategory
}

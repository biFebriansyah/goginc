package repositories

import (
	"biFebriansyah/gogin/config"
	"biFebriansyah/gogin/internal/models"

	"github.com/stretchr/testify/mock"
)

type RepoUserMock struct {
	mock.Mock
}

func (rp *RepoUserMock) CreateUser(data *models.User) (*config.Result, error) {
	args := rp.Mock.Called(data)
	return args.Get(0).(*config.Result), args.Error(1)
}

func (rp *RepoUserMock) GetAllUser() (*config.Result, error) {
	args := rp.Mock.Called()
	return args.Get(0).(*config.Result), args.Error(1)
}

func (rp *RepoUserMock) GetAuthData(user string) (*models.User, error) {
	args := rp.Mock.Called(user)
	return args.Get(0).(*models.User), args.Error(1)
}

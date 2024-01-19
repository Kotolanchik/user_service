package models

import (
	"github.com/google/uuid"
	"strings"
	"user-service/internal/services/validator"
	"user-service/internal/types/consts"
	"user-service/internal/types/interfaces"
	"user-service/internal/types/structs"
)

type UsersModel struct {
	repository interfaces.UsersRepository
	logger     interfaces.Logger
}

func NewUsersModel(repository interfaces.UsersRepository, logger interfaces.Logger) *UsersModel {
	return &UsersModel{
		repository: repository,
		logger:     logger,
	}
}

var validate = validator.CreateValidator()

func (m *UsersModel) Get(userId uuid.UUID) (structs.User, error) {
	user, err := m.repository.Get(userId)
	if err != nil {
		return structs.User{}, err
	}

	return user, nil
}

func (m *UsersModel) GetAll(pagination structs.Pagination, sort structs.Sort) (structs.UserList, error) {
	userList, err := m.repository.GetAll(pagination, sort)
	if err != nil {
		return structs.UserList{}, err
	}
}

func (m *UsersModel) ReplaceOne(fields structs.UserEditingFields) (uuid.UUID, error) {

}

func (m *UsersModel) UpdateOne(userId uuid.UUID, fields structs.UserEditingFields) (bool, error) {

}

func (m *UsersModel) DeleteOne(userID uuid.UUID) error {

}

func normalizePagination(pagination *structs.Pagination) {
	if pagination.Offset < 0 {
		pagination.Offset = consts.UsersPaginationDefaultOffset
	}

	if pagination.Limit < 1 {
		pagination.Limit = consts.UsersPaginationDefaultLimit
	}

	if pagination.Limit > consts.UsersPaginationMaxLimit {
		pagination.Limit = consts.UsersPaginationMaxLimit
	}
}

func parseSort(sort string) (structs.Sort, error) {
	err := validate.Var(sort, "required")
	if err != nil {
		return structs.Sort{}, nil
	}

	sortData := strings.Split(sort, "-")

	if len(sortData) != 2 {

	}

}

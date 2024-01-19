package interfaces

import (
	"github.com/google/uuid"
	"net/http"
	"user-service/internal/types/structs"
)

type UsersController interface {
	Get(http.ResponseWriter, *http.Request)
	GetList(http.ResponseWriter, *http.Request)
	Create(http.ResponseWriter, *http.Request)
	Replace(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
}

type UsersModel interface {
	Get(userId uuid.UUID) (structs.User, error)
	GetAll(pagination structs.Pagination, sort structs.Sort) (structs.UserList, error)
	ReplaceOne(fields structs.UserEditingFields) (uuid.UUID, error)
	UpdateOne(userId uuid.UUID, fields structs.UserEditingFields) (bool, error)
	DeleteOne(userID uuid.UUID) error
}

type UsersRepository interface {
	Get(userId uuid.UUID) (structs.User, error)
	GetAll(pagination structs.Pagination, sort structs.Sort) (structs.UserList, error)
	InsertOne(fields structs.UserEditingFields) (uuid.UUID, error)
	UpdateOne(userId uuid.UUID, fields structs.UserEditingFields) (bool, error)
	DeleteOne(userID uuid.UUID) error
}

type Logger interface {
	Log(level uint32, errorCode int, message string, detailed interface{})
	Debug(errorCode int, message string, detailed interface{})
	Info(errorCode int, message string, detailed interface{})
	Warn(errorCode int, message string, detailed interface{})
	Error(errorCode int, message string, detailed interface{})
	Panic(errorCode int, message string, detailed interface{})
}

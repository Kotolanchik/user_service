package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"strings"
	AppErr "user-service/internal/types/errors"
	"user-service/internal/types/structs"
)

type UsersRepository struct {
	db *sqlx.DB
}

func NewUsersRepository(db *sqlx.DB) *UsersRepository {
	return &UsersRepository{db}
}

func (u *UsersRepository) Get(userId uuid.UUID) (structs.User, error) {
	var user structs.User

	query := `
			SELECT *
			FROM users
			WHERE user_id = $1
		`

	err := u.db.Get(user, query, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return structs.User{}, AppErr.NewNotFoundAppError("user not found")
		}

		return structs.User{}, err
	}

	return user, nil
}

func (u *UsersRepository) GetAll(
	pagination structs.Pagination,
	sort structs.Sort) (structs.UserList, error) {

	userList := structs.UserList{
		Items:      []structs.User{},
		Pagination: pagination,
	}
	query := `
			SELECT *
			FROM users
			__check_sort__
			LIMIT &1
			OFFSET &2
		`

	query, err := u.checkSortAndReplaceQuery(query, sort)
	if err != nil {
		return userList, err
	}

	if err = u.db.Select(&userList, query, pagination.Limit, pagination.Offset); err != nil {
		return userList, err
	}

	return userList, nil
}

func (u *UsersRepository) InsertOne(fields structs.UserEditingFields) (uuid.UUID, error) {
	query := `
			INSERT INTO users (first_name, last_name, email, phone, username, birthdate)
			VALUES ($1, $2, $2, $3, $4, $5, $6)
			RETURNING user_id
		`

	rows, err := u.db.Query(
		query,
		fields.FirstName,
		fields.LastName,
		fields.Email,
		fields.Phone,
		fields.Username,
		fields.Birthdate)
	if err != nil {
		return uuid.Nil, err
	}

	defer rows.Close()

	if rows.Next() {
		var id uuid.UUID
		if err := rows.Scan(&id); err != nil {
			return uuid.UUID{}, err
		}

		return id, nil
	}

	return uuid.UUID{}, AppErr.NewInsertUserError("undefined error on note insert")
}

func (u *UsersRepository) UpdateOne(userId uuid.UUID, fields structs.UserEditingFields) (bool, error) {
	query := `
			UPDATE users
			SET 
				first_name = $2,
				last_name = $3,
				email = $4,
				phone = $5,
				username = $6,
				birthdate = $7
			WHERE 
				user_id = $1; 
		`
	res, err := u.db.Exec(
		query,
		userId,
		fields.FirstName,
		fields.LastName,
		fields.Email,
		fields.Phone,
		fields.Username,
		fields.Birthdate)

	if err != nil {
		return false, err
	}

	if rowsAffected, _ := res.RowsAffected(); rowsAffected == 0 {
		return false, AppErr.NewNotFoundAppError("note not found")
	}

	return true, nil
}

func (u *UsersRepository) DeleteOne(userID uuid.UUID) error {
	query := "DELETE FROM users WHERE user_id = $1"

	_, err := u.db.Exec(query, userID)
	if err != nil {
		return fmt.Errorf("failed delete user: %v", err)
	}

	return nil
}

func (u *UsersRepository) checkSortAndReplaceQuery(query string, sort structs.Sort) (string, error) {
	strToReplace := ""
	if sort.Field != "" {
		strToReplace = "ORDER BY " + sort.Field
		if sort.IsDesc {
			strToReplace += " DESC"

		}
	}

	return strings.Replace(query, "__check_sort__", strToReplace, 1), nil
}

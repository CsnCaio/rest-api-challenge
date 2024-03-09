package repositories

import (
	"rest-api-challenge/config"
	"rest-api-challenge/dtos"
)

func CreateUser(createUserDto dtos.CreateUserDto) (dtos.UserDto, error) {
	insertStatement := `INSERT INTO "users"("name", "email") values($1, $2) RETURNING "id", "name", "email"`
	user := dtos.UserDto{}
	err := config.DbConn.QueryRow(insertStatement, createUserDto.Name, createUserDto.Email).Scan(&user.ID, &user.Name, &user.Email)
	return user, err
}

func UpdateUser(id int, updateUserDto dtos.UpdateUserDto) (dtos.UserDto, error) {
	updateStatement := `UPDATE "users" SET "name" = $1, "email" = $2 WHERE "id" = $3 RETURNING "id", "name", "email"`
	user := dtos.UserDto{}
	err := config.DbConn.QueryRow(updateStatement, updateUserDto.Name, updateUserDto.Email, id).Scan(&user.ID, &user.Name, &user.Email)
	return user, err
}

func GetUser(id int) (dtos.UserDto, error) {
	selectStatement := `SELECT "id", "name", "email" FROM "users" WHERE "id" = $1`
	user := dtos.UserDto{}
	err := config.DbConn.QueryRow(selectStatement, id).Scan(&user.ID, &user.Name, &user.Email)
	return user, err
}

func GetUsers() ([]dtos.UserDto, error) {
	selectStatement := `SELECT "id", "name", "email" FROM "users"`
	rows, err := config.DbConn.Query(selectStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []dtos.UserDto{}
	for rows.Next() {
		user := dtos.UserDto{}
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func DeleteUser(id int) error {
	deleteStatement := `DELETE FROM "users" WHERE "id" = $1`
	_, err := config.DbConn.Exec(deleteStatement, id)
	return err
}

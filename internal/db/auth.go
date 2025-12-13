package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func (user *Users) TranslateRole() string {
	switch user.Role {
	case Role_Guest: return "Гость"
	case Role_Candidate: return "Кандидат"
	case Role_Rescuer: return "Спасатель"
	case Role_Operator: return "Оператор"
	case Role_Admin: return "Администратор"
	}
	return ""
}

func FindUserByLogin(DB *pgx.Conn, login string) (*Users, error) {
	userRow := DB.QueryRow(context.Background(), "SELECT id_user, login, password, role FROM users WHERE login = $1", login)

	var user Users
	err := userRow.Scan(&user.IdUser, &user.Login, &user.Password, &user.Role)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUser(DB *pgx.Conn, id uint64) (*Users, error) {
	userRow := DB.QueryRow(context.Background(), "SELECT id_user, login, password, role FROM users WHERE id_user = $1", id)

	var user Users
	err := userRow.Scan(&user.IdUser, &user.Login, &user.Password, &user.Role)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

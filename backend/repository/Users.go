package repository

import (
	"database/sql"
	"strconv"

	"github.com/ilmiDs/final-project-78-km/models"
	"github.com/ilmiDs/final-project-78-km/db"
)
	type UserRepository struct {
		db *sql.DB
	}
	
	func NewUserRepository(db *sql.DB) UserRepository {
		return UserRepository{db}
	}
/*
type User struct {
	Id int `db:"id"`
	Nim int `db:"nim"`
	Nama string `db:"nama"`
	kelamin string `db:"kelamin"`
	Email string `db:"email"`
	Jurusan string `db:"jurusan"`
	Fakultas string `db:"fakultas"`
	NoHp string `db:"noHp"`
	Password string `db:"password"`
	Role string `db:"role"`
	Logged bool `db:"logged"`
	Gambar string `db:"gambar"`
}*/
//LoadOrCreateUser load or create user
func (u UserRepository) LoadOrCreateUser(nim int, nama string, kelamin string, email string, jurusan string, fakultas string, noHp string, password string, role string, logged bool, gambar string) models.User {
	var user models.User
	user.Nim = nim
	user.Nama = nama
	user.Kelamin = kelamin
	user.Email = email
	user.Jurusan = jurusan
	user.Fakultas = fakultas
	user.NoHp = noHp
	user.Password = password
	user.Role = role
	user.Logged = logged
	user.Gambar = gambar
	return user
}
//LoadUser load user
func (u UserRepository) LoadUser(nim int) models.User {
	var user models.User
	user.Nim = nim
	return user
}
//LoadAllUser load all user
func (u UserRepository) LoadAllUser() []models.User {
	var users []models.User
	db.Db.Select(&users, "SELECT * FROM user")
	return users
}

func (u UserRepository) Login(username string, password string) (*string, error) {
	var user models.User
	err := db.Db.Get(&user, "SELECT * FROM user WHERE email = ? AND password = ?", username, password)
	if err != nil {
		return nil, err
	}
	return &user.Role, nil
}
func (u *UserRepository) FindLoggedinUser() (*string, error) {
	var user models.User
	err := db.Db.Get(&user, "SELECT * FROM user WHERE logged = ?", true)
	if err != nil {
		return nil, err
	}
	return &user.Role, nil
}

func (u *UserRepository) Logout(username string) error {
	_, err := db.Db.Exec("UPDATE user SET logged = ? WHERE email = ?", false, username)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) Save(users models.User) error {
	for _, user := range users {
		_, err := db.Db.Exec("INSERT INTO user (nim, nama, kelamin, email, jurusan, fakultas, noHp, password, role, logged, gambar) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", user.Nim, user.Nama, user.Kelamin, user.Email, user.Jurusan, user.Fakultas, user.NoHp, user.Password, user.Role, user.Logged, user.Gambar)
		if err != nil {
			return err
		}
	}
	return nil
}

func (u *UserRepository) changeStatus(username string, status bool) error {
	_, err := db.Db.Exec("UPDATE user SET logged = ? WHERE email = ?", status, username)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) LogoutAll() error {
	_, err := db.Db.Exec("UPDATE user SET logged = ?", false)
	if err != nil {
		return err
	}
	return nil
}




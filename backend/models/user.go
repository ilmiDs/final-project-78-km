package models

			/*
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			nim INTEGER ,
			nama varchar(255) NOT NULL,
			gender varchar(255) ,
			email varchar(255) NOT NULL,
			jurusan varchar(255) ,
			fakultas varchar(255) ,
			noHp varchar(255),
			password varchar(255) NOT NULL,
			role varchar(255) NOT NULL,
			logged bool NOT NULL,
			gambar varchar(255)
			*/
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
}
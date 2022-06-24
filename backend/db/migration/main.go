package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Run This Script for migration db
func main() {
	db, err := sql.Open("sqlite3", "db/perpus.db")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
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
		);

		CREATE TABLE IF NOT EXISTS buku (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			ISBN INTEGER NOT NULL,
			judul varchar(255) NOT NULL,
			pengarang varchar(255) NOT NULL,
			penerbit varchar(255) NOT NULL,
			tahun varchar(255) NOT NULL,
			stok INTEGER NOT NULL,
			kotaTerbit varchar(255) NOT NULL,
			deskripsi varchar(255) NOT NULL,
			gambar varchar(255) NOT NULL
		);

		CREATE TABLE IF NOT EXISTS peminjaman (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			idUser INTEGER NOT NULL,
			idBuku INTEGER NOT NULL,
			tanggalPeminjaman DATE NOT NULL,
			tanggalPengembalian DATE NOT NULL,
			status varchar(255) NOT NULL,
			foreign key(idUser) references users(id),
			foreign key(idBuku) references buku(id)
		);

		CREATE TABLE IF NOT EXISTS pengembalian (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			idUser INTEGER NOT NULL,
			idBuku INTEGER NOT NULL,
			tanggalPengembalian DATE NOT NULL,
			status varchar(255) NOT NULL,
			foreign key(idUser) references users(id),
			foreign key(idBuku) references buku(id)
		);

		CREATE TABLE IF NOT EXISTS pinjamBuku (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			nama varchar(255) NOT NULL,
			nim INTEGER NOT NULL,
			judul varchar(255) NOT NULL,
			pengarang varchar(255) NOT NULL,
			noBuku INTEGER NOT NULL,
			foreign key(nama) references users(nama),
			foreign key(nim) references users(nim)
			foreign key(noBuku) references buku(id)

		);

		INSERT INTO users (nama, email, password, role, logged) VALUES 
		("admin", "backend@gmail.com", "admin", "admin", false);	

`);

	if err != nil {
		panic(err)
	}
}

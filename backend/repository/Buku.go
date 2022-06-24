package repository

import (
	"database/sql"
	"go/types"

	"github.com/ilmiDs/final-project-78-km/models"
	"github.com/ilmiDs/final-project-78-km/db"
)

	type buku struct {
		db *sql.DB
	}

	func NewBukuRepository (db *sql.DB) buku {
		return buku{db}
	}

	func (b buku) LoadBuku (id int) models.Buku {
		var buku models.Buku
		buku.Id = id
		return buku
	}

	func (b buku) LoadAllBuku () []models.Buku {
		var bukus []models.Buku
		db.Db.Select(&bukus, "SELECT * FROM buku")
		return bukus
	}

	//create buku
	func (b buku) CreateBuku (id int, judul string, penerbit string, pengarang string, tahun string, stok string, harga string, gambar string) models.Buku {
		var buku models.Buku
		buku.Id = id
		buku.Judul = judul
		buku.Penerbit = penerbit
		buku.Pengarang = pengarang
		buku.Tahun = tahun
		buku.Stok = stok
		buku.Harga = harga
		buku.Gambar = gambar
		return buku
	}

	//update buku
	func (b buku) UpdateBuku (id int, judul string, penerbit string, pengarang string, tahun string, stok string, harga string, gambar string) models.Buku {
		var buku models.Buku
		buku.Id = id
		buku.Judul = judul
		buku.Penerbit = penerbit
		buku.Pengarang = pengarang
		buku.Tahun = tahun
		buku.Stok = stok
		buku.Harga = harga
		buku.Gambar = gambar
		return buku
	}

	//delete buku
	func (b buku) DeleteBuku (id int) models.Buku {
		var buku models.Buku
		buku.Id = id
		return buku
	}

	//search buku
	func (b buku) SearchBuku (judul string) []models.Buku {
		var bukus []models.Buku
		db.Db.Select(&bukus, "SELECT * FROM buku WHERE judul LIKE ?", "%"+judul+"%")
		return bukus
	}



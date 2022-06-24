package repository

import (
	"database/sql"
	"go/types"

	"github.com/ilmiDs/final-project-78-km/backend/models"
)

type BukuRepository interface {
	GetAll() ([]models.Buku, error)
	GetById(id int) (models.Buku, error)
	CreateBuke(buku models.Buku) error
	UpdateBuku(buku models.Buku) error
	DeleteBuku(id int) error
	SearchBuku(search string) ([]models.Buku, error)
}

	types buku struct {
		db *sql.DB
	}

	func NewBukuRepository (db *sql.DB) BukuRepository {
		return &buku{db}
	}

	func (buku *buku) GetAll() ([]models.Buku, error) {
		var bukus []models.Buku
		rows, err := buku.db.Query("SELECT * FROM buku")
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		for rows.Next() {
			var buku models.Buku
			err := rows.Scan(&buku.ID, &buku.ISBN, &buku.Judul, &buku.Pengarang, &buku.Penerbit, &buku.Tahun, &buku.Stok, &buku.KotaTerbit, &buku.Deskripsi, &buku.Gambar)
			if err != nil {
				return nil, err
			}
			bukus = append(bukus, buku)
		}
		return bukus, nil
	}

	func (buku *buku) GetById(id int) (models.Buku, error) {
		var buku models.Buku
		err := buku.db.QueryRow("SELECT * FROM buku WHERE id=?", id).Scan(&buku.ID, &buku.ISBN, &buku.Judul, &buku.Pengarang, &buku.Penerbit, &buku.Tahun, &buku.Stok, &buku.KotaTerbit, &buku.Deskripsi, &buku.Gambar)
		if err != nil {
			return buku, err
		}
		return buku, nil
	}

	// create buku
	func (buku *buku) CreateBuke(buku models.Buku) error {
		query := "INSERT INTO buku (isbn, judul, pengarang, penerbit, tahun, stok, kotaTerbit, deskripsi, gambar) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"
		_, err := buku.db.Exec(query, buku.ISBN, buku.Judul, buku.Pengarang, buku.Penerbit, buku.Tahun, buku.Stok, buku.KotaTerbit, buku.Deskripsi, buku.Gambar)
		if err != nil {
			return err
		}
		return nil
	}

	// update buku
	func (buku *buku) UpdateBuku(buku models.Buku) error {
		query := "UPDATE buku SET isbn=?, judul=?, pengarang=?, penerbit=?, tahun=?, stok=?, kotaTerbit=?, deskripsi=?, gambar=? WHERE id=?"
		_, err := buku.db.Exec(query, buku.ISBN, buku.Judul, buku.Pengarang, buku.Penerbit, buku.Tahun, buku.Stok, buku.KotaTerbit, buku.Deskripsi, buku.Gambar, buku.ID)
		if err != nil {
			return err
		}
		return nil
	}

	// delete buku
	func (buku *buku) DeleteBuku(id int) error {
		query := "DELETE FROM buku WHERE id=?"
		_, err := buku.db.Exec(query, id)
		if err != nil {
			return err
		}
		return nil
	}

	// search buku
	func (buku *buku) SearchBuku(search string) ([]models.Buku, error) {
		var bukus []models.Buku
		query := "SELECT * FROM buku WHERE judul LIKE ? OR pengarang LIKE ? OR penerbit LIKE ? OR tahun LIKE ? OR kotaTerbit LIKE ? OR deskripsi LIKE ? OR gambar LIKE ?"
		rows, err := buku.db.Query(query, "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%")
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		for rows.Next() {
			var buku models.Buku
			err := rows.Scan(&buku.ID, &buku.ISBN, &buku.Judul, &buku.Pengarang, &buku.Penerbit, &buku.Tahun, &buku.Stok, &buku.KotaTerbit, &buku.Deskripsi, &buku.Gambar)
			if err != nil {
				return nil, err
			}
			bukus = append(bukus, buku)
		}
		return bukus, nil
	}
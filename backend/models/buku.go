package model

			/*
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
			*/
//struk buku mengambiul data dari db sqlite
type Buku struct {
	ID          int    `db:"id"`
	ISBN        int    `db:"isbn"`
	Judul       string `db:"judul"`
	Pengarang   string `db:"pengarang"`
	Penerbit    string `db:"penerbit"`
	Tahun       string `db:"tahun"`
	Stok        int    `db:"stok"`
	KotaTerbit  string `db:"kotaTerbit"`
	Deskripsi   string `db:"deskripsi"`
	Gambar      string `db:"gambar"`
}
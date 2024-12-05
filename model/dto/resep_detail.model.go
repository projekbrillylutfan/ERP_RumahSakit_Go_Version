package dto

type ResepDetailCreateOrUpdate struct {
	IDObat        int64     `json:"id_obat" validate:"required"`                               // Foreign Key ke tabel Obat
	Jumlah        int       `json:"jumlah" validate:"required"`                                      // Jumlah
}

type ResepDetailFindAllRes struct {
	IDResepDetail int64     `json:"id_resep_detail" validate:"required"`
	IDObat        int64     `json:"id_obat" validate:"required"`                               // Foreign Key ke tabel Obat
	Jumlah        int       `json:"jumlah" validate:"required"`                                      // Jumlah
	Harga         float64   `json:"harga" validate:"required"`                                      // Harga
}
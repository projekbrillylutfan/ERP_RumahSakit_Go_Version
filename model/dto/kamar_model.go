package dto

type KamarCreateOrUpdateReq struct {
	JenisKamar   string    `json:"jenis_kamar" validate:"required"`             // Jenis Kamar (ENUM)
	TarifPerHari float64   `json:"tarif_per_hari" validate:"required"`        // Tarif Per Hari
}
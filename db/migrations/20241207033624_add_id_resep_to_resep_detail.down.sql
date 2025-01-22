-- Hapus foreign key constraint id_resep
ALTER TABLE resep_detail DROP CONSTRAINT fk_resep;

-- Hapus kolom id_resep dari tabel resep_detail
ALTER TABLE resep_detail DROP COLUMN id_resep;

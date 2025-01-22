-- Tambahkan kolom id_resep ke tabel resep_detail
ALTER TABLE resep_detail ADD COLUMN id_resep INT NOT NULL;

-- Tambahkan foreign key constraint untuk id_resep
ALTER TABLE resep_detail 
ADD CONSTRAINT fk_resep FOREIGN KEY (id_resep) REFERENCES resep(id_resep) ON DELETE CASCADE;

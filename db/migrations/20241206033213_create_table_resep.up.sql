CREATE TABLE resep (
    id_resep SERIAL PRIMARY KEY,              -- ID Resep (Primary Key, Auto Increment)
    id_user INT NOT NULL,                   -- ID User (Foreign Key ke tabel User atau Pasien)
    id_dokter INT NOT NULL,                   -- ID Dokter (Foreign Key ke tabel Dokter)
    id_resep_detail INT NOT NULL,             -- ID Resep Detail (Foreign Key ke tabel Resep Detail)
    tanggal DATE NOT NULL,                    -- Tanggal Resep
    total_harga NUMERIC(10, 2),               -- Total Harga (dengan presisi 2 angka desimal)
    created_at TIMESTAMP DEFAULT NOW(),       -- Waktu pembuatan (default waktu sekarang)
    updated_at TIMESTAMP DEFAULT NOW(),       -- Waktu pembaruan (default waktu sekarang)

    -- Foreign Key Constraints
    CONSTRAINT fk_user FOREIGN KEY (id_user) REFERENCES "user"(id_user) ON DELETE CASCADE,
    CONSTRAINT fk_dokter FOREIGN KEY (id_dokter) REFERENCES dokter(id_dokter) ON DELETE CASCADE,
    CONSTRAINT fk_resep_detail FOREIGN KEY (id_resep_detail) REFERENCES resep_detail(id_resep_detail) ON DELETE CASCADE
);

CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER trigger_update_timestamp_resep
BEFORE INSERT OR UPDATE
ON resep
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

-- Trigger function untuk menghitung total_harga
CREATE OR REPLACE FUNCTION update_total_harga_resep()
RETURNS TRIGGER AS $$
BEGIN
    -- Hitung total harga dari semua resep_detail yang berkaitan dengan id_resep
    UPDATE resep
    SET total_harga = (
        SELECT SUM(harga * jumlah) 
        FROM resep_detail 
        WHERE id_resep_detail = NEW.id_resep_detail
    )
    WHERE id_resep = NEW.id_resep_detail;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER trigger_update_total_harga
AFTER INSERT OR UPDATE OR DELETE
ON resep_detail
FOR EACH ROW
EXECUTE FUNCTION update_total_harga_resep();




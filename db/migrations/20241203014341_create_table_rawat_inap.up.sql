CREATE TABLE rawat_inap (
    id_rawat_inap SERIAL PRIMARY KEY,        -- ID Rawat Inap (Primary Key, Auto Increment)
    id_pasien INT NOT NULL,                  -- ID Pasien (Foreign Key ke tabel user atau pasien)
    id_kamar INT NOT NULL,                   -- ID Kamar (Foreign Key ke tabel kamar)
    tanggal_masuk DATE NOT NULL,             -- Tanggal Masuk
    tanggal_keluar DATE,                     -- Tanggal Keluar (opsional, bisa NULL jika belum keluar)
    created_at TIMESTAMP DEFAULT NOW(),      -- Waktu pembuatan (default waktu sekarang)
    updated_at TIMESTAMP DEFAULT NOW(),      -- Waktu pembaruan (default waktu sekarang)

    -- Foreign Key Constraints
    CONSTRAINT fk_pasien FOREIGN KEY (id_pasien) REFERENCES "user"(id_user) ON DELETE CASCADE,
    CONSTRAINT fk_kamar FOREIGN KEY (id_kamar) REFERENCES kamar(id_kamar) ON DELETE CASCADE
);

CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON rawat_inap
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

CREATE TABLE resep_detail (
    id_resep_detail SERIAL PRIMARY KEY,           -- ID Resep Detail (Primary Key, Auto Increment)
    id_obat INT NOT NULL,                         -- ID Obat (Foreign Key ke tabel Obat)
    jumlah INT NOT NULL,                          -- Jumlah
    harga NUMERIC(10, 2) NOT NULL,                -- Harga (dengan 2 angka desimal)
    created_at TIMESTAMP DEFAULT NOW(),           -- Waktu pembuatan (default waktu sekarang)
    updated_at TIMESTAMP DEFAULT NOW(),           -- Waktu pembaruan (default waktu sekarang)

    -- Foreign Key Constraint
    CONSTRAINT fk_obat FOREIGN KEY (id_obat) REFERENCES obat(id_obat) ON DELETE CASCADE
);

-- Membuat fungsi untuk memperbarui kolom updated_at
CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Menambahkan trigger pada tabel resep_detail
CREATE TRIGGER trigger_update_timestamp
BEFORE UPDATE ON resep_detail
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

-- Fungsi untuk menghitung harga total
CREATE OR REPLACE FUNCTION calculate_harga_resep_detail()
RETURNS TRIGGER AS $$
BEGIN
    -- Mengambil harga dari tabel obat dan mengalikan dengan jumlah
    NEW.harga := (SELECT o.harga FROM obat o WHERE o.id_obat = NEW.id_obat) * NEW.jumlah;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger yang dijalankan sebelum insert atau update pada resep_detail
CREATE TRIGGER trigger_calculate_harga
BEFORE INSERT OR UPDATE ON resep_detail
FOR EACH ROW
EXECUTE FUNCTION calculate_harga_resep_detail();


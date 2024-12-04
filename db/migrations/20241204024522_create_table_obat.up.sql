CREATE TABLE obat (
    id_obat SERIAL PRIMARY KEY,              -- ID Obat (Primary Key, Auto Increment)
    nama_obat VARCHAR(255) NOT NULL,         -- Nama Obat (wajib diisi)
    deskripsi TEXT,                          -- Deskripsi Obat (opsional)
    harga NUMERIC(10, 2) NOT NULL,           -- Harga Obat (float dengan presisi 2 desimal)
    created_at TIMESTAMP DEFAULT NOW(),      -- Waktu pembuatan (default: waktu sekarang)
    updated_at TIMESTAMP DEFAULT NOW()       -- Waktu pembaruan (default: waktu sekarang)
);

-- Membuat fungsi untuk memperbarui kolom updated_at
CREATE OR REPLACE FUNCTION update_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW(); -- Atur updated_at ke waktu saat ini
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Membuat trigger untuk tabel obat
CREATE TRIGGER trigger_update_updated_at
BEFORE UPDATE ON obat
FOR EACH ROW
EXECUTE FUNCTION update_updated_at();

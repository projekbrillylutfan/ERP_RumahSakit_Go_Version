-- 1. Buat Enum untuk Jenis Kamar
CREATE TYPE jenis_kamar AS ENUM (
    'STANDARD',
    'DELUXE',
    'VIP'
);

-- 2. Buat Tabel Kamar
CREATE TABLE kamar (
    id_kamar SERIAL PRIMARY KEY,              -- ID Kamar (Primary Key, Auto Increment)
    jenis_kamar jenis_kamar NOT NULL,         -- Jenis Kamar (menggunakan Enum)
    tarif_per_hari NUMERIC(10, 2) NOT NULL,   -- Tarif per Hari (float dengan presisi 2 angka desimal)
    created_at TIMESTAMP DEFAULT NOW(),       -- Waktu pembuatan (default waktu sekarang)
    updated_at TIMESTAMP DEFAULT NOW()        -- Waktu pembaruan (default waktu sekarang)
);

-- Fungsi untuk memperbarui kolom updated_at
CREATE OR REPLACE FUNCTION update_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger untuk tabel kamar
CREATE TRIGGER trigger_update_updated_at
BEFORE UPDATE ON kamar
FOR EACH ROW
EXECUTE FUNCTION update_updated_at();

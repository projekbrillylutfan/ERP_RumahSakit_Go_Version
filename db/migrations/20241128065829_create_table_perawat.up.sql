CREATE TABLE perawat (
    id_perawat SERIAL PRIMARY KEY,           -- ID Perawat (Primary Key, Auto Increment)
    nama VARCHAR(255) NOT NULL,              -- Nama
    email VARCHAR(255) UNIQUE NOT NULL,      -- Email (harus unik)
    username VARCHAR(255) UNIQUE NOT NULL,   -- Username (harus unik)
    password TEXT NOT NULL,                  -- Password
    role VARCHAR(50) NOT NULL DEFAULT 'PERAWAT', -- Role dengan default 'PERAWAT'
    shift VARCHAR(100) NOT NULL,             -- Shift
    nomor_telepon VARCHAR(20),               -- Nomor Telepon
    created_at TIMESTAMP DEFAULT NOW(),      -- Waktu pembuatan (default waktu sekarang)
    updated_at TIMESTAMP DEFAULT NOW()       -- Waktu pembaruan (default waktu sekarang)
);

CREATE OR REPLACE FUNCTION update_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_updated_at
BEFORE UPDATE ON perawat
FOR EACH ROW
EXECUTE FUNCTION update_updated_at();

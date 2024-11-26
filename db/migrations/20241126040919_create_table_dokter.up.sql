CREATE TYPE dokter_role AS ENUM ('ADMIN', 'DOKTER'); -- Enum untuk role

CREATE TABLE dokter (
    id_dokter SERIAL PRIMARY KEY,                  -- ID Dokter dengan auto increment
    nama VARCHAR(255) NOT NULL,                   -- Nama
    email VARCHAR(255) UNIQUE NOT NULL,           -- Email unik
    password VARCHAR(255) NOT NULL,               -- Password
    role dokter_role NOT NULL DEFAULT 'DOKTER',   -- Role dengan default 'DOKTER'
    is_email_verified BOOLEAN NOT NULL DEFAULT FALSE, -- Email terverifikasi
    spesialisasi VARCHAR(255),                    -- Spesialisasi
    nomor_telepon VARCHAR(20),                    -- Nomor Telepon
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Waktu dibuat
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Waktu diperbarui
);

CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_updated_at
BEFORE UPDATE ON dokter
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

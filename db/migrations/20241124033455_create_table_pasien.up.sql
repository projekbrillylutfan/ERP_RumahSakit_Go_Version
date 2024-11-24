CREATE TYPE user_role AS ENUM ('USER', 'ADMIN'); -- Enum untuk Role

CREATE TABLE "user" (
    id_user SERIAL PRIMARY KEY, -- ID User dengan auto increment
    nama VARCHAR(255) NOT NULL, -- Nama
    alamat TEXT, -- Alamat
    username VARCHAR(255) UNIQUE NOT NULL, -- Username unik
    email VARCHAR(255) UNIQUE NOT NULL, -- Email unik
    password VARCHAR(255) NOT NULL, -- Password
    role user_role NOT NULL DEFAULT 'USER', -- Role dengan default 'USER'
    is_email_verified BOOLEAN NOT NULL DEFAULT FALSE, -- Email terverifikasi
    tanggal_lahir DATE, -- Tanggal Lahir
    jenis_kelamin VARCHAR(50), -- Jenis Kelamin
    nomor_telepon VARCHAR(20) UNIQUE, -- Nomor Telepon unik
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Waktu dibuat
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Waktu diperbarui
);

-- Buat fungsi untuk mengupdate kolom updated_at
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = CURRENT_TIMESTAMP;
   RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Buat trigger untuk tabel "user"
CREATE TRIGGER set_updated_at
BEFORE UPDATE ON "user"
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

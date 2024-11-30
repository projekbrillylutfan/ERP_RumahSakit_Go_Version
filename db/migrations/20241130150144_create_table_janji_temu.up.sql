CREATE TABLE janji_temu (
    id_janji_temu SERIAL PRIMARY KEY,         -- ID Janji Temu (Primary Key, Auto Increment)
    id_user INT NOT NULL,                     -- ID User (Foreign Key ke tabel user)
    id_dokter INT NOT NULL,                   -- ID Dokter (Foreign Key ke tabel dokter)
    tanggal DATE NOT NULL,                    -- Tanggal Janji Temu
    waktu VARCHAR(50) NOT NULL,               -- Waktu Janji Temu
    created_at TIMESTAMP DEFAULT NOW(),       -- Waktu pembuatan (default waktu sekarang)
    updated_at TIMESTAMP DEFAULT NOW(),       -- Waktu pembaruan (default waktu sekarang)

    -- Constraint Foreign Key (opsional)
    CONSTRAINT fk_user FOREIGN KEY (id_user) REFERENCES "user"(id_user) ON DELETE CASCADE,
    CONSTRAINT fk_dokter FOREIGN KEY (id_dokter) REFERENCES dokter(id_dokter) ON DELETE CASCADE
);


CREATE OR REPLACE FUNCTION update_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_updated_at
BEFORE UPDATE ON janji_temu
FOR EACH ROW
EXECUTE FUNCTION update_updated_at();

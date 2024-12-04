-- Hapus trigger
DROP TRIGGER IF EXISTS trigger_update_updated_at ON obat;

-- Hapus fungsi
DROP FUNCTION IF EXISTS update_updated_at;

-- Hapus tabel obat
DROP TABLE IF EXISTS obat;

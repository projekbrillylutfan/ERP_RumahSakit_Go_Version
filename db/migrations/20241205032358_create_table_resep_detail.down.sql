-- Menghapus trigger terlebih dahulu
DROP TRIGGER IF EXISTS trigger_update_timestamp ON resep_detail;

-- Menghapus fungsi terkait trigger
DROP FUNCTION IF EXISTS update_timestamp;

-- Menghapus tabel resep_detail
DROP TABLE IF EXISTS resep_detail;

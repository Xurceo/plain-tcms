package queries

const DefectColumns = "id, result_id, external_link, title, severity, status, created_at"

const (
	GetDefectsByResult = "SELECT " + DefectColumns + " FROM defects WHERE result_id = $1"
	GetDefectByID      = "SELECT " + DefectColumns + " FROM defects WHERE id = $1"
	InsertDefect       = "INSERT INTO defects (result_id, external_link, title, severity) VALUES ($1, $2, $3, $4) RETURNING " + DefectColumns
	DeleteDefect       = "DELETE FROM defects WHERE id = $1"
)

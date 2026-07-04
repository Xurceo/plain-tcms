package queries

const ResultAttachmentColumns = "id, result_id, file_url, file_type, created_at"

const (
	GetAttachmentsByResult = "SELECT " + ResultAttachmentColumns + " FROM result_attachments WHERE result_id = $1"
	GetAttachmentByID      = "SELECT " + ResultAttachmentColumns + " FROM result_attachments WHERE id = $1"
	InsertResultAttachment = "INSERT INTO result_attachments (result_id, file_url, file_type) VALUES ($1, $2, $3) RETURNING " + ResultAttachmentColumns
	DeleteResultAttachment = "DELETE FROM result_attachments WHERE id = $1"
)

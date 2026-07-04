package queries

const ProjectColumns = "id, org_id, name, description, created_by, created_at"

const (
	GetAllProjects   = "SELECT " + ProjectColumns + " FROM projects"
	GetProjectByID   = "SELECT " + ProjectColumns + " FROM projects WHERE id = $1"
	GetProjectsByOrg = "SELECT " + ProjectColumns + " FROM projects WHERE org_id = $1"
	InsertProject    = "INSERT INTO projects (org_id, name, description, created_by) VALUES ($1, $2, $3, $4) RETURNING " + ProjectColumns
	DeleteProject    = "DELETE FROM projects WHERE id = $1"
)

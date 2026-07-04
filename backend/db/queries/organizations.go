package queries

const OrganizationColumns = "id, name, created_at"

const (
	GetAllOrganizations = "SELECT " + OrganizationColumns + " FROM organizations"
	GetOrganizationByID = "SELECT " + OrganizationColumns + " FROM organizations WHERE id = $1"
	InsertOrganization  = "INSERT INTO organizations (name) VALUES ($1) RETURNING " + OrganizationColumns
	DeleteOrganization  = "DELETE FROM organizations WHERE id = $1"
)

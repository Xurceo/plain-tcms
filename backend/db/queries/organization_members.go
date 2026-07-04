package queries

const OrgMemberColumns = "org_id, user_id, role"

const (
	GetMembersByOrg     = "SELECT " + OrgMemberColumns + " FROM organization_members WHERE org_id = $1"
	GetMembersByUser    = "SELECT " + OrgMemberColumns + " FROM organization_members WHERE user_id = $1"
	InsertOrgMember     = "INSERT INTO organization_members (org_id, user_id, role) VALUES ($1, $2, $3) RETURNING " + OrgMemberColumns
	DeleteOrgMember     = "DELETE FROM organization_members WHERE org_id = $1 AND user_id = $2"
	UpdateOrgMemberRole = "UPDATE organization_members SET role = $1 WHERE org_id = $2 AND user_id = $3"
)

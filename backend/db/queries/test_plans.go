package queries

const TestPlanColumns = "id, project_id, name, description, created_at"

const (
	GetTestPlansByProject = "SELECT " + TestPlanColumns + " FROM test_plans WHERE project_id = $1"
	GetTestPlanByID       = "SELECT " + TestPlanColumns + " FROM test_plans WHERE id = $1"
	InsertTestPlan        = "INSERT INTO test_plans (project_id, name, description) VALUES ($1, $2, $3) RETURNING " + TestPlanColumns
	DeleteTestPlan        = "DELETE FROM test_plans WHERE id = $1"
)

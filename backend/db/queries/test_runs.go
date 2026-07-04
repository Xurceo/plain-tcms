package queries

const TestRunColumns = "id, project_id, plan_id, name, status, environment, build_version, created_by, created_at, completed_at"

const (
	GetTestRunsByProject = "SELECT " + TestRunColumns + " FROM test_runs WHERE project_id = $1"
	GetTestRunByID       = "SELECT " + TestRunColumns + " FROM test_runs WHERE id = $1"
	InsertTestRun        = "INSERT INTO test_runs (project_id, plan_id, name, environment, build_version, created_by) VALUES ($1, $2, $3, $4, $5, $6) RETURNING " + TestRunColumns
	DeleteTestRun        = "DELETE FROM test_runs WHERE id = $1"
)

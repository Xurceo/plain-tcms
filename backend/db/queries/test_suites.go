package queries

const TestSuiteColumns = "id, project_id, parent_id, name, description, created_at"

const (
	GetTestSuitesByProject = "SELECT " + TestSuiteColumns + " FROM test_suites WHERE project_id = $1"
	GetTestSuiteByID       = "SELECT " + TestSuiteColumns + " FROM test_suites WHERE id = $1"
	InsertTestSuite        = "INSERT INTO test_suites (project_id, parent_id, name, description) VALUES ($1, $2, $3, $4) RETURNING " + TestSuiteColumns
	DeleteTestSuite        = "DELETE FROM test_suites WHERE id = $1"
)

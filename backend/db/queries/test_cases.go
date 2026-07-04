package queries

const TestCaseColumns = "id, project_id, suite_id, title, description, preconditions, steps, expected, status, priority, type, tags, created_by, created_at, updated_at"

const (
	GetTestCasesByProject = "SELECT " + TestCaseColumns + " FROM test_cases WHERE project_id = $1"
	GetTestCasesBySuite   = "SELECT " + TestCaseColumns + " FROM test_cases WHERE suite_id = $1"
	GetTestCaseByID       = "SELECT " + TestCaseColumns + " FROM test_cases WHERE id = $1"
	InsertTestCase        = "INSERT INTO test_cases (project_id, suite_id, title, description, preconditions, expected, priority, type, tags, created_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING " + TestCaseColumns
	DeleteTestCase        = "DELETE FROM test_cases WHERE id = $1"
)

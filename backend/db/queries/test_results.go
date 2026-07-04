package queries

const TestResultColumns = "id, run_id, test_case_id, status, comment, executed_by, duration_ms, created_at"

const (
	GetTestResultsByRun = "SELECT " + TestResultColumns + " FROM test_results WHERE run_id = $1"
	GetTestResultByID   = "SELECT " + TestResultColumns + " FROM test_results WHERE id = $1"
	InsertTestResult    = "INSERT INTO test_results (run_id, test_case_id, status, comment, executed_by, duration_ms) VALUES ($1, $2, $3, $4, $5, $6) RETURNING " + TestResultColumns
	DeleteTestResult    = "DELETE FROM test_results WHERE id = $1"
)

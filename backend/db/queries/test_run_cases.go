package queries

const (
	GetCasesByRun     = "SELECT test_case_id FROM test_run_cases WHERE run_id = $1"
	GetRunsByCase     = "SELECT run_id FROM test_run_cases WHERE test_case_id = $1"
	InsertTestRunCase = "INSERT INTO test_run_cases (run_id, test_case_id) VALUES ($1, $2)"
	DeleteTestRunCase = "DELETE FROM test_run_cases WHERE run_id = $1 AND test_case_id = $2"
	DeleteAllRunCases = "DELETE FROM test_run_cases WHERE run_id = $1"
)

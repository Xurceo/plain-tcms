package queries

const TestCaseHistoryColumns = "id, test_case_id, snapshot, changed_by, changed_at"

const (
	GetHistoryByTestCase  = "SELECT " + TestCaseHistoryColumns + " FROM test_case_history WHERE test_case_id = $1 ORDER BY changed_at DESC"
	InsertTestCaseHistory = "INSERT INTO test_case_history (test_case_id, snapshot, changed_by) VALUES ($1, $2, $3) RETURNING " + TestCaseHistoryColumns
)

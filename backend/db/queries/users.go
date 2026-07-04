package queries

const UserColumns = "id, email, created_at"

const (
	GetAllUsers    = "SELECT " + UserColumns + " FROM users"
	GetUserByID    = "SELECT " + UserColumns + " FROM users WHERE id = $1"
	GetUserByEmail = "SELECT id, email, password_hash, created_at FROM users WHERE email = $1"
	InsertUser     = "INSERT INTO users (email, password_hash) VALUES ($1, $2) RETURNING " + UserColumns
	DeleteUser     = "DELETE FROM users WHERE id = $1"
)

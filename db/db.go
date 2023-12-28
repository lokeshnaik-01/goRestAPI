package db
import(
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)
var DB *sql.DB
// we get pointer to database
func InitDB() {
	DB, err := sql.Open("sqlite3", "api.db")
	if(err!=nil) {
		panic("Could not connect to DB")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTables()
}

func createTables() {
	createEventsTable := `CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name varchar(50) NOT NULL,
		description varchar(50) NOT NULL,
		location text NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER
	)`

	_, err := DB.Exec(createEventsTable)
	if err!=nil {
		panic("events table not executed")
	}
}
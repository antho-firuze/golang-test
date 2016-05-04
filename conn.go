import (
	"database/sql"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "host=192.168.0.7 port=5432 user=postgres password=admin123 dbname=db_genesys sslmode=verify-full")
	if err != nil {
		log.Fatal(err)
	}

	id := 11
	rows, err := db.Query("SELECT name FROM a_user WHERE id = $1", id)
	…
}
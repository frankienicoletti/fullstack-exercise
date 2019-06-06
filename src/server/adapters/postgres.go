package adapters

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/frankienicoletti/fullstack-exercise/src/server/types"
	_ "github.com/go-sql-driver/mysql"
)

var _ types.Store = &MySQLStore{}

// MySQLStore is an instance of Store.
type MySQLStore struct {
	db *sql.DB
}

// NewMySQLStore returns a new instance of MySQLStore.
func NewMySQLStore() *MySQLStore {
	db := pingUntilConnected()

	fmt.Println("database is ready")
	return &MySQLStore{db: db}
}

// FindPeople returns a list of people with an optional text search on name.
func (s *MySQLStore) FindPeople(textFilter string) ([]types.Person, error) {
	tx, err := s.db.BeginTx(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	query := `
		SELECT id
		, name
		, img_url
		, location
		FROM people`

	if textFilter != "" {
		query = fmt.Sprintf(
			"%s WHERE lower(name) LIKE '%%%s%%'",
			query,
			strings.ToLower(textFilter),
		)
	}

	rows, err := tx.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var people []types.Person
	var ids []int
	for rows.Next() {
		var p types.Person
		if err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.ImgURL,
			&p.Location,
		); err != nil {
			return nil, err
		}
		people = append(people, p)
		ids = append(ids, p.ID)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	if len(people) == 0 {
		return nil, nil
	}

	colors, err := s.findColors(tx, ids)
	if err != nil {
		return nil, err
	}

	for i, person := range people {
		if personColors, ok := colors[person.ID]; ok {
			people[i].Colors = personColors
		}
	}

	return people, nil
}

// findColors returns a map of people ids with keys of string arrays with their colors.
func (s *MySQLStore) findColors(tx *sql.Tx, ids []int) (map[int][]string, error) {
	peopleIDs := strings.Trim(strings.Replace(fmt.Sprint(ids), " ", ",", -1), "[]")
	query := fmt.Sprintf(`
		SELECT 
			  people_id
			, color
		FROM people_colors
		WHERE people_id in (%s)`,
		peopleIDs,
	)

	rows, err := tx.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	colors := make(map[int][]string)
	for rows.Next() {
		var id int
		var color string
		if err := rows.Scan(
			&id,
			&color,
		); err != nil {
			return nil, err
		}
		if _, ok := colors[id]; ok {
			colors[id] = append(colors[id], color)
		} else {
			colors[id] = []string{color}
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return colors, nil
}

func pingUntilConnected() *sql.DB {
	port := os.Getenv("MYSQL_PORT")
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	name := os.Getenv("MYSQL_DATABASE")
	host := os.Getenv("MYSQL_HOST")

	for true {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, name)
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			panic(err)
		}
		if err := db.Ping(); err == nil {
			return db
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}

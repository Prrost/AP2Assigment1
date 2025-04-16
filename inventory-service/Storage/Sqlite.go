package Storage

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"inventory-service/config"
	"inventory-service/domain"
	"log"
)

type SqliteStorage struct {
	db *sql.DB
}

func NewSqliteStorage(cfg *config.Config) *SqliteStorage {
	db, err := sql.Open("sqlite3", cfg.DBPath)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS objects (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE,
    amount INTEGER NOT NULL,
    available BOOLEAN NOT NULL
);`)

	if err != nil {
		log.Fatal(err)
	}
	return &SqliteStorage{
		db: db,
	}
}

func (s SqliteStorage) CreateObject(object domain.Object) (domain.Object, error) {
	exists, err := s.IsProductExists(object.Name)
	if err != nil {
		return domain.Object{}, err
	}
	if exists {
		return domain.Object{}, ErrAlreadyExists
	}

	stmt, err := s.db.Prepare(`INSERT INTO objects (name, amount, available) VALUES (?, ?, ?)`)
	if err != nil {
		return domain.Object{}, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(object.Name, object.Amount, object.Available)
	if err != nil {
		return domain.Object{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return domain.Object{}, err
	}

	object.ID = int(id)

	return object, nil
}

func (s SqliteStorage) UpdateObjectByID(id int, object domain.Object) (domain.Object, error) {
	stmt, err := s.db.Prepare(`UPDATE objects SET name = ?, amount = ?, available = ? WHERE id = ?`)
	if err != nil {
		return domain.Object{}, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(object.Name, object.Amount, object.Available, id)
	if err != nil {
		return domain.Object{}, err
	}

	return object, nil
}

func (s SqliteStorage) GetObjectByID(id int) (domain.Object, error) {
	var object domain.Object

	err := s.db.QueryRow(`SELECT * FROM objects WHERE id = ?`, id).Scan(&object.ID, &object.Name, &object.Amount, &object.Available)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Object{}, ErrNotFound
		}
		return domain.Object{}, err
	}

	return object, nil
}

func (s SqliteStorage) DeleteObjectByID(id int) (domain.Object, error) {
	object, err := s.GetObjectByID(id)
	if err != nil {
		return domain.Object{}, err
	}
	if object.Name == "" {
		return domain.Object{}, ErrNotFound
	}

	stmt, err := s.db.Prepare(`DELETE FROM objects WHERE id = ?`)
	if err != nil {
		return domain.Object{}, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return domain.Object{}, err
	}
	return object, nil
}

func (s SqliteStorage) IsProductExists(name string) (bool, error) {
	var id int
	err := s.db.QueryRow("SELECT ID FROM objects WHERE Name = ?", name).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (s SqliteStorage) ListProducts(name string, limit int, offset int) ([]domain.Object, error) {
	var rows *sql.Rows
	var err error

	query := `SELECT ID, Name, Amount, Available FROM objects WHERE 1=1`
	args := []interface{}{}

	if name != "" {
		query += " AND Name LIKE ?"
		args = append(args, "%"+name+"%")
	}

	query += " LIMIT ? OFFSET ?"
	args = append(args, limit, offset)

	rows, err = s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []domain.Object
	for rows.Next() {
		var p domain.Object
		err = rows.Scan(&p.ID, &p.Name, &p.Amount, &p.Available)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

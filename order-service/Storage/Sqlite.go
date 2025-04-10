package Storage

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"order-service/config"
	"order-service/domain"
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

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS orders (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    userid INTEGER NOT NULL,
	productid INTEGER NOT NULL,
    amount INTEGER NOT NULL,
    status TEXT NOT NULL
);`)

	if err != nil {
		log.Fatal(err)
	}
	return &SqliteStorage{
		db: db,
	}
}

func (s SqliteStorage) CreateOrderX(object domain.Order) (domain.Order, error) {

	stmt, err := s.db.Prepare(`INSERT INTO orders (userid, productid, amount, status) VALUES (?, ?, ?, ?)`)
	if err != nil {
		return domain.Order{}, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(object.UserID, object.ProductID, object.Amount, object.Status)
	if err != nil {
		return domain.Order{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return domain.Order{}, err
	}

	object.ID = int(id)

	return object, nil
}

func (s SqliteStorage) UpdateOrderByIDX(id int, object domain.Order) (domain.Order, error) {

	stmt, err := s.db.Prepare(`UPDATE orders SET userid = ?, productid = ?, amount = ?, status = ? WHERE id = ?`)
	if err != nil {
		return domain.Order{}, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(object.UserID, object.ProductID, object.Amount, object.Status, id)
	if err != nil {
		return domain.Order{}, err
	}

	return object, nil
}

func (s SqliteStorage) GetOrderByIDX(id int) (domain.Order, error) {
	var object domain.Order

	err := s.db.QueryRow(`SELECT * FROM orders WHERE id = ?`, id).Scan(&object.ID, &object.UserID, &object.ProductID, &object.Amount, &object.Status)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Order{}, err
		}
		return domain.Order{}, err
	}

	return object, nil
}

func (s SqliteStorage) ListAllOrdersX() ([]domain.Order, error) {

	rows, err := s.db.Query(`SELECT id, userid, productid, amount, status FROM orders`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []domain.Order

	for rows.Next() {
		var o domain.Order
		err := rows.Scan(&o.ID, &o.UserID, &o.ProductID, &o.Amount, &o.Status)
		if err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}

	return orders, nil

}

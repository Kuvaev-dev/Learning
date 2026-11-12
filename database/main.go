package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Human struct {
	ID   int
	Name string
	Age  int
}

func Connect() *sql.DB {
	connStr := "postgres://postgres:<пароль>@127.0.0.1:5432/humandb?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}
	return db
}

func Insert(db *sql.DB) {
	res, err := db.Exec(
		"insert into human(name, age) values ($1, $2), ($3, $4)",
		"Vasya", 10, "Petya", 15,
	)
	if err != nil {
		panic(err)
	}
	n, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("rows inserted:", n)
}

func Select(db *sql.DB) {
	rows, err := db.Query("select id, name, age from human order by id")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	humans := []Human{}
	for rows.Next() {
		var h Human
		if err := rows.Scan(&h.ID, &h.Name, &h.Age); err != nil {
			panic(err)
		}
		humans = append(humans, h)
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}
	for _, h := range humans {
		fmt.Println(h.ID, h.Name, h.Age)
	}
}

func Update(db *sql.DB) {
	res, err := db.Exec("update human set age = $1 where id = $2", 5, 1)
	if err != nil {
		panic(err)
	}
	n, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("rows updated:", n)
}

func Delete(db *sql.DB) {
	res, err := db.Exec("delete from human where id = $1", 1)
	if err != nil {
		panic(err)
	}
	n, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("rows deleted:", n)
}

func UsePostgresDB() {
	db := Connect()
	defer db.Close()

	// Insert(db)
	// Update(db)
	Delete(db)
	Select(db)
}

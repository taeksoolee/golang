package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := connection()
	defer db.Close()

	// queryRow(db)
	// query(db)

	// dmlInsert(db)
	// preparedUpdate(db)

	transaction(db)
}

func connection() *sql.DB {
	db, err := sql.Open("mysql", "root:toor@tcp(127.0.0.1:3306)/golang_db")
	if err != nil {
		log.Fatal(err)
	}

	return db
}

// 1줄 쿼리
func queryRow(db *sql.DB) {
	var title string
	err := db.QueryRow("SELECT b_title title FROM board where idx=2").Scan(&title)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(title)
}

func query(db *sql.DB) {
	var title string
	var content string
	rows, err := db.Query("SELECT b_title title, b_content content FROM board")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&title, &content)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(title, content)
	}
}

func dmlInsert(db *sql.DB) {
	result, err := db.Exec("INSERT INTO board(b_title, b_content) VALUES (?, ?)", "안녕", "하이요")
	if err != nil {
		log.Fatal(err)
	}

	n, err := result.RowsAffected()
	if n == 1 {
		fmt.Println("1 row insert")
	}
}

func preparedUpdate(db *sql.DB) {
	smtp, err := db.Prepare("UPDATE board SET b_content=? WHERE idx=?")
	if err != nil {
		panic(err)
	}

	_, err = smtp.Exec("i like golang", 1)
	if err != nil {
		panic(err)
	}

	_, err = smtp.Exec("golang is boring", 2)
	if err != nil {
		panic(err)
	}
}

func transaction(db *sql.DB) {
	tx, err := db.Begin()
	if err != nil {
		log.Panic(err)
	}
	
	// 중간 에러시 롤백 처리
	defer tx.Rollback()

	_, err =tx.Exec("INSERT INTO board(b_title, b_content) VALUES (?, ?)", "아침", "아침밥")

	if err != nil {
		log.Panic(err)
	}

	_, err =tx.Exec("INSERT INTO board(b_title, b_content) VALUES (?, ?)", "안녕", "안녕 하세요...")
	if err != nil {
		log.Panic(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Panic(err)
	}

}
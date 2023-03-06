package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type address struct {
	s_name string
	i_age  int
}

type db_conn struct {
	db *sql.DB
}

func main() {
	db, err := sql.Open("mysql", "root:RLGH3qjs!!@tcp(localhost:3306)/addrbook")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	conn := &db_conn{db: db}

	//추가
	//conn.insert(address{"sss", 33})
	i_result, i_Err := conn.insert(address{"sss", 33})
	fmt.Println("result:", i_result, "err:", i_Err)
	v, err := i_result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("result:", v, "err:", i_Err)
	//삭제
	//conn.delete("sss")
	//수정
	//conn.update("sss", 22)
	//조회
	result := conn.read()

	for _, v := range result {
		fmt.Printf("name: %s, age: %d\n", v.s_name, v.i_age)
	}

	t_err := conn.truncate()
	fmt.Println(t_err)
}

func (conn *db_conn) insert(address address) (sql.Result, error) {
	query := "insert into info (name,age) values (?,?)"
	v, err := conn.db.Exec(query, address.s_name, address.i_age)
	if err != nil {
		log.Fatal(err)
	}
	return v, err
}

func (conn *db_conn) delete(name string) error {
	query := "delete from info where name = ?"
	_, err := conn.db.Exec(query, name)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func (conn *db_conn) update(name string, age int) error {
	query := "update info set age = ? where name = ?"
	_, err := conn.db.Exec(query, age, name)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func (conn *db_conn) read() []address {
	query := "select * from info"
	rows, err := conn.db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var result []address
	for rows.Next() {
		var a address
		err := rows.Scan(&a.s_name, &a.i_age)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, a)
	}
	return result
}

func (conn *db_conn) truncate() error {
	query := "truncate table info"
	_, err := conn.db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

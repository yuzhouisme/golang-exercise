package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/pkg/errors"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DRIVER = "mysql"
	SOURCE = "root:123456@tcp(127.0.0.1:3306)/test?parseTime=True"
)

type User struct {
	ID           string `db:"id"`
	BusinessName string `db:"business_name"`
	Address      string `db:"address"`
	Phone        string `db:"phone"`
	RoleType     string `db:"role_type"`
	Status       string `db:"status"`
}

var (
	db *sql.DB
)

func main() {
	var err error
	db, err = sql.Open(DRIVER, SOURCE)
	if err != nil {
		fmt.Printf("!!! OPEN DATABASE ERROR, err: %v\n", err)
		panic(err)
	}
	db.SetConnMaxLifetime(100 * time.Second)
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(16)

	clientGetMessage := getClientById("NotExistedId")
	fmt.Printf("%v\n", clientGetMessage)
}

func getClientById(id string) string {
	// mock: call DAO func
	user, err := getOne(id)
	if err != nil {
		// write into log
		fmt.Printf("%+v\n", err)
		//
		if errors.Is(err, sql.ErrNoRows) {
			return "No User"
		} else {
			return err.Error();
		}
	}

	return fmt.Sprintf("SUCCEED!!! User = %v \n", user)
}

func getOne(id string) (*User, error) {
	user := User{ID: id}
	row := db.QueryRow("SELECT * FROM clients WHERE id=?", id)
	if err := row.Scan(&user.ID, &user.BusinessName, &user.Address, &user.Phone, &user.RoleType, &user.Status); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("UserService: getOne id = %v", id))
	}
	return &user, nil
}

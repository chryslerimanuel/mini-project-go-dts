package entities

import (
	"context"
	"log"
	"mini-project-go-dts/configs"
	"time"
)

type User struct {
	Id       int
	Name     string
	RoleId   int
	IsActive bool
}

func (u User) Insert() error {
	db, err := configs.GetConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := "INSERT INTO user (`name`, `roleId`, `isActive`) VALUES ( ?, ?, true );"

	ctx, cancelfunc := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, u.Name, u.RoleId)
	if err != nil {
		log.Printf("Error %s when inserting row into user table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}
	log.Printf("%d user created ", rows)

	return nil
}

func (u User) Update(id int) error {
	db, err := configs.GetConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := "UPDATE user SET name = ?, roleId = ? WHERE Id = ?"

	ctx, cancelfunc := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, u.Name, u.RoleId, id)
	if err != nil {
		log.Printf("Error %s when updating row into user table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}
	log.Printf("%d user updated ", rows)
	return nil
}

func (u User) Delete(id int) error {
	db, err := configs.GetConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := "UPDATE user SET IsActive = false WHERE Id = ?"

	ctx, cancelfunc := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		log.Printf("Error %s when updating row into user table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}
	log.Printf("%d user created ", rows)
	return nil
}

func (u User) FindAll() ([]User, error) {
	db, err := configs.GetConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := "SELECT * FROM user WHERE IsActive = true;"

	ctx, cancelfunc := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		log.Printf("Error %s when getting rows from user table", err)
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		err = rows.Scan(&u.Id, &u.Name, &u.RoleId, &u.IsActive)
		if err != nil {
			panic(err)
		}
		users = append(users, u)
	}

	return users, nil
}

func (u User) Select(id int) (User, error) {
	db, err := configs.GetConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := "SELECT * FROM user WHERE Id = ? ;"

	ctx, cancelfunc := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return u, err
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, id)
	if err != nil {
		log.Printf("Error %s when getting data from user table", err)
		return u, err
	}

	err = row.Scan(&u.Id, &u.Name, &u.RoleId, &u.IsActive)
	if err != nil {
		log.Printf("Error %s when scan data from user table", err)
		return u, err
	}

	return u, nil
}

func (u User) FindAllWithFilter(searchParam string) ([]User, error) {
	db, err := configs.GetConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := "SELECT * FROM user WHERE IsActive = true AND Name LIKE ? ;"

	ctx, cancelfunc := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, "%"+searchParam+"%")
	if err != nil {
		log.Printf("Error %s when getting rows from user table", err)
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		err = rows.Scan(&u.Id, &u.Name, &u.RoleId, &u.IsActive)
		if err != nil {
			panic(err)
		}
		users = append(users, u)
	}

	return users, nil
}

package entities

import (
	"context"
	"database/sql"
	"log"
	"mini-project-go-dts/configs"
	"time"
)

type Task struct {
	Id            int
	TaskDetail    string
	CreatedById   int
	CreatedByName string
	SendToId      int
	SendToName    string
	TaskDeadLine  string
	IsDone        bool
	IsActive      bool
}

type TaskModel struct {
	Id            int
	TaskDetail    sql.NullString
	CreatedById   sql.NullInt64
	CreatedByName sql.NullString
	SendToId      sql.NullInt64
	SendToName    sql.NullString
	TaskDeadLine  sql.NullString
	IsDone        bool
	IsActive      bool
}

func (t Task) Insert() error {
	db, err := configs.GetConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := "INSERT INTO task (`TaskDetail`, `CreatedById`, `CreatedByName`, `SendToId`, `SendToName`, `TaskDeadLine`, `IsDone`, `IsActive`) VALUES ( ?, 1, ?, ?, ?, ?, false, true );"

	ctx, cancelfunc := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, t.TaskDetail, t.CreatedByName, t.SendToId, t.SendToName, t.TaskDeadLine)
	if err != nil {
		log.Printf("Error %s when inserting row into task table", err)
		// return err
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}
	log.Printf("%d task created ", rows)

	return err
}

func (t Task) Update(id int) error {
	db, err := configs.GetConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := "UPDATE task SET TaskDetail = ?, SendToId = ?, SendToName = ?, TaskDeadLine = ? WHERE Id = ?"

	ctx, cancelfunc := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, t.TaskDetail, t.SendToId, t.SendToName, t.TaskDeadLine, id)
	if err != nil {
		log.Printf("Error %s when updating row into task table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}
	log.Printf("%d task updated ", rows)
	return nil
}

func (t Task) Delete(id int) error {
	db, err := configs.GetConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := "UPDATE task SET IsActive = false WHERE Id = ?"

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
		log.Printf("Error %s when updating row into task table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}
	log.Printf("%d task created ", rows)
	return nil
}

func (t Task) FindAll() ([]Task, error) {
	db, err := configs.GetConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := "SELECT * FROM task WHERE IsActive = true;"

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
		log.Printf("Error %s when getting rows from task table", err)
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var t TaskModel
		err = rows.Scan(&t.Id, &t.TaskDetail, &t.CreatedById, &t.CreatedByName, &t.SendToId, &t.SendToName, &t.TaskDeadLine, &t.IsDone, &t.IsActive)
		if err != nil {
			log.Println(err.Error())
			panic(err)
		}

		var t2 Task
		t2.Id = t.Id
		t2.TaskDetail = t.TaskDetail.String
		t2.CreatedByName = t.CreatedByName.String
		t2.CreatedById = int(t.CreatedById.Int64)
		t2.SendToName = t.SendToName.String
		t2.SendToId = int(t.SendToId.Int64)
		t2.TaskDeadLine = t.TaskDeadLine.String
		t2.IsDone = t.IsDone
		t2.IsActive = t.IsActive

		tasks = append(tasks, t2)
	}

	return tasks, nil
}

func (t Task) Select(id int) (Task, error) {
	db, err := configs.GetConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := "SELECT * FROM task WHERE Id = ? ;"

	ctx, cancelfunc := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return t, err
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, id)
	if err != nil {
		log.Printf("Error %s when getting data from task table", err)
		return t, err
	}
	var t2 TaskModel
	err = row.Scan(&t2.Id, &t2.TaskDetail, &t2.CreatedById, &t2.CreatedByName, &t2.SendToId, &t2.SendToName, &t2.TaskDeadLine, &t2.IsDone, &t2.IsActive)
	if err != nil {
		log.Printf("Error %s when scan data from task table", err)
		return t, err
	}

	t.Id = t2.Id
	t.TaskDetail = t2.TaskDetail.String
	t.CreatedByName = t2.CreatedByName.String
	t.CreatedById = int(t2.CreatedById.Int64)
	t.SendToName = t2.SendToName.String
	t.SendToId = int(t2.SendToId.Int64)
	t.TaskDeadLine = t2.TaskDeadLine.String
	t.IsDone = t2.IsDone
	t.IsActive = t2.IsActive

	return t, nil
}

func (t Task) UpdateStatus(id int) error {
	db, err := configs.GetConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := "UPDATE task SET IsDone = true WHERE Id = ?"

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
		log.Printf("Error %s when updating row into task table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}
	log.Printf("%d task updated ", rows)
	return nil
}

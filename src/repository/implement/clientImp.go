package implement

import (
	models "DBdemo/src/model"
	repo "DBdemo/src/repository"
	"database/sql"
	"fmt"
)

type ClientImp struct {
	Db *sql.DB
}

func NewClientRepo(db *sql.DB) repo.ClientRepo {
	return &ClientImp{
		Db: db,
	}
}

func (u *ClientImp) Select() ([]models.Client, error) {
	clients := make([]models.Client, 0)
	rows, err := u.Db.Query("SELECT * FROM CLIENT")
	if err != nil {
		return clients, err
	}

	for rows.Next() {
		client := models.Client{}
		err := rows.Scan(&client.ID, &client.Name, &client.Age, &client.Gender, &client.Email)
		if err != nil {
			break
		}
		clients = append(clients, client)
	}
	err = rows.Err()
	if err != nil {
		return clients, err
	}
	fmt.Println("select")
	return clients, nil
}

func (u *ClientImp) Insert(client models.Client) error {
	insertStatement := `INSERT INTO client (ID, Name, Age, Gender, Email)
						VALUES ($1, $2, $3, $4, $5)`

	_, err := u.Db.Exec(insertStatement, client.ID, client.Name, client.Age, client.Gender, client.Email)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("Added new client", client)
	return nil
}

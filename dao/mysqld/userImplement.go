package mysqld

import (
	"fmt"

	"../../models"
)

type UserImplMysql struct{}

func (dao UserImplMysql) Create(u *models.User) error {

	query := "INSERT INTO users (first_name, last_name, email) values (?,?,?)"

	db := get()
	defer db.Close()

	smtm, err := db.Prepare(query)
	if err != nil {
		return err
	}

	defer smtm.Close()

	fmt.Println(u)

	result, err := smtm.Exec(u.FirstName, u.LastName, u.Email)

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	u.ID = id
	return nil
}

func (dao UserImplMysql) GetAll() ([]models.User, error) {
	query := "SELEC * FROM users"
	users := make([]models.User, 0)

	db := get()
	defer db.Close()

	smtm, err := db.Prepare(query)
	if err != nil {
		return users, err
	}

	defer smtm.Close()

	rows, err := smtm.Query()
	if err != nil {
		return users, err
	}

	for rows.Next() {
		var row models.User
		err := rows.Scan(&row.ID, &row.FirstName, &row.LastName, &row.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, row)
	}

	return users, nil
}

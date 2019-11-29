package implementation

import (
	repo "apiProyecto/database/repository/interface"
	"apiProyecto/models"
	"database/sql"
	"errors"
)

type UserRepoImpl struct {
	Db *sql.DB
}

func (u UserRepoImpl) Select() ([]models.User, error) {
	user_list := make([]models.User, 0)
	rows, err := u.Db.Query("Select id_user, username, password, names, last_names from user")
	if err != nil {
		return user_list, err
	}
	for rows.Next() {
		user := models.User{}
		err := rows.Scan(&user.IdUser, &user.Username, &user.Password, &user.Names, &user.LastNames)
		if err != nil {
			break
		}
		user_list = append(user_list, user)
	}
	err = rows.Err()
	if err != nil {
		return user_list, err
	}
	return user_list, nil
}

func (u UserRepoImpl) Filter(user *models.User) ([]models.User, bool, error) {
	user_list := make([]models.User, 0)
	rows, err := u.Db.Query("Select id_user, username, names, last_names from users where username = $1", user.Username)
	if err != nil {
		return user_list, false, err
	}
	for rows.Next() {
		user := models.User{}
		err := rows.Scan(&user.IdUser, &user.Username, &user.Names, &user.LastNames)
		if err != nil {
			break
		}
		user_list = append(user_list, user)
	}
	err = rows.Err()
	if err != nil {
		return user_list, false, err
	}
	return user_list, true, nil
}

func (u UserRepoImpl) Insert(user *models.User) (bool, error) {
	// new, err := u.Db.Query("Insert into users (username,password) values ($1,$2)", user.Username, user.Password)

	// if err != nil {
	// 	return false, err
	// 	new.Close()
	// }

	query, err := u.Db.Prepare("Insert into users (username,names,last_names,password,state) values ($1,$2,$3,$4,1)")
	if err != nil {
		return false, err
	}

	result, err := query.Exec(user.Username, user.Names, user.LastNames, user.Password)
	if err != nil {
		return false, err
	}

	i, err := result.RowsAffected()
	if i != 1 {
		return false, errors.New("Error: Muchas filas afectadas")
	}

	return true, nil
}

func NewUserRepoImpl(db *sql.DB) repo.UserInterface {
	return &UserRepoImpl{Db: db}
}

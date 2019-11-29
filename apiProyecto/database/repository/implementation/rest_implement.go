package implementation

import (
	repo "apiProyecto/database/repository/interface"
	"apiProyecto/models"
	"database/sql"
	"errors"
	"fmt"
)

type RestRepoImpl struct {
	Db *sql.DB
}

func (r RestRepoImpl) Select() ([]models.Restaurant, error) {
	rest_list := make([]models.Restaurant, 0)
	rows, err := r.Db.Query("Select id_restaurant, rest_name, rest_address, rest_phone, rest_image from restaurant")
	if err != nil {

		fmt.Println(err)
		return rest_list, err
	}
	for rows.Next() {
		restaurant := models.Restaurant{}
		err := rows.Scan(&restaurant.IdRest, &restaurant.RestName, &restaurant.RestAddress, &restaurant.RestPhone, &restaurant.RestImage)
		if err != nil {
			fmt.Println(err)
			break
		}
		rest_list = append(rest_list, restaurant)
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
		return rest_list, err
	}
	return rest_list, nil
}

func (r RestRepoImpl) Filter(rest *models.Restaurant) ([]models.Restaurant, bool, error) {
	rest_list := make([]models.Restaurant, 0)
	rows, err := r.Db.Query("Select id_restaurant, rest_name, rest_address, rest_phone, rest_image from restaurant where restName = $1", rest.RestName)
	if err != nil {
		return rest_list, false, err
	}
	for rows.Next() {
		restaurant := models.Restaurant{}
		err := rows.Scan(&restaurant.IdRest, &restaurant.RestName, &restaurant.RestAddress, &restaurant.RestPhone, &restaurant.RestImage)
		if err != nil {
			break
		}
		rest_list = append(rest_list, restaurant)
	}
	err = rows.Err()
	if err != nil {
		return rest_list, false, err
	}
	return rest_list, true, nil
}

func (r RestRepoImpl) Insert(rest *models.Restaurant) (bool, error) {
	// new, err := u.Db.Query("Insert into users (username,password) values ($1,$2)", user.Username, user.Password)

	// if err != nil {
	// 	return false, err
	// 	new.Close()
	// }

	query, err := r.Db.Prepare("Insert into restaurant (rest_name, rest_image, rest_address, rest_phone, state) values ($1,$2,$3,$4,1)")
	if err != nil {
		return false, err
	}

	result, err := query.Exec(rest.RestName, rest.RestImage, rest.RestAddress, rest.RestPhone)
	if err != nil {
		return false, err
	}

	i, err := result.RowsAffected()
	if i != 1 {
		return false, errors.New("Error: Muchas filas afectadas")
	}

	return true, nil
}

func NewRestRepoImpl(db *sql.DB) repo.RestInterface {
	return &RestRepoImpl{Db: db}
}

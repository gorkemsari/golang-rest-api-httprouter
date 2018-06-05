package repository

import "github.com/gorkemsari/golang-rest-api-httprouter/model"

func GetAllCities() ([]*model.City, error) {
	rows, err := db.Query("SELECT * FROM city")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cities := make([]*model.City, 0)

	for rows.Next() {
		city := new(model.City)
		err := rows.Scan(&city.Id, &city.Name)
		if err != nil {
			return nil, err
		}
		cities = append(cities, city)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return cities, nil
}

func GetCityById(id string) (*model.City, error) {
	city := new(model.City)
	err := db.QueryRow("SELECT * FROM city WHERE id = ?", id).Scan(&city.Id, &city.Name)
	if err != nil {
		return nil, err
	}

	return city, nil
}
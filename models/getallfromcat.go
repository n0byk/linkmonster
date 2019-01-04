package models

import (
	"../config"
)

//получение последних н записей сортировка по дате добавления
func GetAllFromCat(n int) ([]GetifobyUID, error) {

	rows, err := config.DB.Query("SELECT url, title, ico, add_date, description_short, description_full, visits FROM catalog_data  ORDER BY add_date DESC FETCH FIRST  $1 rows ONLY", n)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	newdata := make([]GetifobyUID, 0)
	for rows.Next() {
		new := GetifobyUID{}
		err := rows.Scan(&new.Urllink, &new.Title, &new.Ico, &new.Adddate, &new.Descriptionshort, &new.Descriptionfull, &new.Visits)
		if err != nil {
			return nil, err
		}
		newdata = append(newdata, new)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return newdata, nil
}

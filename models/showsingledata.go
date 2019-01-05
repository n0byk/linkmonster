package models

import (
	"net/http"
	"strings"

	"../config"
)

type GetifobyUID struct {
	Urllink          string
	Title            string
	Ico              string
	Adddate          string
	Descriptionshort string
	Descriptionfull  string
	Visits           int
}

func ShowSingleData(w http.ResponseWriter, r *http.Request) (GetifobyUID, error) {
	singledata := GetifobyUID{}
	getinfo := strings.TrimSpace(r.FormValue("getinfo"))
	/*if getinfo == "" {
		helpers.ShowError(w, "404")
		return singledata, nil
	}
	*/
	row := config.DB.QueryRow("SELECT 	url, title, ico, add_date, description_short, description_full, visits FROM catalog_data WHERE data_id = $1 and active = true", getinfo)

	err := row.Scan(&singledata.Urllink, &singledata.Title, &singledata.Ico, &singledata.Adddate, &singledata.Descriptionshort, &singledata.Descriptionfull, &singledata.Visits)
	if err != nil {
		return singledata, err
	}
	/*	//counter update
		userip := helpers.GetMD5Hash(r.Referer())
		fmt.Println(userip)

		num := 0
		rows := config.DB.QueryRow("SELECT COUNT(data_id) FROM data_visits WHERE  data_id = '$1' AND user_hash = '$2'", getinfo, userip)
		counter := rows.Scan(&num)
		fmt.Println(counter)
		if counter == nil {
			_, err = config.DB.Exec("INSERT INTO data_visits (data_id, user_hash) VALUES ($1, $2)", getinfo, userip)
			if err != nil {
				return singledata, errors.New("500. Internal Server Error." + err.Error())
			}
			_, err = config.DB.Exec("UPDATE catalog_data SET visits = visits + 1 WHERE data_id = $1", getinfo)
			if err != nil {
				return singledata, errors.New("500. Internal Server Error." + err.Error())
			}
		}
	*/
	return singledata, nil
}

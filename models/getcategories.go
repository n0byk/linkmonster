package models

import (
	"../config"
)

type Cat_tree struct {
	Cat_id    int
	Parent_id int
	Cat_name  string
}

func GetCategories() ([]Cat_tree, error) {
	rows, err := config.DB.Query("SELECT cat_id, parent_id, cat_name FROM categories WHERE active = true ORDER BY Parent_id ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := make([]Cat_tree, 0)

	for rows.Next() {
		cat := Cat_tree{}
		err := rows.Scan(&cat.Cat_id, &cat.Parent_id, &cat.Cat_name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, cat)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	//fmt.Println(categories)
	//printTree(categories, 0, 1)
	return categories, nil

}

package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	_ "github.com/lib/pq"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letters = []rune("abcdefghijklmnopqrstuvwxyz ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:qwer@localhost/linkmonster?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database.")
	rand.Seed(time.Now().UnixNano())

	for a := 0; a < 150; a++ {
		_, err = db.Exec("INSERT INTO categories (parent_id, cat_name,keywords,cat_description ) VALUES ($1, $2, $3, $4)", rand.Intn(50), randSeq(15), randSeq(30), randSeq(100))
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	/*
		    for a := 0; a < 100000; a++ {
				_, err = db.Exec("INSERT INTO catalog_data(url,title,ico,description_short,description_full,mark) VALUES ($1, $2, $3, $4, $5, $6)", randSeq(10), randSeq(30), randSeq(30), randSeq(200), randSeq(600), randSeq(50))
				if err != nil {
					fmt.Println(err)
					return
				}
			}
	*/

}

/*
for a := 0; a < 10; a++ {
    fmt.Printf("value of a: %d\n", a)
}
*/

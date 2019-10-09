package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type Menu struct {
	Id       int
	Name     string
	Category string
	Type     string
	Price    int
}

func connectionDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./cocos.db")
	if err != nil {
		panic(err)
	}
	return db
}

func getHamburg(db *sql.DB) *sql.Rows {
	row, err := db.Query("SELECT * FROM menus WHERE Type='hamburg' ORDER BY RANDOM() LIMIT 1")
	if err != nil {
		panic(err)
	}
	return row
}

func getSteak(db *sql.DB) *sql.Rows {
	row, err := db.Query("SELECT * FROM menus WHERE Type='steak' ORDER BY RANDOM() LIMIT 1")
	if err != nil {
		panic(err)
	}
	return row
}

func getSpaghetti(db *sql.DB) *sql.Rows {
	row, err := db.Query("SELECT * FROM menus WHERE Type='spaghetti' ORDER BY RANDOM() LIMIT 1")
	if err != nil {
		panic(err)
	}
	return row
}

func getPlate(db *sql.DB) *sql.Rows {
	row, err := db.Query("SELECT * FROM menus WHERE Type='plate' ORDER BY RANDOM() LIMIT 1")
	if err != nil {
		panic(err)
	}
	return row
}

func getSet(db *sql.DB) *sql.Rows {
	row, err := db.Query("SELECT * FROM menus WHERE Type='set' ORDER BY RANDOM() LIMIT 1")
	if err != nil {
		panic(err)
	}
	return row
}

func getSidedish(db *sql.DB) *sql.Rows {
	row, err := db.Query("SELECT * FROM menus WHERE Type='sidedish' ORDER BY RANDOM() LIMIT 1")
	if err != nil {
		panic(err)
	}
	return row
}

func getSalad(db *sql.DB) *sql.Rows {
	row, err := db.Query("SELECT * FROM menus WHERE Type='salad' ORDER BY RANDOM() LIMIT 1")
	if err != nil {
		panic(err)
	}
	return row
}

func getJapan(db *sql.DB) *sql.Rows {
	row, err := db.Query("SELECT * FROM menus WHERE Type='japan' ORDER BY RANDOM() LIMIT 1")
	if err != nil {
		panic(err)
	}
	return row
}

func getDessert(db *sql.DB) *sql.Rows {
	row, err := db.Query("SELECT * FROM menus WHERE Type='dessert' ORDER BY RANDOM() LIMIT 1")
	if err != nil {
		panic(err)
	}
	return row
}

func main() {
	db := connectionDB()
	defer db.Close()

	router := gin.Default()
	router.LoadHTMLGlob("views/*.html")
	router.Static("/assets", "./assets")

	router.GET("/", func(g *gin.Context) {
		g.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.POST("/result", func(p *gin.Context) {
		err := p.Request.ParseForm()
		if err != nil {
			panic(err)
		}

		fmt.Println(p.Request.Form)

		//0:hamburg, 1:steak, 2:spaghetti, 3:plate, 4:set, 5:sidedish, 6:salad, 7:japan, 8:dessert
		check := p.Request.Form["menu"]
		fmt.Println(check)

		temp := Menu{}
		var result []Menu
		var row *sql.Rows
		total := 0

		for i := 0; i < len(check); i++ {
			switch check[i] {
			case "0":
				row = getHamburg(db)
			case "1":
				row = getSteak(db)
			case "2":
				row = getSpaghetti(db)
			case "3":
				row = getPlate(db)
			case "4":
				row = getSet(db)
			case "5":
				row = getSidedish(db)
			case "6":
				row = getSalad(db)
			case "7":
				row = getJapan(db)
			case "8":
				row = getDessert(db)
			default:
				fmt.Println("defaultが実行されました")
			}

			for row.Next() {
				err := row.Scan(&temp.Id, &temp.Name, &temp.Category, &temp.Type, &temp.Price)
				if err != nil {
					panic(err)
				} else {
					result = append(result, temp) //append(追加先のリスト, 追加したいもの)
					total += temp.Price
				}
			}
		}

		p.HTML(200, "result.html", result)
	})

	router.Run(":8080")
}

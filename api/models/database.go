package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

//postgres://studblogff_user:gEIUkjxgIcqo38tp3T7dYc37pmyqIOID@dpg-ciudbf5gkuvoig806uhg-a.oregon-postgres.render.com/studblogff

const (
	host     = "dpg-ciudbf5gkuvoig806uhg-a"
	port     = 5432
	user     = "studblogff_user"
	password = "gEIUkjxgIcqo38tp3T7dYc37pmyqIOID"
	dbname   = "studblogff"
)

var db *gorm.DB

func init() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=require",
		host, port, user, password, dbname)

	fmt.Println(psqlInfo)
	conn, err := gorm.Open("postgres", "postgres://studblogff_user:gEIUkjxgIcqo38tp3T7dYc37pmyqIOID@dpg-ciudbf5gkuvoig806uhg-a.oregon-postgres.render.com/studblogff")
	if err != nil {
		panic(err)
	}
	db = conn
	db.Debug().AutoMigrate(User{}, Post{},
		Image{}, Tag{}, PostTag{}, Like{}, View{}, Favorite{}, FavoritesList{}, Comment{})
	fmt.Println("Veri Tabanı bağlantısı başarılı!")
}
func GetDB() *gorm.DB {
	return db
}

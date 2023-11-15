package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type dbInfo struct {
	user     string
	pwd      string
	url      string
	engine   string
	database string
}

func getDbInfo() (dbInfo) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return dbInfo{
		os.Getenv("DBUSER"),
		os.Getenv("DBPWD"),
		os.Getenv("DBURL"),
		os.Getenv("DBENGINE"),
		os.Getenv("DBNAME"),
	}
}
func DataSource() string{
	db := getDbInfo()
	return db.user + ":" + db.pwd +
	"@tcp(" + db.url + ")/" + db.database
}
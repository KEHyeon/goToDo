package database

import (
	"os"
)

type dbInfo struct {
	user     string
	pwd      string
	url      string
	engine   string
	database string
}

var db = dbInfo{
	os.Getenv("DBUSER"),
	os.Getenv("DBPWD"),
	os.Getenv("DBURL"),
	os.Getenv("DBENGINE"),
	os.Getenv("DBNAME"),
}

var DataSource = db.user + ":" + db.pwd +
	"@tcp(" + db.url + ")/" + db.database
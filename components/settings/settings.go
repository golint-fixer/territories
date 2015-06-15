package settings

var (
	Port string

	DB struct {
		Engine string
		Source string
	}
)

func init() {
	//DB.Engine = "mysql"
	//DB.Source = "sample:sample@/goblog?charset=utf8&parseTime=true"
	DB.Engine = "sqlite3"
	DB.Source = "/tmp/db.sqlite"
}

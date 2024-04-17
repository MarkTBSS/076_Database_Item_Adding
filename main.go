package main

import (
	"fmt"

	"github.com/MarkTBSS/076_Database_Item_Adding/config"
	"github.com/MarkTBSS/076_Database_Item_Adding/databases"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.Database)
	fmt.Println(db)
}

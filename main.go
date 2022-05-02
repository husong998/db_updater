package main

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/husong998/db_updater/app"
	"github.com/husong998/db_updater/app/db"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Conf struct {
	User     string
	Password string
	DB       string
	CSV      string
}

func main() {
	var conf Conf
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	pflag.String("csv", "", "path to csv file")
	pflag.Parse()
	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		panic(err)
	}
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&conf); err != nil {
		panic(err)
	}

	database, err := sql.Open("mysql", conf.User+":"+conf.Password+"@/"+conf.DB+"?multiStatements=true")
	if err != nil {
		panic(err)
	}

	parser := &app.Parser{}
	upserter := &app.Upserter{DB: &db.Adapter{DB: database}}

	parse, err := parser.Parse(conf.CSV)
	if err != nil {
		panic(err)
	}
	if err = upserter.Upsert(context.Background(), parse); err != nil {
		panic(err)
	}
}

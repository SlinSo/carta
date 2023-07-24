package tests

import (
	"bytes"
	"database/sql"
	"html/template"
	"os"
	"strconv"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

var conn *sql.DB

// Atoi return always an int
func Atoi(s string) int {
	ret, err := strconv.Atoi(s)
	if err != nil {
		return -1
	}
	return ret
}

// multi env setup
// Priority	Filename				Environment
// 1st 		.env.development.local	Development		Local overrides of environment-specific settings.
// 1st		.env.test.local			Test			Local overrides of environment-specific settings.
// 1st		.env.production.local	Production		Local overrides of environment-specific settings.
// 2nd		.env.local				Local 			Local overrides. This file is loaded for all environments except test.
// 3rd		.env.development		Development		Shared environment-specific settings
// 3rd		.env.test				Test			Shared environment-specific settings
// 3rd		.env.production			Production		Shared environment-specific settings
// Last		.env					All Environments
func setupEnv() {
	env := os.Getenv("KNOW_ENV")
	if env == "" {
		env = "dev"
	}

	// 1st spefic locals
	_ = godotenv.Load(".env." + env + ".local")

	// 2nd local generic
	if env != "test" {
		_ = godotenv.Load(".env.local")
	}

	// 3rd shared envs
	_ = godotenv.Load(".env." + env)
	_ = godotenv.Load() // FALLBACK .env
}

func initDB() {
	var err error

	conn, err = Connect(os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		Atoi(os.Getenv("DB_PORT")),
		os.Getenv("DB_SCHEMA"),
		os.Getenv("DB_DSN"))
	if err != nil {
		panic(err)
	}
}

// ConnectSqlx connects to DB or panics
func Connect(user string, password string, host string, port int, schema string, dsn string) (*sql.DB, error) {
	var err error
	var connString bytes.Buffer

	para := map[string]interface{}{}
	para["User"] = user
	para["Pass"] = password
	para["Host"] = host
	para["Port"] = port
	para["Schema"] = schema

	tmpl, err := template.New("dbconn").Option("missingkey=zero").Parse(dsn)
	if err != nil {
		log.Error().Err(err).Msg("tmpl parse")
		return nil, err
	}

	err = tmpl.Execute(&connString, para)
	if err != nil {
		log.Error().Err(err).Msg("tmpl execute")
		return nil, err
	}

	log.Debug().Str("dsn", connString.String()).Msg("connect to db")
	db, err := sql.Open("mysql", connString.String())
	if err != nil {
		log.Error().Err(err).Msg("mysql connect")
		return nil, err
	}

	return db, nil
}

func TestMain(m *testing.M) {
	setupEnv()
	initDB()

	exitVal := m.Run()

	os.Exit(exitVal)
}

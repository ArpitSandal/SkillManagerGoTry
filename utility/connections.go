package utility

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

// initialize the pgx db connection

type dns struct {
	username    string
	password    string
	host        string
	port        string
	dbname      string
	search_path string
	sslmode     string
	//connStr := "postgres://username:password@localhost:5432/mydatabase?search_path=myschema"
}

// func (dnsConfig dns) getDns() string {

// 	return "host=" + dnsConfig.host + "username=" + dnsConfig.username +
// 		"password=" + dnsConfig.password + "port=" + dnsConfig.port + "dbname" + dnsConfig.dbname +
// 		"search_path=" + dnsConfig.search_path + "sslmod=" + dnsConfig.sslmode

// }

func InitializeDB(ctx context.Context) *pgxpool.Pool {

	dnsConfig := dns{
		username:    "admin",
		password:    "admin",
		host:        "localhost",
		port:        "5432",
		dbname:      "postgres",
		search_path: "skillz_db",
		sslmode:     "false",
	}
	conn, err := pgxpool.New(ctx, "postgresql://admin:admin@localhost:5432/postgres?search_path=skillz_db")

	if err != nil {
		slog.Error("Error connecting to the database", "err", err)
		panic("Aborted")
	}

	slog.Info("Successfully connected to db with port ", "port", dnsConfig.port)

	return conn

}

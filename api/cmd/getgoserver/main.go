package main

import (
	"context"
	"database/sql"
	"log"

	"github.com/Quoc-Bao-213/getgo-project/api/internal/appconfig/db/pg"
	"github.com/Quoc-Bao-213/getgo-project/api/internal/controller/products"
	"github.com/Quoc-Bao-213/getgo-project/api/internal/httpserver"
	"github.com/Quoc-Bao-213/getgo-project/api/internal/repository"
	"github.com/Quoc-Bao-213/getgo-project/api/internal/repository/generator"
)

func main() {
	ctx := context.Background()

	conn, err := pg.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	generator.InitSnowflakeGenerators()

	rtr, err := initRouter(ctx, conn)
	if err != nil {
		log.Fatal(err)
	}

	httpserver.Start(httpserver.Handler(ctx, rtr.routes))
}

func initRouter(_ context.Context, db *sql.DB) (router, error) {
	repo := repository.New(db)

	productCtrl := products.New(repo)

	return router{
		productCtrl: productCtrl,
	}, nil
}

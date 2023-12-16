package main

import (
	"os"

	"github.com/matheus-alpe/go-hex-structure-example/internal/adapters/app/api"
	"github.com/matheus-alpe/go-hex-structure-example/internal/adapters/core/arithmetic"
	rpc "github.com/matheus-alpe/go-hex-structure-example/internal/adapters/framework/left/grpc"
	"github.com/matheus-alpe/go-hex-structure-example/internal/adapters/framework/right/db"
	"github.com/matheus-alpe/go-hex-structure-example/internal/ports"
)

func main() {
    dbDriver := os.Getenv("DB_DRIVER")
    dbDSN := os.Getenv("DB_DSN")

    var dbAdapter ports.DbPort = db.NewAdapter(dbDriver, dbDSN)
    defer dbAdapter.CloseDbConnection()

    // core - shouldn't depend on external sources
    var core ports.ArithmeticPort = arithmetic.NewAdapter()

    // application - integrate sources with core functionalities through dependency injection 
    var appAdapter ports.APIPort = api.NewAdapter(dbAdapter, core)

    // framework - integrate with external resources
    var grpcAdapter ports.GRPCPort = rpc.NewAdapter(appAdapter)
    grpcAdapter.Run()
}

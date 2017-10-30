package main

import (
	"github.com/NeuronEvolution/Stock/api/gen/restapi"
	"github.com/NeuronEvolution/Stock/api/gen/restapi/operations"
	"github.com/NeuronEvolution/Stock/cmd/stock-api/handler"
	"github.com/NeuronFramework/log"
	"github.com/NeuronFramework/restful"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"
	"github.com/rs/cors"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"net/http"
)

func main() {
	log.Init(true)

	middleware.Debug = false

	logger := zap.L().Named("main")

	var bind_addr string
	var storageConnectionString = ""

	cmd := cobra.Command{}
	cmd.PersistentFlags().StringVar(&bind_addr, "bind-addr", ":8081", "api server bind addr")
	cmd.PersistentFlags().StringVar(&storageConnectionString, "mysql-connection-string", "root:123456@tcp(127.0.0.1:3307)/fin-stock?parseTime=true", "mysql connection string")
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
		if err != nil {
			return errors.WithStack(err)
		}
		api := operations.NewStockAPI(swaggerSpec)

		h, err := handler.NewStockHandler(&handler.StockHandlerOptions{StorageConnectionString: storageConnectionString})
		if err != nil {
			return err
		}

		api.StocksListHandler = operations.StocksListHandlerFunc(h.List)
		api.StocksGetHandler = operations.StocksGetHandlerFunc(h.Get)

		logger.Info("Start server", zap.String("addr", bind_addr))
		err = http.ListenAndServe(bind_addr,
			restful.Recovery(cors.Default().Handler(api.Serve(nil))))
		if err != nil {
			return errors.WithStack(err)
		}

		return nil
	}
	cmd.Execute()
}

package main

import (
	"github.com/NeuronEvolution/Stock/api/restapi"
	"github.com/NeuronEvolution/Stock/api/restapi/operations"
	"github.com/NeuronEvolution/Stock/cmd/stock-api/handler"
	"github.com/NeuronEvolution/log"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/marshome/i-pkg/httphelper"
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
	cmd.PersistentFlags().StringVar(&storageConnectionString, "mysql-connection-string", "root:123456@tcp(127.0.0.1:3307)/fin-stock", "mysql connection string")
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

		api.StocksListHandler = operations.StocksListHandlerFunc(func(params operations.StocksListParams) middleware.Responder {
			return h.List(params)
		})
		api.StocksGetHandler = operations.StocksGetHandlerFunc(func(params operations.StocksGetParams) middleware.Responder {
			return h.Get(params)
		})

		logger.Info("Start server", zap.String("addr", bind_addr))
		err = http.ListenAndServe(bind_addr,
			httphelper.Recovery(cors.Default().Handler(api.Serve(nil))))
		if err != nil {
			return errors.WithStack(err)
		}

		return nil
	}
	cmd.Execute()
}

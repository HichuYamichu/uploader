package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hichuyamichu-me/goober/app"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "starts goober's http server",
	Run: func(cmd *cobra.Command, args []string) {
		app := app.New()

		go func() {
			done := make(chan os.Signal, 1)
			signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
			<-done
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			app.Shutdown(ctx)
		}()

		host := viper.GetString("app_host")
		port := viper.GetString("app_port")
		log.Fatal(app.Start(host, port))
	},
}

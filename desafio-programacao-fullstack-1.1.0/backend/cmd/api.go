/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/giovanars/desafio-programacao-fullstack/app/api"
	"github.com/giovanars/desafio-programacao-fullstack/config"
	"github.com/spf13/cobra"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Start http listen to provider microservice api",
	Run: func(cmd *cobra.Command, args []string) {
		println(fmt.Printf("Load modules"))

		cfg := config.NewConfig()
		server := api.NewServer(cfg.HTTPAddress)
		//TODO:init database
		initRouters(cfg, server)

		println(fmt.Printf("Starting to load API"))

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		serverShutdown := server.WaitSignal(ctx)
		if err := server.Start(); err != http.ErrServerClosed {
			panic(err)
		}

		<-serverShutdown
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)
}

func initRouters(cfg *config.Config, server *api.Server) {
	_, err := server.Load(
		"",
		cfg,
		api.Routers(),
	)
	if err != nil {
		panic(err)
	}
}

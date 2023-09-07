package cmd

import (
	"github.com/iniudin/demo-ticket-booking/server"
	"github.com/iniudin/demo-ticket-booking/worker"
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use:   "booking",
	Short: "Run service",
}

var runServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the server",
	Run: func(cmd *cobra.Command, args []string) {
		s := server.NewServer()
		s.Run()
	},
}

var runWorkerCmd = &cobra.Command{
	Use:   "worker",
	Short: "Run the worker",
	Run: func(cmd *cobra.Command, args []string) {
		w := worker.NewWorker()
		w.Run()
	},
}

func Execute() {
	rootCmd.AddCommand(runServerCmd)
	rootCmd.AddCommand(runWorkerCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/relrod/adguard-exporter/config"
	"github.com/relrod/adguard-exporter/internal/adguard"
	"github.com/relrod/adguard-exporter/internal/metrics"
	"github.com/relrod/adguard-exporter/internal/server"
)

const (
	name = "adguard-exporter"
)

var (
	s *server.Server
)

func main() {
	conf := config.Load()

	metrics.Init()

	initAdguardClient(conf.AdguardProtocol, conf.AdguardHostname, conf.AdguardPort, conf.AdguardUsername, conf.AdguardPassword, conf.Interval, conf.LogLimit)
	initHttpServer(conf.ServerPort)

	handleExitSignal()
}

func initAdguardClient(protocol, hostname string, port string, username, password string, interval time.Duration, logLimit string) {
	client := adguard.NewClient(protocol, hostname, port, username, password, interval, logLimit)
	go client.Scrape()
}

func initHttpServer(port string) {
	s = server.NewServer(port)
	go s.ListenAndServe()
}

func handleExitSignal() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	s.Stop()
	fmt.Println(fmt.Sprintf("\n%s HTTP server stopped", name))
}

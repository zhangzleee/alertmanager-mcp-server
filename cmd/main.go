// Copyright 2025 The Go MCP SDK Authors. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/zhangzleee/alertmanager-mcp-server/server"
	"github.com/zhangzleee/alertmanager-mcp-server/tools"
)

var (
	host = flag.String("host", "localhost", "host to connect to/listen on")
	port = flag.Int("port", 8000, "port number to connect to/listen on")
)

func main() {
	flag.StringVar(&tools.AlertmanagerUrl, "alertmanager-url", "http://localhost:9093", "Alertmanager URL")
	out := flag.CommandLine.Output()
	flag.Usage = func() {
		fmt.Fprintf(out, "Usage: %s server [-host <host>] [-port <port>] [-alertmanager-url <url>]\n\n", os.Args[0])
		fmt.Fprintf(out, "This program demonstrates MCP over HTTP using the streamable transport.\n")
		fmt.Fprintf(out, "It can run as either a server.\n\n")
		fmt.Fprintf(out, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(out, "\nExamples:\n")
		fmt.Fprintf(out, "  Run as server use default args:  %s server\n", os.Args[0])
		fmt.Fprintf(out, "  Custom host/port: %s server -port 9000 -host 0.0.0.0 \n", os.Args[0])
		os.Exit(1)
	}
	flag.Parse()

	if flag.NArg() < 1 {
		flag.Usage()
	}

	mode := flag.Arg(0)
	switch mode {
	case "help":
		flag.Usage()
	case "server":
		runServer(fmt.Sprintf("%s:%d", *host, *port))
	}
}

func runServer(url string) {
	// Create an MCP server.
	srv := server.New()

	// Create the streamable HTTP handler.
	handler := mcp.NewStreamableHTTPHandler(func(req *http.Request) *mcp.Server {
		return srv
	}, nil)

	handlerWithLogging := server.LoggingHandler(handler)

	log.Printf("MCP server listening on %s", url)
	log.Printf("Available tool: GetAlerts")
	if err := tools.CheckAlertManagerStatus(); err != nil {
		log.Fatal(err)
	}
	log.Printf("Available Alertmanager Url: %s", tools.AlertmanagerUrl)

	// Start the HTTP server with logging handler.
	if err := http.ListenAndServe(url, handlerWithLogging); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

package server

import (
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/zhangzleee/alertmanager-mcp-server/tools"
)

func New() *mcp.Server {
	srv := mcp.NewServer(&mcp.Implementation{
		Name:    "alertmanager-mcp-server",
		Title:   "search and silence alert by alertmanager-mcp-server",
		Version: "v0.0.1",
	}, nil,
	)
	//tools.RegisterToolStatus(srv, amURL)
	RegisterToolAlerts(srv)
	//tools.RegisterToolSilences(srv, amURL)

	return srv
}

func RegisterToolAlerts(srv *mcp.Server) {
	// Add Get Alerts tool.
	mcp.AddTool(srv, &mcp.Tool{
		Name:        "get_alerts",
		Description: "get alerts list from alertmanager",
	}, tools.GetAlerts)

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "get_silences",
		Description: "get silences list from alertmanager",
	}, tools.GetSilences)

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "set_silences",
		Description: "set silence with labels to alertmanager",
	}, tools.SetSilence)
}

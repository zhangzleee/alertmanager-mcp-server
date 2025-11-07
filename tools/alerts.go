package tools

import (
	"context"
	"fmt"
	"net/http"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type GetAlertsParams struct {
	Active      string `json:"active,omitempty" jsonschema:"The status of alerts that are currently active"`
	Silenced    string `json:"silenced,omitempty" jsonschema:"The status of alerts that are currently silenced"`
	Inhibited   string `json:"inhibited,omitempty" jsonschema:"The status of alerts that are currently inhibited"`
	Unprocessed string `json:"unprocessed,omitempty" jsonschema:"The status of alerts that are unprocessed"`
}

func GetAlerts(ctx context.Context, call *mcp.CallToolRequest, arg GetAlertsParams) (*mcp.CallToolResult, any, error) {
	active := arg.Active
	if active == "" {
		active = "false"
	}
	silenced := arg.Silenced
	if silenced == "" {
		silenced = "false"
	}
	inhibited := arg.Inhibited
	if inhibited == "" {
		inhibited = "false"
	}
	unprocessed := arg.Unprocessed
	if unprocessed == "" {
		unprocessed = "false"
	}

	u := fmt.Sprintf("%s/api/v2/alerts", AlertmanagerUrl)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}
	q := req.URL.Query()
	q.Add("active", active)
	q.Add("silenced", silenced)
	q.Add("inhibited", inhibited)
	q.Add("unprocessed", unprocessed)
	return doRequest(req), nil, nil
}

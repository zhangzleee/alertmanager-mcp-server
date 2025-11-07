package tools

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type GetSilenceParams struct {
	Active string `json:"active,omitempty" jsonschema:"Active silence Example: true"`
}

func GetSilences(ctx context.Context, call *mcp.CallToolRequest, arg GetSilenceParams) (*mcp.CallToolResult, any, error) {
	active := arg.Active
	if active == "" {
		active = "false"
	}
	u := fmt.Sprintf("%s/api/v2/silences", AlertmanagerUrl)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}
	q := req.URL.Query()
	q.Add("active", active)
	return doRequest(req), nil, nil
}

type SetSilencesParams struct {
	Comment   string        `json:"comment" jsonschema:"The comment of the silenced alert"`
	CreatedBy string        `json:"createdBy" jsonschema:"who user to the silenced alert"`
	StartsAt  string        `json:"startsAt" jsonschema:"The start time of the silenced alert. ep: '2025-11-07T05:56:07.402Z'"`
	EndsAt    string        `json:"endsAt" jsonschema:"The end time of the silenced alert. ep: '2025-11-07T05:56:07.402Z'"`
	Matchers  []AlertLabels `json:"matchers" jsonschema:"The matchers of the silenced alert Each matcher requires 'name', 'value', 'isRegex', and 'isEqual' fields. Example: [{'name':'11111', 'value':'production', 'isRegex':false, 'isEqual':true}]"`
}

type AlertLabels struct {
	Name    string `json:"name" jsonschema:"The name of the silenced alert"`
	Value   string `json:"value" jsonschema:"The value of the silenced alert"`
	IsRegex bool   `json:"isRegex" jsonschema:"labels value Is regex or not"`
	IsEqual bool   `json:"isEqual" jsonschema:"labels value Is equal or not"`
}

func SetSilence(ctx context.Context, call *mcp.CallToolRequest, arg SetSilencesParams) (*mcp.CallToolResult, any, error) {
	u := fmt.Sprintf("%s/api/v2/silences", AlertmanagerUrl)
	jsonData, err := json.Marshal(arg)
	if err != nil {
		return nil, nil, fmt.Errorf("JSON marshal failed: %w", err)
	}
	reader := bytes.NewReader(jsonData)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u, reader)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return doRequest(req), nil, nil
}

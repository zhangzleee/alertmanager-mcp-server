package tools

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

var AlertmanagerUrl string

func doRequest(req *http.Request) *mcp.CallToolResult {
	resp, err := http.DefaultClient.Do(req)
	result := &mcp.CallToolResult{}
	if err != nil {
		result.IsError = true
		result.Content = []mcp.Content{
			&mcp.TextContent{Text: fmt.Sprintf("alertmanager retrun error: %s", err.Error())},
		}
		return result
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		result.IsError = true
		result.Content = []mcp.Content{
			&mcp.TextContent{Text: fmt.Sprintf("alertmanager retrun error: %s", err.Error())},
		}
		return result
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode != http.StatusOK {
		result.IsError = true
		result.Content = []mcp.Content{
			&mcp.TextContent{Text: fmt.Sprintf("alertmanager retrun error: %s", body)},
		}
		return result
	}
	s := string(body)
	result.Content = []mcp.Content{
		&mcp.TextContent{Text: fmt.Sprintf("alertmanager data: %s", s)},
	}
	return result
}

func CheckAlertManagerStatus() error {
	url := fmt.Sprintf("%s/-/healthy", AlertmanagerUrl)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}
	return nil
}

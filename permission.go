package civogo

import (
	"bytes"
	"encoding/json"
)

// Permission represents a permission and the description for it
type Permission struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

func (c *Client) ListPermissions() ([]Permission, error) {
	resp, err := c.SendGetRequest("/v2/permissions")
	if err != nil {
		return nil, decodeERROR(err)
	}

	permissions := make([]Permission, 0)
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&permissions); err != nil {
		return nil, err
	}

	return permissions, nil
}

package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas() (*LocationAreasResp, error) {
	var locationArea LocationAreasResp
	endPoint := baseURL + "location-area/"

	req, err := http.NewRequest(http.MethodGet, endPoint, nil)
	if err != nil {
		return &LocationAreasResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return &LocationAreasResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return &LocationAreasResp{}, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return &LocationAreasResp{}, err
	}

	err = json.Unmarshal(data, &locationArea)
	if err != nil {
		return &LocationAreasResp{}, err
	}

	return &locationArea, nil
}

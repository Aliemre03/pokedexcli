package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageUrl *string) (*LocationAreasResp, error) {
	var locationArea LocationAreasResp
	endPoint := baseURL + "location-area/"

	if pageUrl != nil {
		endPoint = *pageUrl
	}

	data, ok := c.cache.Get(endPoint)
	if ok {
		err := json.Unmarshal(data, &locationArea)
		if err != nil {
			return &LocationAreasResp{}, err
		}

		return &locationArea, nil
	}

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

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return &LocationAreasResp{}, err
	}

	err = json.Unmarshal(data, &locationArea)
	if err != nil {
		return &LocationAreasResp{}, err
	}

	c.cache.Add(endPoint, data)

	return &locationArea, nil
}

package apileek

import (
    "encoding/json"
)

type leekwarsService struct {
    apiService
}

type version struct {
    Version uint `json:"version"`
}

// Get version of leekwars
func (s *leekwarsService) Version(
) (uint, error) {
    data := ""
    resp, body, err := s.apiRequest("POST", s.url + "version/", &data)
    if err != nil {
        return 0, err
    }

    if resp.StatusCode != 200 {
        return 0, newApiError(resp, body)
    }

    var obj = version{}
    err = json.Unmarshal([]byte(body), &obj)
    if err != nil {
        return 0, err
    }

    return obj.Version, nil
}


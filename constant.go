package apileek

import (
    "encoding/json"
)

type constantService struct {
    apiService
}

type Constant struct {
    Name string `json:"name"`
    Value string `json:"value"`
    Type uint `json:"type"`
    Category uint `json:"category"`
    Deprecated uint `json:"deprecated"`
}

type constList struct {
    Constants []Constant `json:"constants"`
}

// Get all constant definition
func (s *constantService) GetAll(
) ([]Constant, error) {
    data := ""
    resp, body, err := s.apiRequest("POST", s.url + "get-all/", &data)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode != 200 {
        return nil, newApiError(resp, body)
    }

    var obj = constList{}
    err = json.Unmarshal([]byte(body), &obj)
    if err != nil {
        return nil, err
    }

    return obj.Constants, nil
}


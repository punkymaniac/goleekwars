package apileek

import (
    "encoding/json"
)

type countryService struct {
    apiService
}

type Country struct {
    Code string `json:"code"`
}

type countryList struct {
    Countries []Country `json:"countries"`
}

// Get all country code to deal with
func (s *countryService) GetAll(
) ([]Country, error) {
    data := ""
    resp, body, err := s.apiRequest("POST", s.url + "get-all/", &data)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode != 200 {
        return nil, newApiError(resp, body)
    }

    var obj = countryList{}
    err = json.Unmarshal([]byte(body), &obj)
    if err != nil {
        return nil, err
    }

    return obj.Countries, nil
}


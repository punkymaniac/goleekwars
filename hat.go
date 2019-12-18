package apileek

import (
    "encoding/json"
)

type hatService struct {
    apiService
}


type Hat struct {
    Id uint `json:"id"`
    Name string `json:"name"`
    Level uint `json:"level"`
    Width float32 `json:"width"`
    Height float32 `json:"height"`
}

type HatTemplate struct {
    Id uint `json:"id"`
    Item uint `json:"item"`
}

type hatList struct {
    Hats map[string]Hat `json:"hats"`
}

type hatTemplateList struct {
    HatTemplate map[string]HatTemplate `json:"hat_templates"`
}

// Get all hat
func (s *hatService) GetAll(
) (map[string]Hat, error) {
    data := ""
    resp, body, err := s.apiRequest("POST", s.url + "get-all/", &data)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode != 200 {
        return nil, newApiError(resp, body)
    }

    var obj = hatList{}
    err = json.Unmarshal([]byte(body), &obj)
    if err != nil {
        return nil, err
    }

    return obj.Hats, nil
}

// Get hat template
func (s *hatService) GetTemplates(
) (map[string]HatTemplate, error) {
    data := ""
    resp, body, err := s.apiRequest("POST", s.url + "get-templates/", &data)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode != 200 {
        return nil, newApiError(resp, body)
    }

    var obj = hatTemplateList{}
    err = json.Unmarshal([]byte(body), &obj)
    if err != nil {
        return nil, err
    }

    return obj.HatTemplate, nil
}


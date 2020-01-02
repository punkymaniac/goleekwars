package apileek

import (
    "encoding/json"
)

type chipService struct {
    apiService
}

type Chip struct {
    Id uint `json:"id"`
    Name string `json:"name"`
    Level uint `json:"level"`
    MinRange uint `json:"min_range"`
    MaxRange uint `json:"max_range"`
    LaunchType uint `json:"launch_type"`
    Effects []effect `json:"effects"`
    Cost uint `json:"cost"`
    Area uint `json:"area"`
    Cooldown int `json:"cooldown"`
    Los uint `json:"los"`
    TeamCooldown int `json:"team_cooldown"`
    InitialCooldown int `json:"initial_cooldown"`
    Template uint `json:"template"`
}

type ChipTemplate struct {
    Id uint `json:"id"`
    Item uint `json:"item"`
}

type chipRef struct {
    Template uint `json:"template"`
    Id uint `json:"id"`
}

type effect struct {
    Id uint `json:"id"`
    Value1 float32 `json:"value1"`
    Value2 float32 `json:"value2"`
    Turns uint `json:"turns"`
    Targets uint `json:"targets"`
    Type uint `json:"type"`
}

type chipList struct {
    Chips map[string]Chip `json:"chips"`
}

type chipTemplateList struct {
    ChipTemplate map[string]ChipTemplate `json:"chip_templates"`
}

// Get all chip
func (s *chipService) GetAll(
) (map[string]Chip, error) {
    data := ""
    resp, body, err := s.apiRequest("POST", s.url + "get-all/", &data)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode != 200 {
        return nil, newApiError(resp, body)
    }

    var obj = chipList{}
    err = json.Unmarshal([]byte(body), &obj)
    if err != nil {
        return nil, err
    }

    return obj.Chips, nil
}

// Get chip template
func (s *chipService) GetTemplates(
) (map[string]ChipTemplate, error) {
    data := ""
    resp, body, err := s.apiRequest("POST", s.url + "get-templates/", &data)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode != 200 {
        return nil, newApiError(resp, body)
    }

    var obj = chipTemplateList{}
    err = json.Unmarshal([]byte(body), &obj)
    if err != nil {
        return nil, err
    }

    return obj.ChipTemplate, nil
}


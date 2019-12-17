package apileek

import (
    "encoding/json"
)

type weaponService struct {
    apiService
}

type Weapon struct {
    Id uint `json:"id"`
    Name string `json:"name"`
    Level uint `json:"level"`
    MinRange uint `json:"min_range"`
    MaxRange uint `json:"max_range"`
    LaunchType uint `json:"launch_type"`
    Effects []effect `json:"effects"`
    Cost uint `json:"cost"`
    Area uint `json:"area"`
    Los uint `json:"los"`
    Template uint `json:"template"`
}

type WeaponTemplate struct {
    Id uint `json:"id"`
    Item uint `json:"item"`
}

type weaponList struct {
    Weapons map[string]Weapon `json:"weapons"`
}

type weaponTemplateList struct {
    WeaponTemplate map[string]WeaponTemplate `json:"weapon_templates"`
}

// Get all weapon
func (s *weaponService) GetAll(
) (map[string]Weapon, error) {
    data := ""
    resp, body, err := s.apiRequest("POST", s.url + "get-all/", &data)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode != 200 {
        return nil, newApiError(resp, body)
    }

    var obj = weaponList{}
    err = json.Unmarshal([]byte(body), &obj)
    if err != nil {
        return nil, err
    }

    return obj.Weapons, nil
}

// Get weapon template
func (s *weaponService) GetTemplates(
) (map[string]WeaponTemplate, error) {
    data := ""
    resp, body, err := s.apiRequest("POST", s.url + "get-templates/", &data)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode != 200 {
        return nil, newApiError(resp, body)
    }

    var obj = weaponTemplateList{}
    err = json.Unmarshal([]byte(body), &obj)
    if err != nil {
        return nil, err
    }

    return obj.WeaponTemplate, nil
}


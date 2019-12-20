package apileek

import (
    "strconv"
    "encoding/json"
)

type historyService struct {
    apiService
}

type FarmerHistory struct {
    Fights []farmerFight `json:"fights"`
    Entity farmerEntity `json:"entity"`
}

type LeekHistory struct {
    Fights []leekFight `json:"fights"`
}

type farmerFight struct {
    Id uint `json:"id"`
    Leeks1 []leek5 `json:"leeks1"`
    Leeks2 []leek5 `json:"leeks2"`
    Winner uint `json:"winner"`
    Status uint `json:"status"`
    Date timestamp `json:"date"`
    Context uint `json:"context"`
    Type uint `json:"type"`
    Farmer1 int `json:"farmer1"`
    Farmer2 int `json:"farmer2"`
    Team1 int `json:"team1"`
    Team2 int `json:"team2"`
    Result string `json:"result"`
    Farmer1Name string `json:"farmer1_name"`
    Farmer2Name string `json:"farmer2_name"`
}

type leekFight struct {
    Id uint `json:"id"`
    Leeks1 []leek5 `json:"leeks1"`
    Leeks2 []leek5 `json:"leeks2"`
    Winner uint `json:"winner"`
    Status uint `json:"status"`
    Date timestamp `json:"date"`
    Context uint `json:"context"`
    Type uint `json:"type"`
    Farmer1 int `json:"farmer1"`
    Farmer2 int `json:"farmer2"`
    Team1 int `json:"team1"`
    Team2 int `json:"team2"`
    Result string `json:"result"`
    Team1Name string `json:"team1_name"`
    Team2Name string `json:"team2_name"`
}

type farmerEntity struct {
    Id uint `json:"id"`
    Name string `json:"name"`
    AvatarChanged timestamp `json:"avatar_changed"`
    Talent uint `json:"talent"`
    TalentMore int `json:"talent_more"`
}

type leek5 struct {
    Id uint `json:"id"`
    Name string `json:"name"`
}

// Get farmer history
func (s *historyService) GetFarmerHistory(
    farmerId uint, // Farmer id
) (*FarmerHistory, error) {
    data := "farmer_id=" + strconv.FormatUint(uint64(farmerId), 10)
    resp, body, err := s.apiRequest("POST", s.url + "get-farmer-history/", &data)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode != 200 {
        return nil, newApiError(resp, body)
    }

    var obj = FarmerHistory{}
    err = json.Unmarshal([]byte(body), &obj)
    if err != nil {
        return nil, err
    }

    return &obj, nil
}

// Get leek history
func (s *historyService) GetLeekHistory(
    leekId uint, // Leek id
) (*LeekHistory, error) {
    data := "leek_id=" + strconv.FormatUint(uint64(leekId), 10)
    resp, body, err := s.apiRequest("POST", s.url + "get-leek-history/", &data)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode != 200 {
        return nil, newApiError(resp, body)
    }

    var obj = LeekHistory{}
    err = json.Unmarshal([]byte(body), &obj)
    if err != nil {
        return nil, err
    }

    return &obj, nil
}


package apileek

import (
    "strconv"
    "encoding/json"
)

type fightService struct {
    apiService
}

type Fight struct {
    Id uint `json:"id"`
    Date timestamp `json:"date"`
    Year uint `json:"year"`
    Type uint `json:"type"`
    Context uint `json:"context"`
    Status uint `json:"status"`
    Winner uint `json:"winner"`
    Leeks1 []leek1 `json:"leeks1"`
    Leeks2 []leek1 `json:"leeks2"`
    Farmers1 map[string]farmer1 `json:"farmers1"`
    Farmers2 map[string]farmer1 `json:"farmers2"`
    Data fightData `json:"data"`
    Dead map[string]uint `json:"dead"`
    Leeks []leek2 `json:"leeks"`
    Map fightMap `json:"map"`
    Report struct {
        Bonus uint `json:"bonus"`
        Duration uint `json:"duration"`
        Leeks1 []leek3 `json:"leeks1"`
        Leeks2 []leek3 `json:"leeks2"`
        Win uint `json:"win"`
    } `json:"report"`
    Comments []string `json:"comments"`
    Tournament int `json:"tournament"`
    Views uint `json:"views"`
    Starter uint `json:"starter"`
    Team1Name string `json:"team1_name"`
    Team2Name string `json:"team2_name"`
}

type farmer1 struct {
    Id uint `json:"id"`
    Name string `json:"name"`
    AvatarChanged timestamp `json:"avatar_changed"`
}

type leek1 struct {
    Id uint `json:"id"`
    Name string `json:"name"`
    Level uint `json:"level"`
    Talent uint `json:"talent"`
    Hat *uint `json:"hat"`
    Skin uint `json:"skin"`
}

type leek2 struct {
    Id uint `json:"id"`
    Level uint `json:"level"`
    Name string `json:"name"`
    Farmer uint `json:"farmer"`
    CellPos uint `json:"cellPos"`
    Agility uint `json:"agility"`
    Frequency uint `json:"frequency"`
    Life uint `json:"life"`
    Magic uint `json:"magic"`
    Mp uint `json:"mp"`
    Tp uint `json:"tp"`
    Resistance uint `json:"resistance"`
    Science uint `json:"science"`
    Strength uint `json:"strength"`
    Wisdom uint `json:"wisdom"`
    Skin uint `json:"skin"`
    Summon bool `json:"summon"`
    Team uint `json:"team"`
    Type uint `json:"type"`
}

type leek3 struct {
    Apperance uint `json:"apperance"`
    CurXp uint `json:"cur_xp"`
    Dead bool `json:"dead"`
    Id uint `json:"id"`
    Level uint `json:"level"`
    Money uint `json:"money"`
    Name string `json:"name"`
    NextXp uint `json:"next_xp"`
    PrevXp uint `json:"prev_xp"`
    Talent uint `json:"talent`
    TalentGain int `json:"talent_gain"`
    Tb uint `json:"tb"`
    Td int `json:"td"`
    Xp uint `json:"xp"`
}

type fightData struct {
    Actions [][]interface{} `json:"actions"`
}

type fightMap struct {
    Height uint `json:"height"`
    Width uint `json:"width"`
    Obstacles map[string][]uint `json:"obstacles"`
    Type uint `json:"type"`
}

type fightRoot struct {
    Fight Fight `json:"fight"`
}

type FightLog struct {
    Logs map[string][][]interface{} `json:"logs"`
    Habs uint `json:"habs"`
    Talent uint `json:"talent"`
    LeekTalents map[string]uint `json:"leek_talents"`
}

// Comment a fight
func (s *fightService) Comment(
    fightId uint, // The fight id
    comment string, // The comment to made on the fight
) error {
    data := "fight_id=" + strconv.FormatUint(uint64(fightId), 10) + "&comment=" + comment
    resp, body, err := s.apiRequest("POST", s.url + "comment/", &data)
    if err != nil {
        return err
    }

    if resp.StatusCode != 200 {
        return newApiError(resp, body)
    }
    return nil
}

// Get fight data
func (s *fightService) Get(
    fightId uint, // The fight id
) (*Fight, error) {
    data := "fight_id=" + strconv.FormatUint(uint64(fightId), 10)
    resp, body, err := s.apiRequest("POST", s.url + "get/", &data)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode != 200 {
        return nil, newApiError(resp, body)
    }

    var obj = fightRoot{}
    err = json.Unmarshal([]byte(body), &obj)
    if err != nil {
        return nil, err
    }
    return &obj.Fight, nil
}

// Get logs of a fight
func (s *fightService) GetLogs(
    fightId uint, // The fight id
) (*FightLog, error) {
    data := "fight_id=" + strconv.FormatUint(uint64(fightId), 10)
    resp, body, err := s.apiRequest("POST", s.url + "get-logs/", &data)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode != 200 {
        return nil, newApiError(resp, body)
    }

    var obj = FightLog{}
    err = json.Unmarshal([]byte(body), &obj)
    if err != nil {
        return nil, err
    }
    return &obj, nil
}


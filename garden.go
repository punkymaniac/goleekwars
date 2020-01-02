package apileek

import (
    "encoding/json"
    "strconv"
)

type gardenService struct {
    apiService
}

type gardenData struct {
    Garden Garden `json:"garden"`
}

type Garden struct {
    Fights uint `json:"fights"`
    MaxFights uint `json:"max_fights"`
    FarmerEnabled bool `json:"farmer_enabled"`
    TeamEnabled bool `json:"team_enabled"`
    MyCompositions []teamComposition `json:"my_compositions"`
    MyTeam teamInfo `json:"my_team"`
    SoloFights json.RawMessage `json:"solo_fights"` // missing information
    TotalSoloFights uint `json:"total_solo_fights"`
    MaxSoloFights uint `json:"max_solo_fights"`
    TeamFights uint `json:"team_fights"`
    MaxTeamFights uint `json:"max_team_fights"`
    BattleRoyalFights uint `json:"battle_royal_fights"`
    MaxBattleRoyalFights uint `json:""max_battle_royal_fights`
    BattleRoyalEnabled bool `json:"battle_royal_enabled"`
}

type fightId struct {
    Fight uint `json:"fight"`
}

type compositionOpponentList struct {
    Opponents []CompositionOpponent `json:"opponents"`
}

type CompositionOpponent struct {
    Id uint `json;"id"`
    Team uint `json:"team"`
    Talent uint `json:"talent"`
    Name string `json:"name"`
    Level uint `json:"level"`
    EmblemChanged timestamp `json:"emblem_changed"`
    TeamId uint `json:"team_id"`
    TotalLevel uint `json:"total_level"`
    TotalPower uint `json:"total_power"`
}

type farmerOpponentList struct {
    Opponents []FarmerOpponent `json:"opponents"`
}

type FarmerOpponent struct {
    Id uint `json:"id"`
    Name string `json:"name"`
    AvatarChanged timestamp `json:"avatar_changed"`
    Talent uint `json:"talent"`
    TotalLevel uint `json:"total_level"`
    LeekCount uint `json:"leek_count"`
}

type leekOpponentList struct {
    Opponents []LeekOpponent `json:"opponents"`
}

type LeekOpponent struct {
    Id uint `json:"id"`
    Name string `json:"name"`
    Level uint `json:"level"`
    Talent uint `json:"talent"`
    Skin uint `json:"skin"`
    Hat uint `json:"hat"`
    Fights uint `json:"fights"`
}

type FarmerChallenge struct {
    Farmer farmerInfo02 `json:"farmer"`
    Challenges uint `json:"challenges"`
}

type farmerInfo02 struct {
    Id uint `json:"id"`
    Name string `json:"name"`
    AvatarChanged timestamp `json:"avatar_changed"`
    TotalLevel uint `json:"total_level"`
    LeekCount uint `json:"leek_count"`
    Talent uint `json:"talent"`
}

type SoloChallenge struct {
    Leek leekInfo `json:"leek"`
    Challenges uint `json:"challenges"`
}

type leekInfo struct {
    Id uint `json:"id"`
    Name string `json:"name"`
    Farmer farmerInfo01 `json:"farmer"`
    Talent uint `json:"talent"`
    TalentMore int `json:"talent_more"`
    Level uint `json:"level"`
    Xp uint `json:"xp"`
    UpXp uint `json:"up_xp"`
    DownXp uint `json:"down_xp"`
    RemainingXp uint `json:"remaining_xp"`
    Life uint `json:"life"`
    Strength uint `json:"strength"`
    Wisdom uint `json:"wisdom"`
    Agility uint `json:"agility"`
    Resistance uint `json:"resistance"`
    Science uint `json:"science"`
    Magic uint `json:"magic"`
    Tp uint `json:"tp"`
    Mp uint `json:"mp"`
    Frequency uint `json:"frequency"`
    Victories uint `json:"victories"`
    Draws uint `json:"draws"`
    Defeats uint `json:"defeats"`
    Ratio string `json:"ratio"`
    Chips []chipRef `json:"chips"`
    Weapons []weaponRef `json:"weapons"`
    Ai aiBase `json:"ai"`
    MaxWeapons uint `json:"max_weapons"`
    MaxChips uint `json:"max_chips"`
    InGarden bool `json:"in_garden"`
    Fights []fightLog `json:"fights"`
    Skin uint `json:"skin"`
    Hat uint `json:"hat"`
    TalentHistory []uint `json:"talent"`
    Tournaments []tournamentsRef `json:"tournaments"`
}

type tournamentsRef struct {
    Id uint `json:"id"`
    Date timestamp `json:"date"`
}

type fightLog struct {
    Id uint `json:"id"`
    Date timestamp `json:"date"`
    Year uint `json:"year"`
    Type uint `json:"type"`
    Context uint `json:"context"`
    Status uint `json:"status"`
    Winner uint `json:"winner"`
    LeekTeam uint `json:"leek_team"`
    Result string `json:"result"`
    Leeks1 []leek4 `json:"leeks1"`
    Leeks2 []leek4 `json:"leeks2"`
    Team1 uint `json:"team1"`
    Team2 uint `json:"team2"`
    Team1Name string `json:"team1_name"`
    Team2Name string `json:"team2_name"`
}

type leek4 struct {
    Id string `json:"id"`
    Name string `json:"name"`
}

type farmerInfo01 struct {
    Id uint `json:"id"`
    Name string `json:"name"`
}

type teamInfo struct {
    Id uint `json:"id"`
    Name string `json:"name"`
    Level uint `json:""level`
    EmblemChanged timestamp `json:"emblem_changed"`
}

type teamComposition struct {
    Id uint `json:"id"`
    TeamId uint `json:"team_id"`
    Name string `json:"name"`
    Level uint `json:"level"`
    TotalLevel uint `json:"total_level"`
    Talent uint `json:"talent"`
    Fights uint `json:"fights"`
    EmblemChanged timestamp `json:"emblem_changed"`
}

// Get garden information
func (s *gardenService) Get(
) (*Garden, error) {
    data := ""
    resp, body, err := s.apiRequest("POST", s.url + "get/", &data)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode != 200 {
        return nil, newApiError(resp, body)
    }

    var obj = gardenData{}
    err = json.Unmarshal([]byte(body), &obj)
    if err != nil {
        return nil, err
    }

    return &obj.Garden, nil
}

// Get the composition opponents
func (s *gardenService) GetCompositionOpponents(
    composition uint, // The opponents composition number
) ([]CompositionOpponent, error) {
    data := "composition=" + strconv.FormatUint(uint64(composition), 10)
    resp, body, err := s.apiRequest("POST", s.url + "get-composition-opponents/", &data)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode != 200 {
        return nil, newApiError(resp, body)
    }

    var obj = compositionOpponentList{}
    err = json.Unmarshal([]byte(body), &obj)
    if err != nil {
        return nil, err
    }

    return obj.Opponents, nil
}

// Get farmer challenge details
func (s *gardenService) GetFarmerChallenge(
    target uint, // The farmer id
) (*FarmerChallenge, error) {
    data := "target=" + strconv.FormatUint(uint64(target), 10)
    resp, body, err := s.apiRequest("POST", s.url + "get-farmer-challenge/", &data)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode != 200 {
        return nil, newApiError(resp, body)
    }

    var obj = FarmerChallenge{}
    err = json.Unmarshal([]byte(body), &obj)
    if err != nil {
        return nil, err
    }

    return &obj, nil
}

// Get the farmer opponents
func (s *gardenService) GetFarmerOpponents(
) ([]FarmerOpponent, error) {
    data := ""
    resp, body, err := s.apiRequest("POST", s.url + "get-farmer-opponents/", &data)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode != 200 {
        return nil, newApiError(resp, body)
    }

    var obj = farmerOpponentList{}
    err = json.Unmarshal([]byte(body), &obj)
    if err != nil {
        return nil, err
    }

    return obj.Opponents, nil
}

// Get the leek opponents
func (s *gardenService) GetLeekOpponents(
    leekId uint, // The leek id
) ([]LeekOpponent, error) {
    data := "leek_id=" + strconv.FormatUint(uint64(leekId), 10)
    resp, body, err := s.apiRequest("POST", s.url + "get-leek-opponents/", &data)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode != 200 {
        return nil, newApiError(resp, body)
    }

    var obj = leekOpponentList{}
    err = json.Unmarshal([]byte(body), &obj)
    if err != nil {
        return nil, err
    }

    return obj.Opponents, nil
}

// Get leek challenge details
func (s *gardenService) GetSoloChallenge(
    leekId uint, // The leek id
) (*SoloChallenge, error) {
    data := "leek_id=" + strconv.FormatUint(uint64(leekId), 10)
    resp, body, err := s.apiRequest("POST", s.url + "get-solo-challenge/", &data)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode != 200 {
        return nil, newApiError(resp, body)
    }

    var obj = SoloChallenge{}
    err = json.Unmarshal([]byte(body), &obj)
    if err != nil {
        return nil, err
    }

    return &obj, nil
}

// Start a farmer challenge
func (s *gardenService) StartFarmerChallenge(
    targetId uint, // The farmer id to challenge
) (uint, error) {
    data := "target_id=" + strconv.FormatUint(uint64(targetId), 10)
    resp, body, err := s.apiRequest("POST", s.url + "start-farmer-challenge/", &data)
    if err != nil {
        return 0, err
    }

    if resp.StatusCode != 200 {
        return 0, newApiError(resp, body)
    }

    var obj = fightId{}
    err = json.Unmarshal([]byte(body), &obj)
    if err != nil {
        return 0, err
    }

    return obj.Fight, nil
}

// Start a farmer fight
func (s *gardenService) StartFarmerFight(
    targetId uint, // The farmer id to fight
) (uint, error) {
    data := "target_id=" + strconv.FormatUint(uint64(targetId), 10)
    resp, body, err := s.apiRequest("POST", s.url + "start-farmer-fight/", &data)
    if err != nil {
        return 0, err
    }

    if resp.StatusCode != 200 {
        return 0, newApiError(resp, body)
    }

    var obj = fightId{}
    err = json.Unmarshal([]byte(body), &obj)
    if err != nil {
        return 0, err
    }

    return obj.Fight, nil
}

// Start a solo challenge
func (s *gardenService) StartSoloChallenge(
    leekId uint, // The leek id with which challenge
    targetId uint, // The leek id to challenge
) (uint, error) {
    data := "leek_id=" + strconv.FormatUint(uint64(leekId), 10) +
            "&target_id=" + strconv.FormatUint(uint64(targetId), 10)
    resp, body, err := s.apiRequest("POST", s.url + "start-solo-challenge/", &data)
    if err != nil {
        return 0, err
    }

    if resp.StatusCode != 200 {
        return 0, newApiError(resp, body)
    }

    var obj = fightId{}
    err = json.Unmarshal([]byte(body), &obj)
    if err != nil {
        return 0, err
    }

    return obj.Fight, nil
}

// Start a solo fight
func (s *gardenService) StartSoloFight(
    leekId uint, // The leek id with which fight
    targetId uint, // The leek id to fight
) (uint, error) {
    data := "leek_id=" + strconv.FormatUint(uint64(leekId), 10) +
            "&target_id=" + strconv.FormatUint(uint64(targetId), 10)
    resp, body, err := s.apiRequest("POST", s.url + "start-solo-fight/", &data)
    if err != nil {
        return 0, err
    }

    if resp.StatusCode != 200 {
        return 0, newApiError(resp, body)
    }

    var obj = fightId{}
    err = json.Unmarshal([]byte(body), &obj)
    if err != nil {
        return 0, err
    }

    return obj.Fight, nil
}

// Start a team fight
func (s *gardenService) StartTeamFight(
    compositionId uint, // The composition id with which fight
    targetId uint, // The composition id to fight
) (uint, error) {
    data := "composition_id=" + strconv.FormatUint(uint64(compositionId), 10) +
            "&target_id=" + strconv.FormatUint(uint64(targetId), 10)
    resp, body, err := s.apiRequest("POST", s.url + "start-team-fight/", &data)
    if err != nil {
        return 0, err
    }

    if resp.StatusCode != 200 {
        return 0, newApiError(resp, body)
    }

    var obj = fightId{}
    err = json.Unmarshal([]byte(body), &obj)
    if err != nil {
        return 0, err
    }

    return obj.Fight, nil
}


package apileek

import (
    "encoding/json"
)

type FarmerService struct {
    apiService
}

type timestamp uint64

type LoginFarmer struct {
    farmer
    Token string `json:"token"`
}

type farmer struct {
    Farmer farmerInfo `json:"farmer"`
}

type farmerInfo struct {
    Id uint64 `json:"id"`
    Login string `json:"login"`
    Team team `json:"team"`
    Name string `json:"name"`
    Talent uint64 `json:"talent"`
    Leeks map[uint64]leek `json:"leeks"`
    AvatarChanged timestamp `json:"avatar_changed"`
    TalentMore uint64 `json:"talent_more"`
    Victories uint64 `json:"victories"`
    Draws uint64 `json:"draws"`
    Defeats uint64 `json:"defeats"`
    Ratio string `json:"ratio"`
    Connected bool `json:"connected"`
    LastConnection timestamp `json:"last_connection"`
    RegisterDate timestamp `json:"register_date"`
    FightHistory json.RawMessage `json:"fight_history"` // missing information
    Tournaments json.RawMessage `json:"tournaments"` // missing information
    Admin bool `json:"admin"`
    Moderator bool `json:"moderator"`
    Country json.RawMessage `json:"country"` //missing information
    Godfather json.RawMessage `json:"godfather"` // missing information
    Godsons json.RawMessage `json:"godsons"` // missing information
    Color string `json:"color"`
    Banned uint64 `json:"banned"`
    WonSoloTournaments uint64 `json:"won_solo_tournaments"`
    WonFarmerTournaments uint64 `json:"won_farmer_tournaments"`
    WonTeamTournaments uint64 `json:"won_team_tournaments"`
    TotalLevel uint64 `json:"total_level"`
    LeekCount uint64 `json:"leek_count"`
    InGarden uint64 `json:"in_garden"`
    Fights uint64 `json:"fights"`
    Github json.RawMessage `json:"github"` // missing information
    Website json.RawMessage `json:"website"` // missing information
    ForumMessages uint64 `json:"forum_messages"`
    DidacticielSeen uint64 `json:"didacticiel_seen"`
    Contributor bool `json:"contributor"`
    Trophies uint64 `json:"trophies"`
    Habs uint64 `json:"habs"`
    Crystals uint64 `json:"crystals"`
    Weapons []weapon `json:"weapons"`
    Chips []chip `json:"chips"`
    Ais []aiBase `json:"ais"`
    Potions []potion `json:"potions"`
    Hats []hat `json:"hats"`
    Tournament tournament `json:"tournament"`
    Candidacy json.RawMessage `json:"candidacy"` // missing information
}

type weapon struct {
    Id uint64 `json:"id"`
    Template uint64 `json:"template"`
}

type chip struct {
    Id uint64 `json:"id"`
    Template uint64 `json:"template"`
}

type potion struct {
    Id uint64 `json:"id"`
    Template uint64 `json:"template"`
    Quantity uint64 `json:"quantity"`
}

type tournament struct {
    Registered bool `json:"registered"`
    Current json.RawMessage `json:"current"` // missing information
}

type team struct {
    Id uint64 `json:"id"`
    Name string `json:"name"`
    Level uint64 `json:"level"`
    EmblemChanged timestamp `json:"emblem_changed"`
}

type leek struct {
    Id uint64 `json:"id"`
    Name string `json:"name"`
    Color string `json:"color"`
    Capital uint64 `json:"capital"`
    Level uint64 `json:"level"`
    Talent uint64 `json:"talent"`
    Skin uint64 `json:"skin"`
    Hat *uint64 `json:"Hat"`
}

type hat struct {
    Id uint64 `json:"id"`
    Template uint64 `json:"template"`
    Level uint64 `json:"level"`
    Name string `json:"name"`
    HatTemplate uint64 `json:"hat_template"`
}

// Get a auth token from the api
func (s *FarmerService) LoginToken(
    username string, // Username of the account
    password string, // Password of the account
) (*LoginFarmer, error) {
    data := "login=" + username + "&password=" + password
    resp, body, err := s.apiRequest("POST", s.url + "login-token/", &data)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode != 200 {
        return nil, newApiError(resp, body)
    }

    var obj = LoginFarmer{}
    err = json.Unmarshal([]byte(body), &obj)
    if err != nil {
        return nil, err
    }

    return &obj, nil
}


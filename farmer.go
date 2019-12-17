package apileek

import (
    "encoding/json"
)

type farmerService struct {
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
    Id uint `json:"id"`
    Login string `json:"login"`
    Team team `json:"team"`
    Name string `json:"name"`
    Talent uint `json:"talent"`
    Leeks map[uint]leek `json:"leeks"`
    AvatarChanged timestamp `json:"avatar_changed"`
    TalentMore uint `json:"talent_more"`
    Victories uint `json:"victories"`
    Draws uint `json:"draws"`
    Defeats uint `json:"defeats"`
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
    Banned uint `json:"banned"`
    WonSoloTournaments uint `json:"won_solo_tournaments"`
    WonFarmerTournaments uint `json:"won_farmer_tournaments"`
    WonTeamTournaments uint `json:"won_team_tournaments"`
    TotalLevel uint `json:"total_level"`
    LeekCount uint `json:"leek_count"`
    InGarden uint `json:"in_garden"`
    Fights uint `json:"fights"`
    Github json.RawMessage `json:"github"` // missing information
    Website json.RawMessage `json:"website"` // missing information
    ForumMessages uint `json:"forum_messages"`
    DidacticielSeen uint `json:"didacticiel_seen"`
    Contributor bool `json:"contributor"`
    Trophies uint `json:"trophies"`
    Habs uint `json:"habs"`
    Crystals uint `json:"crystals"`
    Weapons []weapon `json:"weapons"`
    Chips []chip `json:"chips"`
    Ais []aiBase `json:"ais"`
    Potions []potion `json:"potions"`
    Hats []hat `json:"hats"`
    Tournament tournament `json:"tournament"`
    Candidacy json.RawMessage `json:"candidacy"` // missing information
}

type weapon struct {
    Id uint `json:"id"`
    Template uint `json:"template"`
}

type chip struct {
    Id uint `json:"id"`
    Template uint `json:"template"`
}

type potion struct {
    Id uint `json:"id"`
    Template uint `json:"template"`
    Quantity uint `json:"quantity"`
}

type tournament struct {
    Registered bool `json:"registered"`
    Current json.RawMessage `json:"current"` // missing information
}

type team struct {
    Id uint `json:"id"`
    Name string `json:"name"`
    Level uint `json:"level"`
    EmblemChanged timestamp `json:"emblem_changed"`
}

type leek struct {
    Id uint `json:"id"`
    Name string `json:"name"`
    Color string `json:"color"`
    Capital uint `json:"capital"`
    Level uint `json:"level"`
    Talent uint `json:"talent"`
    Skin uint `json:"skin"`
    Hat *uint `json:"Hat"`
}

type hat struct {
    Id uint `json:"id"`
    Template uint `json:"template"`
    Level uint `json:"level"`
    Name string `json:"name"`
    HatTemplate uint `json:"hat_template"`
}

// Get a auth token from the api
func (s *farmerService) LoginToken(
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


package apileek

import (
    "fmt"
    "encoding/json"
    "strconv"
)

type leekService struct {
    apiService
}

type leekBase struct {
    Id string `json:"id"`
    Name string `json:"name"`
}

type Leek struct {
    Id uint `json:"id"`
    Name string `json:"name"`
    Farmer farmerInfoBase `json:"farmer"`
    Talent uint `json:"talent"`
    TalentMore uint `json:"talent_more"`
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
    Chips []chipRef `json:"chip"`
    Weapons []weaponRef `json:"weapons"`
    Ai aiBase `json:"ai"`
    MaxWeapons uint `json:"max_weapons"`
    MaxChips uint `json:"max_chips"`
    InGarden bool `json:"in_garden"`
    Fights []fight `json:"fights"`
    Skin uint `json:"skin"`
    Hat uint `json:"hat"`
    TalentHistory []uint `json:"talent_history"`
    Tournaments []tournamentHist `json:"tournaments"`
    Capital uint `json:"capital"`
    LevelSeen uint `json:"level_seen"`
    Registers
    Tournament tournament `json:"tournament"`
}

type Registers struct {
    Registers json.RawMessage `json:"registers"` // missing information
}

type leekCount struct {
    Leeks uint `json:"leeks"`
}

type Price struct {
    Price uint `json:"price"`
}

type popup struct {
    Popup levelPopup `json:"popup"`
}

type levelPopup struct {
    Level uint `json:"level"`
    Weapons []string `json:"weapons"`
    Chips []string `json:"chips"`
    Functions []string `json:"functions"`
    Hats json.RawMessage `json:"hats"` // missing information
    Gains gain `json:"gains"`
}

type gain struct {
    Life uint `json:"life"`
    Capital uint `json:"capital"`
}

type tournamentHist struct {
    Id uint `json:"id"`
    Date timestamp `json:"date"`
}

type fight struct {
    Id uint `json:"id"`
    Date timestamp `json:"date"`
    Type uint `json:"type"`
    Context uint `json:"context"`
    Status uint `json:"status"`
    Winner int `json:"winner"`
    LeekTeam uint `json:"leek_team"`
    Result string `json:"result"`
    Leeks1 []leekBase `json:"leeks1"`
    Leeks2 []leekBase `json:"leeks2"`
}

type farmerInfoBase struct {
    Id uint `json:"id"`
    Name string `json:"name"`
}

// Add a chip to a leek
func (s *leekService) AddChip(
    leekId uint, // Leek id
    chipId uint, // Chip id
) (uint, error) {
    data := "leek_id=" + strconv.FormatUint(uint64(leekId), 10) +
            "&chip_id=" + strconv.FormatUint(uint64(chipId), 10)
    resp, body, err := s.apiRequest("POST", s.url + "add-chip/", &data)
    if err != nil {
        return 0, err
    }

    if resp.StatusCode != 200 {
        return 0, newApiError(resp, body)
    }

    var obj = id{}
    err = json.Unmarshal([]byte(body), &obj)
    if err != nil {
        return 0, err
    }

    return obj.Id, nil
}

// Add a weapon to a leek
func (s *leekService) AddWeapon(
    leekId uint, // Leek id
    weaponId uint, // Weapon id
) (uint, error) {
    data := "leek_id=" + strconv.FormatUint(uint64(leekId), 10) +
            "&weapon_id=" + strconv.FormatUint(uint64(weaponId), 10)
    resp, body, err := s.apiRequest("POST", s.url + "add-weapon/", &data)
    if err != nil {
        return 0, err
    }

    if resp.StatusCode != 200 {
        return 0, newApiError(resp, body)
    }

    var obj = id{}
    err = json.Unmarshal([]byte(body), &obj)
    if err != nil {
        return 0, err
    }

    return obj.Id, nil
}

// Create a new leek
// TODO NOT TESTED
func (s *leekService) Create(
    name string, // Leek name
) (uint, error) {
    data := "name=" + name
    resp, body, err := s.apiRequest("POST", s.url + "create/", &data)
    if err != nil {
        return 0, err
    }

    if resp.StatusCode != 200 {
        return 0, newApiError(resp, body)
    }

    var obj = id{}
    err = json.Unmarshal([]byte(body), &obj)
    if err != nil {
        return 0, err
    }

    return obj.Id, nil
}

// Delete a leek register
func (s *leekService) DeleteRegister(
    leekId uint, // Leek id
    key string, // The register key
) error {
    data := "leek_id=" + strconv.FormatUint(uint64(leekId), 10) +
            "&key=" + key
    resp, body, err := s.apiRequest("POST", s.url + "delete-register/", &data)
    if err != nil {
        return err
    }

    if resp.StatusCode != 200 {
        return newApiError(resp, body)
    }

    return nil
}

// Get leek information
func (s *leekService) Get(
    leekId uint, // Leek id
) (*Leek, error) {
    data := "leek_id=" + strconv.FormatUint(uint64(leekId), 10)
    resp, body, err := s.apiRequest("POST", s.url + "get/", &data)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode != 200 {
        return nil, newApiError(resp, body)
    }

    var obj = Leek{}
    err = json.Unmarshal([]byte(body), &obj)
    if err != nil {
        return nil, err
    }

    return &obj, nil
}

// Get leek count
func (s *leekService) GetCount(
) (uint, error) {
    data := ""
    resp, body, err := s.apiRequest("POST", s.url + "get-count/", &data)
    if err != nil {
        return 0, err
    }

    if resp.StatusCode != 200 {
        return 0, newApiError(resp, body)
    }

    var obj = leekCount{}
    err = json.Unmarshal([]byte(body), &obj)
    if err != nil {
        return 0, err
    }

    return obj.Leeks, nil
}

// Get leek image as svg string
func (s *leekService) GetImage(
    leek uint, // Leek id
    scale uint, // Image scale
) (string, error) {
    data := "leek=" + strconv.FormatUint(uint64(leek), 10) +
            "&scale=" + strconv.FormatUint(uint64(scale), 10)
    resp, body, err := s.apiRequest("POST", s.url + "get-image/", &data)
    if err != nil {
        return "", err
    }

    if resp.StatusCode != 200 {
        return "", newApiError(resp, body)
    }

    var obj string = body
    return obj, nil
}

// Get leek level popup information
func (s *leekService) GetLevelPopup(
    leekId uint, // Leek id
) (*levelPopup, error) {
    data := "leek_id=" + strconv.FormatUint(uint64(leekId), 10)
    resp, body, err := s.apiRequest("POST", s.url + "get-level-popup/", &data)
    if err != nil {
        return nil, err
    }

    fmt.Println(body)
    if resp.StatusCode != 200 {
        return nil, newApiError(resp, body)
    }

    var obj = popup{}
    err = json.Unmarshal([]byte(body), &obj)
    if err != nil {
        return nil, err
    }

    return &obj.Popup, nil
}

// Get next price
func (s *leekService) GetNextPrice(
) (*Price, error) {
    data := ""
    resp, body, err := s.apiRequest("POST", s.url + "get-next-price/", &data)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode != 200 {
        return nil, newApiError(resp, body)
    }

    var obj = Price{}
    err = json.Unmarshal([]byte(body), &obj)
    if err != nil {
        return nil, err
    }

    return &obj, nil
}

// Get leek private information. Same as Leek.Get
func (s *leekService) GetPrivate(
    leekId uint, // Leek id
) (*Leek, error) {
    data := "leek_id=" + strconv.FormatUint(uint64(leekId), 10)
    resp, body, err := s.apiRequest("POST", s.url + "get-private/", &data)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode != 200 {
        return nil, newApiError(resp, body)
    }

    var obj = Leek{}
    err = json.Unmarshal([]byte(body), &obj)
    if err != nil {
        return nil, err
    }

    return &obj, nil
}

// Get registers
func (s *leekService) GetRegisters(
    leekId uint, // Leek id
) (*Registers, error) {
    data := "leek_id=" + strconv.FormatUint(uint64(leekId), 10)
    resp, body, err := s.apiRequest("POST", s.url + "get-registers/", &data)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode != 200 {
        return nil, newApiError(resp, body)
    }

    var obj = Registers{}
    err = json.Unmarshal([]byte(body), &obj)
    if err != nil {
        return nil, err
    }

    return &obj, nil
}

// Registers to the tournament
func (s *leekService) RegisterTournament(
    leekId uint, // Leek id
) error {
    data := "leek_id=" + strconv.FormatUint(uint64(leekId), 10)
    resp, body, err := s.apiRequest("POST", s.url + "register-tournament/", &data)
    if err != nil {
        return err
    }

    if resp.StatusCode != 200 {
        return newApiError(resp, body)
    }

    return nil
}

// Remove the ai equiped by the leek
func (s *leekService) RemoveAi(
    leekId uint, // Leek id
) error {
    data := "leek_id=" + strconv.FormatUint(uint64(leekId), 10)
    resp, body, err := s.apiRequest("POST", s.url + "remove-ai/", &data)
    if err != nil {
        return err
    }

    if resp.StatusCode != 200 {
        return newApiError(resp, body)
    }

    return nil
}

// Remove a chip from leek equipement
// This method work only for your main leek
func (s *leekService) RemoveChip(
    chipId uint, // Chip id
) error {
    data := "chip_id=" + strconv.FormatUint(uint64(chipId), 10)
    resp, body, err := s.apiRequest("POST", s.url + "remove-chip/", &data)
    if err != nil {
        return err
    }

    if resp.StatusCode != 200 {
        return newApiError(resp, body)
    }

    return nil
}

// Remove the leek hat
// TODO seem to not work
func (s *leekService) RemoveHat(
    leekId uint, // Leek id
) error {
    data := "leek_id=" + strconv.FormatUint(uint64(leekId), 10)
    resp, body, err := s.apiRequest("POST", s.url + "remove-ai/", &data)
    if err != nil {
        return err
    }

    fmt.Println(resp)
    fmt.Println(body)
    if resp.StatusCode != 200 {
        return newApiError(resp, body)
    }

    return nil
}

// Remove a weapon from leek equipement
// This method work only for your main leek
func (s *leekService) RemoveWeapon(
    weaponId uint, // Weapon id
) error {
    data := "weapon_id=" + strconv.FormatUint(uint64(weaponId), 10)
    resp, body, err := s.apiRequest("POST", s.url + "remove-weapon/", &data)
    if err != nil {
        return err
    }

    if resp.StatusCode != 200 {
        return newApiError(resp, body)
    }

    return nil
}

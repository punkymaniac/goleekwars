package apileek

import (
    "encoding/json"
    "strconv"
)

type marketService struct {
    apiService
}

type rootItemTemplate struct {
    Items map[uint]ItemTemplate `json:"items"`
}

type ItemTemplate struct {
    Id uint `json:"id"`
    Type uint `json:"type"`
    Name string `json:"name"`
    Level uint `json:"level"`
    PriveHabs uint `json:"price_habs"`
    PriceCrystals uint `json:"price_crystals"`
    FarmerCount uint `json:"farmer_count"`
    LeekCount uint `json:"leek_count"`
    Sellable uint `json:"sellable"`
    SellPrice uint `json:"sell_price"`
    Leeks []uint `json:"leeks"`
}

type moneyValue struct {
    Money uint `json:"money"`
}

type Buy struct {
    moneyValue
    Item uint `json:"item"`
}

// Buy one item with crystals
func (s *marketService) BuyCrystals(
    itemId uint, // Item id to buy
) (*Buy, error) {
    data := "item_id=" + strconv.FormatUint(uint64(itemId), 10)
    resp, body, err := s.apiRequest("POST", s.url + "buy-crystals/", &data)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode != 200 {
        return nil, newApiError(resp, body)
    }

    var obj = Buy{}
    err = json.Unmarshal([]byte(body), &obj)
    if err != nil {
        return nil, err
    }

    return &obj, nil
}

// Buy one item with habs
func (s *marketService) BuyHabs(
    itemId uint, // Item id to buy
) (*Buy, error) {
    data := "item_id=" + strconv.FormatUint(uint64(itemId), 10)
    resp, body, err := s.apiRequest("POST", s.url + "buy-habs/", &data)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode != 200 {
        return nil, newApiError(resp, body)
    }

    var obj = Buy{}
    err = json.Unmarshal([]byte(body), &obj)
    if err != nil {
        return nil, err
    }

    return &obj, nil
}

// Get item templates
func (s *marketService) GetItemTemplates(
) (map[uint]ItemTemplate, error) {
    data := ""
    resp, body, err := s.apiRequest("POST", s.url + "get-item-templates/", &data)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode != 200 {
        return nil, newApiError(resp, body)
    }

    var obj = rootItemTemplate{}
    err = json.Unmarshal([]byte(body), &obj)
    if err != nil {
        return nil, err
    }

    return obj.Items, nil
}

// Sell one item against habs
func (s *marketService) SellHabs(
    itemId uint, // Item id to sell
) (uint, error) {
    data := "item_id=" + strconv.FormatUint(uint64(itemId), 10)
    resp, body, err := s.apiRequest("POST", s.url + "sell-habs/", &data)
    if err != nil {
        return 0, err
    }

    if resp.StatusCode != 200 {
        return 0, newApiError(resp, body)
    }

    var obj = moneyValue{}
    err = json.Unmarshal([]byte(body), &obj)
    if err != nil {
        return 0, err
    }

    return obj.Money, nil
}


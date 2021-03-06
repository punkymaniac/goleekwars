package apileek

import (
    "strconv"
    "encoding/json"
)

type AiService struct {
    apiService
}

type Folder struct {
    Id uint64 `json:"id"`
    Name string `json:"name"`
    Folder uint64 `json:"folder"`
}

type Ais struct {
    Ais []aiAllInfo `json:"ais"`
    Folders []Folder `json:"folders"`
    LeekAis map[uint64]uint64 `json:"leek_ais"`
}

type ai struct {
    Ai aiInfo
}

type aiBase struct {
    Id uint64 `json:"id"`
    Name string `json:"name"`
    Level uint64 `json:"level"`
}

type aiInfo struct {
    aiBase
    Code string `json:"code"`
    Folder uint64 `json:"folder"`
}

type aiAllInfo struct {
    aiInfo
    Valid bool `json:"valid"`
    V2 uint64 `json:"v2"`
}

type aiList struct {
    Ai aiElem `json:"ai"`
}

type aiElem struct {
    aiBase
    Code string `json:"code"`
    Valid bool `json:"valid"`
    Owner uint64 `json:"owner"`
}

// Get a AI
func (s *AiService) Get(
    aiId uint64, // AI id
) (*aiElem, error) {
    data := "ai_id=" + strconv.FormatUint(aiId, 10)
    resp, body, err := s.apiRequest("POST", s.url + "get/", &data)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode != 200 {
        return nil, newApiError(resp, body)
    }

    var obj = aiList{}
    err = json.Unmarshal([]byte(body), &obj)
    if err != nil {
        return nil, err
    }

    return &obj.Ai, nil
}

// Get all farmer AI
func (s *AiService) GetFarmerAis(
) (*Ais, error) {
    data := ""
    resp, body, err := s.apiRequest("POST", s.url + "get-farmer-ais/", &data)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode != 200 {
        return nil, newApiError(resp, body)
    }

    var obj = Ais{}
    err = json.Unmarshal([]byte(body), &obj)
    if err != nil {
        return nil, err
    }

    return &obj, nil
}

// Create a new AI
func (s *AiService) New(
    folderId uint64, // Folder id where to create the AI
    v2 bool, // Id true, create a V2 AI
) (*aiInfo, error) {
    data := "folder_id=" + strconv.FormatUint(folderId, 10) + "&v2=" + strconv.FormatBool(v2)
    resp, body, err := s.apiRequest("POST", s.url + "new/", &data)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode != 200 {
        return nil, newApiError(resp, body)
    }

    var obj = ai{}
    err = json.Unmarshal([]byte(body), &obj)
    if err != nil {
        return nil, err
    }

    return &obj.Ai, nil
}

// Rename a AI
func (s *AiService) Rename(
    aiId uint64, // AI id
    name string, // The new name of the AI
) (error) {
    data := "ai_id=" + strconv.FormatUint(aiId, 10) + "&new_name=" + name
    resp, body, err := s.apiRequest("POST", s.url + "rename/", &data)
    if err != nil {
        return err
    }

    if resp.StatusCode != 200 {
        return newApiError(resp, body)
    }

    return nil
}

// Save a code on the AI
func (s *AiService) Save(
    aiId uint64, // AI id
    code string, // Code to save in the AI
) (error) {
    data := "ai_id=" + strconv.FormatUint(aiId, 10) + "&code=" + code
    resp, body, err := s.apiRequest("POST", s.url + "save/", &data)
    if err != nil {
        return err
    }

    if resp.StatusCode != 200 {
        return newApiError(resp, body)
    }

    return nil
}


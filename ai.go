package apileek

import (
    "strconv"
    "encoding/json"
)

type aiService struct {
    apiService
}

type Folder struct {
    Id uint `json:"id"`
    Name string `json:"name"`
    Folder uint `json:"folder"`
}

type Ais struct {
    Ais []aiAllInfo `json:"ais"`
    Folders []Folder `json:"folders"`
    LeekAis map[uint]uint `json:"leek_ais"`
}

type ai struct {
    Ai aiInfo
}

type aiBase struct {
    Id uint `json:"id"`
    Name string `json:"name"`
    Level uint `json:"level"`
}

type aiInfo struct {
    aiBase
    Code string `json:"code"`
    Folder uint `json:"folder"`
}

type aiAllInfo struct {
    aiInfo
    Valid bool `json:"valid"`
    V2 uint `json:"v2"`
}

type aiList struct {
    Ai aiElem `json:"ai"`
}

type aiElem struct {
    aiBase
    Code string `json:"code"`
    Valid bool `json:"valid"`
    Owner uint `json:"owner"`
}

// Change the folder of the AI
func (s *aiService) ChangeFolder(
    aiId uint, // AI id
    folderId uint, // Folder id to move the AI
) error {
    data := "ai_id=" + strconv.FormatUint(uint64(aiId), 10) + "&folder_id=" + strconv.FormatUint(uint64(folderId), 10)
    resp, body, err := s.apiRequest("POST", s.url + "change-folder/", &data)
    if err != nil {
        return err
    }

    if resp.StatusCode != 200 {
        return newApiError(resp, body)
    }

    return nil
}

// Delete a AI
func (s *aiService) Delete(
    aiId uint, // Ai id
) error {
    data := "ai_id=" + strconv.FormatUint(uint64(aiId), 10)
    resp, body, err := s.apiRequest("POST", s.url + "delete/", &data)
    if err != nil {
        return err
    }

    if resp.StatusCode != 200 {
        return newApiError(resp, body)
    }

    return nil
}

// Get a AI
func (s *aiService) Get(
    aiId uint, // AI id
) (*aiElem, error) {
    data := "ai_id=" + strconv.FormatUint(uint64(aiId), 10)
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
func (s *aiService) GetFarmerAis(
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
func (s *aiService) New(
    folderId uint, // Folder id where to create the AI
    v2 bool, // Id true, create a V2 AI
) (*aiInfo, error) {
    data := "folder_id=" + strconv.FormatUint(uint64(folderId), 10) + "&v2=" + strconv.FormatBool(v2)
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
func (s *aiService) Rename(
    aiId uint, // AI id
    name string, // The new name of the AI
) (error) {
    data := "ai_id=" + strconv.FormatUint(uint64(aiId), 10) + "&new_name=" + name
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
func (s *aiService) Save(
    aiId uint, // AI id
    code string, // Code to save in the AI
) (error) {
    data := "ai_id=" + strconv.FormatUint(uint64(aiId), 10) + "&code=" + code
    resp, body, err := s.apiRequest("POST", s.url + "save/", &data)
    if err != nil {
        return err
    }

    if resp.StatusCode != 200 {
        return newApiError(resp, body)
    }

    return nil
}

// Test the ai
// TODO Incomplete, missing documentation
// Always return a error
func (s *aiService) Test(
    aiId uint, // AI id
    leekId uint, // Leek id
    bots []string, // Bots
    Type string, // TODO Missing documentation
) (error) {
    return newError("Api not implemented: Missing documentation")
}

//
// TODO Incomplete, missing documentation
// Always return a error
func (s *aiService) TestNew(
    jsonData string, // TODO Missing documentation
) (error) {
    return newError("Api not implemented: Missing documentation")
}

//
// TODO Incomplete, missing documentation
// Always return a error
func (s *aiService) TestV2(
    jsonData string, // TODO Missing documentation
) (error) {
    return newError("Api not implemented: Missing documentation")
}


package apileek

import (
    "encoding/json"
)

type functionService struct {
    apiService
}

type Function struct {
    Id uint `json:"id"`
    Name string `json:"name"`
    Level uint `json:"level"`
    Category uint `json:"category"`
    Operations int `json:"operations"`
    ArgumentsNames []string `json:"arguments_names"`
    ArgumentsTypes []string `json:"arguments_types"`
    ReturnType int `json:"return_type"`
    ReturnName string `json:"return_name"`
    Deprecated uint `json:"deprecated"`
}

type FunctionCategories struct {
    Id uint `json:"id"`
    Name string `json:"name"`
}

type FunctionOp struct {
    Op int `json:"op"`
    VarOp json.RawMessage `json:"var_op"` // TODO Bad format of the received data
}

type functionCatList struct {
    Categories map[string]FunctionCategories `json:"categories"`
}

type functionList struct {
    Functions []Function `json:"functions"`
}

// Get all function
func (s *functionService) GetAll(
) ([]Function, error) {
    data := ""
    resp, body, err := s.apiRequest("POST", s.url + "get-all/", &data)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode != 200 {
        return nil, newApiError(resp, body)
    }

    var obj = functionList{}
    err = json.Unmarshal([]byte(body), &obj)
    if err != nil {
        return nil, err
    }
    return obj.Functions, nil
}

// Get function categories
func (s *functionService) GetCategories(
) (map[string]FunctionCategories, error) {
    data := ""
    resp, body, err := s.apiRequest("POST", s.url + "get-categories/", &data)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode != 200 {
        return nil, newApiError(resp, body)
    }

    var obj = functionCatList{}
    err = json.Unmarshal([]byte(body), &obj)
    if err != nil {
        return nil, err
    }
    return obj.Categories, nil
}

// Get function operations
func (s *functionService) Operations(
) (map[string]FunctionOp, error) {
    data := ""
    resp, body, err := s.apiRequest("POST", s.url + "operations/", &data)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode != 200 {
        return nil, newApiError(resp, body)
    }

    var obj = map[string]FunctionOp{}
    err = json.Unmarshal([]byte(body), &obj)
    if err != nil {
        return nil, err
    }
    return obj, nil
}


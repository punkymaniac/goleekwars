package apileek

import (
    "encoding/json"
)

type changelogService struct {
    apiService
}

type Changelog struct {
    Version uint `json:"version"`
    VersionName string `json:"version_name"`
    Date string `json:"date"`
    Data string `json:"data"`
    ForumTopic uint `json:"forum_topic"`
    ForumCategory uint `json:"forum_category"`
    Image bool `json:"image"`
}

type changelogList struct {
    Changelog []Changelog `json:"changelog"`
}

type changelogLast struct {
    Changelog Changelog `json:"changelog"`
}

// Get all changelog
func (s *changelogService) Get(
    language string, // The country code of the language to get the changelog
) ([]Changelog, error) {
    data := "language=" + language
    resp, body, err := s.apiRequest("POST", s.url + "get/", &data)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode != 200 {
        return nil, newApiError(resp, body)
    }

    var obj = changelogList{}
    err = json.Unmarshal([]byte(body), &obj)
    if err != nil {
        return nil, err
    }

    return obj.Changelog, nil
}

// Get the last changelog
func (s *changelogService) GetLast(
    language string, // The country code of the language to get the changelog
) (*Changelog, error) {
    data := "language=" + language
    resp, body, err := s.apiRequest("POST", s.url + "get-last/", &data)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode != 200 {
        return nil, newApiError(resp, body)
    }

    var obj = changelogLast{}
    err = json.Unmarshal([]byte(body), &obj)
    if err != nil {
        return nil, err
    }

    return &obj.Changelog, nil
}


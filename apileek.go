package apileek

import (
    "time"
    "bytes"
    "net/http"
    "io/ioutil"
)

type ApiLeek struct {
    client apiClient
    Ai aiService
    AiFolder aiFolderService
    Changelog changelogService
    Farmer farmerService
}

type apiClient interface {
    ApiRequest(string, string, *string) (*http.Response, string, error)
    SetToken(string)
}

type apiService struct {
    client apiClient
    url string
}

type leekClient struct {
    httpcli *http.Client
    token *string
}

// Create a new ApiLeek object
func NewApi() ApiLeek {
    apiUrl := "https://leekwars.com/api/"
    leekCli := leekClient{
                   httpcli: &http.Client{Timeout: 10 * time.Second},
                   token: nil,
               }
    api := ApiLeek{
                client: &leekCli,
                Ai: aiService{apiService{client: &leekCli, url: apiUrl + "ai/"}},
                AiFolder: aiFolderService{apiService{client: &leekCli, url: apiUrl + "ai-folder/"}},
                Changelog: changelogService{apiService{client: &leekCli, url: apiUrl + "changelog/"}},
                Farmer: farmerService{apiService{client: &leekCli, url: apiUrl + "farmer/"}},
           }

    return api
}

// Made a API request over the apiClient
func (c *apiService) apiRequest(
    method string, // HTTP method od the request
    uri string, // Uri of the request
    data *string, // If not nil, string of the post data
) (*http.Response, string, error) {
    return c.client.ApiRequest(method, uri, data)
}

// Made a request to the API
func (c *leekClient) ApiRequest(
    method string, // HTTP method od the request
    uri string, // Uri of the request
    data *string, // If not nil, string of the post data
) (*http.Response, string, error) {
    var err error
    var req *http.Request

    if data != nil {
        req, err = http.NewRequest(method, uri, bytes.NewBuffer([]byte(*data)))
    } else {
        req, err = http.NewRequest(method, uri, nil)
    }
    if err != nil {
        return nil, "", err
    }

    if method == "POST" {
        req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    }

    // Use auth token
    if c.token != nil {
        req.Header.Set("Authorization", "Bearer " + *c.token)
    }

    resp, err := c.httpcli.Do(req)
    if err != nil {
        return nil, "", err
    }
    defer resp.Body.Close()

    raw, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, "", err
    }

    return resp, string(raw), nil
}

// Set a token to auth with the api
func (c *leekClient) SetToken(
    token string, // Token used to auth with the api
) {
    c.token = &token
}

// Auth to the api
func (l *ApiLeek) Auth(
    username string, // Username of the account
    password string, // Password of the account
) (*farmerInfo, error) {
    farmer, err := l.Farmer.LoginToken(username, password)
    if err != nil {
        return nil, err
    }

    if (farmer != nil) {
        l.client.SetToken(farmer.Token)
        return &farmer.Farmer, nil
    }

    return nil, newError("Nil token received")
}


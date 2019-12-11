package apileek

import (
    "time"
    "bytes"
    "net/http"
    "io/ioutil"
)

type ApiLeek struct {
    url string
    client *leekClient
    Ai AiService
    Farmer FarmerService
}

type apiService struct {
    client *leekClient
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
                url: apiUrl,
                client: &leekCli,
                Ai: AiService{apiService{client: &leekCli, url: apiUrl + "ai/"}},
                Farmer: FarmerService{apiService{client: &leekCli, url: apiUrl + "farmer/"}},
           }
    return api
}


func (c *apiService) apiRequest(
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
    if c.client.token != nil {
        req.Header.Set("Authorization", "Bearer " + *c.client.token)
    }

    resp, err := c.client.httpcli.Do(req)
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


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
    Chip chipService
    Constant constantService
    Country countryService
    Farmer farmerService
    Fight fightService
    Forum forumService
    Function functionService
    Garden gardenService
    Hat hatService
    History historyService
    Lang langService
    Leekwars leekwarsService
    Weapon weaponService
}

type apiClient interface {
    ApiRequest(string, string, *string) (*http.Response, string, error)
    SetToken(string)
    SetSession(string)
}

type apiService struct {
    client apiClient
    url string
}

type leekClient struct {
    httpcli *http.Client
    token *string
    session *string
}

// Create a new ApiLeek object
func NewApi() ApiLeek {
    apiUrl := "https://leekwars.com/api/"

    leekCli := leekClient{
                   httpcli: &http.Client{Timeout: 10 * time.Second},
                   token: nil,
                   session: nil,
               }

    // Get and keep the session id, if exist
    resp, _, _ := leekCli.ApiRequest("POST", apiUrl, nil)
    if resp != nil {
        for _, cookie := range resp.Cookies() {
            if cookie.Name == "PHPSESSID" {
                leekCli.SetSession(cookie.Value)
                break
            }
        }
    }

    api := ApiLeek{
                client: &leekCli,
                Ai: aiService{apiService{client: &leekCli, url: apiUrl + "ai/"}},
                AiFolder: aiFolderService{apiService{client: &leekCli, url: apiUrl + "ai-folder/"}},
                Changelog: changelogService{apiService{client: &leekCli, url: apiUrl + "changelog/"}},
                Chip: chipService{apiService{client: &leekCli, url: apiUrl + "chip/"}},
                Constant: constantService{apiService{client: &leekCli, url: apiUrl + "constant/"}},
                Country: countryService{apiService{client: &leekCli, url: apiUrl + "country/"}},
                Farmer: farmerService{apiService{client: &leekCli, url: apiUrl + "farmer/"}},
                Fight: fightService{apiService{client: &leekCli, url: apiUrl + "fight/"}},
                Forum: forumService{apiService{client: &leekCli, url: apiUrl + "forum/"}},
                Function: functionService{apiService{client: &leekCli, url: apiUrl + "function/"}},
                Garden: gardenService{apiService{client: &leekCli, url: apiUrl + "garden/"}},
                Hat: hatService{apiService{client: &leekCli, url: apiUrl + "hat/"}},
                History: historyService{apiService{client: &leekCli, url: apiUrl + "history/"}},
                Lang: langService{apiService{client: &leekCli, url: apiUrl + "lang/"}},
                Leekwars: leekwarsService{apiService{client: &leekCli, url: apiUrl + "leek-wars/"}},
                Weapon: weaponService{apiService{client: &leekCli, url: apiUrl + "weapon/"}},
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

    // Use the saved session
    if c.session != nil {
        req.AddCookie(&http.Cookie{Name: "PHPSESSID", Value: *c.session})
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

// Set the session id
func (c *leekClient) SetSession(
    session string, // Current session id
) {
    c.session = &session
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


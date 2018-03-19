package go_league_of_legends

import (
    "fmt"
    "net/http"
    "errors"
    "encoding/json"
    "io/ioutil"
)

// add prefix
const (
    baseUrl           = ".api.riotgames.com/lol/"
    urlSummonerByName = "summoner/v3/summoners/by-name/"
    urlMatchLists     = "match/v3/matchlists"
)

type Client struct {
    token    string
    platform string
}

func (c *Client) Init(token string, platform string) {
    c.token = token
    c.platform = platform
}

func (c *Client) getUrl(url string) ([]byte, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode != 200 {
        return nil, errors.New(fmt.Sprintf("API refused with code %v", resp.Status))
    }

    return body, nil
}

func (c *Client) GetSummonerByName(name string) (Summoner, error) {
    var (
        res  Summoner
        url  string
        err  error
        body []byte
    )

    url = fmt.Sprintf("https://%s%s%s%s?api_key=%s", c.platform, baseUrl, urlSummonerByName, name, c.token)

    if body, err = c.getUrl(url); err != nil {
        return Summoner{}, err
    }

    if err = json.Unmarshal(body, &res); err != nil {
        return Summoner{}, err
    }

    return res, nil
}

func (c *Client) GetMatchListsByAccount(name string, beginIndex int) (MatchLists, error) {
    var (
        res  MatchLists
        url  string
        err  error
        body []byte
    )

    url = fmt.Sprintf("https://%s%s%s%s?api_key=%s&beginIndex=%s", c.platform, baseUrl, urlMatchLists, name, c.token, beginIndex)

    if body, err = c.getUrl(url); err != nil {
        return MatchLists{}, err
    }

    if err = json.Unmarshal(body, &res); err != nil {
        return MatchLists{}, err
    }

    return res, nil
}

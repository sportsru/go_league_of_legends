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
    baseUrl                = ".api.riotgames.com/lol/"
    urlSummonerByName      = "summoner/v3/summoners/by-name/"
    urlMatchListsByAccount = "match/v3/matchlists/by-account/"
)

type Client struct {
    token    string
}

func (c *Client) Init(token string) {
    c.token = token
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

func (c *Client) GetSummonerByName(platform string, name string) (Summoner, error) {
    var (
        res  Summoner
        url  string
        err  error
        body []byte
    )

    url = fmt.Sprintf("https://%s%s%s%s?api_key=%s", platform, baseUrl, urlSummonerByName, name, c.token)

    if body, err = c.getUrl(url); err != nil {
        return Summoner{}, err
    }

    if err = json.Unmarshal(body, &res); err != nil {
        return Summoner{}, err
    }

    return res, nil
}

func (c *Client) GetMatchListsByAccount(platform string, playerId int, beginIndex int) (MatchLists, error) {
    var (
        res  MatchLists
        url  string
        err  error
        body []byte
    )

    url = fmt.Sprintf("https://%s%s%s%d?api_key=%s&beginIndex=%d", platform, baseUrl, urlMatchListsByAccount, playerId, c.token, beginIndex)
    
    if body, err = c.getUrl(url); err != nil {
        return MatchLists{}, err
    }

    if err = json.Unmarshal(body, &res); err != nil {
        return MatchLists{}, err
    }

    return res, nil
}

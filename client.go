package iksm

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"path"
	"strconv"
	"strings"
)

const (
	ROOT = "https://app.splatoon2.nintendo.net"
)

type IksmClient struct {
	client *http.Client
}

func Client(token string) (*IksmClient, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{Jar: jar}

	req, err := http.NewRequest(`GET`, ROOT, nil)
	if err != nil {
		return nil, err
	}
	req.AddCookie(&http.Cookie{Name: `iksm_session`, Value: token})

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return &IksmClient{client: client}, nil
}

func (client *IksmClient) Get(p string) (r *http.Response, err error) {
	u, err := url.Parse(ROOT)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, p)
	return client.client.Get(u.String())
}

func (client *IksmClient) Parse(r io.Reader, v interface{}) error {
	if err := json.NewDecoder(r).Decode(&v); err != nil {
		return err
	}
	return nil
}

func (client *IksmClient) GetJson(p string, v interface{}) error {
	res, err := client.Get(p)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if err := client.Parse(res.Body, &v); err != nil {
		return err
	}
	return nil
}

func (client *IksmClient) GetStages() (*stages, error) {
	var v stages
	if err := client.GetJson(`api/data/stages`, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

func (client *IksmClient) GetResult(num int) (*result, error) {
	var v result
	if err := client.GetJson(`api/results/`+strconv.Itoa(num), &v); err != nil {
		return nil, err
	}
	return &v, nil
}

func (client *IksmClient) UpdateResults(dbname string, overwrite bool) error {
	indices, err := client.GetAvailableResultNumbers()
	if err != nil {
		return err
	}
	for _, i := range indices {
		result, err := client.GetResult(i)
		if err != nil {
			return err
		}
		if err := result.SaveToDB(dbname, overwrite); err != nil {
			if strings.Contains(err.Error(), `UNIQUE constraint failed:`) {
				log.Printf("Warning. Skip battle_number %d, because overwrite is false", i)
				continue
			}
			return err
		}
	}
	return nil
}

func (client *IksmClient) GetAvailableResultNumbers() ([]int, error) {
	var v struct {
		Results []struct {
			Id string `json:"battle_number"`
		} `json:"results"`
	}
	if err := client.GetJson(`api/results/`, &v); err != nil {
		return nil, err
	}
	s_indices := v.Results
	i_indices := make([]int, len(s_indices), len(s_indices))
	for i, s := range s_indices {
		index, err := strconv.Atoi(s.Id)
		if err != nil {
			return nil, err
		}
		i_indices[i] = index
	}
	return i_indices, nil
}

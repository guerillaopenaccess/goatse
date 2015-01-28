package doiutils

import (
    "fmt"
	"io/ioutil"
	"net/http"
)

func GetDOIMeta(doi string) (*DoiMeta, error) {
	fetchUrl := fmt.Sprintf("http://api.crossref.org/works/%s", doi)
	body, err := HttpGetBody(fetchUrl)
	if err != nil {
		return nil, err
	}
	returned := new(CrossRefWorksSingleton)
	returned.DecodeFromAPI(string(body))
	return &returned.Message, nil
}

func getListQuery(fetchUrl string) ([]DoiMeta, error) {
	var (
		body []byte
		returned *CrossRefWorksList
		err error
	)
	body, err = HttpGetBody(fetchUrl)
	if err != nil {
		return nil, err
	}	
	returned = new(CrossRefWorksList)
	returned.DecodeFromAPI(string(body))
	return returned.Message.Items, nil
}

func GetSample(n int) ([]DoiMeta, error) {
	return getListQuery(fmt.Sprintf("http://api.crossref.org/works?sample=%d", n))
}

func GetMemberSample(n int, member int16) ([]DoiMeta, error) {
	return getListQuery(fmt.Sprintf("http://api.crossref.org/works?sample=%d&filter=member:%d", n, member))
}

func GetMemberSampleByName(n int, membername string) (sample []DoiMeta, err error) {
	var member *PublisherMeta
	member, err = GetPublisherByName(membername)
	if err != nil {
		return nil, err
	}
	return GetMemberSample(n, int16(member.Id))
}

func GetNotOpenSample(n int) (notopen []DoiMeta, err error) {
	var (
		sample []DoiMeta
	)
	notopen = make([]DoiMeta, 0, n)
	for {
		sample, err = GetSample(n*4)
		if err != nil {
			return nil, err
		}
		for _, m := range sample {
			if !(m.IsOpen() || m.IsOnLibgen()) {
				notopen = append(notopen, m)
			}
			if len(notopen) >= n {
				break
			}
		}
		if len(notopen) >= n {
			break
		}
	}
	return notopen, nil
}

func HttpGetBody(fetchUrl string) ([]byte, error) {
	res, err := http.Get(fetchUrl)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

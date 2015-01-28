package doiutils

import (
	"time"
	"log"
	"strings"
	"net/url"
	"net/http"
	"encoding/json"
)

type DateParts struct{
	DateParts [][3]int `json:"date-parts"`
	UnixMS int64 `json:"timestamp"`
}

func (self *DateParts) TimeStruct() time.Time {
	unix_seconds := self.UnixMS / 1000
	return time.Unix(unix_seconds, 0)
}

type DoiAuthor struct{
	Given string `json:",omitempty"`
	Family string `json:",omitempty"`
}

type DoiLicense struct{
	ContentVersion string `json:"content-version"`

	Start DateParts

	DelayInDays int `json:"delay-in-days"`
	Url string
}

func (self *DoiLicense) IsValid() bool {
	return self.Start.TimeStruct().Before(time.Now())
}

func (self *DoiLicense) IsOpen() bool {
	return LicenseIsOpen(self.Url)
}

type DoiMeta struct{
	Issued DateParts
	Publisher string
	License []DoiLicense
	Source string
	Url string
	Doi string
	ISSN []string
	Score float32
	Page string
	Issue string
	Deposited DateParts
	ReferenceCount int `json:"reference-count"`
	Author []DoiAuthor
	Title []string
	Type string
	Volume string
	ContainerTitle []string `json:"container-title"`
	Subtitle []string
}

func (self DoiMeta) Journal() string {
	for _, J := range self.ContainerTitle {
		return J
	}
	return ""
}

func (self DoiMeta) MainTitle() string {
	for _, T := range self.Title {
		return T
	}
	return ""
}

func (self DoiMeta) IsOpen() bool {
	for _, L := range self.License {
		if L.IsOpen() && L.IsValid() {
			return true
		}
	}
	if strings.Contains(self.Publisher, "PLoS") {
		return true
	}
	return false
}

func (self DoiMeta) IsOnLibgen() bool {
	querystring := url.Values{}
	querystring.Add("doi", self.Doi)
	url := "http://libgen.in/scimag/get.php?" + querystring.Encode()
	res, err := http.Head(url)
	if err != nil {
		log.Printf("Error on HEAD to Libgen: %s", err.Error())
		return false
	}
	if res.StatusCode == 200 {
		return true
	}
	return false
}

type DoiMetaContainer struct{
	ItemsPerPage int `json:"items-per-page"`
	TotalResults int64 `json:"total-results"`
	Items []DoiMeta
}

type CrossRefWorksSingleton struct{
	Status string
	MessageVersion string `json:"message-version"`
	Message DoiMeta
	MessageType string `json:"message-type"`
}

func (self *CrossRefWorksSingleton) DecodeFromAPI(apiresponse string) {
	json.Unmarshal([]byte(apiresponse), self)
}

type CrossRefWorksList struct{	
	Status string
	MessageVersion string `json:"message-version"`
	Message DoiMetaContainer
	MessageType string `json:"message-type"`
}

func (self *CrossRefWorksList) DecodeFromAPI(apiresponse string) {
	json.Unmarshal([]byte(apiresponse), self)
}

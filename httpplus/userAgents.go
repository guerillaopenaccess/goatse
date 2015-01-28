package httpplus

import (
	"crypto/rand"
	"math/big"
)

// Select randomly at init.
var userAgents = [...]string{
	"Mac / Firefox 29: Mozilla/5.0 (Macintosh; Intel Mac OS X 10.9; rv:29.0) Gecko/20100101 Firefox/29.0",
	"Mac / Chrome 34: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/34.0.1847.137 Safari/537.36",
	"Mac / Safari 7: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_3) AppleWebKit/537.75.14 (KHTML, like Gecko) Version/7.0.3 Safari/537.75.14",
	"Windows / Firefox 29: Mozilla/5.0 (Windows NT 6.1; WOW64; rv:29.0) Gecko/20100101 Firefox/29.0",
	"Windows / Chrome 34: Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/34.0.1847.137 Safari/537.36",
	"Windows / IE 9: Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; WOW64; Trident/5.0)",
	"Windows / IE 10: Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.1; WOW64; Trident/6.0)",
	"Windows / IE 11: Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; rv:11.0) like Gecko",
}

func randomUserAgent() (ua string, err error) {
	var (
		maxn     int64
		bigindex *big.Int
		index    uint64
	)
	maxn = int64(len(userAgents))
	bigindex, err = rand.Int(rand.Reader, big.NewInt(maxn))
	if err != nil {
		return "", err
	}
	index = bigindex.Uint64()
	ua = userAgents[index]
	return ua, nil
}

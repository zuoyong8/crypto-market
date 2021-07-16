package xhttp

import (
	"crypto-market/common"
	"io/ioutil"
	"net/http"
)

//获取数据
func Get(url string, headers map[string]string) ([]byte, error) {
	req, err := http.NewRequest(common.GET, url, nil)
	if err != nil {
		return nil, err
	}

	transport := http.Transport{
		DisableKeepAlives: true, //禁止
	}

	client := http.Client{
		Transport: &transport,
	}

	//
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	info, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return info, nil
}

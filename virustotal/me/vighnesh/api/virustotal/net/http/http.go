package http

import (
	"bytes"
	"io/ioutil"
	"me/vighnesh/api/virustotal/net/multipart"
	"net/http"
)

func RequestPost(me *multipart.MultipartEntity, url string) ([]byte, error) {
	request, e := http.NewRequest("POST", url, bytes.NewBuffer(me.Build()))
	if e == nil {
		request.Header.Set("Content-Type", multipart.CONTENT_TYPE_FORM_DATA_WITH_BOUNDARY)
		request.Header.Add("User-Agent", multipart.USER_AGENT)
		return Execute(request)
	} else {
		return nil, e
	}
}

func RequestGet(apikey, url, key, value string) ([]byte, error) {
	request, e := http.NewRequest("GET", url, nil)
	if e == nil {
		request.Header.Add("User-Agent", multipart.USER_AGENT)
		query := request.URL.Query()
		query.Add(key, value)
		query.Add("apikey", apikey)
		request.URL.RawQuery = query.Encode()
		return Execute(request)
	} else {
		return nil, e
	}
}

func Execute(request *http.Request) ([]byte, error) {
	client := &http.Client{}
	response, e := client.Do(request)
	defer response.Body.Close()
	if e == nil {
		return ioutil.ReadAll(response.Body)
	} else {
		return nil, e
	}
}

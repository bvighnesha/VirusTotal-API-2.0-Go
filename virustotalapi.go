package virustotal

import (
	"errors"
	"io"
	"os"
	"vighnesh.org/virustotal/json"
	"vighnesh.org/virustotal/net/http"
	"vighnesh.org/virustotal/net/multipart"
	"vighnesh.org/virustotal/util"
)

const (
	URL_FILE_SCAN        string = "https://www.virustotal.com/vtapi/v2/file/scan"
	URL_FILE_RESCAN      string = "https://www.virustotal.com/vtapi/v2/file/rescan"
	URL_FILE_SCAN_REPORT string = "https://www.virustotal.com/vtapi/v2/file/report"
	URL_URL_SCAN         string = "https://www.virustotal.com/vtapi/v2/url/scan"
	URL_URL_SCAN_REPORT  string = "http://www.virustotal.com/vtapi/v2/url/report"
	URL_IP_SCAN_REPORT   string = "http://www.virustotal.com/vtapi/v2/ip-address/report"
	URL_DOMAIN_REPORT    string = "http://www.virustotal.com/vtapi/v2/domain/report"
	URL_COMMENTS         string = "https://www.virustotal.com/vtapi/v2/comments/put"
)

var (
	API_KEY_ERROR     error = errors.New("API Key is Not Valid")
	FILE_ERROR        error = errors.New("File Not Found or Can not Access")
	FILE_STREAM_ERROR error = errors.New("File Stream is nil")
	COMMENT_ERROR     error = errors.New("Some Thing Went Wrong")
	DOMAIN_ERROR      error = errors.New("Domain is Not Valid")
	FILE_NAME_ERROR   error = errors.New("File Name Error")
	IP_ADDRESS_ERROR  error = errors.New("IP Address is Not Valid")
)

type VirusTotalApi interface {
	ScanFile(file string) (*json.Response, error)
	ScanFileStream(filename string, reader io.Reader) (*json.Response, error)
	ScanURL(url string) (*json.Response, error)
	ReScanFile(resource string) (*json.ReScanResponse, error)
	FileReport(resource string) (*json.Report, error)
	URLReport(resource string) (*json.URLReport, error)
	IPReport(ip string) (*json.IPReport, error)
	DomainReport(domain string) (*json.DomainReport, error)
	Comment(resource, comment string) (*json.CommentStatus, error)
}

type virustotal struct {
	apiKey string
}

func Configure(apiKey string) (VirusTotalApi, error) {
	if apiKey != "" && len(apiKey) > 0 {
		return virustotal{apiKey}, nil
	}
	return nil, API_KEY_ERROR

}

func (virustotal virustotal) ScanFile(file string) (*json.Response, error) {

	f, e := os.Open(file)
	defer f.Close()
	if e != nil {
		return nil, e
	}
	return virustotal.ScanFileStream(f.Name(), f)

}

func (virustotal virustotal) ScanFileStream(filename string, reader io.Reader) (*json.Response, error) {
	if filename == "" {
		return nil, FILE_NAME_ERROR
	}

	if reader == nil {
		return nil, FILE_STREAM_ERROR
	}
	me := &multipart.MultipartEntity{}
	me.AddTextBody("apikey", virustotal.apiKey)
	me.AddBinaryBody("file", filename, reader)
	response, e := http.RequestPost(me, URL_FILE_SCAN)
	if e != nil {
		return nil, e
	}
	var report json.Response
	e = util.To(response, &report)
	if e != nil {
		return nil, e
	}
	return &report, e

}

func (virustotal virustotal) ScanURL(url string) (*json.Response, error) {
	me := &multipart.MultipartEntity{}
	me.AddTextBody("apikey", virustotal.apiKey)
	me.AddTextBody("url", url)
	response, e := http.RequestPost(me, URL_URL_SCAN)
	if e != nil {
		return nil, e
	}
	var report json.Response
	e = util.To(response, &report)
	if e != nil {
		return nil, e
	}
	return &report, e
}

func (virustotal virustotal) ReScanFile(resource string) (*json.ReScanResponse, error) {
	me := &multipart.MultipartEntity{}
	me.AddTextBody("apikey", virustotal.apiKey)
	me.AddTextBody("resource", resource)
	response, e := http.RequestPost(me, URL_FILE_RESCAN)
	if e != nil {
		return nil, e
	}
	var report json.ReScanResponse
	e = util.To(response, &report)
	if e != nil {
		return nil, e
	}
	return &report, e
}

func (virustotal virustotal) FileReport(resource string) (*json.Report, error) {
	me := &multipart.MultipartEntity{}
	me.AddTextBody("apikey", virustotal.apiKey)
	me.AddTextBody("resource", resource)
	response, e := http.RequestPost(me, URL_FILE_SCAN_REPORT)
	if e != nil {
		return nil, e
	}
	var report json.Report
	e = util.To(response, &report)
	if e != nil {
		return nil, e
	}
	return &report, e
}

func (virustotal virustotal) URLReport(resource string) (*json.URLReport, error) {
	me := &multipart.MultipartEntity{}
	me.AddTextBody("apikey", virustotal.apiKey)
	me.AddTextBody("resource", resource)
	me.AddTextBody("scan", "1")
	response, e := http.RequestPost(me, URL_URL_SCAN_REPORT)
	if e != nil {
		return nil, e
	}
	var report json.URLReport
	e = util.To(response, &report)
	if e != nil {
		return nil, e
	}
	return &report, e
}
func (virustotal virustotal) IPReport(ip string) (*json.IPReport, error) {
	response, e := http.RequestGet(virustotal.apiKey, URL_IP_SCAN_REPORT, "ip", ip)
	if e != nil {
		return nil, e
	}
	var report json.IPReport
	e = util.To(response, &report)
	if e != nil {
		return nil, e
	}
	return &report, e
}
func (virustotal virustotal) DomainReport(domain string) (*json.DomainReport, error) {
	response, e := http.RequestGet(virustotal.apiKey, URL_DOMAIN_REPORT, "domain", domain)
	if e != nil {
		return nil, e
	}
	var report json.DomainReport
	e = util.To(response, &report)
	if e != nil {
		return nil, e
	}
	return &report, e
}

func (virustotal virustotal) Comment(resource, comment string) (*json.CommentStatus, error) {
	me := &multipart.MultipartEntity{}
	me.AddTextBody("resource", resource)
	me.AddTextBody("comment", comment)
	me.AddTextBody("apikey", virustotal.apiKey)
	response, e := http.RequestPost(me, URL_COMMENTS)
	if e != nil {
		return nil, e
	}
	var report json.CommentStatus
	e = util.To(response, &report)
	if e != nil {
		return nil, e
	}
	return &report, e
}

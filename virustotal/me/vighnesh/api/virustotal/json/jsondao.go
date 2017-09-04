package json

import "fmt"

type VirusTotalResponse struct {
	ResponseCode int    `json:"response_code"`
	Message      string `json:"verbose_msg"`
}

type Response struct {
	VirusTotalResponse

	ScanId    string `json:"scan_id"`
	Sha1      string `json:"sha1"`
	Resource  string `json:"resource"`
	Sha256    string `json:"sha256"`
	Permalink string `json:"permalink"`
	Md5       string `json:"md5"`
}

type FileReport struct {
	Detected bool   `json:"detected"`
	Version  string `json:"version"`
	Malware  string `json:"result"`
	Update   string `json:"update"`
}

type Report struct {
	VirusTotalResponse
	Resource  string                `json:"resource"`
	ScanId    string                `json:"scan_id"`
	Sha1      string                `json:"sha1"`
	Sha256    string                `json:"sha256"`
	Md5       string                `json:"md5"`
	Scandate  string                `json:"scan_date"`
	Positives int                   `json:"positives"`
	Total     int                   `json:"total"`
	Permalink string                `json:"permalink"`
	Scans     map[string]FileReport `json:"scans"`
}

type DetectedUrl struct {
	ScanDate  string `json:"scan_date"`
	Url       string `json:"url"`
	Positives int    `json:"positives"`
	Total     int    `json:"total"`
}

type Resolution struct {
	LastResolved string `json:"last_resolved"`
	Hostname     string `json:"hostname"`
}

type IPReport struct {
	VirusTotalResponse
	Resolutions  []Resolution  `json:"resolutions"`
	DetectedUrls []DetectedUrl `json:"detected_urls"`
}

type DomainReport struct {
	VirusTotalResponse
	Resolutions  []Resolution  `json:"resolutions"`
	DetectedUrls []DetectedUrl `json:"detected_urls"`
}

type URLReport struct {
	Report
}

type ReScanResponse struct {
	Response
}

type CommentStatus struct {
	VirusTotalResponse
}

func (report *Response) String() string {
	return fmt.Sprintf("scanid: %s, resource: %s, permalink: %s, md5: %s", report.ScanId, report.Resource, report.Permalink, report.Md5)
}

func (report *ReScanResponse) String() string {
	return fmt.Sprintf("scanid: %s, resource: %s, permalink: %s, md5: %s", report.ScanId, report.Resource, report.Permalink, report.Md5)
}

package constants

import "net/url"

var kosikSearchUrl = url.URL{
	Scheme: "https",
	Host:   "www.kosik.cz",
	Path:   "/vyhledavani",
}

func GetKosikSearchUrl() url.URL {
	return kosikSearchUrl
}

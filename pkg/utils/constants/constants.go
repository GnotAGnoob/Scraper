package constants

import "net/url"

var KOSIK_SEARCH_ENDPOINT = "/vyhledavani"
var KOSIK_HOST = "www.kosik.cz"

func GetKosikUrl() url.URL {
	return url.URL{
		Scheme: "https",
		Host:   KOSIK_HOST,
	}
}

func GetKosikSearchUrl() url.URL {
	return url.URL{
		Scheme: "https",
		Host:   KOSIK_HOST,
		Path:   KOSIK_SEARCH_ENDPOINT,
	}
}

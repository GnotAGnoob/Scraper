package constants

import "net/url"

var kosikUrl = url.URL{
	Scheme: "https",
	Host:   "www.kosik.cz",
}

var KOSIK_SEARCH_ENDPOINT = "/vyhledavani"

func GetKosikUrl() url.URL {
	return kosikUrl
}

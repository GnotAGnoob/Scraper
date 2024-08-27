package constants

import "net/url"

const kosikSearchEndpoint = "/vyhledavani"
const kosikHost = "www.kosik.cz"

func GetKosikUrl() url.URL {
	return url.URL{
		Scheme: "https",
		Host:   kosikHost,
	}
}

func GetKosikSearchUrl() url.URL {
	return url.URL{
		Scheme: "https",
		Host:   kosikHost,
		Path:   kosikSearchEndpoint,
	}
}

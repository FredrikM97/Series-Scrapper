package utils

import (
	"strings"
)

func Address2string(address []string) string {
	/*
		Extract name from list to string and remove parameters that shouldnt be included
		Params:	[]string
		Return: string
	*/
	r := strings.NewReplacer("_", " ", "  ", " ")
	serieAddress := address[1]
	oldAddress := ""

	for oldAddress != serieAddress {
		oldAddress = serieAddress
		serieAddress = r.Replace(serieAddress)
	}

	serieAddress = strings.ToLower(serieAddress)
	return serieAddress
}

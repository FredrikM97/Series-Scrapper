package module

import "strings"

func address2string(address []string) string {
	/*
		Extract name from list to string and remove parameters that shouldnt be included
		Params:	[]string
		Return: string
	*/
	r := strings.NewReplacer("_", " ", "  ", " ")
	serie_address := address[1]
	old_address := ""

	for old_address != serie_address {
		old_address = serie_address
		serie_address = r.Replace(serie_address)
	}

	serie_address = strings.ToLower(serie_address)
	return serie_address
}

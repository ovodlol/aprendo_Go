package main

import (
	"github.com/k0kubun/pp/v3"
)

func main() {
	m := map[string]string{"Log": "Não", "Print bonito": "Sim"}
	scheme := pp.ColorScheme{
		Integer:         pp.Yellow,
		StructName:      pp.Magenta,
		StringQuotation: pp.Green | pp.Bold,
		String:          pp.Green | pp.Bold,
	}

	mypp := pp.New()
	mypp.SetExportedOnly(true)
	mypp.SetOmitEmpty(true)
	mypp.SetColorScheme(scheme)
	mypp.Println(m)
}

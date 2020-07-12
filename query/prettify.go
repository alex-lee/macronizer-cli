package query

import "regexp"

var accentRE = regexp.MustCompile(`[aeiouy][_^+]`)

var accentMap = map[string]string{
	"a_": "ā",
	"a^": "ă",
	"a+": "ä",

	"e_": "ē",
	"e^": "ĕ",
	"e+": "ë",

	"i_": "ī",
	"i^": "ĭ",
	"i+": "ï",

	"o_": "ō",
	"o^": "ŏ",
	"o+": "ö",

	"u_": "ū",
	"u^": "ŭ",
	"u+": "ü",

	"y_": "ȳ",
	"y^": "y̆",
	"y+": "ÿ",

	"A_": "Ā",
	"A^": "Ă",
	"A+": "Ä",

	"E_": "Ē",
	"E^": "Ĕ",
	"E+": "Ë",

	"I_": "Ī",
	"I^": "Ĭ",
	"I+": "Ï",

	"O_": "Ō",
	"O^": "Ŏ",
	"O+": "Ö",

	"U_": "Ū",
	"U^": "Ŭ",
	"U+": "Ü",

	"Y_": "Ȳ",
	"Y^": "Y̆",
	"Y+": "Ÿ",
}

func prettify(accented string) string {
	return accentRE.ReplaceAllStringFunc(accented, func(s string) string {
		if p, ok := accentMap[s]; ok {
			return p
		}
		return s
	})
}

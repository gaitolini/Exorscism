package leap

func IsLeapYear(year int) bool {
	// Um ano bissexto (no calendário gregoriano) ocorre:

	// Em cada ano isso é igualmente divisível por 4.
	// A menos que o ano seja igualmente divisível por
	//100, caso em que é apenas um ano bissexto
	//se o ano também for igualmente divisível por 400.
	// Alguns exemplos:

	// 1997 não foi um ano bissexto, pois não é divisível por 4.
	// 1900 não foi um ano bissexto, pois não é divisível por 400.
	// 2000 foi um ano bissexto!
	// 2004 foi um ano bissexto, pois é divisível por 4.

	return (year%400 == 0) || (year%100 != 0 && year%4 == 0)
}

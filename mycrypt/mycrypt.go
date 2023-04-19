package mycrypt

import "unicode"

var ALF_SEM03 []rune = []rune("abcdefghijklmnopqrstuvwxyzæøå0123456789.,:; ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func Krypter(melding []rune, alphabet []rune, chiffer int, decrypt bool) []rune {
	kryptertMelding := make([]rune, len(melding))
	for i := 0; i < len(melding); i++ {
		indeks := sokIAlfabetet(melding[i], alphabet)
		if decrypt {
			indeks -= chiffer
			if indeks < 0 {
				indeks += len(alphabet)
			}
		} else {
			indeks += chiffer
			if indeks >= len(alphabet) {
				indeks -= len(alphabet)
			}
		}
		kryptertMelding[i] = alphabet[indeks]
	}
	return kryptertMelding
}

func sokIAlfabetet(symbol rune, alfabet []rune) int {
	for i := 0; i < len(alfabet); i++ {
		if unicode.ToLower(symbol) == unicode.ToLower(alfabet[i]) {
			return i
		}
	}
	return -1
}

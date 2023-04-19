package mycrypt

import "unicode"

var ALF_SEM03 []rune = []rune("abcdefghijklmnopqrstuvwxyzæøå0123456789.,:; ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func Krypter(melding []rune, alphabet []rune, chiffer int) []rune {
	kryptertMelding := make([]rune, len(melding))
	for i := 0; i < len(melding); i++ {
		if melding[i] == ' ' {
			kryptertMelding[i] = rune(63)
			continue
		}
		indeks := sokIAlfabetet(melding[i], alphabet)
		if indeks+chiffer >= len(alphabet) {
			kryptertMelding[i] = alphabet[indeks+chiffer-len(alphabet)]
		} else {
			kryptertMelding[i] = alphabet[indeks+chiffer]
		}
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

func Dekrypter(melding []rune, alphabet []rune, chiffer int) []rune {
	dekryptertMelding := make([]rune, len(melding))
	for i := 0; i < len(melding); i++ {
		if melding[i] == rune(63) {
			dekryptertMelding[i] = ' '
			continue
		}
		indeks := sokIAlfabetet(melding[i], alphabet)
		if indeks-chiffer < 0 {
			dekryptertMelding[i] = alphabet[indeks-chiffer+len(alphabet)]
		} else {
			dekryptertMelding[i] = alphabet[indeks-chiffer]
		}
	}
	return dekryptertMelding
}

package mycrypt

var ALF_SEM03 []rune = []rune("abcdefghijklmnopqrstuvwxyzæøå0123456789.,:; ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func Krypter(melding []rune, alphabet []rune, chiffer int) []rune {
	kryptertMelding := make([]rune, len(melding))
	for i := 0; i < len(melding); i++ {
		// Convert any uppercase letters to lowercase
		if melding[i] >= 'A' && melding[i] <= 'Z' {
			melding[i] = melding[i] - 'A' + 'a'
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
	// Convert any uppercase letters to lowercase
	if symbol >= 'A' && symbol <= 'Z' {
		symbol = symbol - 'A' + 'a'
	}

	for i := 0; i < len(alfabet); i++ {
		if symbol == alfabet[i] {
			return i
		}
	}
	return -1
}

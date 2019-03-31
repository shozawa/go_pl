package ex06

import "unicode"

func DeleteSpace(str string) string {
	slice := []byte(str)
	for i := 0; i < len(slice)-1; i++ {
		a, b := slice[i], slice[i+1]
		if unicode.IsSpace(rune(a)) && unicode.IsSpace(rune(b)) {
			slice = remove(slice, i)
			i--
		}
	}
	return string(slice)
}

func remove(slice []byte, i int) []byte {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

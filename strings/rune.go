package main


func main() {
	s := "тест"
	println(s[0]) // t
	r := []rune(s)
	R := []rune{rune('R')}
	println(R)
	r = append(R, r[1:]...)

	println(string(r)) // Rест
}

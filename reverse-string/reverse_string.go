// Package reverse deals with all thing esrever
package reverse

// Reverse helps to return a reversed string
func Reverse(input string) string {
	length := len([]rune(input))
	reversed := ""
	for i := length - 1; i >= 0; i-- {
		reversed += string(input[i])
	}
	return reversed
}

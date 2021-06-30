/*
	other title:最大公约数和最小公倍数
*/

package math

// 最大公约数
func gcd(a, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}

// 最小公倍数
func lcd(a, b int) int {
	return a * b / gcd(a, b)
}

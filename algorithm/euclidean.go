package main

func gcd(a int, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}
func main() {
	println("%d", gcd(3425, 1233))
}

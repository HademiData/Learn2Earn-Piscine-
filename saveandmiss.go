package main

func SaveAndMiss(s string, n int) string {
	if n <= 0 {
		return s
	}

	ans := ""
	save := true
	for i := 0; i < len(s); i += n {
		end := i + n
		if end > len(s) {
			end = len(s)
		}

		if save {
			ans += s[i:end]
		}
		save = !save
	}
	return ans
}

func main() {
	println(SaveAndMiss("123456789", 3))
}

package arrayshashing

func IsSubsequence(s string, t string) bool {
	j, length := 0, 0

	for i := 0; i < len(t); i++ {
		if j < len(s) && t[i] == s[j] {
			length++
			j++
		}
	}

	return length == len(s)
}

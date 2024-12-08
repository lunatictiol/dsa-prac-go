package questions

func RepeatedCharacter(s string) byte {
	m := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		v, ok := m[s[i]]
		if ok {
			return v
		} else {
			m[s[i]] = s[i]
		}
	}

	return ' '

}

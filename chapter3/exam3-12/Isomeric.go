package comma

// Use bytes.Buffer
func Isomeric(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	mp := make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		mp[s1[i]]++
	}
	for i := 0; i < len(s1); i++ {
		mp[s2[i]]--
	}
	for _, v := range mp {
		if v != 0 {
			return false
		}
	}
	return true
}

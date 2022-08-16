package main

func isIsomorphic(s string, t string) bool {
	stmap, tmap := make(map[byte]byte), make(map[byte]byte)

	for i := range s {
		if val, ok := stmap[t[i]]; ok {
			if val != s[i] {
				return false
			}
		} else {
			stmap[t[i]] = s[i]
			if val, ok := tmap[s[i]]; ok && val != t[i] {
				return false
			}
		}
		tmap[s[i]] = t[i]
	}
	return true
}

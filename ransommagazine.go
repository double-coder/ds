package main

func canConstruct(r, m string) bool {

	lista := [26]int{}
	for _, v := range m {
		lista[v-'a']++
	}

	for i := 0; i < len(r); i++ {
		if lista[r[i]-'a'] == 0 {
			return false
		}

		lista[r[i]-'a']--
	}
	return true
}

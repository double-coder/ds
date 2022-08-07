package main

func lengthOfLongestSubstring(s string) int {
	max, mp := 0, make(map[byte]int)

	for i, j := 0, 0; i < len(s); i++ {
		if _, ok := mp[s[i]]; ok {
			j = maxi(j, mp[s[i]]+1)
		}
		mp[s[i]] = i
		max = maxi(max, i-j+1)
	}
	return max
}

func maxi(i, j int) int {
	if i > j {
		return i
	}
	return j
}

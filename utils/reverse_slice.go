package utils

//ReverseSlice utility function reverses order of the elements in a slice
func ReverseSlice(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

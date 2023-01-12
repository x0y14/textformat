package textformat

func format(in string) string {
	return string(tokenize([]rune(in)))
}

func contains(s1 []string, s2 string) bool {
	for _, s := range s1 {
		if s == s2 {
			return true
		}
	}
	return false
}

func isLowerAlpha(s string) bool {

}

func tokenize(in []rune) []rune {
	return nil
}

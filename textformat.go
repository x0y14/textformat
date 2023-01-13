package textformat

import (
	"fmt"
	"strconv"
	"strings"
)

type TokenKind int

const (
	Text TokenKind = iota
	Specifier
)

// Token 辞書型の代わり
type Token struct {
	Kind                 TokenKind
	Value                string
	CurrentPosition      int
	ReplaceZeroPositions []int
	Specifiers           []string
}

type SimpleToken struct {
	Kind  TokenKind
	Value string
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
	lowers := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	return contains(lowers, s)
}

func isUpperAlpha(s string) bool {
	uppers := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	return contains(uppers, s)
}

func isNumber(s string) bool {
	numbers := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	return contains(numbers, s)
}

func isWhite(s string) bool {
	whites := []string{"\n", "\t", " "}
	return contains(whites, s)
}

func isSharp(s string) bool {
	return s == "#"
}

func consumeSpecifier(in []string, pos int, replaceZero []int, specs []string) Token {
	var specifier string
	// #を消費
	// ここがzeroPos
	//replaceZero = append(replaceZero, pos)
	specifier += in[pos]
	pos++
	// 指定子本体、例えば文字列ならs, 整数ならd, 少数ならf, ...
	specs = append(specs, in[pos])
	specifier += in[pos]
	pos++

	return Token{
		Kind:            Specifier,
		Value:           specifier,
		CurrentPosition: pos,
		//ReplaceZeroPositions: replaceZero,
		Specifiers: specs,
	}
}

func consumeText(in []string, pos int) Token {
	t := in[pos]
	pos++
	return Token{
		Kind:            Text,
		Value:           t,
		CurrentPosition: pos,
	}
}

func tokenize(format string) ([]SimpleToken, []string) {
	input := strings.Split(format, "")
	pos := 0

	var replaceZeroPositions []int
	var specs []string

	var simpleTokens []SimpleToken

	var cur string
	for pos < len(input) {
		cur = input[pos]
		if isSharp(cur) {
			token := consumeSpecifier(input, pos, replaceZeroPositions, specs)
			simpleTokens = append(simpleTokens, SimpleToken{
				Kind:  token.Kind,
				Value: token.Value,
			})
			pos = token.CurrentPosition
			//replaceZeroPositions = token.ReplaceZeroPositions
			if len(token.Specifiers) != 0 {
				specs = token.Specifiers
			}
		} else {
			token := consumeText(input, pos)
			simpleTokens = append(simpleTokens, SimpleToken{
				Kind:  token.Kind,
				Value: token.Value,
			})
			pos = token.CurrentPosition
		}
	}

	return simpleTokens, specs
}

func format(in string, values []any) (string, error) {
	tokens, specs := tokenize(in)
	//replacePos := tokens[len(tokens)-1].ReplaceZeroPositions

	formatted := ""
	no := 0
	for _, tok := range tokens {
		if tok.Kind == Text {
			formatted += tok.Value
		} else {
			//formatted += values[no]
			var repVal string
			// 個別にキャストしてかなきゃいけない
			if specs[no] == "d" {
				repVal = strconv.Itoa(values[no].(int))
			} else if specs[no] == "f" {
				repVal = strconv.FormatFloat(values[no].(float64), 'f', -1, 64)
			} else if specs[no] == "s" {
				repVal = values[no].(string)
			} else {
				return "", fmt.Errorf("unsupported specifier: #%v <- %v", specs[no], values[no])
			}
			formatted += repVal
			no++
		}
	}

	return formatted, nil
}

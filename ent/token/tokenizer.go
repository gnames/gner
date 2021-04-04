package token

// space chars that indicate new line have value true
var spaceChr = map[rune]bool{
	'\n': true,
	'\r': true,
	'\v': false,
	'\t': false,
	' ':  false,
}

// Tokenize creates a slice containing tokens for every word in the document.
func Tokenize(text []rune) []Token {
	var line int
	var res []Token
	start := 0
	dashToken := Token{}
	for i, v := range text {
		if _, ok := spaceChr[v]; ok {
			if _, isSpace := spaceChr[text[start]]; !isSpace {
				t := NewToken(text[start:i], start, i)
				t.Line = line
				if dashToken.Start > 0 {
					t = concatenateTokens(dashToken, t)
					dashToken.Start = 0
					res = append(res, t)
				} else {
					if lineEndsWithDash(text, i, t) {
						dashToken = t
					} else {
						res = append(res, t)
					}
				}
			}
			start = i + 1
			if v == '\n' {
				line++
			}
		}
	}
	if len(text)-start > 0 {
		t := NewToken(text[start:], start, len(text))
		t.Line = line
		res = append(res, t)
	}
	return res
}

func lineEndsWithDash(text []rune, i int, t Token) bool {
	dash := rune('-')
	l := len(t.Raw)
	if l > 1 && t.Raw[l-1] == dash && lastWordForLine(text, i) {
		return true
	}
	return false
}

func lastWordForLine(text []rune, i int) bool {
	for {
		if isNewLineChr, ok := spaceChr[text[i]]; ok {
			if isNewLineChr {
				return true
			}
		} else {
			return false
		}
		i++
		if i >= len(text) {
			return false
		}
	}
}

func concatenateTokens(t1 Token, t2 Token) Token {
	var v []rune
	t1Raw := make([]rune, len(t1.Raw))
	copy(t1Raw, t1.Raw)
	if t2.Raw[0] >= rune('a') && t2.Raw[0] <= rune('z') {
		v = append(t1Raw[0:len(t1Raw)-1], t2.Raw...)
	} else {
		v = append(t1Raw, t2.Raw...)
	}
	t := NewToken(v, t1.Start, t2.End)
	t.Line = t1.Line
	return t
}

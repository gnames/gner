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
func Tokenize(text []rune, wrapToken func(TokenNER) TokenNER) []TokenNER {
	var line int
	var res []TokenNER
	start := 0
	dashToken := tokenNER{}
	for i, v := range text {
		if _, ok := spaceChr[v]; ok {
			if _, isSpace := spaceChr[text[start]]; !isSpace {
				t := tokenNER{
					raw:   text[start:i],
					start: start,
					end:   i,
				}
				t.line = line
				if dashToken.start > 0 {
					t := concatenateTokens(dashToken, t)
					dashToken.start = 0
					res = addToken(res, &t, wrapToken)
				} else {
					if lineEndsWithDash(text, i, t) {
						dashToken = t
					} else {
						res = addToken(res, &t, wrapToken)
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
		var t tokenNER = tokenNER{
			line:  line,
			raw:   text[start:],
			start: start,
			end:   len(text),
		}
		res = addToken(res, &t, wrapToken)
	}
	return res
}

func addToken(
	tokens []TokenNER,
	token TokenNER,
	wrapToken func(TokenNER) TokenNER,
) []TokenNER {
	tWrapped := wrapToken(token)
	tWrapped.ProcessRaw()
	tokens = append(tokens, tWrapped)
	return tokens
}

func lineEndsWithDash(text []rune, i int, t tokenNER) bool {
	dash := rune('-')
	l := len(t.raw)
	if l > 1 && t.raw[l-1] == dash && lastWordForLine(text, i) {
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

func concatenateTokens(t1 tokenNER, t2 tokenNER) tokenNER {
	var v []rune
	t1Raw := make([]rune, len(t1.raw))
	copy(t1Raw, t1.raw)
	if t2.raw[0] >= rune('a') && t2.raw[0] <= rune('z') {
		v = append(t1Raw[0:len(t1Raw)-1], t2.raw...)
	} else {
		v = append(t1Raw, t2.raw...)
	}
	t := tokenNER{
		raw:   v,
		start: t1.start,
		end:   t2.end,
	}
	t.line = t1.line
	return t
}

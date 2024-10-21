package lexer

import "photon/src/token"

type FileByte byte

func (b *FileByte) isDelimiter() bool {
	delimits := []byte{';', ' ', '\n', '\t'}
	for _, d := range delimits {
		if byte(*b) == d {
			return true
		}
	}
	return false
}

func Tokenize(fileContent []byte) ([]token.Token, error) {
	var tokens []token.Token
	end := 0

	for start := 0; start < len(fileContent); {
		if start > end {
			end = start
		}

		if fileContent[start] == ' ' {
			start++
			continue
		}

		fb := (*FileByte)(&fileContent[start])
		if fb.isDelimiter() {
			delimiter := string(fileContent[start])
			tokenKey := token.Lookup(delimiter)

			if tokenKey == token.T_ZERO_MEASURE {
				return nil, &token.TokenizeError{Message: "Invalid delimiter: " + delimiter}
			}

			if tokenKey != token.T_SPACE {
				tokenObj := token.Token{}
				tokenObj.New(tokenKey, delimiter)
				tokens = append(tokens, tokenObj)
			}

			start++
			end = start
			continue
		}

		for end < len(fileContent) {
			fb = (*FileByte)(&fileContent[end])
			if fb.isDelimiter() {
				break
			}
			end++
		}

		if start < end {
			word := string(fileContent[start:end])
			tokenKey := token.Lookup(word)

			if tokenKey == token.T_ZERO_MEASURE {
				return nil, &token.TokenizeError{Message: "Invalid token: " + word}
			}

			if tokenKey != token.T_SPACE {
				tokenObj := token.Token{}
				tokenObj.New(tokenKey, word)
				tokens = append(tokens, tokenObj)
			}

			start = end
		}
	}

	return tokens, nil
}

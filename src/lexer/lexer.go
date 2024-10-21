package lexer

import "photon/src/token"

type FileByte byte

var delimiters = []byte{';', ' ', '\n', '\t'}
var phonyTokens = []token.TType{token.T_SPACE, token.T_NEWLINE, token.T_EOF}
var variableAntecedentTokens = []token.TType{token.KW_LET, token.KW_CONST, token.KW_FUNC, token.KW_RETURN}

func (b *FileByte) isDelimiter() bool {
	for _, d := range delimiters {
		if byte(*b) == d {
			return true
		}
	}
	return false
}

func isPhonyToken(t *token.TType) bool {
	for _, pt := range phonyTokens {
		if *t == pt {
			return true
		}
	}
	return false
}

func isVariableAntecedentToken(t *token.TType) bool {
	for _, vat := range variableAntecedentTokens {
		if *t == vat {
			return true
		}
	}
	return false
}

func getToken(word string, prevToken *token.Token) (*token.Token, error) {

	if prevToken != nil && isVariableAntecedentToken(&prevToken.Type) {
		tokenObj := token.Token{}
		tokenObj.New(token.T_IDENT, word)
		return &tokenObj, nil
	}

	tokenKey := token.Lookup(word)

	if tokenKey == token.T_ZERO_MEASURE {
		return nil, &token.TokenizeError{Message: "Invalid token: " + word}
	}

	if !isPhonyToken(&tokenKey) {
		tokenObj := token.Token{}
		tokenObj.New(tokenKey, word)
		return &tokenObj, nil
	}

	return nil, nil

}

func Tokenize(fileContent []byte) ([]token.Token, error) {
	var tokens []token.Token
	var prevToken *token.Token = nil
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
			tokenObj, err := getToken(delimiter, prevToken)
			if err != nil {
				return nil, err
			}
			if tokenObj != nil {
				tokens = append(tokens, *tokenObj)
			}
			prevToken = tokenObj
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
			tokenObj, err := getToken(word, prevToken)
			if err != nil {
				return nil, err
			}
			if tokenObj != nil {
				tokens = append(tokens, *tokenObj)
			}
			prevToken = tokenObj
			start = end
		}
	}

	return tokens, nil
}

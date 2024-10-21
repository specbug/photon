package lexer

import "photon/src/token"

type FileByte byte

var phonyTokens = []token.TType{token.T_SPACE, token.T_NEWLINE, token.T_EOF}
var variableAntecedentTokens = []token.TType{token.KW_LET, token.KW_CONST, token.KW_FUNC, token.KW_RETURN, token.SYM_LPAREN, token.SYM_ASSIGN, token.SYM_COMMA}

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
		word := string(*fb)
		charTokenType := token.LookupMap[word]
		if charTokenType != token.T_ZERO_MEASURE {
			tokenObj, err := getToken(word, prevToken)
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
			charTokenType = token.LookupMap[string(*fb)]
			if charTokenType != token.T_ZERO_MEASURE {
				break
			}
			end++
		}

		if start < end {
			word = string(fileContent[start:end])
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

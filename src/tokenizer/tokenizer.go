package tokenizer

import (
	"fmt"
	"strconv"
)

type TokenKey string
type FileByte byte
type TokenType struct {
	Key   TokenKey
	Value interface{} // TODO: Change this to a more specific type
}

const (
	P_EXIT     = "EXIT"
	P_SEMI     = "SEMI"
	P_INT_LTRL = "INT_LTRL"
	P_SPACE    = "SPACE"
)

func (tt *TokenType) String() string {
	fmtString := ""
	if tt.Value != nil {
		fmtString = fmt.Sprintf("{Key: %s, Value: %v}", tt.Key, tt.Value)
	} else {
		fmtString = fmt.Sprintf("{Key: %s}", tt.Key)
	}
	return fmtString
}

func mapStringToTokenKey(s *string) TokenKey {
	switch *s {
	case ";":
		return P_SEMI
	case "exit":
		return P_EXIT
	case " ":
		return P_SPACE
	default:
		// TODO: Support Unicode characters
		if _, err := strconv.Atoi(*s); err == nil {
			return P_INT_LTRL
		}
	}
	return ""
}

func (tt *TokenType) new(key *TokenKey, value *string) {
	tt.Key = *key

	switch *key {
	case P_INT_LTRL:
		intValue, _ := strconv.Atoi(*value)
		tt.Value = intValue
	default:
		tt.Value = nil
	}

}

func (b *FileByte) isDelimiter() bool {
	delimits := []byte{';', ' ', '\n', '\t'}
	for _, d := range delimits {
		if byte(*b) == d {
			return true
		}
	}
	return false
}

type TokenizeError struct {
	Message string
}

func (te *TokenizeError) Error() string {
	return te.Message
}

func Tokenize(fileContent *[]byte) ([]TokenType, error) {
	var tokens []TokenType
	var token *TokenType
	var tokenKey TokenKey
	var fb *FileByte
	var word string
	var start, end int32 = 0, 0

	for i, b := range *fileContent {
		fb = (*FileByte)(&b)
		if fb.isDelimiter() {
			word = string((*fileContent)[start : end+1])
			tokenKey = mapStringToTokenKey(&word)
			if tokenKey == "" {
				return nil, &TokenizeError{Message: "Invalid token: " + word}
			}
			if tokenKey != P_SPACE {
				token = &TokenType{}
				token.new(&tokenKey, &word)
				tokens = append(tokens, *token)
			}
			if *fb == ' ' {
				start = int32(i) + 1
			} else {
				start = int32(i)
			}
		}
		end = int32(i)
	}

	if start < int32(len(*fileContent)) {
		word = string((*fileContent)[start:])
		tokenKey = mapStringToTokenKey(&word)
		if tokenKey == "" {
			return nil, &TokenizeError{Message: "Invalid token: " + word}
		} else if tokenKey == P_SPACE {
			return tokens, nil
		}
		token = &TokenType{}
		token.new(&tokenKey, &word)
		tokens = append(tokens, *token)
	}

	return tokens, nil
}

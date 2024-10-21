package tokenizer

import (
	"fmt"
	"strconv"
)

type TokenKey int32
type FileByte byte
type TokenType struct {
	Key   TokenKey
	Value interface{} // TODO: Change this to a more specific type
}

const (
	_PEXIT TokenKey = iota
	_PSEMI
	_PINT_LTRL
	_PSPACE
)

func (tk *TokenKey) String() string {
	return [...]string{"_PEXIT", "_PSEMI", "_PINT_LTRL", "_PSPACE"}[*tk]
}

func (tt *TokenType) String() string {
	fmtString := ""
	if tt.Value != nil {
		fmtString = fmt.Sprintf("{Key: %s, Value: %v}", &tt.Key, tt.Value)
	} else {
		fmtString = fmt.Sprintf("{Key: %s}", &tt.Key)
	}
	return fmtString
}

func mapStringToTokenKey(s *string) TokenKey {
	switch *s {
	case ";":
		return _PSEMI
	case "exit":
		return _PEXIT
	case " ":
		return _PSPACE
	default:
		// TODO: Support Unicode characters
		if _, err := strconv.Atoi(*s); err == nil {
			return _PINT_LTRL
		}
	}
	return -1
}

func (tt *TokenType) new(key *TokenKey, value *string) {
	tt.Key = *key

	switch *key {
	case _PINT_LTRL:
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
			if tokenKey == -1 {
				return nil, &TokenizeError{Message: "Invalid token: " + word}
			}
			if tokenKey != _PSPACE {
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
		if tokenKey == -1 {
			return nil, &TokenizeError{Message: "Invalid token: " + word}
		} else if tokenKey == _PSPACE {
			return tokens, nil
		}
		token = &TokenType{}
		token.new(&tokenKey, &word)
		tokens = append(tokens, *token)
	}

	return tokens, nil
}

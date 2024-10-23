package main

import (
	"os"
	"photon/src/lexer"
	"strings"
	"testing"
)

func loadLexerTestFile(filename string) ([]byte, []string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")

	var codeLines []string
	var expectedTokens []string
	parsingTokens := false

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "// Expected tokens:") {
			parsingTokens = true
			continue
		} else if parsingTokens {
			tokensSplit := strings.Split(strings.TrimPrefix(line, "// "), "} ")
			for _, token := range tokensSplit {
				if !strings.HasSuffix(token, "}") {
					token += "}"
				}
				expectedTokens = append(expectedTokens, string(strings.TrimSpace(token)))
			}
		} else if line != "" {
			codeLines = append(codeLines, line)
		}
	}

	code := []byte(strings.Join(codeLines, "\n"))
	return code, expectedTokens
}

func TestLexer(t *testing.T) {
	// load all files from photon/test/lexer
	// and append them to testFiles
	testFiles, err := os.ReadDir("../test/lexer")
	if err != nil {
		t.Fatal(err)
	}

	for _, file := range testFiles {
		filename := file.Name()
		if !strings.HasSuffix(filename, ".hv") {
			continue
		}

		println("=====================================")
		println("Testing lexer for file:", filename)
		code, expectedTokens := loadLexerTestFile("../test/lexer/" + filename)
		println("Code:")
		println(string(code))
		println("Expected tokens:")
		for _, token := range expectedTokens {
			println(token)
		}

		outputTokens, err := lexer.Tokenize(code)
		if err != nil {
			t.Fatal(err)
		}

		println("Output tokens:")
		for _, token := range outputTokens {
			if token.Literal != "" {
				println("{" + string(token.Type) + ", " + string(token.Literal) + "}")
			} else {
				println("{" + string(token.Type) + "}")
			}
		}

		if len(outputTokens) != len(expectedTokens) {
			t.Fatalf("Expected %d tokens, got %d", len(expectedTokens), len(outputTokens))
		}

		for i, token := range outputTokens {
			sToken := ""
			if token.Literal != "" {
				sToken = "{" + string(token.Type) + ", " + string(token.Literal) + "}"
			} else {
				sToken = "{" + string(token.Type) + "}"
			}
			if sToken != expectedTokens[i] {
				t.Fatalf("Expected token %d to be %s, got %s", i, expectedTokens[i], sToken)
			}
		}

		println("=====================================")
		println()
	}
}

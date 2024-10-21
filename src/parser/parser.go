package parser

type Parser interface {
	Parse() error
}

type BaseParser struct {
	FileContent *[]byte
}

func (bp *BaseParser) Parse() error {
	// Print the file content
	println("File content:")
	println(string(*bp.FileContent))
	return nil
}

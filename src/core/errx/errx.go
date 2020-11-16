package errx

import "fmt"

type errorCustom struct {
	code string 
	message string
}

func (e errorCustom) Error() string {
	//TODO: implement internationalized messages
	return fmt.Sprintf("(%s) %s", e.code, e.message)
}

func New(code string, message string, args ...interface{}) *errorCustom {
	formatted := fmt.Sprintf(message, args)
	return &errorCustom{code: code, message: formatted}
}

var (
	ErrTaxonomyAlreadyExists = New("TxAlEx", "Taxonomy already exists")
	ErrWrongSpecie = New("WroSpe", "Specie must be specified")
	ErrWrongGenus = New("WroGen", "Genus must be specified")
)
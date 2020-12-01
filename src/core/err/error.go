package err

import "fmt"

type BasicError struct {
	code string
	message string
}

type GatewayError struct {
	BasicError
	DevMessage string
}

// Validation creates basic error with code and default message.
func Basic(code string, message string, args ...interface{}) *BasicError {
	formatted := fmt.Sprintf(message, args)
	return &BasicError{
		code:    code,
		message: formatted,
	}
}

// Gateway creates gateway error with dev message
func Gateway(code string, message string, devMessage string) *GatewayError {
	return &GatewayError{
		BasicError: BasicError{
			code:    code,
			message: message,
		},
		DevMessage: devMessage,
	}
}

func (e BasicError) Error() string {
	return e.message
}



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
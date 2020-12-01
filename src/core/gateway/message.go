package gateway

type MessageProvider interface {
	Get(code string, defaultMessage string, args ...map[string]string) string
}

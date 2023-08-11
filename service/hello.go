package service

// Hello パッケージの外部に公開する関数
func Hello(message string) string {
	return hello(message)
}

// hello パッケージ内でのみ利用できる関数
func hello(message string) string {
	return "Hello " + message
}

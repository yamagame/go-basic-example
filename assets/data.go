// Package assets 埋め込みテキストデータ
package assets

import (
	"embed"
)

//go:embed hello.txt
var HelloTextBytes []byte

//go:embed *.txt
var EmbedTextFile embed.FS

//go:embed *.csv
var EmbedCSVFile embed.FS

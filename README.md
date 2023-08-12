# Go言語の基本

## Goプロジェクトの作成

プロジェクトを配置するディレクトリを作成し、そのディレクトリをカレントディレクトリにして以下のコマンドを実行する。

```sh
$ go mod init sample/go-basic-example
```
上記例ではパッケージ名を「sample/go-basic-example」としている。パッケージ名はプロジェクトに合わせて任意の名前を指定できる。

## 実行とビルド

```sh
# ビルドせず実行
$ go run ./main.go

# ビルド
$ go build -o ./build/main ./main.go
```

## パッケージ

### パッケージ名の基本

```go
package main   // <=== main パッケージの持ち物

import (
	"fmt"

	"sample/go-basic-example/service"
)

func main() {
	fmt.Println(service.Hello("World"))
}
```

Go言語のソースは必ず package を指定する。package はそのソースが入っているディレクトリ名。一つのディレクトリに異なるパッケージを含めることはできない。
例外的にUT用の xxxx_test パッケージは含めることができる。

パッケージには必ずパッケージの説明文をコメントとして記載する。

```go
// Package service アプリケーションのサービス
package service
```

説明文の配置はそのパッケージのソースであればどのソースでもよい。慣例としてパッケージの説明文用の doc.go ソースを作成することがある。

```sh
$ ls -1 ./service
doc.go        # <=== パッケージの説明を記入するソース
hello.g
```

外部に公開する関数やクラス、定数は名前を大文字から記載する。

```go
package service

// Hello パッケージの外部に公開する関数
func Hello(message string) string {
	return hello(message)
}

// hello パッケージ内でのみ利用できる関数
func hello(message string) string {
	return "Hello " + message
}
```

小文字から始まる名前でも、同一パッケージ内ではアクセス可能であることに注意。

### パッケージ名の命名規則

よいパッケージ名は以下の３つの条件に従う。

- 短く
- 関係
- 明快

パッケージ名は「小文字」の「1単語」であり、「priority_queue」や「computeServiceClient」のように区切り文字を使用してはいけない。
区切り文字を使用している場合はLintでエラーになることもあるため、区切り文字を使用する必要が出てきる場合はパッケージの構成を見直す必要がある。

参考：[Golangのパッケージ名はどうするのが正解か](https://zenn.dev/_kazuya/scraps/fdc65096b0d1d7)

パッケージ名に「-」区切りの名前は使用できない。

util や common、misc などの用途がはっきりしないパッケージ名は利用してはいけない。

例えば、

```go
package util
func NewStringSet(...string) map[string]bool {...}
func SortStringSet(map[string]bool) []string {...}
```

の場合、クライアントコードは

```go
set := util.NewStringSet("c", "a", "b")
fmt.Println(util.SortStringSet(set))
```

となるが、

```go
package stringset
func New(...string) map[string]bool {...}
func Sort(map[string]bool) []string {...}
```

であれば、以下のようにより簡潔に記述することができる。

```go
set := stringset.New("c", "a", "b")
fmt.Println(stringset.Sort(set))
```

参考：[よくないパッケージ名](https://zenn.dev/link/comments/e6b6b90b422a0d)

また、以下のようにパッケージ名と同じ名前のプレフィックスを付けるとLintエラーとなる。

```go
package stringset
func StringSetNew(...string) map[string]bool {...}
func StringSetSort(map[string]bool) []string {...}
```


### パッケージのインポート

同一ディレクトリ外のソースコードを import するときは、必ずフルパスで指定する。相対パスは使用できない。

```go
package main

import (
	"fmt"

	"sample/go-basic-example/service"   // <=== パッケージ名を含むフルパス
)

func main() {
	fmt.Println(service.Hello("World"))
}
```

## 変数と定数

```go
// 定数
const constant = "定数"

// 変数の定義と代入
variable := "変数の作成と代入"

// 変数の代入
variable = "変数の代入"

// 変数の定義
var defval string
defval = "定義と代入"
```

## 構造体

Go言語にはクラスがない。代わりに構造体をクラスのように利用する。

```go
// Person 構造体定義の基本
type Person struct {
    Name string     // パブリック変数
    Age int
    secret string   // プライベート変数
}

// インスタンス化
person := Person{Name:"山田太郎",Age: 23}
personptr := &Person{Name:"山田太郎",Age: 23}

// Runner Runner構造体 Person構造体を継承
type Runner struct {
    Person
    Speed int
}

// インスタンス化
runner := Runner{
    Person: Person{Name: "山田太郎", Age: 23},
    Speed:  13,
}
runnerptr := &Runner{
    Person: Person{Name: "山田太郎", Age: 23},
    Speed:  13,
}
```

```go
// Say パブリックメソッド
func (x *Person) Say(text string) string {
    return x.Name + ":" + text
}

// プライベートメソッド
func (x *Person) say(text string) string {
    return x.Name + ":" + text
}

// SayAll クラスメソッド
func (Person) SayAll(text string) string {
    return "Person:" + text
}
```

```go
// 無名構造体
person := struct {
    Name string
    Age  int
}{
    Name: "山田太郎",
    Age:  23,
}
```

## go:embed

参考：[Go 1.16からリリースされたgo:embedとは](https://future-architect.github.io/articles/20210208/)

embed はファイルをそのままビルドされたバイナリに埋め込み、ファイルや定数として扱える機能である。

```sh
# assets ディレクトリに data.go と hello.txt を配置
$ ls -1 assets
data.go
hello.txt
```

data.go は下記の通り。コメントの「//go:embed hello.txt」が埋め込みファイルであり、次の HelloTextBytes が embed hello.txt を指すデータになる。

```go
// Package assets 埋め込みテキストデータ
package assets

import (
	_ "embed"
)

//go:embed hello.txt
var HelloTextBytes []byte
```

hello.txt は「Embed Hello World!」テキストを含むファイル。HelloTextBytes がこのテキストデータになる。

```text
Embed Hello World!
```

HelloTextBytes の使用例。

```go
package main

import (
	"fmt"

	"sample/go-basic-example/assets"
)

func main() {
	fmt.Println(string(assets.HelloTextBytes))
}
```

変数型を embed.FS にするとファイルとして扱うことができる。

```go
// Package assets 埋め込みテキストデータ
package assets

import "embed"

//go:embed *.txt
var EmbedFile embed.FS
```

```go
package main

import (
	"fmt"

	"sample/go-basic-example/assets"
)

func main() {
    // hello.txt をファイルとして読み込む
	data, err := assets.EmbedFile.ReadFile("hello.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
```

## Goroutine と Channel

参考：[【Go言語入門】goroutineとは？ 実際に手を動かしながら goroutineの基礎を理解しよう！](https://www.ariseanalytics.com/activities/report/20221005/)  
参考：[初心者がGo言語のcontextを爆速で理解する ~ cancel編　~](https://qiita.com/yoshinori_hisakawa/items/a6608b29059a945fbbbd)  
参考：[Go言語でチャネルとselect](https://qiita.com/najeira/items/71a0bcd079c9066347b4)

関数呼び出しの前に go を付けることで並行処理(goroutine)となる。

```go
package main

import (
  "fmt"
)

func Say(s string) {
  fmt.Println(s)
}

func main() {
  go Say("hello")
  go Say("world")
}
```

goroutine 間でデータのやりとりを行う場合は channel を使用する。

```go
package main

import "fmt"

func Say(s string, ch chan string) {
    // channel にデータを入れる
	ch <- s
}

func main() {
    // channel の作成、扱うデータ型は string 型
	ch := make(chan string)

    go Say("hello", ch)
	go Say("world", ch)

	var ret string
    // channel からデータを取り出す
	ret = <-ch
	fmt.Println(ret)
    // channel からデータを取り出す
	ret = <-ch
	fmt.Println(ret)
}
```

## Context

Context を使用すると Channel を利用せずに簡単に goroutine 間で情報の伝達ができる。

参考：[よくわかるcontextの使い方](https://zenn.dev/hsaki/books/golang-context)

context を使用した値の受け渡しは goroutine セーフである。

参考：[Goでスレッド（goroutine）セーフなプログラムを書くために必ず注意しなければいけない点](https://qiita.com/ruiu/items/54f0dbdec0d48082a5b1)

context.WithValue() で値を設定し、ctx.Value() で取り出す。下記に例を示す。
context.WithValue() の第２引数はキーであるが、独自の型定義をしないと Lint エラーとなる。また、context は関数の第１引数で渡さなければ、こちらも Lint エラーとなる。

```go
package main

import (
	"context"
	"fmt"
)

type myContextKey int

const (
	ContextKeyName myContextKey = iota
)

func Say(ctx context.Context, s string, ch chan string) {
	// channel にデータを入れる
	ch <- s + "+" + ctx.Value("name").(string)
}

func main() {
	// channel の作成、扱うデータ型は string 型
	ch := make(chan string)

	ctx := context.Background()
	ctx = context.WithValue(ctx, ContextKeyName, "sample app")

	go Say(ctx, "hello", ch)
	go Say(ctx, "world", ch)

	var ret string
	// channel からデータを取り出す
	ret = <-ch
	fmt.Println(ret)
	// channel からデータを取り出す
	ret = <-ch
	fmt.Println(ret)
}
```

## Generics

参考：[Goの標準ライブラリに学ぶジェネリクス](https://gihyo.jp/article/2023/05/tukinami-go-07)

下記例の IntValue[T] 構造体は int、int32、int64 が利用できる。

```go
package main

import (
	"fmt"
)

type IntValue[T int | int32 | int64] struct {
	Value T
}

func NewIntValue[T int | int32 | int64](in T) IntValue[T] {
	return IntValue[T]{
		Value: in,
	}
}

func (x *IntValue[T]) Add(v T) T {
	x.Value = x.Value + v
	return x.Value
}

func main() {
	intval := NewIntValue[int](3)
	fmt.Println(intval.Add(5))

	int32val := NewIntValue[int32](3)
	fmt.Println(int32val.Add(5))

	int64val := NewIntValue[int64](3)
	fmt.Println(int64val.Add(5))
}
```

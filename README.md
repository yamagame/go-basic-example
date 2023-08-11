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

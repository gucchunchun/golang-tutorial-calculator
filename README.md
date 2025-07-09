# golang-tutorial-calculator


## 学習内容

### bufioパッケージ
バッファ付きI/Oを扱うための標準ライブラリ
キーボード入力やファイルなどから行単位・単語単位で読み取るために使用する

#### 使用例
```
scanner := bufio.NewScanner(os.Stdin)
scanner.Scan() // 1行読み取る
text := scanner.Text() // 入力された文字列を取得
```

#### メリット
- `fmt.Scanln()`はスペース区切りに弱く、空白や改行で止まってしまう可能性がある
- `bufio.Scanner`は改行までを一行全体として扱うことが可能


### 標準入力/出力

`os.Stdin`
: ユーザーの標準入力（キーボードなど）を受け取る

`os.Stdout`
: ユーザーに出力を（ターミナルなどで）見せる

`io.MultiWriter`
: 複数の出力先に同時に書き出すことが可能


### ファイル操作

`os.Create`
: 作成、または既存のファイルを上書き。読み取り不可。

`os.OpenFile`
: 読み取り・上書き・追記を選択することが可能。

`os.OpenFile`の引数の説明

| flag | 説明 |
| --- | --- |
| `os.O_RDONLY` | 読み取り専用 |
| `os.O_WRONLY` | 書き込み専用 |
| `os.O_RDWR` | 読み書き |
| `os.O_APPEND` | ファイルの最後に追記する |
| `os.O_CREATE` | ファイルが存在しない場合、作成する |
| `os.O_TRUNC` | ファイルが存在する場合、元内容をクリアする |
| `os.O_EXCL` | ファイルが存在する場合、エラーが起こる |




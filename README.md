# go-weather

golangでlivedoorのお天気APIを叩いて結果を取得する処理を書いてみるテスト

### 要素

- HTTP Clientで外部APIをキック
- JSONパース
- 文字コード変換
- ファイル出力

### 使い方

```
$ git clone git@github.com:wb55yokota/go-weather.git
$ cd go-weather
$ dep ensure
$ go build

$ ./go-weather
東京の天気
明日 (雨のち曇) 最高：31度 / 最低：26度

$ cat /tmp/output.txt
?????̓V?C
???? (?J?̂???) ?ō??F31?x / ?Œ?F26?x

（ShiftJISなので環境によっては文字化けする）
```
# grpc-practice

## gRPC

### RFC（Remote Procedure Call）とは
- 外部のプログラムが提供するProcedureを呼び出す仕組み
- 異なる言語間でRFCとして呼び出す関数を定義するために、特定の言語に依らない定義方法（IDL：Interface Description Language）が必要になる
- IDLにはコンパイラが付属しており、サポートしている言語用に関数を呼び出すためのコードを自動生成できる

### gRPCについて

#### 概要
- 異なる言語で書かれたアプリケーション同士がgRPCにより自動生成されたインターフェースを通じて通信することが可能になる
- フロントは別マシーンのメソッドを簡単に呼び出せるようになるのでマイクロサービスを簡単に構築できる
- サーバー側はインターフェースを実装し、フロント側はサーバーと同じメソッドを提供するスタブを持つ

![grpc](https://user-images.githubusercontent.com/53222150/131676182-fab89e6d-6a45-4fb2-a7d6-ff14efb1321f.png)

#### 特徴
- RFCの実装の1つで、IDLにProtocol Buffersを採用
  - バイナリ形式のシリアライズフォーマットなので、JSONなどのようなテキスト形式より処理が軽い
- 通信プロトコルとしてHTTP2を採用
  - ストリーム処理を提供
    - クライアント・サーバー間のリクエストとレスポンスが1:1だけでなくN:Nも可能になる
    - 詳しい内容についてはこちらを参照
  　  https://qiita.com/tomo0/items/310d8ffe82749719e029j
- リクエストチェーン全体に渡るタイムアウトやキャンセル処理をプロトコルレベルでサポート
  - 複数のマイクロサービスを跨ぐリクエストを制御するのに長けている

#### 実装
- protoファイルにデータの構造やサービスを定義
- protoc（コンパイラ）で、シリアライズなどを含むデータアクセスするためのコードを好きな言語で自動生成
  - Protocol Buffersのメッセージやシリアライズに protocolbuffers/protobuf-go のprotoc-gen-go
  - gPRCのサーバ/クライアントに grpc/grpc-goのprotoc-gen-go-grpc


## 並行処理

### goroutineについて
golangのプログラムで並行に実行されるもの

関数やメソッドの呼び出しの前にgoを付けると、異なるgoroutineで関数を実行することができる

runtime.NumGroutine()で現在起動しているgoroutineの数を取得できる

メインgoroutineが終わったら他のgoroutineの終了を待たずにプログラム全体が終わる.
これを防ぐために以下の2つの方法がある.
- sync.WaitGroupを使う
- channelを使う

### channelについて
channelを使うことで異なるgoroutine間で連携できる

channelとgoroutineを使うことで、「何かデータを受信するまで待って、受信したら処理を開始する」というような処理を簡単に実装できる

```
ch := make(chan [型])
ch := make(chan [型], バッファサイズ)
```

```
ch := make(chan<- 型) // 読み込み専用
ch := make(<-chan 型) // 書き込み専用
```


```
ch <- data // 送信
var := <-ch // 受信
```

受信側では受信可能なデータが来るまでブロック（goroutineが停止）される

送信側はchannelがいっぱいの場合、空きができるまでブロックされる

channelはデータを書き込む送信側と受け取る受信側が存在していないとエラーになる

channelはcloseすることでデータを読み込むことを

```
select {
case num := <-gen1:  // gen1から受信できるとき
	fmt.Println(num)
case channel<-1: // channelに送信できるとき
	fmt.Println("write channel to 1")
default:  // どっちも受信できないとき
	fmt.Println("neither chan cannot use")
}
```
複数のcaseが成り立つときはどちらかがランダムで選ばれる

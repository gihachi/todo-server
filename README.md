# todo-server

todoのPOSTと登録したtodoのGETをするサーバー.

## 動作環境

Golang 1.11以上

## 起動方法

```
$ git clone git@github.com:gihachi/todo-server.git
$ export GO111MODULE=on 
$ go mod download
$ go run server.go
```

起動するとlocalhost:8080で待ち受けを行う

## API

JSONでやり取りする

### todoのPOST

todoをサーバーに登録する.

- エンドポイント: api/v1/event

#### POSTパラメータ

| パラメータ名 | data type | 説明        | 必須項目か  |
|----------|-----------|--------------|------------|
| deadline | string    | 締め切りの時間(RFC3339 形式)| 必須       |
|----------|-----------|--------------|------------|
| title    | string    | todoのタイトル | 必須       |
|----------|-----------|--------------|------------|
| memo     | string    | todoの備考    | 必須ではない |

#### レスポンス

<dl>
    <dt>正しくPOSTできた場合</dt>
    <dd>200 OK</dd>
    <dt>正しくPOSTできなかった場合</dt>
    <dd>400 Bad Request</dd>
</dl>
|         | 200 OKの時  | 400 Bad Requestの時 | 
|---------|------------|---------------------|
| status  | success    | failure             |
|---------|------------|---------------------|
| message | registered | invalid date format |
|---------|------------|---------------------|
| id      | todoのid   |                     | 

### 全イベントのGET

登録されているtodoを全て取得する.

<dl>
    <dt>リクエストメソッド</dt>
    <dd>GET</dd>
    <dt>エンドポイント<dt>
    <dd>api/v1/event</dd>

</dl>

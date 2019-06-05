# todo-server

todoのPOSTと登録したtodoのGETをするサーバー.

## 開発環境

- 使用言語 : Golang
- フレームワーク : [echo](https://echo.labstack.com/)
- ORM : [GORM](https://gorm.io/)

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

## テストの実行方法

あらかじめサーバを起動しておく必要がある

```
$ go test
```

## API

JSONでやり取りする

### todoのPOST

todoをサーバーに登録する.

- エンドポイント: api/v1/event

#### POSTパラメータ

| パラメータ名 | data type | 説明        | 必須項目か  |
|----------|-----------|--------------|------------|
| deadline | string    | 締め切りの時間(RFC3339 形式)| 必須       |
| title    | string    | todoのタイトル | 必須       |
| memo     | string    | todoの備考    | 必須ではない |

#### レスポンス

<dl>
    <dt>正しくPOSTできた場合</dt>
    <dd>200 OK</dd>
    <dt>正しくPOSTできなかった場合</dt>
    <dd>400 Bad Request</dd>
</dl>

|  JSONプロパティ| data type | 200 OKの時  | 400 Bad Requestの時 | 
|---------|--------------|------------|---------------------|
| status  | string     |success    | failure             |
| message | string     | registered   |nvalid date format |
| id      | int        | todoのid                    | 

### 全todoのGET

登録されているtodoを全て取得する.

- エンドポイント : api/v1/event

#### レスポンス

- ステータス : 200 OK

レスポンスのbody

| JSONプロパティ | 説明  |
|---------|------------|
| events  | todoの配列  | 

eventsの配列内の各todoのプロパティは[個別todoの取得](#個別todoの取得)で説明する.

### 個別todoの取得

検索idに該当するtodoを取得する

- エンドポイント : /api/v1/event/${id}
  - ${id}にはtodoのidが入る

#### レスポンス

- ステータス : 200 OK
  - idに該当するtodoがなければ 404 Not Found

レスポンスのbody

|  JSONプロパティ | data type | 説明    |  
|----------|--------------|-----------|
| id       | int         | todoのid     |
| deadline | string      | todoの締め切り |
| title    | string      | todoのタイトル |
| memo     | string      | todoの備考    |


## todoデータの保存

POSTされたtodoはsqliteのファイルに保存している.
/dbディレクトリ直下にtodo.dbというファイルが生成され,todoの内容はこのファイルに保存される.DB接続は全て/util/ConnectionUtil.goの`GetDB()`で行なっているので,`gorm.Open()`の引数を変えれば別のデータベースに接続することも可能.



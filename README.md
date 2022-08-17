# NNE-GO

graphql のプロジェクト

フレームワークは、[gin](https://github.com/gin-gonic/gin)及び[gqlgen](https://gqlgen.com/getting-started/)を利用

# 環境構築

```sh
$ git clone https://github.com/Yuki-TU/nne-go
$ cd ./nne-go
$ docker-compose up -d --build
$ docker-compose exec go sh
# 以下コンテナ内
$ go generate ./...
$ go run server.go
```

[http://localhost:8081/](http://localhost:8081/)にアクセス

# 開発手順

### 1. スキーマファイル`*.graphqls`を編集

### 2. 以下のコマンド実行でリゾルバーファイル`*.resolvers.go`, 型ファイル`model/*.go`が更新

```sh
$ go generate ./...
# または
$ go run -mod=mod github.com/99designs/gqlgen generate
```

`$ go generate ./...`のエイリアスは`./graph/resolver.go`の`//go:generate go run -mod=mod github.com/99designs/gqlgen generate`のコメントである。

詳しくは、[gqlgen](https://gqlgen.com/getting-started/)を参照

### 3. リゾルバーである`*.resolvers.go`を編集して、返却するデータなどを定義

### 4. `$ go run server.go`を実行し、`http://localhost:8081/`にアクセスして確認

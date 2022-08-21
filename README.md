# Gin+gqlgenでGraphQLサーバ

* 使用言語
  * Go
* フレームワーク
  * Gin
* ORM
  * GORM
* GraphQLライブラリ
  * gqlgen

## セットアップ

```
# フロントエンドのセットアップ
mkdir frontend
cd frontend
git clone https://github.com/takuyapv/calendar.git
cd calendar
yarn install
yarn generate
cd ../../
# サーバー起動 http://localhost:8080
go run server.go
```

## 各パスについて

* `/` フロントエンドトップ
  * frontend/calendar/dist/index.htmlを読み込んでいます
* `/ide` GraphiQL (GraphQLのブラウザIDE)
* `/query` GraphQL PlayGround

## 作ってみた感想

Goで一通りサーバーを作ってみたのは初めてですが、それなりに情報も多いこともあり学習コストが低く、非常に簡易的ではありますが短時間で作成することができました。

単純にPWA用バックエンドの開発効率として考えるとPythonの **FastAPI** やPHPの**API Platform**の方が良さそうなので、実行パフォーマンスとの兼ね合いでの判断となりそうです。


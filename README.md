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

Goで一通りサーバーを作ってみたのは初めてですが、それなりに情報も多いこともあり学習コストが非常に低く、簡易的ではありますが短時間で作成することができました。

ただ、GinについてはPythonのdjangoやphpのsymfonyのようなフルスタックフレームワークと比べると、どこに何を書くかなどのルールがないので、気を付けないと大規模化したときに苦労しそうです。

そういう意味ではgqlgenはある程度雛型を作ってくれるので、Ginと組み合わせることで良いアシストになってくれている印象です。

単純にPWA用バックエンドの開発効率として考えるとPythonのFastAPIやPHPのAPI Platformの方が良さそうなので、実行パフォーマンスとの兼ね合いでの判断となりそうです。


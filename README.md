# cli-sample-built-with-golang-following-clean-architecture

本リポジトリは、クリーンアーキテクチャを適用した Golang ベースの CLI ツールサンプルを提供します。

## 前提条件

本 CLI ツールを利用するためには、以下のソフトウェアがインストールされている必要があります。

* Golang 1.16 以上

また、 CLI ツールの動作検証用のスタブサーバーを利用するためには、以下のソフトウェアがインストールされている必要があります。

* Node.js v14.17.1 以上

## 提供機能

本 CLI ツールは、以下の機能を提供します。

* `http://localhost:8080/api/users` に対して GET リクエストを実行するサブコマンドを提供します。

## 実行方法

本 CLI ツールは、以下の手順で実行します。

1. 本プロジェクトのディレクトリに移動します。
   ```sh
   $ cd cli-sample-built-with-golang-following-clean-architecture
   ```
1. GET リクエストを実行するためのスタブサーバーを起動します。
   ```sh
   $ node stub_server/index.js

   listening on port 8080
   ```
1. 以下のコマンドを実行します。
   ```sh
   $ go run main.go get

   [{1 foo} {2 bar}]
   ```

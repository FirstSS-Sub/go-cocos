# COCO's ガチャ
- サイゼリヤガチャのパクリで作ってみました
- Golangで作成 フレームワークはginを使用
- データベースはSQLiteを使用

# ローカルでの実行方法（作者の場合）
- GO言語をインストール http://golang.jp/install
- SQLite3のドライバをインストール
    - Linux, Mac OS
        - ターミナルで以下のコマンドを実行
        > $ go get github.com/mattn/go-sqlite3
    - Windows
        - https://text.baldanders.info/golang/sqlite-with-golang-in-windows/ を参照
- このリポジトリを go/src 下にclone
- フォルダ go-cocos 下でターミナルにて以下のコマンドを実行
    > $ go run main.go
- ブラウザで http://127.0.0.1:8080 にアクセス

# 今後の改善
- Golangでhtmlに値を複数渡すことが出来ない？ようなので、値段の合計値をどうやって出力するか悩み中
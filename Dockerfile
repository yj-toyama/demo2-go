# ベースイメージとして公式のGoイメージを使用
FROM golang:1.17

# 作業ディレクトリを設定
WORKDIR /app

# ソースコードをコンテナにコピー
COPY . .

# Goのアプリケーションをビルド
RUN go build -o main .

# ポート8001を公開
EXPOSE 8001

# コンテナ起動時に実行するコマンド
CMD [ "./main" ]

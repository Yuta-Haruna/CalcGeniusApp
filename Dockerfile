
# ベースとなるGoのイメージを指定
FROM golang:latest

# コンテナ内に作業ディレクトリを作成
WORKDIR /app

# ホストOSのファイルをコンテナ内にコピー
COPY . .

# Goプロジェクトをビルド
RUN go build -o CalcGeniusApp

# 実行可能ファイルを起動
CMD ["./CalcGeniusApp"]
# ベースとなるイメージ
FROM golang:1.13.7-buster

# プロジェクトルートに移動する
WORKDIR /usr/local/mplatmm

# delve インストール
RUN go get -u github.com/go-delve/delve/cmd/dlv \
  && go build -o /go/bin/dlv github.com/go-delve/delve/cmd/dlv

# Air インストール
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin \
  && chmod +x /go/bin/air

# コンテナ実行時のデフォルトを設定する
# ライブリロードを実行する
CMD air -c .air.conf

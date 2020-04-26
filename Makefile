ROOT := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))
SKAFFOLD_YAML := $(ROOT)/skaffold.yaml

.PHONY: protobuf
protobuf:  ## protobuf の定義をもとにGoのコードを自動生成する
	protoc --go_out=plugins=grpc:. protobuf/*.proto


.PHONY: gazelle
gazelle:  ## gazelle によって設定ファイルを自動生成する
	bazel run //:gazelle -- update-repos -from_file ./go.mod
	bazel run //:gazelle 


.PHONY: skaffold
skaffold: ## 自動リロードを有効にしてローカルでサーバーを起動する
	skaffold dev --filename $(SKAFFOLD_YAML)

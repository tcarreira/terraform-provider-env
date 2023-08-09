default: testacc

NAME=env
BINARY=terraform-provider-${NAME}

# Run acceptance tests
.PHONY: testacc
testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m

build:
	go build -o ${BINARY} .

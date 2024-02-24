PROTOS := Device.proto Automate.proto
PROTO_DIR := ./proto
GO_PROTO_DIR := $(PROTO_DIR)/go
GO_PROTO_OUTS := $(addprefix $(GO_PROTO_DIR)/, $(subst .proto,.pb.go,$(PROTOS)))
GO_SRC := src/elgato/*.go src/govee/*.go main.go $(GO_PROTO_OUTS)
NAME := home-automation

all: $(NAME)

$(NAME): $(GO_PROTO_OUTS) $(GO_SRC)
	go build .

install: $(NAME)
	cp $(NAME) $(GOPATH)/bin

# $(GO_PROTOS):
# 	protoc -I=$(PROTO_DIR) --go_out=$(PROTO_DIR) $(PROTOS)

$(GO_PROTO_DIR)/%.pb.go: $(PROTO_DIR)/%.proto
	protoc -I=$(PROTO_DIR) --go_out=$(PROTO_DIR) $<

clean:
	$(RM) $(GO_PROTO_DIR)/*.pb.go $(NAME)

re: clean all

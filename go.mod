module github.com/furynet/xfury

go 1.18

require (
	github.com/cosmos/cosmos-sdk v0.45.1
	github.com/cosmos/ibc-go/v3 v3.0.0-rc2
	github.com/gogo/protobuf v1.3.3
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/golang-jwt/jwt/v4 v4.5.0
	github.com/golang/protobuf v1.5.3
	github.com/golangci/golangci-lint v1.50.1
	github.com/google/uuid v1.3.0
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/rakyll/statik v0.1.7
	github.com/spf13/cast v1.5.0
	github.com/spf13/cobra v1.6.1
	github.com/stretchr/testify v1.8.1
	github.com/tendermint/tendermint v0.34.24
	github.com/tendermint/tm-db v0.6.8-0.20220506192307-f628bb5dc95b
	google.golang.org/genproto v0.0.0-20230223222841-637eb2293923
	google.golang.org/grpc v1.53.0
	gopkg.in/yaml.v2 v2.4.0
	mvdan.cc/gofumpt v0.4.0
)

require (
	github.com/google/btree v1.1.2 // indirect
	github.com/spf13/viper v1.14.0 // indirect
)

require (
	github.com/99designs/keyring v1.2.1 // indirect
	github.com/bgentry/speakeasy v0.1.1-0.20220910012023-760eaf8b6816 // indirect
	github.com/btcsuite/btcd v0.22.2 // indirect
	github.com/coinbase/rosetta-sdk-go v0.7.9 // indirect
	github.com/confio/ics23/go v0.9.0 // indirect
	github.com/cosmos/btcutil v1.0.5 // indirect
	github.com/cosmos/iavl v0.19.0 // indirect
	github.com/cosmos/ledger-cosmos-go v0.12.2 // indirect
	github.com/dustin/go-humanize v1.0.1-0.20200219035652-afde56e7acac // indirect
	github.com/golang/glog v1.0.0 // indirect
	github.com/hdevalence/ed25519consensus v0.0.0-20220222234857-c00d1f31bab3 // indirect
	github.com/improbable-eng/grpc-web v0.15.0 // indirect
	github.com/mattn/go-runewidth v0.0.10 // indirect
	github.com/prometheus/client_golang v1.14.0 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	golang.org/x/crypto v0.2.0 // indirect
	nhooyr.io/websocket v1.8.7 // indirect
)

replace (
	github.com/confio/ics23/go => github.com/cosmos/cosmos-sdk/ics23/go v0.8.0
	// use cosmos-compatible protobufs
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	// use grpc compatible with cosmos protobufs
	google.golang.org/grpc => google.golang.org/grpc v1.33.2
)

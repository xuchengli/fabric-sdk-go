module github.com/xuchengli/fabric-sdk-go

go 1.14

require (
	github.com/Knetic/govaluate v3.0.1-0.20171022003610-9aa49832a739+incompatible
	github.com/cloudflare/cfssl v1.6.1
	github.com/go-kit/kit v0.12.0
	github.com/gogo/protobuf v1.3.2
	github.com/golang/mock v1.6.0
	github.com/golang/protobuf v1.5.2
	github.com/hyperledger/fabric-lib-go v1.0.0
	github.com/hyperledger/fabric-protos-go v0.0.0-20190821180310-6b6ac9042dfd
	github.com/miekg/pkcs11 v1.1.1
	github.com/mitchellh/mapstructure v1.4.3
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.12.1
	github.com/spf13/cast v1.4.1
	github.com/spf13/viper v1.10.1
	github.com/stretchr/testify v1.7.1
	github.com/sykesm/zap-logfmt v0.0.4 // indirect
	github.com/xuchengli/gm-crypto v0.0.0-20220408103036-316db6bcf11a
	github.com/xuchengli/gm-plugins v0.0.0-20220408102215-54f855ff7927
	github.com/xuchengli/http v0.0.0-20220408140247-23e0c75ec4b4
	go.uber.org/zap v1.21.0
	golang.org/x/crypto v0.0.0-20220408190544-5352b0902921
	golang.org/x/net v0.0.0-20220407224826-aac1ed45d8e3
	google.golang.org/grpc v1.45.0
	gopkg.in/yaml.v2 v2.4.0
)

replace github.com/hyperledger/fabric-sdk-go => github.com/xuchengli/fabric-sdk-go v0.0.0

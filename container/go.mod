module github.com/cosmos/cosmos-sdk/container

go 1.17

require (
	github.com/cosmos/cosmos-sdk/api v0.1.0-alpha3
	github.com/goccy/go-graphviz v0.0.9
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.7.0
	google.golang.org/protobuf v1.27.1
	sigs.k8s.io/yaml v1.3.0
)

require (
	github.com/cosmos/cosmos-proto v1.0.0-alpha6 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/fogleman/gg v1.3.0 // indirect
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/image v0.0.0-20200119044424-58c23975cae1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c // indirect
)

replace github.com/cosmos/cosmos-sdk/api => ../api

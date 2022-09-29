module cosmossdk.io/client/v2

go 1.19

require (
	cosmossdk.io/api v0.2.1
	cosmossdk.io/tx v0.0.0
	cosmossdk.io/crypto/v2 v2.0.0
	github.com/99designs/keyring v1.2.1
	github.com/cosmos/cosmos-proto v1.0.0-alpha7
	github.com/iancoleman/strcase v0.2.0
	github.com/spf13/cobra v1.5.0
	github.com/spf13/pflag v1.0.5
	google.golang.org/grpc v1.49.0
	google.golang.org/protobuf v1.28.1
	gotest.tools/v3 v3.3.0
)

require (
	github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4 // indirect
	github.com/cosmos/gogoproto v1.4.2 // indirect
	github.com/danieljoos/wincred v1.1.2 // indirect
	github.com/dvsekhvalnov/jose2go v1.5.0 // indirect
	github.com/godbus/dbus v0.0.0-20190726142602-4481cbc300e2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.8 // indirect
	github.com/gsterjov/go-libsecret v0.0.0-20161001094733-a6f4afe4910c // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/mtibben/percent v0.2.1 // indirect
	golang.org/x/net v0.0.0-20220726230323-06994584191e // indirect
	golang.org/x/sys v0.0.0-20220811171246-fbc7d0a398ab // indirect
	golang.org/x/term v0.0.0-20210927222741-03fcf44c2211 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220725144611-272f38e5d71b // indirect
)

replace cosmossdk.io/tx => ../../tx

replace cosmossdk.io/api => ../../api

replace cosmossdk.io/crypto/v2 => ../../crypto/v2

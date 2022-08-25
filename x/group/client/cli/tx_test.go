package cli_test

import (
	"context"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/cosmos/cosmos-sdk/testutil"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	testutilmod "github.com/cosmos/cosmos-sdk/types/module/testutil"

	cli "github.com/cosmos/cosmos-sdk/x/group/client/cli"

	groupmodule "github.com/cosmos/cosmos-sdk/x/group/module"
	abci "github.com/tendermint/tendermint/abci/types"
	rpcclientmock "github.com/tendermint/tendermint/rpc/client/mock"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"
	tmtypes "github.com/tendermint/tendermint/types"
)

const validMetadata = "metadata"

var tooLongMetadata = strings.Repeat("A", 256)

var commonFlags = []string{
	fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
	fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
	// fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
}

var _ client.TendermintRPC = (*mockTendermintRPC)(nil)

type mockTendermintRPC struct {
	rpcclientmock.Client

	responseQuery abci.ResponseQuery
}

func (_ mockTendermintRPC) BroadcastTxCommit(_ context.Context, _ tmtypes.Tx) (*coretypes.ResultBroadcastTxCommit, error) {
	return &coretypes.ResultBroadcastTxCommit{}, nil
}

func TestTxCreateGroup(t *testing.T) {
	encCfg := testutilmod.MakeTestEncodingConfig(groupmodule.AppModuleBasic{})
	kr := keyring.NewInMemory(encCfg.Codec)
	baseCtx := client.Context{}.
		WithKeyring(kr).
		WithTxConfig(encCfg.TxConfig).
		WithCodec(encCfg.Codec).
		WithClient(mockTendermintRPC{Client: rpcclientmock.Client{}}).
		WithAccountRetriever(client.MockAccountRetriever{}).
		WithOutput(io.Discard).
		WithChainID("test-chain")

	accounts := testutil.CreateKeyringAccounts(t, kr, 1)

	validMembers := fmt.Sprintf(`{"members": [{
		"address": "%s",
		  "weight": "1",
		  "metadata": "%s"
	  }]}`, accounts[0].Address.String(), validMetadata)
	validMembersFile := testutil.WriteToNewTempFile(t, validMembers)

	invalidMembersAddress := `{"members": [{
		  "address": "",
		  "weight": "1"
	  }]}`
	invalidMembersAddressFile := testutil.WriteToNewTempFile(t, invalidMembersAddress)

	invalidMembersWeight := fmt.Sprintf(`{"members": [{
		"address": "%s",
		  "weight": "0"
	  }]}`, accounts[0].Address.String())
	invalidMembersWeightFile := testutil.WriteToNewTempFile(t, invalidMembersWeight)

	// invalidMembersMetadata := fmt.Sprintf(`{"members": [{
	// 	"address": "%s",
	// 	  "weight": "1",
	// 	  "metadata": "%s"
	//   }]}`, accounts[0].Address.String(), tooLongMetadata)
	// invalidMembersMetadataFile := testutil.WriteToNewTempFile(t, invalidMembersMetadata)

	testCases := []struct {
		name         string
		args         []string
		expectErr    bool
		expectErrMsg string
		expectedCode uint32
	}{
		{
			"correct data",
			[]string{
				accounts[0].Address.String(),
				"",
				validMembersFile.Name(),
			},
			false,
			"",
			0,
		},
		{
			"with amino-json",
			[]string{
				accounts[0].Address.String(),
				"",
				validMembersFile.Name(),
				fmt.Sprintf("--%s=%s", flags.FlagSignMode, flags.SignModeLegacyAminoJSON),
			},
			false,
			"",
			0,
		},
		// {
		// 	"group metadata too long",
		// 	[]string{
		// 		accounts[0].Address.String(),
		// 		strings.Repeat("a", 256),
		// 		"",
		// 	},
		// 	true,
		// 	"group metadata: limit exceeded",
		// 	0,
		// },
		{
			"invalid members address",
			[]string{
				accounts[0].Address.String(),
				"null",
				invalidMembersAddressFile.Name(),
			},
			true,
			"message validation failed: address: empty address string is not allowed",
			0,
		},
		{
			"invalid members weight",
			[]string{
				accounts[0].Address.String(),
				"null",
				invalidMembersWeightFile.Name(),
			},
			true,
			"expected a positive decimal, got 0: invalid decimal string",
			0,
		},
		// {
		// 	"members metadata too long",
		// 	[]string{
		// 		accounts[0].Address.String(),
		// 		"null",
		// 		invalidMembersMetadataFile.Name(),
		// 	},
		// 	true,
		// 	"member metadata: limit exceeded",
		// 	0,
		// },
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := svrcmd.CreateExecuteContext(context.Background())

			cmd := cli.MsgCreateGroupCmd()
			cmd.SetOut(io.Discard)
			assert.NotNil(t, cmd)

			cmd.SetContext(ctx)
			cmd.SetArgs(tc.args)
			cmd.SetArgs(append(tc.args, commonFlags...))

			assert.NoError(t, client.SetCmdClientContextHandler(baseCtx, cmd))

			err := cmd.Execute()
			if tc.expectErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectErrMsg)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

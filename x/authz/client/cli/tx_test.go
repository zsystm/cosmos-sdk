package cli_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	abci "github.com/tendermint/tendermint/abci/types"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
	rpcclientmock "github.com/tendermint/tendermint/rpc/client/mock"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/testutil"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	testutilmod "github.com/cosmos/cosmos-sdk/types/module/testutil"
	"github.com/cosmos/cosmos-sdk/x/authz/client/cli"
	authzmodule "github.com/cosmos/cosmos-sdk/x/authz/module"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
	govcli "github.com/cosmos/cosmos-sdk/x/gov/client/cli"
	govtestutil "github.com/cosmos/cosmos-sdk/x/gov/client/testutil"
	govv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

var (
	typeMsgSend           = bank.SendAuthorization{}.MsgTypeURL()
	typeMsgVote           = sdk.MsgTypeURL(&govv1.MsgVote{})
	typeMsgSubmitProposal = sdk.MsgTypeURL(&govv1.MsgSubmitProposal{})
)

var _ client.TendermintRPC = (*mockTendermintRPC)(nil)

type mockTendermintRPC struct {
	rpcclientmock.Client

	responseQuery abci.ResponseQuery
}

func newMockTendermintRPC(respQuery abci.ResponseQuery) mockTendermintRPC {
	return mockTendermintRPC{responseQuery: respQuery}
}

func (mockTendermintRPC) BroadcastTxSync(context.Context, tmtypes.Tx) (*coretypes.ResultBroadcastTx, error) {
	return &coretypes.ResultBroadcastTx{Code: 0}, nil
}

func (m mockTendermintRPC) ABCIQueryWithOptions(
	_ context.Context,
	_ string, _ tmbytes.HexBytes,
	_ rpcclient.ABCIQueryOptions,
) (*coretypes.ResultABCIQuery, error) {
	return &coretypes.ResultABCIQuery{Response: m.responseQuery}, nil
}

type CLITestSuite struct {
	suite.Suite

	kr        keyring.Keyring
	encCfg    testutilmod.TestEncodingConfig
	baseCtx   client.Context
	clientCtx client.Context
	grantee   []sdk.AccAddress
}

func TestCLITestSuite(t *testing.T) {
	suite.Run(t, new(CLITestSuite))
}

func (s *CLITestSuite) SetupSuite() {
	s.encCfg = testutilmod.MakeTestEncodingConfig(authzmodule.AppModuleBasic{})
	s.kr = keyring.NewInMemory(s.encCfg.Codec)
	s.baseCtx = client.Context{}.
		WithKeyring(s.kr).
		WithTxConfig(s.encCfg.TxConfig).
		WithCodec(s.encCfg.Codec).
		WithClient(mockTendermintRPC{Client: rpcclientmock.Client{}}).
		WithAccountRetriever(client.MockAccountRetriever{}).
		WithOutput(io.Discard).
		WithChainID("test-chain")

	var outBuf bytes.Buffer
	ctxGen := func() client.Context {
		bz, _ := s.encCfg.Codec.Marshal(&sdk.TxResponse{})
		c := newMockTendermintRPC(abci.ResponseQuery{
			Value: bz,
		})
		return s.baseCtx.WithClient(c)
	}
	s.clientCtx = ctxGen().WithOutput(&outBuf)

	val := testutil.CreateKeyringAccounts(s.T(), s.kr, 1)
	s.grantee = make([]sdk.AccAddress, 6)

	// Send some funds to the new account.
	// Create new account in the keyring.
	s.grantee[0] = s.createAccount("grantee1")
	s.msgSendExec(s.grantee[0])

	// create a proposal with deposit
	_, err := govtestutil.MsgSubmitLegacyProposal(s.clientCtx, val[0].Address.String(),
		"Text Proposal 1", "Where is the title!?", govv1beta1.ProposalTypeText,
		fmt.Sprintf("--%s=%s", govcli.FlagDeposit, sdk.NewCoin("stake", govv1.DefaultMinDepositTokens).String()))
	s.Require().NoError(err)
}

func (s *CLITestSuite) createAccount(uid string) sdk.AccAddress {
	// Create new account in the keyring.
	k, _, err := s.clientCtx.Keyring.NewMnemonic(uid, keyring.English, sdk.FullFundraiserPath, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
	s.Require().NoError(err)

	addr, err := k.GetAddress()
	s.Require().NoError(err)

	return addr
}

func (s *CLITestSuite) msgSendExec(grantee sdk.AccAddress) {
	val := testutil.CreateKeyringAccounts(s.T(), s.kr, 1)
	// Send some funds to the new account.
	out, err := clitestutil.MsgSendExec(
		s.clientCtx,
		val[0].Address,
		grantee,
		sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(200))), fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10))).String()),
	)
	s.Require().NoError(err)
	s.Require().Contains(out.String(), `"code":0`)
}

func (s *CLITestSuite) TestCLITxGrantAuthorization() {
	val := testutil.CreateKeyringAccounts(s.T(), s.kr, 1)

	grantee := s.grantee[0]

	twoHours := time.Now().Add(time.Minute * 120).Unix()
	pastHour := time.Now().Add(-time.Minute * 60).Unix()

	testCases := []struct {
		name         string
		args         []string
		expectedCode uint32
		expectErr    bool
		expErrMsg    string
	}{
		{
			"Invalid granter Address",
			[]string{
				"grantee_addr",
				"send",
				fmt.Sprintf("--%s=100stake", cli.FlagSpendLimit),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, "granter"),
				fmt.Sprintf("--%s=true", flags.FlagGenerateOnly),
				fmt.Sprintf("--%s=%d", cli.FlagExpiration, twoHours),
			},
			0,
			true,
			"key not found",
		},
		{
			"Invalid grantee Address",
			[]string{
				"grantee_addr",
				"send",
				fmt.Sprintf("--%s=100stake", cli.FlagSpendLimit),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val[0].Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagGenerateOnly),
				fmt.Sprintf("--%s=%d", cli.FlagExpiration, twoHours),
			},
			0,
			true,
			"invalid separator index",
		},
		{
			"Invalid expiration time",
			[]string{
				grantee.String(),
				"send",
				fmt.Sprintf("--%s=100stake", cli.FlagSpendLimit),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val[0].Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagBroadcastMode),
				fmt.Sprintf("--%s=%d", cli.FlagExpiration, pastHour),
			},
			0,
			true,
			"",
		},
		{
			"fail with error invalid msg-type",
			[]string{
				grantee.String(),
				"generic",
				fmt.Sprintf("--%s=invalid-msg-type", cli.FlagMsgType),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val[0].Address.String()),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10))).String()),
				fmt.Sprintf("--%s=%d", cli.FlagExpiration, twoHours),
			},
			0x1d,
			false,
			"",
		},
		{
			"failed with error both validators not allowed",
			[]string{
				grantee.String(),
				"delegate",
				fmt.Sprintf("--%s=100stake", cli.FlagSpendLimit),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val[0].Address.String()),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf("--%s=%d", cli.FlagExpiration, twoHours),
				fmt.Sprintf("--%s=%s", cli.FlagAllowedValidators, val.ValAddress.String()),
				fmt.Sprintf("--%s=%s", cli.FlagDenyValidators, val.ValAddress.String()),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10))).String()),
			},
			0,
			true,
			"cannot set both allowed & deny list",
		},
		{
			"invalid bond denom for tx delegate authorization allowed validators",
			[]string{
				grantee.String(),
				"delegate",
				fmt.Sprintf("--%s=100xyz", cli.FlagSpendLimit),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val[0].Address.String()),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf("--%s=%d", cli.FlagExpiration, twoHours),
				fmt.Sprintf("--%s=%s", cli.FlagAllowedValidators, val.ValAddress.String()),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10))).String()),
			},
			0,
			true,
			"invalid denom",
		},
		{
			"invalid bond denom for tx delegate authorization deny validators",
			[]string{
				grantee.String(),
				"delegate",
				fmt.Sprintf("--%s=100xyz", cli.FlagSpendLimit),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val[0].Address.String()),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf("--%s=%d", cli.FlagExpiration, twoHours),
				fmt.Sprintf("--%s=%s", cli.FlagDenyValidators, val.ValAddress.String()),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10))).String()),
			},
			0,
			true,
			"invalid denom",
		},
		{
			"invalid bond denom for tx undelegate authorization",
			[]string{
				grantee.String(),
				"unbond",
				fmt.Sprintf("--%s=100xyz", cli.FlagSpendLimit),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val[0].Address.String()),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf("--%s=%d", cli.FlagExpiration, twoHours),
				fmt.Sprintf("--%s=%s", cli.FlagAllowedValidators, val.ValAddress.String()),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10))).String()),
			},
			0,
			true,
			"invalid denom",
		},
		{
			"invalid bond denon for tx redelegate authorization",
			[]string{
				grantee.String(),
				"redelegate",
				fmt.Sprintf("--%s=100xyz", cli.FlagSpendLimit),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val[0].Address.String()),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf("--%s=%d", cli.FlagExpiration, twoHours),
				fmt.Sprintf("--%s=%s", cli.FlagAllowedValidators, val.ValAddress.String()),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10))).String()),
			},
			0,
			true,
			"invalid denom",
		},
		{
			"invalid decimal coin expression with more than single coin",
			[]string{
				grantee.String(),
				"delegate",
				fmt.Sprintf("--%s=100stake,20xyz", cli.FlagSpendLimit),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val[0].Address.String()),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf("--%s=%d", cli.FlagExpiration, twoHours),
				fmt.Sprintf("--%s=%s", cli.FlagAllowedValidators, val.ValAddress.String()),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10))).String()),
			},
			0,
			true,
			"invalid decimal coin expression",
		},
		{
			"valid tx delegate authorization allowed validators",
			[]string{
				grantee.String(),
				"delegate",
				fmt.Sprintf("--%s=100stake", cli.FlagSpendLimit),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val[0].Address.String()),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf("--%s=%d", cli.FlagExpiration, twoHours),
				fmt.Sprintf("--%s=%s", cli.FlagAllowedValidators, val.ValAddress.String()),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10))).String()),
			},
			0,
			false,
			"",
		},
		{
			"valid tx delegate authorization deny validators",
			[]string{
				grantee.String(),
				"delegate",
				fmt.Sprintf("--%s=100stake", cli.FlagSpendLimit),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val[0].Address.String()),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf("--%s=%d", cli.FlagExpiration, twoHours),
				fmt.Sprintf("--%s=%s", cli.FlagDenyValidators, val.ValAddress.String()),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10))).String()),
			},
			0,
			false,
			"",
		},
		{
			"valid tx undelegate authorization",
			[]string{
				grantee.String(),
				"unbond",
				fmt.Sprintf("--%s=100stake", cli.FlagSpendLimit),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val[0].Address.String()),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf("--%s=%d", cli.FlagExpiration, twoHours),
				fmt.Sprintf("--%s=%s", cli.FlagAllowedValidators, val.ValAddress.String()),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10))).String()),
			},
			0,
			false,
			"",
		},
		{
			"valid tx redelegate authorization",
			[]string{
				grantee.String(),
				"redelegate",
				fmt.Sprintf("--%s=100stake", cli.FlagSpendLimit),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val[0].Address.String()),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf("--%s=%d", cli.FlagExpiration, twoHours),
				fmt.Sprintf("--%s=%s", cli.FlagAllowedValidators, val.ValAddress.String()),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10))).String()),
			},
			0,
			false,
			"",
		},
		{
			"Valid tx send authorization",
			[]string{
				grantee.String(),
				"send",
				fmt.Sprintf("--%s=100stake", cli.FlagSpendLimit),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val[0].Address.String()),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf("--%s=%d", cli.FlagExpiration, twoHours),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10))).String()),
			},
			0,
			false,
			"",
		},
		{
			"Valid tx send authorization with allow list",
			[]string{
				grantee.String(),
				"send",
				fmt.Sprintf("--%s=100stake", cli.FlagSpendLimit),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val[0].Address.String()),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf("--%s=%d", cli.FlagExpiration, twoHours),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10))).String()),
				fmt.Sprintf("--%s=%s", cli.FlagAllowList, s.grantee[1]),
			},
			0,
			false,
			"",
		},
		{
			"Invalid tx send authorization with duplicate allow list",
			[]string{
				grantee.String(),
				"send",
				fmt.Sprintf("--%s=100stake", cli.FlagSpendLimit),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val[0].Address.String()),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf("--%s=%d", cli.FlagExpiration, twoHours),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10))).String()),
				fmt.Sprintf("--%s=%s", cli.FlagAllowList, fmt.Sprintf("%s,%s", s.grantee[1], s.grantee[1])),
			},
			0,
			true,
			"duplicate entry",
		},
		{
			"Valid tx generic authorization",
			[]string{
				grantee.String(),
				"generic",
				fmt.Sprintf("--%s=%s", cli.FlagMsgType, typeMsgVote),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val[0].Address.String()),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf("--%s=%d", cli.FlagExpiration, twoHours),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10))).String()),
			},
			0,
			false,
			"",
		},
		{
			"fail when granter = grantee",
			[]string{
				grantee.String(),
				"generic",
				fmt.Sprintf("--%s=%s", cli.FlagMsgType, typeMsgVote),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, grantee.String()),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf("--%s=%d", cli.FlagExpiration, twoHours),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10))).String()),
			},
			0,
			true,
			"grantee and granter should be different",
		},
		{
			"Valid tx with amino",
			[]string{
				grantee.String(),
				"generic",
				fmt.Sprintf("--%s=%s", cli.FlagMsgType, typeMsgVote),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val[0].Address.String()),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf("--%s=%d", cli.FlagExpiration, twoHours),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10))).String()),
				fmt.Sprintf("--%s=%s", flags.FlagSignMode, flags.SignModeLegacyAminoJSON),
			},
			0,
			false,
			"",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			out, err := s.createGrant(
				tc.args,
			)
			if tc.expectErr {
				s.Require().Error(err, out)
				s.Require().Contains(err.Error(), tc.expErrMsg)
			} else {
				var txResp sdk.TxResponse
				s.Require().NoError(err)
				s.Require().NoError(s.clientCtx.Codec.UnmarshalJSON(out.Bytes(), &txResp), out.String())
				// s.Require().NoError(clitestutil.CheckTxCode(s.network, val.ClientCtx, txResp.TxHash, tc.expectedCode))
			}
		})
	}
}

func (s *CLITestSuite) createGrant(args []string) (testutil.BufferWriter, error) {
	cmd := cli.NewCmdGrantAuthorization()
	return clitestutil.ExecTestCLICmd(s.clientCtx, cmd, args)
}

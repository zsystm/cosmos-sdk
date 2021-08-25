package rest_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	kmultisig "github.com/cosmos/cosmos-sdk/crypto/keys/multisig"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/testutil"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authclient "github.com/cosmos/cosmos-sdk/x/auth/client"
	authtest "github.com/cosmos/cosmos-sdk/x/auth/client/testutil"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
	bankcli "github.com/cosmos/cosmos-sdk/x/bank/client/testutil"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
)

type IntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network

	stdTx    legacytx.StdTx
	stdTxRes *sdk.TxResponse
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite")

	cfg := network.DefaultConfig()
	cfg.NumValidators = 2

	s.cfg = cfg
	s.network = network.New(s.T(), cfg)

	kb := s.network.Validators[0].ClientCtx.Keyring
	_, _, err := kb.NewMnemonic("newAccount", keyring.English, sdk.FullFundraiserPath, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
	s.Require().NoError(err)

	account1, _, err := kb.NewMnemonic("newAccount1", keyring.English, sdk.FullFundraiserPath, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
	s.Require().NoError(err)

	account2, _, err := kb.NewMnemonic("newAccount2", keyring.English, sdk.FullFundraiserPath, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
	s.Require().NoError(err)

	multi := kmultisig.NewLegacyAminoPubKey(2, []cryptotypes.PubKey{account1.GetPubKey(), account2.GetPubKey()})
	_, err = kb.SaveMultisig("multi", multi)
	s.Require().NoError(err)

	_, err = s.network.WaitForHeight(1)
	s.Require().NoError(err)

	// Broadcast a StdTx used for tests.
	s.stdTx = s.createTestStdTx(s.network.Validators[0], 0, 1)
	s.stdTxRes, err = s.broadcastReq(s.stdTx, "block")
	s.Require().NoError(err)

	s.Require().NoError(s.network.WaitForNextBlock())
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
	s.network.Cleanup()
}

func (s *IntegrationTestSuite) mkStdTx() legacytx.StdTx {
	val := s.network.Validators[0]
	// NOTE: this uses StdTx explicitly, don't migrate it!
	return legacytx.StdTx{
		Msgs: []sdk.Msg{&types.MsgSend{
			FromAddress: val.Address.String(),
		}},
		Fee: legacytx.StdFee{
			Amount: sdk.Coins{sdk.NewInt64Coin("foo", 10)},
			Gas:    10000,
		},
		Memo: "FOOBAR",
		Signatures: []legacytx.StdSignature{{
			PubKey:    val.PubKey,
			Signature: []byte{42},
		}},
	}
}

func (s *IntegrationTestSuite) TestQueryAccountWithColon() {
	val := s.network.Validators[0]
	// This address is not a valid simapp address! It is only used to test that addresses with
	// colon don't 501. See
	// https://github.com/cosmos/cosmos-sdk/issues/8650
	addrWithColon := "cosmos:1m4f6lwd9eh8e5nxt0h00d46d3fr03apfh8qf4g"

	res, err := rest.GetRequest(fmt.Sprintf("%s/cosmos/auth/v1beta1/accounts/%s", val.APIAddress, addrWithColon))
	s.Require().NoError(err)
	s.Require().Contains(string(res), "decoding bech32 failed")
}

// Helper function to test querying txs. We will use it to query StdTx and service `Msg`s.
func (s *IntegrationTestSuite) testQueryTx(txHeight int64, txHash, txRecipient string) {
	val0 := s.network.Validators[0]

	testCases := []struct {
		desc     string
		malleate func() *sdk.TxResponse
	}{
		{
			"Query by hash",
			func() *sdk.TxResponse {
				txJSON, err := rest.GetRequest(fmt.Sprintf("%s/txs/%s", val0.APIAddress, txHash))
				s.Require().NoError(err)

				var txResAmino sdk.TxResponse
				s.Require().NoError(val0.ClientCtx.LegacyAmino.UnmarshalJSON(txJSON, &txResAmino))
				return &txResAmino
			},
		},
		{
			"Query by height",
			func() *sdk.TxResponse {
				txJSON, err := rest.GetRequest(fmt.Sprintf("%s/txs?limit=10&page=1&tx.height=%d", val0.APIAddress, txHeight))
				s.Require().NoError(err)

				var searchtxResult sdk.SearchTxsResult
				s.Require().NoError(val0.ClientCtx.LegacyAmino.UnmarshalJSON(txJSON, &searchtxResult))
				s.Require().Len(searchtxResult.Txs, 1)
				return searchtxResult.Txs[0]
			},
		},
		{
			"Query by event (transfer.recipient)",
			func() *sdk.TxResponse {
				txJSON, err := rest.GetRequest(fmt.Sprintf("%s/txs?transfer.recipient=%s", val0.APIAddress, txRecipient))
				s.Require().NoError(err)

				var searchtxResult sdk.SearchTxsResult
				s.Require().NoError(val0.ClientCtx.LegacyAmino.UnmarshalJSON(txJSON, &searchtxResult))
				s.Require().Len(searchtxResult.Txs, 1)
				return searchtxResult.Txs[0]
			},
		},
	}

	for _, tc := range testCases {
		s.Run(fmt.Sprintf("Case %s", tc.desc), func() {
			txResponse := tc.malleate()

			// Check that the height is correct.
			s.Require().Equal(txHeight, txResponse.Height)

			// Check that the events are correct.
			s.Require().Contains(
				txResponse.RawLog,
				fmt.Sprintf("{\"key\":\"recipient\",\"value\":\"%s\"}", txRecipient),
			)

			// Check that the Msg is correct.
			stdTx, ok := txResponse.Tx.GetCachedValue().(legacytx.StdTx)
			s.Require().True(ok)
			msgs := stdTx.GetMsgs()
			s.Require().Equal(len(msgs), 1)
			msg, ok := msgs[0].(*types.MsgSend)
			s.Require().True(ok)
			s.Require().Equal(txRecipient, msg.ToAddress)
		})
	}
}

func (s *IntegrationTestSuite) TestQueryLegacyStdTx() {
	val0 := s.network.Validators[0]

	// We broadcasted a StdTx in SetupSuite.
	// We just check for a non-empty TxHash here, the actual hash will depend on the underlying tx configuration
	s.Require().NotEmpty(s.stdTxRes.TxHash)

	s.testQueryTx(s.stdTxRes.Height, s.stdTxRes.TxHash, val0.Address.String())
}

func (s *IntegrationTestSuite) TestQueryTx() {
	val := s.network.Validators[0]

	sendTokens := sdk.NewInt64Coin(s.cfg.BondDenom, 10)
	_, _, addr := testdata.KeyTestPubAddr()

	// Might need to wait a block to refresh sequences from previous setups.
	s.Require().NoError(s.network.WaitForNextBlock())

	out, err := bankcli.MsgSendExec(
		val.ClientCtx,
		val.Address,
		addr,
		sdk.NewCoins(sendTokens),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
		fmt.Sprintf("--gas=%d", flags.DefaultGasLimit),
	)

	s.Require().NoError(err)
	var txRes sdk.TxResponse
	s.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &txRes))
	s.Require().Equal(uint32(0), txRes.Code)

	s.Require().NoError(s.network.WaitForNextBlock())

	s.testQueryTx(txRes.Height, txRes.TxHash, addr.String())
}

func (s *IntegrationTestSuite) createTestStdTx(val *network.Validator, accNum, sequence uint64) legacytx.StdTx {
	txConfig := legacytx.StdTxConfig{Cdc: s.cfg.LegacyAmino}

	msg := &types.MsgSend{
		FromAddress: val.Address.String(),
		ToAddress:   val.Address.String(),
		Amount:      sdk.Coins{sdk.NewInt64Coin(fmt.Sprintf("%stoken", val.Moniker), 100)},
	}

	// prepare txBuilder with msg
	txBuilder := txConfig.NewTxBuilder()
	feeAmount := sdk.Coins{sdk.NewInt64Coin(s.cfg.BondDenom, 10)}
	gasLimit := testdata.NewTestGasLimit()
	txBuilder.SetMsgs(msg)
	txBuilder.SetFeeAmount(feeAmount)
	txBuilder.SetGasLimit(gasLimit)
	txBuilder.SetMemo("foobar")

	// setup txFactory
	txFactory := tx.Factory{}.
		WithChainID(val.ClientCtx.ChainID).
		WithKeybase(val.ClientCtx.Keyring).
		WithTxConfig(txConfig).
		WithSignMode(signing.SignMode_SIGN_MODE_LEGACY_AMINO_JSON).
		WithAccountNumber(accNum).
		WithSequence(sequence)

	// sign Tx (offline mode so we can manually set sequence number)
	err := authclient.SignTx(txFactory, val.ClientCtx, val.Moniker, txBuilder, true, true)
	s.Require().NoError(err)

	stdTx := txBuilder.GetTx().(legacytx.StdTx)

	return stdTx
}

func (s *IntegrationTestSuite) broadcastReq(stdTx legacytx.StdTx, mode string) (*sdk.TxResponse, error) {
	val := s.network.Validators[0]

	var err error
	accSeqs := make([]uint64, len(stdTx.GetSigners()))
	for i, signer := range stdTx.GetSigners() {
		_, accSeqs[i], err = val.ClientCtx.AccountRetriever.GetAccountNumberSequence(val.ClientCtx, signer)
		if err != nil {
			return nil, err
		}
	}

	txBz, err := convertAndEncodeStdTx(val.ClientCtx.TxConfig, stdTx, accSeqs)
	if err != nil {
		return nil, err
	}

	val.ClientCtx = val.ClientCtx.WithBroadcastMode(mode)

	return val.ClientCtx.BroadcastTx(txBz)
}

// TestLegacyMultiSig creates a legacy multisig transaction, and makes sure
// we can query it via the legacy REST endpoint.
// ref: https://github.com/cosmos/cosmos-sdk/issues/8679
func (s *IntegrationTestSuite) TestLegacyMultisig() {
	val1 := *s.network.Validators[0]

	// Generate 2 accounts and a multisig.
	account1, err := val1.ClientCtx.Keyring.Key("newAccount1")
	s.Require().NoError(err)

	account2, err := val1.ClientCtx.Keyring.Key("newAccount2")
	s.Require().NoError(err)

	multisigInfo, err := val1.ClientCtx.Keyring.Key("multi")
	s.Require().NoError(err)

	// Send coins from validator to multisig.
	sendTokens := sdk.NewInt64Coin(s.cfg.BondDenom, 1000)
	_, err = bankcli.MsgSendExec(
		val1.ClientCtx,
		val1.Address,
		multisigInfo.GetAddress(),
		sdk.NewCoins(sendTokens),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
		fmt.Sprintf("--gas=%d", flags.DefaultGasLimit),
	)

	s.Require().NoError(s.network.WaitForNextBlock())

	// Generate multisig transaction to a random address.
	_, _, recipient := testdata.KeyTestPubAddr()
	multiGeneratedTx, err := bankcli.MsgSendExec(
		val1.ClientCtx,
		multisigInfo.GetAddress(),
		recipient,
		sdk.NewCoins(
			sdk.NewInt64Coin(s.cfg.BondDenom, 5),
		),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
		fmt.Sprintf("--%s=true", flags.FlagGenerateOnly),
	)
	s.Require().NoError(err)

	// Save tx to file
	multiGeneratedTxFile := testutil.WriteToNewTempFile(s.T(), multiGeneratedTx.String())

	// Sign with account1
	val1.ClientCtx.HomeDir = strings.Replace(val1.ClientCtx.HomeDir, "simd", "simcli", 1)
	account1Signature, err := authtest.TxSignExec(val1.ClientCtx, account1.GetAddress(), multiGeneratedTxFile.Name(), "--multisig", multisigInfo.GetAddress().String())
	s.Require().NoError(err)

	sign1File := testutil.WriteToNewTempFile(s.T(), account1Signature.String())

	// Sign with account1
	account2Signature, err := authtest.TxSignExec(val1.ClientCtx, account2.GetAddress(), multiGeneratedTxFile.Name(), "--multisig", multisigInfo.GetAddress().String())
	s.Require().NoError(err)

	sign2File := testutil.WriteToNewTempFile(s.T(), account2Signature.String())

	// Does not work in offline mode.
	_, err = authtest.TxMultiSignExec(val1.ClientCtx, multisigInfo.GetName(), multiGeneratedTxFile.Name(), "--offline", sign1File.Name(), sign2File.Name())
	s.Require().EqualError(err, fmt.Sprintf("couldn't verify signature for address %s", account1.GetAddress()))

	val1.ClientCtx.Offline = false
	multiSigWith2Signatures, err := authtest.TxMultiSignExec(val1.ClientCtx, multisigInfo.GetName(), multiGeneratedTxFile.Name(), sign1File.Name(), sign2File.Name())
	s.Require().NoError(err)

	// Write the output to disk
	signedTxFile := testutil.WriteToNewTempFile(s.T(), multiSigWith2Signatures.String())

	_, err = authtest.TxValidateSignaturesExec(val1.ClientCtx, signedTxFile.Name())
	s.Require().NoError(err)

	val1.ClientCtx.BroadcastMode = flags.BroadcastBlock
	out, err := authtest.TxBroadcastExec(val1.ClientCtx, signedTxFile.Name())
	s.Require().NoError(err)

	s.Require().NoError(s.network.WaitForNextBlock())

	var txRes sdk.TxResponse
	err = val1.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &txRes)
	s.Require().NoError(err)
	s.Require().Equal(uint32(0), txRes.Code)

	s.testQueryTx(txRes.Height, txRes.TxHash, recipient.String())
}

// convertAndEncodeStdTx encodes the stdTx as a transaction in the format
// specified by txConfig. Since stdTx doesn't contain the signers' sequence
// explicitly, we pass an additional accSeqs array and set them in the proto
// tx's signer_infos.
func convertAndEncodeStdTx(txConfig client.TxConfig, stdTx legacytx.StdTx, accSeqs []uint64) ([]byte, error) {
	builder := txConfig.NewTxBuilder()

	var theTx sdk.Tx

	// check if we need a StdTx anyway, in that case don't copy
	if _, ok := builder.GetTx().(legacytx.StdTx); ok {
		theTx = stdTx
	} else {
		err := tx.CopyTx(stdTx, builder, false)
		if err != nil {
			return nil, err
		}

		sigs, err := stdTx.GetSignaturesV2()
		if err != nil {
			return nil, err
		}

		// Set the correct account sequences on the proto tx signer_infos.
		if len(sigs) != len(accSeqs) {
			return nil, fmt.Errorf("expected %d accSeqs, got %d", len(sigs), len(accSeqs))
		}
		correctSigs := make([]signing.SignatureV2, len(sigs))
		for i, sig := range sigs {
			correctSigs[i] = signing.SignatureV2{
				PubKey:   sig.PubKey,
				Data:     sig.Data,
				Sequence: accSeqs[i],
			}
		}
		builder.SetSignatures(correctSigs...)

		theTx = builder.GetTx()
	}

	return txConfig.TxEncoder()(theTx)
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

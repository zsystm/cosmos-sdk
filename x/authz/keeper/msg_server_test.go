package keeper_test

import (
	"fmt"
	"time"

	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

var typeMsgDeposit = sdk.MsgTypeURL(&govtypes.MsgDeposit{})

func (s *TestSuite) TestCollectAllExecResponses() {

	require := s.Require()
	granterAddr := s.addrs[0]
	granteeAddr := s.addrs[1]
	recipientAddr := s.addrs[2]

	coinsLarge := sdk.NewCoins(sdk.NewInt64Coin("stake", 1000))
	coinsSmall := sdk.NewCoins(sdk.NewInt64Coin("stake", 10))

	sendAuthorization := &banktypes.SendAuthorization{SpendLimit: coinsLarge}
	genericAuthorization := authz.NewGenericAuthorization(typeMsgDeposit)
	sendAny, err := cdctypes.NewAnyWithValue(sendAuthorization)
	require.NoError(err)
	genericAny, err := cdctypes.NewAnyWithValue(genericAuthorization)
	require.NoError(err)

	res, err := s.msgSrvr.Grant(s.ctx, &authz.MsgGrant{
		Granter: granterAddr.String(),
		Grantee: granteeAddr.String(),
		Grant: authz.Grant{
			Authorization: sendAny,
			Expiration:    s.sdkCtx.BlockTime().Add(time.Hour * 2),
		},
	})
	require.NoError(err)
	require.NotNil(res)

	res, err = s.msgSrvr.Grant(s.ctx, &authz.MsgGrant{
		Granter: granterAddr.String(),
		Grantee: granteeAddr.String(),
		Grant: authz.Grant{
			Authorization: genericAny,
			Expiration:    s.sdkCtx.BlockTime().Add(time.Hour * 2),
		},
	})
	require.NoError(err)
	require.NotNil(res)

	anyBankSend, err := cdctypes.NewAnyWithValue(&banktypes.MsgSend{
		FromAddress: granterAddr.String(),
		ToAddress:   recipientAddr.String(),
		Amount:      coinsSmall,
	})
	require.NoError(err)

	tp := govtypes.NewTextProposal("title", "description")
	propsalRes, err := s.app.GovKeeper.SubmitProposal(s.sdkCtx, tp)
	require.NoError(err)

	anyMsgDeposit, err := cdctypes.NewAnyWithValue(&govtypes.MsgDeposit{
		ProposalId: propsalRes.ProposalId,
		Depositor:  granterAddr.String(),
		Amount:     coinsLarge,
	})
	require.NoError(err)

	execResult, err := s.msgSrvr.Exec(s.ctx, &authz.MsgExec{
		Grantee: granteeAddr.String(),
		Msgs: []*cdctypes.Any{
			anyBankSend,
			anyMsgDeposit,
		},
	})
	require.NoError(err)
	fmt.Println(execResult.Results)
	// `execResult.Results` always returning [[][]] 
	require.Equal(len(execResult.Results), 2)
	require.True(false)
}

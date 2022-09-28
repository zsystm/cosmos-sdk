package clientv2

import authzv1beta1 "cosmossdk.io/api/cosmos/authz/v1beta1"

func (c *Client) StartAuthzExec(grantee string) error {
	c.authzExec = &authzv1beta1.MsgExec{Grantee: grantee}
	return c.addMsg(c.authzExec)
}

func (c *Client) EndAuthzExec() {

}

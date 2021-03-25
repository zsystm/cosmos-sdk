package base

type MsgRequest interface {
	GetSigners() []string
	ValidateBasic() error
}

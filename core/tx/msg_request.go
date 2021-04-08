package tx

type Msg interface {
	GetSigners() []string
	ValidateBasic() error
}

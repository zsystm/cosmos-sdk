package authnv1

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (x *MsgSetCredential) ValidateBasic() error {
	if x.Address == "" {
		return status.Errorf(codes.InvalidArgument, "missing address")
	}

	if x.NewCredential == nil {
		return status.Errorf(codes.InvalidArgument, "missing new_credential")
	}

	return nil
}

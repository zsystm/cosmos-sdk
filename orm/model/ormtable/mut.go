package ormtable

import "google.golang.org/protobuf/proto"

type Updater struct {
	Value    proto.Message
	updateFn func(proto.Message) error
}

func (u *Updater) Save(message proto.Message) error {
	return u.updateFn(message)
}

func (u *Updater) Delete() error {
	return u.updateFn(nil)
}

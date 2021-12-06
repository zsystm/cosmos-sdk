package ormsql

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"reflect"
)

func (s *schema) Save(message proto.Message) error {
	cdc, err := s.getMessageCodec(message)
	if err != nil {
		return err
	}
	cdc.save(s.gormDb, message.ProtoReflect())
	return nil
}

var protoMessageType = reflect.TypeOf((*proto.Message)(nil)).Elem()

func (s *schema) Where(query interface{}, args ...interface{}) *schema {
	if protoMsg, ok := query.(proto.Message); ok {
		cdc, err := s.getMessageCodec(protoMsg)
		if err != nil {
			s.Error = err
			return s
		}

		val := cdc.encode(protoMsg.ProtoReflect())
		query = val.Interface()
	}
	s.gormDb.Where(query, args)
	if s.gormDb.Error != nil {
		s.Error = s.gormDb.Error
	}
	return s
}

func (s *schema) Find(dest interface{}) *schema {
	typ := reflect.TypeOf(dest).Elem()
	if typ.Kind() != reflect.Slice {
		s.Error = fmt.Errorf("expected a slice, got %T", dest)
		return s
	}

	elem := typ.Elem()
	if !elem.AssignableTo(protoMessageType) {
		s.Error = fmt.Errorf("expected a proto.Message slice type, got %T", dest)
		return s
	}

	msg := reflect.Zero(elem).Interface().(proto.Message)
	cdc, err := s.getMessageCodec(msg)
	if err != nil {
		s.Error = err
		return s
	}

	structSliceType := reflect.SliceOf(cdc.structType)
	structSlicePtr := reflect.New(structSliceType)
	s.gormDb.Table(cdc.tableName).Find(structSlicePtr.Interface())
	if s.gormDb.Error != nil {
		s.Error = s.gormDb.Error
	}
	structSlice := structSlicePtr.Elem()
	n := structSlice.Len()
	destVal := reflect.ValueOf(dest)
	resSlice := reflect.MakeSlice(typ, n, n)
	destVal.Elem().Set(resSlice)
	for i := 0; i < n; i++ {
		msg := cdc.msgType.New()
		err = cdc.decode(structSlice.Index(i), msg)
		if err != nil {
			s.Error = err
			return s
		}
		resSlice.Index(i).Set(reflect.ValueOf(msg.Interface()))
	}
	return s
}

func (s *schema) First(message proto.Message) *schema {
	msgCdc, err := s.messageCodecForType(message.ProtoReflect().Type())
	if err != nil {
		s.Error = err
		return s
	}

	ptr := reflect.New(msgCdc.structType)
	s.gormDb.Table(msgCdc.tableName).First(ptr.Interface())
	if s.gormDb.Error != nil {
		s.Error = s.gormDb.Error
		return s
	}

	err = msgCdc.decode(ptr.Elem(), message.ProtoReflect())
	if err != nil {
		s.Error = err
	}

	return s
}

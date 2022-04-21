package graphql

import (
	"fmt"
	"strings"

	"github.com/graphql-go/graphql"
	"github.com/iancoleman/strcase"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
)

type Builder struct {
	objects map[string]*graphql.Object
	enums   map[string]*graphql.Enum
	query   graphql.Fields
}

func NewBuilder() *Builder {
	return &Builder{
		objects: map[string]*graphql.Object{},
		enums:   map[string]*graphql.Enum{},
		query:   graphql.Fields{},
	}
}

//func (b Builder) buildTable(tableDesc *ormpb.TableDescriptor, desc protoreflect.MessageDescriptor) (*graphql.Field, error) {
//	name := descriptorName(desc)
//	objType, err := b.protoMessageToGraphqlObject(desc)
//	if err != nil {
//		return nil, err
//	}
//
//	return &graphql.Field{
//		Name:              name,
//		Type:              objType,
//		Args:              nil,
//		Resolve:           nil,
//		DeprecationReason: "",
//		Description:       getDocComments(desc),
//	}, nil
//}

func (b *Builder) RegisterFileDescriptor(descriptor protoreflect.FileDescriptor) error {
	return b.registerMessages(descriptor.Messages())
}

func (b *Builder) RegisterTable(view ormtable.View) error {
	descriptor := view.MessageType().Descriptor()
	msgObj, err := b.protoMessageToGraphqlObject(descriptor)
	if err != nil {
		return err
	}
	pkgObj := b.getProtoPackageObject(descriptor)
	allMessagesFieldName := fmt.Sprintf("all%s", descriptor.Name())
	connTyp, err := b.messageConnectionType(descriptor)
	if err != nil {
		return err
	}
	pkgObj.AddFieldConfig(allMessagesFieldName, &graphql.Field{
		Name: allMessagesFieldName,
		Type: connTyp,
	})
	for _, index := range view.Indexes() {
		if uniq, ok := index.(ormtable.UniqueIndex); ok {
			getterName := fmt.Sprintf("get%sBy%s", descriptor.Name(), fieldsToCamelCase(uniq.Fields()))
			pkgObj.AddFieldConfig(getterName, &graphql.Field{
				Name: getterName,
				Type: msgObj,
			})
		}
	}
	return err
}

func fieldsToCamelCase(fields string) string {
	splitFields := strings.Split(fields, ",")
	camelFields := make([]string, len(splitFields))
	for i, field := range splitFields {
		camelFields[i] = strcase.ToCamel(field)
	}
	return strings.Join(camelFields, "")
}

func (b *Builder) messageConnectionType(desc protoreflect.MessageDescriptor) (graphql.Output, error) {
	// TODO
	connTypName := fmt.Sprintf("%s_Connection", descriptorName(desc))
	return graphql.NewObject(graphql.ObjectConfig{
		Name: connTypName,
		Fields: graphql.Fields{
			"pageInfo": &graphql.Field{
				Name: "pageInfo",
				Type: pageInfo,
			},
		},
	}), nil
}

func (b *Builder) getProtoPackageObject(descriptor protoreflect.Descriptor) *graphql.Object {
	pkgName := string(descriptor.ParentFile().Package())
	if existing, ok := b.query[pkgName]; ok {
		return existing.Type.(*graphql.Object)
	}

	obj := graphql.NewObject(graphql.ObjectConfig{
		Name:   strings.ReplaceAll(pkgName, ".", "_"),
		Fields: graphql.Fields{},
	})
	b.query[pkgName] = &graphql.Field{
		Name:              pkgName,
		Type:              obj,
		Args:              nil,
		Resolve:           nil,
		Subscribe:         nil,
		DeprecationReason: "",
		Description:       "",
	}
	return obj
}

func (b *Builder) Build() (graphql.Schema, error) {
	return graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:        "Query",
			Fields:      b.query,
			Description: "The root query object.",
		}),
	})
}

//func Execute(schema graphql.Schema, store kv.ReadMultiStore, requestString string) {
//	graphql.Do(graphql.Params{
//		Schema:         schema,
//		RequestString:  "",
//		RootObject:     nil,
//		VariableValues: nil,
//		OperationName:  "",
//		Context:        nil,
//	})
//}

type Schema interface {
	ormtable.Schema
	GetTables() []ormtable.Table
}

//func Execute(schema graphql.Schema, ctx context.Context, requestString string) *graphql.Result {
//	return graphql.Do(graphql.Params{
//		Schema:         schema,
//		RequestString:  requestString,
//		RootObject:     nil,
//		VariableValues: nil,
//		OperationName:  "",
//		Context:        ctx,
//	})
//}

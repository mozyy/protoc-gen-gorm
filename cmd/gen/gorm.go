package gen

import (
	gorm "github.com/mozyy/protoc-gen-gorm/options"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"gorm.io/gorm/schema"
)

const (
	DeletedAt_message_fullname = "type.DeletedAt"

	contextPackage = protogen.GoImportPath("context")
)

func getMessageOrmable(message *protogen.Message) bool {
	if message == nil || message.Desc == nil {
		return false
	}
	options := message.Desc.Options()
	return proto.GetExtension(options, gorm.E_Ormable).(bool)
}
func getMessageTable(message *protogen.Message) string {
	if message == nil || message.Desc == nil {
		return ""
	}
	options := message.Desc.Options()
	return proto.GetExtension(options, gorm.E_Table).(string)
}

func getFieldTag(field *protogen.Field) string {
	options := field.Desc.Options()
	return proto.GetExtension(options, gorm.E_Tag).(string)
}

func genMessageGormMethods(g *protogen.GeneratedFile, f *fileInfo, m *messageInfo) {
	genMessageTableMethods(g, f, m)
	genMessageToPBMethods(g, f, m)
	genMessageToGORMMethods(g, f, m)
}

func genMessageTableMethods(g *protogen.GeneratedFile, f *fileInfo, m *messageInfo) {
	table := getMessageTable(m.Message)
	tableName := ""
	if table != "" {
		tableName = table
	} else {
		ns := schema.NamingStrategy{}
		tableName = ns.TableName(m.Message.GoIdent.GoName)
	}
	g.P("// TableName overrides the default tablename generated by GORM")
	g.P("func (", m.GoIdent.GoName, ") TableName() string {")
	g.P(`return "`, tableName, `"`)
	g.P("}")
}
func genMessageToPBMethods(g *protogen.GeneratedFile, f *fileInfo, m *messageInfo) {
	g.P("// ToPB  converts the fields of this object to PB object")
	g.P("func (m *", m.GoIdent.GoName, ") ToPB(ctx ", contextPackage.Ident("Context"), ") *", m.Message.GoIdent.GoName, " {")
	g.P("	to := ", m.Message.GoIdent.GoName, "{}")
	//if m.CreatedAt != nil {
	// 	to.CreatedAt = timestamppb.New(*m.CreatedAt)
	// }
	for _, field := range m.Fields {
		goType, pointer := fieldGoType(g, f, field, convOToG)
		if pointer {
			goType = "*" + goType
		}
		name := field.GoName
		switch goType {
		case "*time.Time":
			g.P("to.", name, "=", field.Message.GoIdent.GoImportPath.Ident("New"), "(*m.", name, ")")
		// case "*gorm.DeletedAt":
		// 	g.P(`deletedAtValue, _ := o.DeletedAt.Value()`)
		// 	g.P(`if deletedAt, ok := deletedAtValue.(time.Time); ok {`)
		// 	g.P(`value.DeletedAt = timestamppb.New(deletedAt)`)
		// 	g.P("}")
		default:
			if getMessageOrmable(field.Message) {
				if field.Desc.IsList() {
					g.P("for _, ", name, ":= range m.", name, " {")
					g.P("to.", name, "= append(to.", name, ", ", name, ".ToPB(ctx))")
					g.P("}")
				} else {
					g.P("to.", name, "= m.", name, ".ToPB(ctx)")
				}
			} else {
				g.P("to.", name, "= m.", name)
			}

		}
	}
	g.P("return &to")
	g.P("}")
}
func genMessageToGORMMethods(g *protogen.GeneratedFile, f *fileInfo, m *messageInfo) {
	g.P("// ToORM  converts the fields of this object to PB object")
	g.P("func (m *", m.Message.GoIdent.GoName, ") ToORM(ctx ", contextPackage.Ident("Context"), ") *", m.GoIdent.GoName, " {")
	g.P("	to := ", m.GoIdent.GoName, "{}")
	//if m.CreatedAt != nil {
	// 	to.CreatedAt = timestamppb.New(*m.CreatedAt)
	// }
	for _, field := range m.Fields {
		goType, pointer := fieldGoType(g, f, field, convOToG)
		if pointer {
			goType = "*" + goType
		}
		name := field.GoName
		switch goType {
		case "*time.Time":

			g.P(name, ":= m.Get", name, "().AsTime()")
			g.P("to.", name, "=&", name)
		// case "*gorm.DeletedAt":
		// 	g.P(`deletedAtValue, _ := o.DeletedAt.Value()`)
		// 	g.P(`if deletedAt, ok := deletedAtValue.(time.Time); ok {`)
		// 	g.P(`value.DeletedAt = timestamppb.New(deletedAt)`)
		// 	g.P("}")
		default:
			if getMessageOrmable(field.Message) {
				if field.Desc.IsList() {
					g.P("for _, ", name, ":= range m.Get", name, "() {")
					g.P("to.", name, "= append(to.", name, ", ", name, ".ToORM(ctx))")
					g.P("}")
				} else {
					g.P("to.", name, "= m.Get", name, "().ToORM(ctx)")
				}
			} else {
				g.P("to.", name, "= m.Get", name, "()")
			}

		}
	}
	g.P("return &to")
	g.P("}")
}

type fieldSort []*protogen.Field

// Len is the number of elements in the collection.
func (f fieldSort) Len() int {
	return len(f)
}

// Less reports whether the element with index i
func (f fieldSort) Less(i int, j int) bool {
	sortStrs := []string{"CreatedAt", "UpdatedAt", "DeletedAt"}
	ii := f[i].GoName
	jj := f[j].GoName
	if ii == "ID" {
		return true
	}
	if jj == "ID" {
		return false
	}
	return sliceIndex(sortStrs, ii) < sliceIndex(sortStrs, jj)
}

// Swap swaps the elements with indexes i and j.
func (f fieldSort) Swap(i int, j int) {
	f[i], f[j] = f[j], f[i]
}

func sliceIndex(strs []string, str string) int {
	for i, s := range strs {
		if s == str {
			return i
		}
	}
	return -1
}

func convGToO(m *protogen.Message) protogen.GoIdent {
	switch m.Desc.FullName() {
	// case genid.Timestamp_message_fullname:
	case "google.protobuf.Timestamp":
		return timePackage.Ident("Time")
	}
	if getMessageOrmable(m) {
		return toGORMGoIdent(m.GoIdent)
	}
	return m.GoIdent
}

func convOToG(m *protogen.Message) protogen.GoIdent {
	switch m.Desc.FullName() {
	case "google.protobuf.Timestamp":
		return timePackage.Ident("Time")
	}
	return m.GoIdent
}

func toGORMGoIdent(m protogen.GoIdent) protogen.GoIdent {
	return protogen.GoIdent{
		GoName:       m.GoName + "GORM",
		GoImportPath: m.GoImportPath,
	}
}
func fieldIsMessage(field *protogen.Field) *protogen.Message {
	switch field.Desc.Kind() {
	case protoreflect.MessageKind, protoreflect.GroupKind:
		return field.Message
	default:
		return nil
	}
}

func fieldGORMTagValue(field *protogen.Field) string {
	return getFieldTag(field)
	// options := getFieldTag(field)
	// tag := options.GetTag()
	// tags := []string{}
	// // https://gorm.io/zh_CN/docs/models.html#%E5%AD%97%E6%AE%B5%E6%A0%87%E7%AD%BE
	// if tag.GetColumn() != "" {
	// 	tags = append(tags, "column:"+tag.GetColumn())
	// }
	// if tag.GetType() != "" {
	// 	tags = append(tags, "type:"+tag.GetType())
	// }
	// if tag.GetSize() != 0 {
	// 	tags = append(tags, "size:"+strconv.Itoa(int(tag.GetSize())))
	// }
	// if tag.GetPrimaryKey() {
	// 	tags = append(tags, "primaryKey")
	// }
	// if tag.GetUnique() {
	// 	tags = append(tags, "unique")
	// }
	// if tag.GetDefault() != "" {
	// 	tags = append(tags, "default:"+tag.GetDefault())
	// }
	// if tag.GetPrecision() != "" {
	// 	tags = append(tags, "precision:"+tag.GetPrecision())
	// }
	// if tag.GetScale() != 0 {
	// 	tags = append(tags, "scale:"+strconv.Itoa(int(tag.GetScale())))
	// }
	// // TODO: add tags
	// if tag.GetIndex() {
	// 	tags = append(tags, "index")
	// }
	// if tag.GetTags() != "" {
	// 	tags = append(tags, tag.GetTags())
	// }

	// // TODO: add tags
	// return strings.Join(tags, ";")
}

// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

import (
	"sort"

	"google.golang.org/protobuf/compiler/protogen"
)

type fileInfo struct {
	*protogen.File
	allMessages      []*messageInfo
	allMessagesByPtr map[*messageInfo]int // value is index into allMessages
}

func newFileInfo(file *protogen.File) *fileInfo {
	f := &fileInfo{File: file}

	// Collect all enums, messages, and extensions in "flattened ordering".
	// See filetype.TypeBuilder.
	var walkMessages func([]*protogen.Message, func(*protogen.Message))
	walkMessages = func(messages []*protogen.Message, f func(*protogen.Message)) {
		for _, m := range messages {
			f(m)
			walkMessages(m.Messages, f)
		}
	}

	initMessageInfos := func(messages []*protogen.Message) {
		for _, message := range messages {
			if getMessageOrmable(message) {
				f.allMessages = append(f.allMessages, newMessageInfo(f, message))
			}
		}
	}

	initMessageInfos(f.Messages)
	walkMessages(f.Messages, func(m *protogen.Message) {
		initMessageInfos(m.Messages)
	})

	// Derive a reverse mapping of enum and message pointers to their index
	// in allEnums and allMessages.

	if len(f.allMessages) > 0 {
		f.allMessagesByPtr = make(map[*messageInfo]int)
		for i, m := range f.allMessages {
			f.allMessagesByPtr[m] = i
		}
	}

	return f
}

type messageInfo struct {
	*protogen.Message

	GoIdent protogen.GoIdent // name of the generated Go ORM type
	fields  []*fieldInfo
}

func newMessageInfo(f *fileInfo, message *protogen.Message) *messageInfo {
	m := &messageInfo{Message: message}
	m.GoIdent = toGORMGoIdent(message.GoIdent)
	for _, field := range m.Fields {
		if getFieldDrop(field) {
			continue
		}
		m.fields = append(m.fields, newFieldInfo(f, field))
	}
	sort.Stable(fieldSort(m.fields))
	return m
}

type fieldInfo struct {
	*protogen.Field
}

func newFieldInfo(f *fileInfo, field *protogen.Field) *fieldInfo {
	if oneof := field.Oneof; oneof != nil && !oneof.Desc.IsSynthetic() {
		// TODO
		panic("unsport oneof")
	}
	m := &fieldInfo{Field: field}
	return m
}

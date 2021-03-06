// Copyright 2014 The ql Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSES/QL-LICENSE file.

// Copyright 2015 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package types

import (
	"io"

	"github.com/pingcap/errors"
	"github.com/pingcap/parser/charset"
	"github.com/pingcap/parser/mysql"
	"github.com/pingcap/parser/opcode"
	"github.com/pingcap/parser/terror"
	ast "github.com/pingcap/parser/types"
	"github.com/pingcap/tidb/trace_util_0"
)

// IsTypeBlob returns a boolean indicating whether the tp is a blob type.
var IsTypeBlob = ast.IsTypeBlob

// IsTypeChar returns a boolean indicating
// whether the tp is the char type like a string type or a varchar type.
var IsTypeChar = ast.IsTypeChar

// IsTypeVarchar returns a boolean indicating
// whether the tp is the varchar type like a varstring type or a varchar type.
func IsTypeVarchar(tp byte) bool {
	trace_util_0.Count(_etc_00000, 0)
	return tp == mysql.TypeVarString || tp == mysql.TypeVarchar
}

// IsTypeUnspecified returns a boolean indicating whether the tp is the Unspecified type.
func IsTypeUnspecified(tp byte) bool {
	trace_util_0.Count(_etc_00000, 1)
	return tp == mysql.TypeUnspecified
}

// IsTypePrefixable returns a boolean indicating
// whether an index on a column with the tp can be defined with a prefix.
func IsTypePrefixable(tp byte) bool {
	trace_util_0.Count(_etc_00000, 2)
	return IsTypeBlob(tp) || IsTypeChar(tp)
}

// IsTypeFractionable returns a boolean indicating
// whether the tp can has time fraction.
func IsTypeFractionable(tp byte) bool {
	trace_util_0.Count(_etc_00000, 3)
	return tp == mysql.TypeDatetime || tp == mysql.TypeDuration || tp == mysql.TypeTimestamp
}

// IsTypeTime returns a boolean indicating
// whether the tp is time type like datetime, date or timestamp.
func IsTypeTime(tp byte) bool {
	trace_util_0.Count(_etc_00000, 4)
	return tp == mysql.TypeDatetime || tp == mysql.TypeDate || tp == mysql.TypeTimestamp
}

// IsTypeNumeric returns a boolean indicating whether the tp is numeric type.
func IsTypeNumeric(tp byte) bool {
	trace_util_0.Count(_etc_00000, 5)
	switch tp {
	case mysql.TypeBit, mysql.TypeTiny, mysql.TypeInt24, mysql.TypeLong, mysql.TypeLonglong, mysql.TypeNewDecimal,
		mysql.TypeDecimal, mysql.TypeFloat, mysql.TypeDouble, mysql.TypeShort:
		trace_util_0.Count(_etc_00000, 7)
		return true
	}
	trace_util_0.Count(_etc_00000, 6)
	return false
}

// IsTemporalWithDate returns a boolean indicating
// whether the tp is time type with date.
func IsTemporalWithDate(tp byte) bool {
	trace_util_0.Count(_etc_00000, 8)
	return IsTypeTime(tp)
}

// IsBinaryStr returns a boolean indicating
// whether the field type is a binary string type.
func IsBinaryStr(ft *FieldType) bool {
	trace_util_0.Count(_etc_00000, 9)
	if ft.Collate == charset.CollationBin && IsString(ft.Tp) {
		trace_util_0.Count(_etc_00000, 11)
		return true
	}
	trace_util_0.Count(_etc_00000, 10)
	return false
}

// IsNonBinaryStr returns a boolean indicating
// whether the field type is a non-binary string type.
func IsNonBinaryStr(ft *FieldType) bool {
	trace_util_0.Count(_etc_00000, 12)
	if ft.Collate != charset.CollationBin && IsString(ft.Tp) {
		trace_util_0.Count(_etc_00000, 14)
		return true
	}
	trace_util_0.Count(_etc_00000, 13)
	return false
}

// IsString returns a boolean indicating
// whether the field type is a string type.
func IsString(tp byte) bool {
	trace_util_0.Count(_etc_00000, 15)
	return IsTypeChar(tp) || IsTypeBlob(tp) || IsTypeVarchar(tp) || IsTypeUnspecified(tp)
}

var kind2Str = map[byte]string{
	KindNull:          "null",
	KindInt64:         "bigint",
	KindUint64:        "unsigned bigint",
	KindFloat32:       "float",
	KindFloat64:       "double",
	KindString:        "char",
	KindBytes:         "bytes",
	KindBinaryLiteral: "bit/hex literal",
	KindMysqlDecimal:  "decimal",
	KindMysqlDuration: "time",
	KindMysqlEnum:     "enum",
	KindMysqlBit:      "bit",
	KindMysqlSet:      "set",
	KindMysqlTime:     "datetime",
	KindInterface:     "interface",
	KindMinNotNull:    "min_not_null",
	KindMaxValue:      "max_value",
	KindRaw:           "raw",
	KindMysqlJSON:     "json",
}

// TypeStr converts tp to a string.
var TypeStr = ast.TypeStr

// KindStr converts kind to a string.
func KindStr(kind byte) (r string) {
	trace_util_0.Count(_etc_00000, 16)
	return kind2Str[kind]
}

// TypeToStr converts a field to a string.
// It is used for converting Text to Blob,
// or converting Char to Binary.
// Args:
//	tp: type enum
//	cs: charset
var TypeToStr = ast.TypeToStr

// EOFAsNil filtrates errors,
// If err is equal to io.EOF returns nil.
func EOFAsNil(err error) error {
	trace_util_0.Count(_etc_00000, 17)
	if terror.ErrorEqual(err, io.EOF) {
		trace_util_0.Count(_etc_00000, 19)
		return nil
	}
	trace_util_0.Count(_etc_00000, 18)
	return errors.Trace(err)
}

// InvOp2 returns an invalid operation error.
func InvOp2(x, y interface{}, o opcode.Op) (interface{}, error) {
	trace_util_0.Count(_etc_00000, 20)
	return nil, errors.Errorf("Invalid operation: %v %v %v (mismatched types %T and %T)", x, o, y, x, y)
}

// overflow returns an overflowed error.
func overflow(v interface{}, tp byte) error {
	trace_util_0.Count(_etc_00000, 21)
	return ErrOverflow.GenWithStack("constant %v overflows %s", v, TypeStr(tp))
}

// IsTypeTemporal checks if a type is a temporal type.
func IsTypeTemporal(tp byte) bool {
	trace_util_0.Count(_etc_00000, 22)
	switch tp {
	case mysql.TypeDuration, mysql.TypeDatetime, mysql.TypeTimestamp,
		mysql.TypeDate, mysql.TypeNewDate:
		trace_util_0.Count(_etc_00000, 24)
		return true
	}
	trace_util_0.Count(_etc_00000, 23)
	return false
}

var _etc_00000 = "types/etc.go"

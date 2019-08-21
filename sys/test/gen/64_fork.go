// AUTOGENERATED FILE
// +build !codeanalysis
// +build !syz_target syz_target,syz_os_test,syz_arch_64_fork

package gen

import . "github.com/google/syzkaller/prog"
import . "github.com/google/syzkaller/sys/test"

func init() {
	RegisterTarget(&Target{OS: "test", Arch: "64_fork", Revision: revision_64_fork, PtrSize: 8, PageSize: 8192, NumPages: 2048, DataOffset: 536870912, Syscalls: syscalls_64_fork, Resources: resources_64_fork, Structs: structDescs_64_fork, Consts: consts_64_fork}, InitTarget)
}

var resources_64_fork = []*ResourceDesc(nil)

var structDescs_64_fork = []*KeyedStruct{
	{Key: StructKey{Name: "align0"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "align0", TypeSize: 24}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int16", FldName: "f0", TypeSize: 2}}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "pad", TypeSize: 2}}, IsPad: true},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "f1", TypeSize: 4}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f2", TypeSize: 1}}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "pad", TypeSize: 1}}, IsPad: true},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int16", FldName: "f3", TypeSize: 2}}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "pad", TypeSize: 4}}, IsPad: true},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int64", FldName: "f4", TypeSize: 8}}},
	}}},
	{Key: StructKey{Name: "compare_data"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "compare_data", IsVarlen: true}, Fields: []Type{
		&StructType{Key: StructKey{Name: "align0"}, FldName: "align0"},
		&StructType{Key: StructKey{Name: "syz_bf_struct0"}, FldName: "bf0"},
		&StructType{Key: StructKey{Name: "syz_bf_struct1"}, FldName: "bf1"},
		&StructType{Key: StructKey{Name: "syz_bf_struct2"}, FldName: "bf2"},
		&StructType{Key: StructKey{Name: "syz_bf_struct3"}, FldName: "bf3"},
		&BufferType{TypeCommon: TypeCommon{TypeName: "string", FldName: "str", IsVarlen: true}, Kind: 2},
		&BufferType{TypeCommon: TypeCommon{TypeName: "array", FldName: "blob", IsVarlen: true}},
		&ArrayType{TypeCommon: TypeCommon{TypeName: "array", FldName: "arr16be", IsVarlen: true}, Type: &IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int16be", TypeSize: 2}, ArgFormat: 1}}},
	}}},
	{Key: StructKey{Name: "syz_bf_struct0"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "syz_bf_struct0", TypeSize: 32}, Fields: []Type{
		&FlagsType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "syz_bf_flags", FldName: "f0", TypeSize: 2}, BitfieldLen: 10}, Vals: []uint64{0, 1, 2}, BitMask: true},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "pad", TypeSize: 6}}, IsPad: true},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int64", FldName: "f1", TypeSize: 8}}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "const", FldName: "f2", TypeSize: 2}, BitfieldLen: 5, BitfieldMdl: true}, Val: 66},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int16", FldName: "f3", TypeSize: 2}, BitfieldOff: 5, BitfieldLen: 6}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "pad", TypeSize: 2}}, IsPad: true},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "const", FldName: "f4", TypeSize: 4}, BitfieldLen: 15}, Val: 66},
		&LenType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "len", FldName: "f5", TypeSize: 2}, BitfieldLen: 11}, Path: []string{"parent"}},
		&LenType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "len", FldName: "f6", TypeSize: 2}, ArgFormat: 1, BitfieldLen: 11}, Path: []string{"parent"}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f7", TypeSize: 1}}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "pad", TypeSize: 3}}, IsPad: true},
	}}},
	{Key: StructKey{Name: "syz_bf_struct1"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "syz_bf_struct1", TypeSize: 8}, Fields: []Type{
		&StructType{Key: StructKey{Name: "syz_bf_struct1_internal"}, FldName: "f0"},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f1", TypeSize: 1}}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "pad", TypeSize: 3}}, IsPad: true},
	}}},
	{Key: StructKey{Name: "syz_bf_struct1_internal"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "syz_bf_struct1_internal", TypeSize: 4}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "f0", TypeSize: 4}, BitfieldLen: 10, BitfieldMdl: true}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "f1", TypeSize: 4}, BitfieldOff: 10, BitfieldLen: 10, BitfieldMdl: true}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "f2", TypeSize: 4}, BitfieldOff: 20, BitfieldLen: 10}},
	}}},
	{Key: StructKey{Name: "syz_bf_struct2"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "syz_bf_struct2", TypeSize: 8}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int64", FldName: "f0", TypeSize: 8}, BitfieldLen: 4, BitfieldMdl: true}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int64", FldName: "f1", TypeSize: 8}, BitfieldOff: 4, BitfieldLen: 8, BitfieldMdl: true}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int64", FldName: "f2", TypeSize: 8}, BitfieldOff: 12, BitfieldLen: 12, BitfieldMdl: true}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int64", FldName: "f3", TypeSize: 8}, BitfieldOff: 24, BitfieldLen: 20, BitfieldMdl: true}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int64", FldName: "f4", TypeSize: 8}, BitfieldOff: 44, BitfieldLen: 16}},
	}}},
	{Key: StructKey{Name: "syz_bf_struct3"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "syz_bf_struct3", TypeSize: 8}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int64be", FldName: "f0", TypeSize: 8}, ArgFormat: 1, BitfieldLen: 4, BitfieldMdl: true}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int64be", FldName: "f1", TypeSize: 8}, ArgFormat: 1, BitfieldOff: 4, BitfieldLen: 8, BitfieldMdl: true}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int64be", FldName: "f2", TypeSize: 8}, ArgFormat: 1, BitfieldOff: 12, BitfieldLen: 12, BitfieldMdl: true}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int64be", FldName: "f3", TypeSize: 8}, ArgFormat: 1, BitfieldOff: 24, BitfieldLen: 20, BitfieldMdl: true}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int64be", FldName: "f4", TypeSize: 8}, ArgFormat: 1, BitfieldOff: 44, BitfieldLen: 16}},
	}}},
}

var syscalls_64_fork = []*Syscall{
	{Name: "syz_compare", CallName: "syz_compare", Args: []Type{
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "want", TypeSize: 8}, Type: &BufferType{TypeCommon: TypeCommon{TypeName: "string", IsVarlen: true}, Kind: 2}},
		&LenType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "bytesize", FldName: "want_len", TypeSize: 8}}, BitSize: 8, Path: []string{"want"}},
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "got", TypeSize: 8}, Type: &UnionType{Key: StructKey{Name: "compare_data"}}},
		&LenType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "bytesize", FldName: "got_len", TypeSize: 8}}, BitSize: 8, Path: []string{"got"}},
	}},
	{Name: "syz_compare_int$2", CallName: "syz_compare_int", MissingArgs: 2, Args: []Type{
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "const", FldName: "n", TypeSize: 8}}, Val: 2},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "intptr", FldName: "v0", TypeSize: 8}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "intptr", FldName: "v1", TypeSize: 8}}},
	}},
	{Name: "syz_compare_int$3", CallName: "syz_compare_int", MissingArgs: 1, Args: []Type{
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "const", FldName: "n", TypeSize: 8}}, Val: 3},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "intptr", FldName: "v0", TypeSize: 8}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "intptr", FldName: "v1", TypeSize: 8}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "intptr", FldName: "v2", TypeSize: 8}}},
	}},
	{Name: "syz_compare_int$4", CallName: "syz_compare_int", Args: []Type{
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "const", FldName: "n", TypeSize: 8}}, Val: 4},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "intptr", FldName: "v0", TypeSize: 8}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "intptr", FldName: "v1", TypeSize: 8}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "intptr", FldName: "v2", TypeSize: 8}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "intptr", FldName: "v3", TypeSize: 8}}},
	}},
	{Name: "syz_errno", CallName: "syz_errno", Args: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "v", TypeSize: 4}}},
	}},
	{Name: "syz_execute_func", CallName: "syz_execute_func", Args: []Type{
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "text", TypeSize: 8}, Type: &BufferType{TypeCommon: TypeCommon{TypeName: "text", IsVarlen: true}, Kind: 4}},
	}},
	{Name: "syz_exit", CallName: "syz_exit", Args: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "status", TypeSize: 4}}},
	}},
	{Name: "syz_mmap", CallName: "syz_mmap", Args: []Type{
		&VmaType{TypeCommon: TypeCommon{TypeName: "vma", FldName: "addr", TypeSize: 8}},
		&LenType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "len", FldName: "len", TypeSize: 8}}, Path: []string{"addr"}},
	}},
}

var consts_64_fork = []ConstValue{
	{Name: "IPPROTO_ICMPV6", Value: 58},
	{Name: "IPPROTO_TCP", Value: 6},
	{Name: "IPPROTO_UDP", Value: 17},
}

const revision_64_fork = "6eaf5712193b05fd314fb6e483f11d9937923636"

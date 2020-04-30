// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCES:
 *     lsst.alert.avsc
 *     lsst.alert.cutout.avsc
 *     lsst.alert.diaForcedSource.avsc
 *     lsst.alert.diaNondetectionLimit.avsc
 *     lsst.alert.diaSource.avsc
 *     lsst.diaObject.avsc
 *     lsst.ssObject.avsc
 */
package schema

import (
	"io"
	"fmt"

	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/vm/types"
)


type UnionNullArrayDiaSourceTypeEnum int
const (

	 UnionNullArrayDiaSourceTypeEnumNull UnionNullArrayDiaSourceTypeEnum = 0

	 UnionNullArrayDiaSourceTypeEnumArrayDiaSource UnionNullArrayDiaSourceTypeEnum = 1

)

type UnionNullArrayDiaSource struct {

	Null *types.NullVal

	ArrayDiaSource []*DiaSource

	UnionType UnionNullArrayDiaSourceTypeEnum
}

func writeUnionNullArrayDiaSource(r *UnionNullArrayDiaSource, w io.Writer) error {
	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType{
	
	case UnionNullArrayDiaSourceTypeEnumNull:
		return vm.WriteNull(r.Null, w)
        
	case UnionNullArrayDiaSourceTypeEnumArrayDiaSource:
		return writeArrayDiaSource(r.ArrayDiaSource, w)
        
	}
	return fmt.Errorf("invalid value for *UnionNullArrayDiaSource")
}

func NewUnionNullArrayDiaSource() *UnionNullArrayDiaSource {
	return &UnionNullArrayDiaSource{}
}

func (_ *UnionNullArrayDiaSource) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *UnionNullArrayDiaSource) SetInt(v int32) { panic("Unsupported operation") }
func (_ *UnionNullArrayDiaSource) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *UnionNullArrayDiaSource) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullArrayDiaSource) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *UnionNullArrayDiaSource) SetString(v string) { panic("Unsupported operation") }
func (r *UnionNullArrayDiaSource) SetLong(v int64) { 
	r.UnionType = (UnionNullArrayDiaSourceTypeEnum)(v)
}
func (r *UnionNullArrayDiaSource) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
		
		return r.Null
		
	
	case 1:
		
		r.ArrayDiaSource = make([]*DiaSource, 0)
		
		
		return (*ArrayDiaSourceWrapper)(&r.ArrayDiaSource)
		
	
	}
	panic("Unknown field index")
}
func (_ *UnionNullArrayDiaSource) SetDefault(i int) { panic("Unsupported operation") }
func (_ *UnionNullArrayDiaSource) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullArrayDiaSource) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *UnionNullArrayDiaSource) Finalize()  { }
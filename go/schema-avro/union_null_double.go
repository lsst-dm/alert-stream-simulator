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


type UnionNullDoubleTypeEnum int
const (

	 UnionNullDoubleTypeEnumNull UnionNullDoubleTypeEnum = 0

	 UnionNullDoubleTypeEnumDouble UnionNullDoubleTypeEnum = 1

)

type UnionNullDouble struct {

	Null *types.NullVal

	Double float64

	UnionType UnionNullDoubleTypeEnum
}

func writeUnionNullDouble(r *UnionNullDouble, w io.Writer) error {
	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType{
	
	case UnionNullDoubleTypeEnumNull:
		return vm.WriteNull(r.Null, w)
        
	case UnionNullDoubleTypeEnumDouble:
		return vm.WriteDouble(r.Double, w)
        
	}
	return fmt.Errorf("invalid value for *UnionNullDouble")
}

func NewUnionNullDouble() *UnionNullDouble {
	return &UnionNullDouble{}
}

func (_ *UnionNullDouble) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *UnionNullDouble) SetInt(v int32) { panic("Unsupported operation") }
func (_ *UnionNullDouble) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *UnionNullDouble) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullDouble) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *UnionNullDouble) SetString(v string) { panic("Unsupported operation") }
func (r *UnionNullDouble) SetLong(v int64) { 
	r.UnionType = (UnionNullDoubleTypeEnum)(v)
}
func (r *UnionNullDouble) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
		
		return r.Null
		
	
	case 1:
		
		
		return (*types.Double)(&r.Double)
		
	
	}
	panic("Unknown field index")
}
func (_ *UnionNullDouble) SetDefault(i int) { panic("Unsupported operation") }
func (_ *UnionNullDouble) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullDouble) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *UnionNullDouble) Finalize()  { }
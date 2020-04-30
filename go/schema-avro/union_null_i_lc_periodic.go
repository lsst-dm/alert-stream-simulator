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


type UnionNullILcPeriodicTypeEnum int
const (

	 UnionNullILcPeriodicTypeEnumNull UnionNullILcPeriodicTypeEnum = 0

	 UnionNullILcPeriodicTypeEnumILcPeriodic UnionNullILcPeriodicTypeEnum = 1

)

type UnionNullILcPeriodic struct {

	Null *types.NullVal

	ILcPeriodic *ILcPeriodic

	UnionType UnionNullILcPeriodicTypeEnum
}

func writeUnionNullILcPeriodic(r *UnionNullILcPeriodic, w io.Writer) error {
	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType{
	
	case UnionNullILcPeriodicTypeEnumNull:
		return vm.WriteNull(r.Null, w)
        
	case UnionNullILcPeriodicTypeEnumILcPeriodic:
		return writeILcPeriodic(r.ILcPeriodic, w)
        
	}
	return fmt.Errorf("invalid value for *UnionNullILcPeriodic")
}

func NewUnionNullILcPeriodic() *UnionNullILcPeriodic {
	return &UnionNullILcPeriodic{}
}

func (_ *UnionNullILcPeriodic) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *UnionNullILcPeriodic) SetInt(v int32) { panic("Unsupported operation") }
func (_ *UnionNullILcPeriodic) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *UnionNullILcPeriodic) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullILcPeriodic) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *UnionNullILcPeriodic) SetString(v string) { panic("Unsupported operation") }
func (r *UnionNullILcPeriodic) SetLong(v int64) { 
	r.UnionType = (UnionNullILcPeriodicTypeEnum)(v)
}
func (r *UnionNullILcPeriodic) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
		
		return r.Null
		
	
	case 1:
		
		r.ILcPeriodic = NewILcPeriodic()
		
		
		return r.ILcPeriodic
		
	
	}
	panic("Unknown field index")
}
func (_ *UnionNullILcPeriodic) SetDefault(i int) { panic("Unsupported operation") }
func (_ *UnionNullILcPeriodic) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullILcPeriodic) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *UnionNullILcPeriodic) Finalize()  { }

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


type UnionNullGLcNonPeriodicTypeEnum int
const (

	 UnionNullGLcNonPeriodicTypeEnumNull UnionNullGLcNonPeriodicTypeEnum = 0

	 UnionNullGLcNonPeriodicTypeEnumGLcNonPeriodic UnionNullGLcNonPeriodicTypeEnum = 1

)

type UnionNullGLcNonPeriodic struct {

	Null *types.NullVal

	GLcNonPeriodic *GLcNonPeriodic

	UnionType UnionNullGLcNonPeriodicTypeEnum
}

func writeUnionNullGLcNonPeriodic(r *UnionNullGLcNonPeriodic, w io.Writer) error {
	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType{
	
	case UnionNullGLcNonPeriodicTypeEnumNull:
		return vm.WriteNull(r.Null, w)
        
	case UnionNullGLcNonPeriodicTypeEnumGLcNonPeriodic:
		return writeGLcNonPeriodic(r.GLcNonPeriodic, w)
        
	}
	return fmt.Errorf("invalid value for *UnionNullGLcNonPeriodic")
}

func NewUnionNullGLcNonPeriodic() *UnionNullGLcNonPeriodic {
	return &UnionNullGLcNonPeriodic{}
}

func (_ *UnionNullGLcNonPeriodic) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *UnionNullGLcNonPeriodic) SetInt(v int32) { panic("Unsupported operation") }
func (_ *UnionNullGLcNonPeriodic) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *UnionNullGLcNonPeriodic) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullGLcNonPeriodic) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *UnionNullGLcNonPeriodic) SetString(v string) { panic("Unsupported operation") }
func (r *UnionNullGLcNonPeriodic) SetLong(v int64) { 
	r.UnionType = (UnionNullGLcNonPeriodicTypeEnum)(v)
}
func (r *UnionNullGLcNonPeriodic) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
		
		return r.Null
		
	
	case 1:
		
		r.GLcNonPeriodic = NewGLcNonPeriodic()
		
		
		return r.GLcNonPeriodic
		
	
	}
	panic("Unknown field index")
}
func (_ *UnionNullGLcNonPeriodic) SetDefault(i int) { panic("Unsupported operation") }
func (_ *UnionNullGLcNonPeriodic) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullGLcNonPeriodic) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *UnionNullGLcNonPeriodic) Finalize()  { }
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


type UnionNullPm_parallax_CovTypeEnum int
const (

	 UnionNullPm_parallax_CovTypeEnumNull UnionNullPm_parallax_CovTypeEnum = 0

	 UnionNullPm_parallax_CovTypeEnumPm_parallax_Cov UnionNullPm_parallax_CovTypeEnum = 1

)

type UnionNullPm_parallax_Cov struct {

	Null *types.NullVal

	Pm_parallax_Cov *Pm_parallax_Cov

	UnionType UnionNullPm_parallax_CovTypeEnum
}

func writeUnionNullPm_parallax_Cov(r *UnionNullPm_parallax_Cov, w io.Writer) error {
	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType{
	
	case UnionNullPm_parallax_CovTypeEnumNull:
		return vm.WriteNull(r.Null, w)
        
	case UnionNullPm_parallax_CovTypeEnumPm_parallax_Cov:
		return writePm_parallax_Cov(r.Pm_parallax_Cov, w)
        
	}
	return fmt.Errorf("invalid value for *UnionNullPm_parallax_Cov")
}

func NewUnionNullPm_parallax_Cov() *UnionNullPm_parallax_Cov {
	return &UnionNullPm_parallax_Cov{}
}

func (_ *UnionNullPm_parallax_Cov) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *UnionNullPm_parallax_Cov) SetInt(v int32) { panic("Unsupported operation") }
func (_ *UnionNullPm_parallax_Cov) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *UnionNullPm_parallax_Cov) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullPm_parallax_Cov) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *UnionNullPm_parallax_Cov) SetString(v string) { panic("Unsupported operation") }
func (r *UnionNullPm_parallax_Cov) SetLong(v int64) { 
	r.UnionType = (UnionNullPm_parallax_CovTypeEnum)(v)
}
func (r *UnionNullPm_parallax_Cov) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
		
		return r.Null
		
	
	case 1:
		
		r.Pm_parallax_Cov = NewPm_parallax_Cov()
		
		
		return r.Pm_parallax_Cov
		
	
	}
	panic("Unknown field index")
}
func (_ *UnionNullPm_parallax_Cov) SetDefault(i int) { panic("Unsupported operation") }
func (_ *UnionNullPm_parallax_Cov) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullPm_parallax_Cov) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *UnionNullPm_parallax_Cov) Finalize()  { }

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


type UnionNullDip_CovTypeEnum int
const (

	 UnionNullDip_CovTypeEnumNull UnionNullDip_CovTypeEnum = 0

	 UnionNullDip_CovTypeEnumDip_Cov UnionNullDip_CovTypeEnum = 1

)

type UnionNullDip_Cov struct {

	Null *types.NullVal

	Dip_Cov *Dip_Cov

	UnionType UnionNullDip_CovTypeEnum
}

func writeUnionNullDip_Cov(r *UnionNullDip_Cov, w io.Writer) error {
	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType{
	
	case UnionNullDip_CovTypeEnumNull:
		return vm.WriteNull(r.Null, w)
        
	case UnionNullDip_CovTypeEnumDip_Cov:
		return writeDip_Cov(r.Dip_Cov, w)
        
	}
	return fmt.Errorf("invalid value for *UnionNullDip_Cov")
}

func NewUnionNullDip_Cov() *UnionNullDip_Cov {
	return &UnionNullDip_Cov{}
}

func (_ *UnionNullDip_Cov) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *UnionNullDip_Cov) SetInt(v int32) { panic("Unsupported operation") }
func (_ *UnionNullDip_Cov) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *UnionNullDip_Cov) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullDip_Cov) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *UnionNullDip_Cov) SetString(v string) { panic("Unsupported operation") }
func (r *UnionNullDip_Cov) SetLong(v int64) { 
	r.UnionType = (UnionNullDip_CovTypeEnum)(v)
}
func (r *UnionNullDip_Cov) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
		
		return r.Null
		
	
	case 1:
		
		r.Dip_Cov = NewDip_Cov()
		
		
		return r.Dip_Cov
		
	
	}
	panic("Unknown field index")
}
func (_ *UnionNullDip_Cov) SetDefault(i int) { panic("Unsupported operation") }
func (_ *UnionNullDip_Cov) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullDip_Cov) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *UnionNullDip_Cov) Finalize()  { }

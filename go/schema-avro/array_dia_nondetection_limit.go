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

	"github.com/actgardner/gogen-avro/vm/types"
	"github.com/actgardner/gogen-avro/vm"
)

func writeArrayDiaNondetectionLimit(r []*DiaNondetectionLimit, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)),w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = writeDiaNondetectionLimit(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0,w)
}



type ArrayDiaNondetectionLimitWrapper []*DiaNondetectionLimit

func (_ *ArrayDiaNondetectionLimitWrapper) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *ArrayDiaNondetectionLimitWrapper) SetInt(v int32) { panic("Unsupported operation") }
func (_ *ArrayDiaNondetectionLimitWrapper) SetLong(v int64) { panic("Unsupported operation") }
func (_ *ArrayDiaNondetectionLimitWrapper) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *ArrayDiaNondetectionLimitWrapper) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *ArrayDiaNondetectionLimitWrapper) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *ArrayDiaNondetectionLimitWrapper) SetString(v string) { panic("Unsupported operation") }
func (_ *ArrayDiaNondetectionLimitWrapper) SetUnionElem(v int64) { panic("Unsupported operation") }
func (_ *ArrayDiaNondetectionLimitWrapper) Get(i int) types.Field { panic("Unsupported operation") }
func (_ *ArrayDiaNondetectionLimitWrapper) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *ArrayDiaNondetectionLimitWrapper) Finalize() { }
func (_ *ArrayDiaNondetectionLimitWrapper) SetDefault(i int) { panic("Unsupported operation") }
func (r *ArrayDiaNondetectionLimitWrapper) AppendArray() types.Field {
	var v *DiaNondetectionLimit
	
	v = NewDiaNondetectionLimit()

 	
	*r = append(*r, v)
        
        return (*r)[len(*r)-1]
        
}

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
	"github.com/actgardner/gogen-avro/compiler"
)


type DiaForcedSource struct {

	
	
		DiaForcedSourceId int64
	

	
	
		CcdVisitId int64
	

	
	
		DiaObjectId int64
	

	
	
		MidPointTai float64
	

	
	
		FilterName string
	

	
	
		PsFlux float32
	

	
	
		PsFluxErr float32
	

	
	
		TotFlux float32
	

	
	
		TotFluxErr float32
	

}

const DiaForcedSourceAvroCRC64Fingerprint = "\xecB\x9d\x82\xef(͌"

func NewDiaForcedSource() (*DiaForcedSource) {
	return &DiaForcedSource{}
}

func DeserializeDiaForcedSource(r io.Reader) (*DiaForcedSource, error) {
	t := NewDiaForcedSource()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err
	}
	return t, err
}

func DeserializeDiaForcedSourceFromSchema(r io.Reader, schema string) (*DiaForcedSource, error) {
	t := NewDiaForcedSource()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err
	}
	return t, err
}

func writeDiaForcedSource(r *DiaForcedSource, w io.Writer) error {
	var err error
	
	err = vm.WriteLong( r.DiaForcedSourceId, w)
	if err != nil {
		return err
	}
	
	err = vm.WriteLong( r.CcdVisitId, w)
	if err != nil {
		return err
	}
	
	err = vm.WriteLong( r.DiaObjectId, w)
	if err != nil {
		return err
	}
	
	err = vm.WriteDouble( r.MidPointTai, w)
	if err != nil {
		return err
	}
	
	err = vm.WriteString( r.FilterName, w)
	if err != nil {
		return err
	}
	
	err = vm.WriteFloat( r.PsFlux, w)
	if err != nil {
		return err
	}
	
	err = vm.WriteFloat( r.PsFluxErr, w)
	if err != nil {
		return err
	}
	
	err = vm.WriteFloat( r.TotFlux, w)
	if err != nil {
		return err
	}
	
	err = vm.WriteFloat( r.TotFluxErr, w)
	if err != nil {
		return err
	}
	
	return err
}

func (r *DiaForcedSource) Serialize(w io.Writer) error {
	return writeDiaForcedSource(r, w)
}

func (r *DiaForcedSource) Schema() string {
	return "{\"fields\":[{\"name\":\"diaForcedSourceId\",\"type\":\"long\"},{\"name\":\"ccdVisitId\",\"type\":\"long\"},{\"name\":\"diaObjectId\",\"type\":\"long\"},{\"name\":\"midPointTai\",\"type\":\"double\"},{\"name\":\"filterName\",\"type\":\"string\"},{\"name\":\"psFlux\",\"type\":\"float\"},{\"name\":\"psFluxErr\",\"type\":\"float\"},{\"name\":\"totFlux\",\"type\":\"float\"},{\"name\":\"totFluxErr\",\"type\":\"float\"}],\"name\":\"lsst.alert.diaForcedSource\",\"type\":\"record\"}"
}

func (r *DiaForcedSource) SchemaName() string {
	return "lsst.alert.diaForcedSource"
}

func (_ *DiaForcedSource) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *DiaForcedSource) SetInt(v int32) { panic("Unsupported operation") }
func (_ *DiaForcedSource) SetLong(v int64) { panic("Unsupported operation") }
func (_ *DiaForcedSource) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *DiaForcedSource) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *DiaForcedSource) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *DiaForcedSource) SetString(v string) { panic("Unsupported operation") }
func (_ *DiaForcedSource) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *DiaForcedSource) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
		
			return (*types.Long)(&r.DiaForcedSourceId)
		
	
	case 1:
		
		
			return (*types.Long)(&r.CcdVisitId)
		
	
	case 2:
		
		
			return (*types.Long)(&r.DiaObjectId)
		
	
	case 3:
		
		
			return (*types.Double)(&r.MidPointTai)
		
	
	case 4:
		
		
			return (*types.String)(&r.FilterName)
		
	
	case 5:
		
		
			return (*types.Float)(&r.PsFlux)
		
	
	case 6:
		
		
			return (*types.Float)(&r.PsFluxErr)
		
	
	case 7:
		
		
			return (*types.Float)(&r.TotFlux)
		
	
	case 8:
		
		
			return (*types.Float)(&r.TotFluxErr)
		
	
	}
	panic("Unknown field index")
}

func (r *DiaForcedSource) SetDefault(i int) {
	switch (i) {
	
        
	
        
	
        
	
        
	
        
	
        
	
        
	
        
	
        
	
	}
	panic("Unknown field index")
}

func (_ *DiaForcedSource) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *DiaForcedSource) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *DiaForcedSource) Finalize() { }


func (_ *DiaForcedSource) AvroCRC64Fingerprint() []byte {
  return []byte(DiaForcedSourceAvroCRC64Fingerprint)
}

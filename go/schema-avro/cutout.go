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


type Cutout struct {

	
	
		FileName string
	

	
	
		StampData []byte
	

}

const CutoutAvroCRC64Fingerprint = "\xbb\xd0$\xe1 \xd9:\x82"

func NewCutout() (*Cutout) {
	return &Cutout{}
}

func DeserializeCutout(r io.Reader) (*Cutout, error) {
	t := NewCutout()
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

func DeserializeCutoutFromSchema(r io.Reader, schema string) (*Cutout, error) {
	t := NewCutout()

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

func writeCutout(r *Cutout, w io.Writer) error {
	var err error
	
	err = vm.WriteString( r.FileName, w)
	if err != nil {
		return err
	}
	
	err = vm.WriteBytes( r.StampData, w)
	if err != nil {
		return err
	}
	
	return err
}

func (r *Cutout) Serialize(w io.Writer) error {
	return writeCutout(r, w)
}

func (r *Cutout) Schema() string {
	return "{\"fields\":[{\"name\":\"fileName\",\"type\":\"string\"},{\"name\":\"stampData\",\"type\":\"bytes\"}],\"name\":\"lsst.alert.cutout\",\"type\":\"record\"}"
}

func (r *Cutout) SchemaName() string {
	return "lsst.alert.cutout"
}

func (_ *Cutout) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *Cutout) SetInt(v int32) { panic("Unsupported operation") }
func (_ *Cutout) SetLong(v int64) { panic("Unsupported operation") }
func (_ *Cutout) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *Cutout) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *Cutout) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *Cutout) SetString(v string) { panic("Unsupported operation") }
func (_ *Cutout) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Cutout) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
		
			return (*types.String)(&r.FileName)
		
	
	case 1:
		
		
			return (*types.Bytes)(&r.StampData)
		
	
	}
	panic("Unknown field index")
}

func (r *Cutout) SetDefault(i int) {
	switch (i) {
	
        
	
        
	
	}
	panic("Unknown field index")
}

func (_ *Cutout) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *Cutout) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *Cutout) Finalize() { }


func (_ *Cutout) AvroCRC64Fingerprint() []byte {
  return []byte(CutoutAvroCRC64Fingerprint)
}

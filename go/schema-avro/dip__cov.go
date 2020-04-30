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


type Dip_Cov struct {

	
	
		DipMeanFluxSigma float32
	

	
	
		DipFluxDiffSigma float32
	

	
	
		DipRaSigma float32
	

	
	
		DipDeclSigma float32
	

	
	
		DipLengthSigma float32
	

	
	
		DipAngleSigma float32
	

	
	
		DipMeanFlux_dipFluxDiff_Cov float32
	

	
	
		DipMeanFlux_dipRa_Cov float32
	

	
	
		DipMeanFlux_dipDecl_Cov float32
	

	
	
		DipMeanFlux_dipLength_Cov float32
	

	
	
		DipMeanFlux_dipAngle_Cov float32
	

	
	
		DipFluxDiff_dipRa_Cov float32
	

	
	
		DipFluxDiff_dipDecl_Cov float32
	

	
	
		DipFluxDiff_dipLength_Cov float32
	

	
	
		DipFluxDiff_dipAngle_Cov float32
	

	
	
		DipRa_dipDecl_Cov float32
	

	
	
		DipRa_dipLength_Cov float32
	

	
	
		DipRa_dipAngle_Cov float32
	

	
	
		DipDecl_dipLength_Cov float32
	

	
	
		DipDecl_dipAngle_Cov float32
	

	
	
		DipLength_dipAngle_Cov float32
	

}

const Dip_CovAvroCRC64Fingerprint = "\x8c\x8a\x80Q\x89\xe6R\x8a"

func NewDip_Cov() (*Dip_Cov) {
	return &Dip_Cov{}
}

func DeserializeDip_Cov(r io.Reader) (*Dip_Cov, error) {
	t := NewDip_Cov()
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

func DeserializeDip_CovFromSchema(r io.Reader, schema string) (*Dip_Cov, error) {
	t := NewDip_Cov()

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

func writeDip_Cov(r *Dip_Cov, w io.Writer) error {
	var err error
	
	err = vm.WriteFloat( r.DipMeanFluxSigma, w)
	if err != nil {
		return err
	}
	
	err = vm.WriteFloat( r.DipFluxDiffSigma, w)
	if err != nil {
		return err
	}
	
	err = vm.WriteFloat( r.DipRaSigma, w)
	if err != nil {
		return err
	}
	
	err = vm.WriteFloat( r.DipDeclSigma, w)
	if err != nil {
		return err
	}
	
	err = vm.WriteFloat( r.DipLengthSigma, w)
	if err != nil {
		return err
	}
	
	err = vm.WriteFloat( r.DipAngleSigma, w)
	if err != nil {
		return err
	}
	
	err = vm.WriteFloat( r.DipMeanFlux_dipFluxDiff_Cov, w)
	if err != nil {
		return err
	}
	
	err = vm.WriteFloat( r.DipMeanFlux_dipRa_Cov, w)
	if err != nil {
		return err
	}
	
	err = vm.WriteFloat( r.DipMeanFlux_dipDecl_Cov, w)
	if err != nil {
		return err
	}
	
	err = vm.WriteFloat( r.DipMeanFlux_dipLength_Cov, w)
	if err != nil {
		return err
	}
	
	err = vm.WriteFloat( r.DipMeanFlux_dipAngle_Cov, w)
	if err != nil {
		return err
	}
	
	err = vm.WriteFloat( r.DipFluxDiff_dipRa_Cov, w)
	if err != nil {
		return err
	}
	
	err = vm.WriteFloat( r.DipFluxDiff_dipDecl_Cov, w)
	if err != nil {
		return err
	}
	
	err = vm.WriteFloat( r.DipFluxDiff_dipLength_Cov, w)
	if err != nil {
		return err
	}
	
	err = vm.WriteFloat( r.DipFluxDiff_dipAngle_Cov, w)
	if err != nil {
		return err
	}
	
	err = vm.WriteFloat( r.DipRa_dipDecl_Cov, w)
	if err != nil {
		return err
	}
	
	err = vm.WriteFloat( r.DipRa_dipLength_Cov, w)
	if err != nil {
		return err
	}
	
	err = vm.WriteFloat( r.DipRa_dipAngle_Cov, w)
	if err != nil {
		return err
	}
	
	err = vm.WriteFloat( r.DipDecl_dipLength_Cov, w)
	if err != nil {
		return err
	}
	
	err = vm.WriteFloat( r.DipDecl_dipAngle_Cov, w)
	if err != nil {
		return err
	}
	
	err = vm.WriteFloat( r.DipLength_dipAngle_Cov, w)
	if err != nil {
		return err
	}
	
	return err
}

func (r *Dip_Cov) Serialize(w io.Writer) error {
	return writeDip_Cov(r, w)
}

func (r *Dip_Cov) Schema() string {
	return "{\"fields\":[{\"name\":\"dipMeanFluxSigma\",\"type\":\"float\"},{\"name\":\"dipFluxDiffSigma\",\"type\":\"float\"},{\"name\":\"dipRaSigma\",\"type\":\"float\"},{\"name\":\"dipDeclSigma\",\"type\":\"float\"},{\"name\":\"dipLengthSigma\",\"type\":\"float\"},{\"name\":\"dipAngleSigma\",\"type\":\"float\"},{\"name\":\"dipMeanFlux_dipFluxDiff_Cov\",\"type\":\"float\"},{\"name\":\"dipMeanFlux_dipRa_Cov\",\"type\":\"float\"},{\"name\":\"dipMeanFlux_dipDecl_Cov\",\"type\":\"float\"},{\"name\":\"dipMeanFlux_dipLength_Cov\",\"type\":\"float\"},{\"name\":\"dipMeanFlux_dipAngle_Cov\",\"type\":\"float\"},{\"name\":\"dipFluxDiff_dipRa_Cov\",\"type\":\"float\"},{\"name\":\"dipFluxDiff_dipDecl_Cov\",\"type\":\"float\"},{\"name\":\"dipFluxDiff_dipLength_Cov\",\"type\":\"float\"},{\"name\":\"dipFluxDiff_dipAngle_Cov\",\"type\":\"float\"},{\"name\":\"dipRa_dipDecl_Cov\",\"type\":\"float\"},{\"name\":\"dipRa_dipLength_Cov\",\"type\":\"float\"},{\"name\":\"dipRa_dipAngle_Cov\",\"type\":\"float\"},{\"name\":\"dipDecl_dipLength_Cov\",\"type\":\"float\"},{\"name\":\"dipDecl_dipAngle_Cov\",\"type\":\"float\"},{\"name\":\"dipLength_dipAngle_Cov\",\"type\":\"float\"}],\"name\":\"lsst.alert.dip_Cov\",\"type\":\"record\"}"
}

func (r *Dip_Cov) SchemaName() string {
	return "lsst.alert.dip_Cov"
}

func (_ *Dip_Cov) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *Dip_Cov) SetInt(v int32) { panic("Unsupported operation") }
func (_ *Dip_Cov) SetLong(v int64) { panic("Unsupported operation") }
func (_ *Dip_Cov) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *Dip_Cov) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *Dip_Cov) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *Dip_Cov) SetString(v string) { panic("Unsupported operation") }
func (_ *Dip_Cov) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Dip_Cov) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
		
			return (*types.Float)(&r.DipMeanFluxSigma)
		
	
	case 1:
		
		
			return (*types.Float)(&r.DipFluxDiffSigma)
		
	
	case 2:
		
		
			return (*types.Float)(&r.DipRaSigma)
		
	
	case 3:
		
		
			return (*types.Float)(&r.DipDeclSigma)
		
	
	case 4:
		
		
			return (*types.Float)(&r.DipLengthSigma)
		
	
	case 5:
		
		
			return (*types.Float)(&r.DipAngleSigma)
		
	
	case 6:
		
		
			return (*types.Float)(&r.DipMeanFlux_dipFluxDiff_Cov)
		
	
	case 7:
		
		
			return (*types.Float)(&r.DipMeanFlux_dipRa_Cov)
		
	
	case 8:
		
		
			return (*types.Float)(&r.DipMeanFlux_dipDecl_Cov)
		
	
	case 9:
		
		
			return (*types.Float)(&r.DipMeanFlux_dipLength_Cov)
		
	
	case 10:
		
		
			return (*types.Float)(&r.DipMeanFlux_dipAngle_Cov)
		
	
	case 11:
		
		
			return (*types.Float)(&r.DipFluxDiff_dipRa_Cov)
		
	
	case 12:
		
		
			return (*types.Float)(&r.DipFluxDiff_dipDecl_Cov)
		
	
	case 13:
		
		
			return (*types.Float)(&r.DipFluxDiff_dipLength_Cov)
		
	
	case 14:
		
		
			return (*types.Float)(&r.DipFluxDiff_dipAngle_Cov)
		
	
	case 15:
		
		
			return (*types.Float)(&r.DipRa_dipDecl_Cov)
		
	
	case 16:
		
		
			return (*types.Float)(&r.DipRa_dipLength_Cov)
		
	
	case 17:
		
		
			return (*types.Float)(&r.DipRa_dipAngle_Cov)
		
	
	case 18:
		
		
			return (*types.Float)(&r.DipDecl_dipLength_Cov)
		
	
	case 19:
		
		
			return (*types.Float)(&r.DipDecl_dipAngle_Cov)
		
	
	case 20:
		
		
			return (*types.Float)(&r.DipLength_dipAngle_Cov)
		
	
	}
	panic("Unknown field index")
}

func (r *Dip_Cov) SetDefault(i int) {
	switch (i) {
	
        
	
        
	
        
	
        
	
        
	
        
	
        
	
        
	
        
	
        
	
        
	
        
	
        
	
        
	
        
	
        
	
        
	
        
	
        
	
        
	
        
	
	}
	panic("Unknown field index")
}

func (_ *Dip_Cov) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *Dip_Cov) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *Dip_Cov) Finalize() { }


func (_ *Dip_Cov) AvroCRC64Fingerprint() []byte {
  return []byte(Dip_CovAvroCRC64Fingerprint)
}

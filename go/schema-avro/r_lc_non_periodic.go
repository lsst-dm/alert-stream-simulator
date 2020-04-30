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


type RLcNonPeriodic struct {

	
	
		RLcNonPeriodic01 *UnionNullFloat
	

	
	
		RLcNonPeriodic02 *UnionNullFloat
	

	
	
		RLcNonPeriodic03 *UnionNullFloat
	

	
	
		RLcNonPeriodic04 *UnionNullFloat
	

	
	
		RLcNonPeriodic05 *UnionNullFloat
	

	
	
		RLcNonPeriodic06 *UnionNullFloat
	

	
	
		RLcNonPeriodic07 *UnionNullFloat
	

	
	
		RLcNonPeriodic08 *UnionNullFloat
	

	
	
		RLcNonPeriodic09 *UnionNullFloat
	

	
	
		RLcNonPeriodic10 *UnionNullFloat
	

	
	
		RLcNonPeriodic11 *UnionNullFloat
	

	
	
		RLcNonPeriodic12 *UnionNullFloat
	

	
	
		RLcNonPeriodic13 *UnionNullFloat
	

	
	
		RLcNonPeriodic14 *UnionNullFloat
	

	
	
		RLcNonPeriodic15 *UnionNullFloat
	

	
	
		RLcNonPeriodic16 *UnionNullFloat
	

	
	
		RLcNonPeriodic17 *UnionNullFloat
	

	
	
		RLcNonPeriodic18 *UnionNullFloat
	

	
	
		RLcNonPeriodic19 *UnionNullFloat
	

	
	
		RLcNonPeriodic20 *UnionNullFloat
	

}

const RLcNonPeriodicAvroCRC64Fingerprint = "\x95\xc0\x1b\x99\x13\xba(\x13"

func NewRLcNonPeriodic() (*RLcNonPeriodic) {
	return &RLcNonPeriodic{}
}

func DeserializeRLcNonPeriodic(r io.Reader) (*RLcNonPeriodic, error) {
	t := NewRLcNonPeriodic()
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

func DeserializeRLcNonPeriodicFromSchema(r io.Reader, schema string) (*RLcNonPeriodic, error) {
	t := NewRLcNonPeriodic()

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

func writeRLcNonPeriodic(r *RLcNonPeriodic, w io.Writer) error {
	var err error
	
	err = writeUnionNullFloat( r.RLcNonPeriodic01, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.RLcNonPeriodic02, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.RLcNonPeriodic03, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.RLcNonPeriodic04, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.RLcNonPeriodic05, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.RLcNonPeriodic06, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.RLcNonPeriodic07, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.RLcNonPeriodic08, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.RLcNonPeriodic09, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.RLcNonPeriodic10, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.RLcNonPeriodic11, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.RLcNonPeriodic12, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.RLcNonPeriodic13, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.RLcNonPeriodic14, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.RLcNonPeriodic15, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.RLcNonPeriodic16, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.RLcNonPeriodic17, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.RLcNonPeriodic18, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.RLcNonPeriodic19, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.RLcNonPeriodic20, w)
	if err != nil {
		return err
	}
	
	return err
}

func (r *RLcNonPeriodic) Serialize(w io.Writer) error {
	return writeRLcNonPeriodic(r, w)
}

func (r *RLcNonPeriodic) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"rLcNonPeriodic01\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"rLcNonPeriodic02\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"rLcNonPeriodic03\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"rLcNonPeriodic04\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"rLcNonPeriodic05\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"rLcNonPeriodic06\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"rLcNonPeriodic07\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"rLcNonPeriodic08\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"rLcNonPeriodic09\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"rLcNonPeriodic10\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"rLcNonPeriodic11\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"rLcNonPeriodic12\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"rLcNonPeriodic13\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"rLcNonPeriodic14\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"rLcNonPeriodic15\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"rLcNonPeriodic16\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"rLcNonPeriodic17\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"rLcNonPeriodic18\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"rLcNonPeriodic19\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"rLcNonPeriodic20\",\"type\":[\"null\",\"float\"]}],\"name\":\"lsst.rLcNonPeriodic\",\"type\":\"record\"}"
}

func (r *RLcNonPeriodic) SchemaName() string {
	return "lsst.rLcNonPeriodic"
}

func (_ *RLcNonPeriodic) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *RLcNonPeriodic) SetInt(v int32) { panic("Unsupported operation") }
func (_ *RLcNonPeriodic) SetLong(v int64) { panic("Unsupported operation") }
func (_ *RLcNonPeriodic) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *RLcNonPeriodic) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *RLcNonPeriodic) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *RLcNonPeriodic) SetString(v string) { panic("Unsupported operation") }
func (_ *RLcNonPeriodic) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *RLcNonPeriodic) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
			r.RLcNonPeriodic01 = NewUnionNullFloat()

		
		
			return r.RLcNonPeriodic01
		
	
	case 1:
		
			r.RLcNonPeriodic02 = NewUnionNullFloat()

		
		
			return r.RLcNonPeriodic02
		
	
	case 2:
		
			r.RLcNonPeriodic03 = NewUnionNullFloat()

		
		
			return r.RLcNonPeriodic03
		
	
	case 3:
		
			r.RLcNonPeriodic04 = NewUnionNullFloat()

		
		
			return r.RLcNonPeriodic04
		
	
	case 4:
		
			r.RLcNonPeriodic05 = NewUnionNullFloat()

		
		
			return r.RLcNonPeriodic05
		
	
	case 5:
		
			r.RLcNonPeriodic06 = NewUnionNullFloat()

		
		
			return r.RLcNonPeriodic06
		
	
	case 6:
		
			r.RLcNonPeriodic07 = NewUnionNullFloat()

		
		
			return r.RLcNonPeriodic07
		
	
	case 7:
		
			r.RLcNonPeriodic08 = NewUnionNullFloat()

		
		
			return r.RLcNonPeriodic08
		
	
	case 8:
		
			r.RLcNonPeriodic09 = NewUnionNullFloat()

		
		
			return r.RLcNonPeriodic09
		
	
	case 9:
		
			r.RLcNonPeriodic10 = NewUnionNullFloat()

		
		
			return r.RLcNonPeriodic10
		
	
	case 10:
		
			r.RLcNonPeriodic11 = NewUnionNullFloat()

		
		
			return r.RLcNonPeriodic11
		
	
	case 11:
		
			r.RLcNonPeriodic12 = NewUnionNullFloat()

		
		
			return r.RLcNonPeriodic12
		
	
	case 12:
		
			r.RLcNonPeriodic13 = NewUnionNullFloat()

		
		
			return r.RLcNonPeriodic13
		
	
	case 13:
		
			r.RLcNonPeriodic14 = NewUnionNullFloat()

		
		
			return r.RLcNonPeriodic14
		
	
	case 14:
		
			r.RLcNonPeriodic15 = NewUnionNullFloat()

		
		
			return r.RLcNonPeriodic15
		
	
	case 15:
		
			r.RLcNonPeriodic16 = NewUnionNullFloat()

		
		
			return r.RLcNonPeriodic16
		
	
	case 16:
		
			r.RLcNonPeriodic17 = NewUnionNullFloat()

		
		
			return r.RLcNonPeriodic17
		
	
	case 17:
		
			r.RLcNonPeriodic18 = NewUnionNullFloat()

		
		
			return r.RLcNonPeriodic18
		
	
	case 18:
		
			r.RLcNonPeriodic19 = NewUnionNullFloat()

		
		
			return r.RLcNonPeriodic19
		
	
	case 19:
		
			r.RLcNonPeriodic20 = NewUnionNullFloat()

		
		
			return r.RLcNonPeriodic20
		
	
	}
	panic("Unknown field index")
}

func (r *RLcNonPeriodic) SetDefault(i int) {
	switch (i) {
	
        
	case 0:
       	 	r.RLcNonPeriodic01 = NewUnionNullFloat()

		return
	
	
        
	case 1:
       	 	r.RLcNonPeriodic02 = NewUnionNullFloat()

		return
	
	
        
	case 2:
       	 	r.RLcNonPeriodic03 = NewUnionNullFloat()

		return
	
	
        
	case 3:
       	 	r.RLcNonPeriodic04 = NewUnionNullFloat()

		return
	
	
        
	case 4:
       	 	r.RLcNonPeriodic05 = NewUnionNullFloat()

		return
	
	
        
	case 5:
       	 	r.RLcNonPeriodic06 = NewUnionNullFloat()

		return
	
	
        
	case 6:
       	 	r.RLcNonPeriodic07 = NewUnionNullFloat()

		return
	
	
        
	case 7:
       	 	r.RLcNonPeriodic08 = NewUnionNullFloat()

		return
	
	
        
	case 8:
       	 	r.RLcNonPeriodic09 = NewUnionNullFloat()

		return
	
	
        
	case 9:
       	 	r.RLcNonPeriodic10 = NewUnionNullFloat()

		return
	
	
        
	case 10:
       	 	r.RLcNonPeriodic11 = NewUnionNullFloat()

		return
	
	
        
	case 11:
       	 	r.RLcNonPeriodic12 = NewUnionNullFloat()

		return
	
	
        
	case 12:
       	 	r.RLcNonPeriodic13 = NewUnionNullFloat()

		return
	
	
        
	case 13:
       	 	r.RLcNonPeriodic14 = NewUnionNullFloat()

		return
	
	
        
	case 14:
       	 	r.RLcNonPeriodic15 = NewUnionNullFloat()

		return
	
	
        
	case 15:
       	 	r.RLcNonPeriodic16 = NewUnionNullFloat()

		return
	
	
        
	case 16:
       	 	r.RLcNonPeriodic17 = NewUnionNullFloat()

		return
	
	
        
	case 17:
       	 	r.RLcNonPeriodic18 = NewUnionNullFloat()

		return
	
	
        
	case 18:
       	 	r.RLcNonPeriodic19 = NewUnionNullFloat()

		return
	
	
        
	case 19:
       	 	r.RLcNonPeriodic20 = NewUnionNullFloat()

		return
	
	
	}
	panic("Unknown field index")
}

func (_ *RLcNonPeriodic) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *RLcNonPeriodic) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *RLcNonPeriodic) Finalize() { }


func (_ *RLcNonPeriodic) AvroCRC64Fingerprint() []byte {
  return []byte(RLcNonPeriodicAvroCRC64Fingerprint)
}

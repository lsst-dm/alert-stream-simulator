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


type ZLcPeriodic struct {

	
	
		ZLcPeriodic01 *UnionNullFloat
	

	
	
		ZLcPeriodic02 *UnionNullFloat
	

	
	
		ZLcPeriodic03 *UnionNullFloat
	

	
	
		ZLcPeriodic04 *UnionNullFloat
	

	
	
		ZLcPeriodic05 *UnionNullFloat
	

	
	
		ZLcPeriodic06 *UnionNullFloat
	

	
	
		ZLcPeriodic07 *UnionNullFloat
	

	
	
		ZLcPeriodic08 *UnionNullFloat
	

	
	
		ZLcPeriodic09 *UnionNullFloat
	

	
	
		ZLcPeriodic10 *UnionNullFloat
	

	
	
		ZLcPeriodic11 *UnionNullFloat
	

	
	
		ZLcPeriodic12 *UnionNullFloat
	

	
	
		ZLcPeriodic13 *UnionNullFloat
	

	
	
		ZLcPeriodic14 *UnionNullFloat
	

	
	
		ZLcPeriodic15 *UnionNullFloat
	

	
	
		ZLcPeriodic16 *UnionNullFloat
	

	
	
		ZLcPeriodic17 *UnionNullFloat
	

	
	
		ZLcPeriodic18 *UnionNullFloat
	

	
	
		ZLcPeriodic19 *UnionNullFloat
	

	
	
		ZLcPeriodic20 *UnionNullFloat
	

	
	
		ZLcPeriodic21 *UnionNullFloat
	

	
	
		ZLcPeriodic22 *UnionNullFloat
	

	
	
		ZLcPeriodic23 *UnionNullFloat
	

	
	
		ZLcPeriodic24 *UnionNullFloat
	

	
	
		ZLcPeriodic25 *UnionNullFloat
	

	
	
		ZLcPeriodic26 *UnionNullFloat
	

	
	
		ZLcPeriodic27 *UnionNullFloat
	

	
	
		ZLcPeriodic28 *UnionNullFloat
	

	
	
		ZLcPeriodic29 *UnionNullFloat
	

	
	
		ZLcPeriodic30 *UnionNullFloat
	

	
	
		ZLcPeriodic31 *UnionNullFloat
	

	
	
		ZLcPeriodic32 *UnionNullFloat
	

}

const ZLcPeriodicAvroCRC64Fingerprint = "Q\nJ\x0fE\"\xb7\x0e"

func NewZLcPeriodic() (*ZLcPeriodic) {
	return &ZLcPeriodic{}
}

func DeserializeZLcPeriodic(r io.Reader) (*ZLcPeriodic, error) {
	t := NewZLcPeriodic()
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

func DeserializeZLcPeriodicFromSchema(r io.Reader, schema string) (*ZLcPeriodic, error) {
	t := NewZLcPeriodic()

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

func writeZLcPeriodic(r *ZLcPeriodic, w io.Writer) error {
	var err error
	
	err = writeUnionNullFloat( r.ZLcPeriodic01, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.ZLcPeriodic02, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.ZLcPeriodic03, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.ZLcPeriodic04, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.ZLcPeriodic05, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.ZLcPeriodic06, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.ZLcPeriodic07, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.ZLcPeriodic08, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.ZLcPeriodic09, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.ZLcPeriodic10, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.ZLcPeriodic11, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.ZLcPeriodic12, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.ZLcPeriodic13, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.ZLcPeriodic14, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.ZLcPeriodic15, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.ZLcPeriodic16, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.ZLcPeriodic17, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.ZLcPeriodic18, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.ZLcPeriodic19, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.ZLcPeriodic20, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.ZLcPeriodic21, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.ZLcPeriodic22, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.ZLcPeriodic23, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.ZLcPeriodic24, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.ZLcPeriodic25, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.ZLcPeriodic26, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.ZLcPeriodic27, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.ZLcPeriodic28, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.ZLcPeriodic29, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.ZLcPeriodic30, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.ZLcPeriodic31, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.ZLcPeriodic32, w)
	if err != nil {
		return err
	}
	
	return err
}

func (r *ZLcPeriodic) Serialize(w io.Writer) error {
	return writeZLcPeriodic(r, w)
}

func (r *ZLcPeriodic) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"zLcPeriodic01\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"zLcPeriodic02\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"zLcPeriodic03\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"zLcPeriodic04\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"zLcPeriodic05\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"zLcPeriodic06\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"zLcPeriodic07\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"zLcPeriodic08\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"zLcPeriodic09\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"zLcPeriodic10\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"zLcPeriodic11\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"zLcPeriodic12\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"zLcPeriodic13\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"zLcPeriodic14\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"zLcPeriodic15\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"zLcPeriodic16\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"zLcPeriodic17\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"zLcPeriodic18\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"zLcPeriodic19\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"zLcPeriodic20\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"zLcPeriodic21\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"zLcPeriodic22\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"zLcPeriodic23\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"zLcPeriodic24\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"zLcPeriodic25\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"zLcPeriodic26\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"zLcPeriodic27\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"zLcPeriodic28\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"zLcPeriodic29\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"zLcPeriodic30\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"zLcPeriodic31\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"zLcPeriodic32\",\"type\":[\"null\",\"float\"]}],\"name\":\"lsst.zLcPeriodic\",\"type\":\"record\"}"
}

func (r *ZLcPeriodic) SchemaName() string {
	return "lsst.zLcPeriodic"
}

func (_ *ZLcPeriodic) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *ZLcPeriodic) SetInt(v int32) { panic("Unsupported operation") }
func (_ *ZLcPeriodic) SetLong(v int64) { panic("Unsupported operation") }
func (_ *ZLcPeriodic) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *ZLcPeriodic) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *ZLcPeriodic) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *ZLcPeriodic) SetString(v string) { panic("Unsupported operation") }
func (_ *ZLcPeriodic) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ZLcPeriodic) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
			r.ZLcPeriodic01 = NewUnionNullFloat()

		
		
			return r.ZLcPeriodic01
		
	
	case 1:
		
			r.ZLcPeriodic02 = NewUnionNullFloat()

		
		
			return r.ZLcPeriodic02
		
	
	case 2:
		
			r.ZLcPeriodic03 = NewUnionNullFloat()

		
		
			return r.ZLcPeriodic03
		
	
	case 3:
		
			r.ZLcPeriodic04 = NewUnionNullFloat()

		
		
			return r.ZLcPeriodic04
		
	
	case 4:
		
			r.ZLcPeriodic05 = NewUnionNullFloat()

		
		
			return r.ZLcPeriodic05
		
	
	case 5:
		
			r.ZLcPeriodic06 = NewUnionNullFloat()

		
		
			return r.ZLcPeriodic06
		
	
	case 6:
		
			r.ZLcPeriodic07 = NewUnionNullFloat()

		
		
			return r.ZLcPeriodic07
		
	
	case 7:
		
			r.ZLcPeriodic08 = NewUnionNullFloat()

		
		
			return r.ZLcPeriodic08
		
	
	case 8:
		
			r.ZLcPeriodic09 = NewUnionNullFloat()

		
		
			return r.ZLcPeriodic09
		
	
	case 9:
		
			r.ZLcPeriodic10 = NewUnionNullFloat()

		
		
			return r.ZLcPeriodic10
		
	
	case 10:
		
			r.ZLcPeriodic11 = NewUnionNullFloat()

		
		
			return r.ZLcPeriodic11
		
	
	case 11:
		
			r.ZLcPeriodic12 = NewUnionNullFloat()

		
		
			return r.ZLcPeriodic12
		
	
	case 12:
		
			r.ZLcPeriodic13 = NewUnionNullFloat()

		
		
			return r.ZLcPeriodic13
		
	
	case 13:
		
			r.ZLcPeriodic14 = NewUnionNullFloat()

		
		
			return r.ZLcPeriodic14
		
	
	case 14:
		
			r.ZLcPeriodic15 = NewUnionNullFloat()

		
		
			return r.ZLcPeriodic15
		
	
	case 15:
		
			r.ZLcPeriodic16 = NewUnionNullFloat()

		
		
			return r.ZLcPeriodic16
		
	
	case 16:
		
			r.ZLcPeriodic17 = NewUnionNullFloat()

		
		
			return r.ZLcPeriodic17
		
	
	case 17:
		
			r.ZLcPeriodic18 = NewUnionNullFloat()

		
		
			return r.ZLcPeriodic18
		
	
	case 18:
		
			r.ZLcPeriodic19 = NewUnionNullFloat()

		
		
			return r.ZLcPeriodic19
		
	
	case 19:
		
			r.ZLcPeriodic20 = NewUnionNullFloat()

		
		
			return r.ZLcPeriodic20
		
	
	case 20:
		
			r.ZLcPeriodic21 = NewUnionNullFloat()

		
		
			return r.ZLcPeriodic21
		
	
	case 21:
		
			r.ZLcPeriodic22 = NewUnionNullFloat()

		
		
			return r.ZLcPeriodic22
		
	
	case 22:
		
			r.ZLcPeriodic23 = NewUnionNullFloat()

		
		
			return r.ZLcPeriodic23
		
	
	case 23:
		
			r.ZLcPeriodic24 = NewUnionNullFloat()

		
		
			return r.ZLcPeriodic24
		
	
	case 24:
		
			r.ZLcPeriodic25 = NewUnionNullFloat()

		
		
			return r.ZLcPeriodic25
		
	
	case 25:
		
			r.ZLcPeriodic26 = NewUnionNullFloat()

		
		
			return r.ZLcPeriodic26
		
	
	case 26:
		
			r.ZLcPeriodic27 = NewUnionNullFloat()

		
		
			return r.ZLcPeriodic27
		
	
	case 27:
		
			r.ZLcPeriodic28 = NewUnionNullFloat()

		
		
			return r.ZLcPeriodic28
		
	
	case 28:
		
			r.ZLcPeriodic29 = NewUnionNullFloat()

		
		
			return r.ZLcPeriodic29
		
	
	case 29:
		
			r.ZLcPeriodic30 = NewUnionNullFloat()

		
		
			return r.ZLcPeriodic30
		
	
	case 30:
		
			r.ZLcPeriodic31 = NewUnionNullFloat()

		
		
			return r.ZLcPeriodic31
		
	
	case 31:
		
			r.ZLcPeriodic32 = NewUnionNullFloat()

		
		
			return r.ZLcPeriodic32
		
	
	}
	panic("Unknown field index")
}

func (r *ZLcPeriodic) SetDefault(i int) {
	switch (i) {
	
        
	case 0:
       	 	r.ZLcPeriodic01 = NewUnionNullFloat()

		return
	
	
        
	case 1:
       	 	r.ZLcPeriodic02 = NewUnionNullFloat()

		return
	
	
        
	case 2:
       	 	r.ZLcPeriodic03 = NewUnionNullFloat()

		return
	
	
        
	case 3:
       	 	r.ZLcPeriodic04 = NewUnionNullFloat()

		return
	
	
        
	case 4:
       	 	r.ZLcPeriodic05 = NewUnionNullFloat()

		return
	
	
        
	case 5:
       	 	r.ZLcPeriodic06 = NewUnionNullFloat()

		return
	
	
        
	case 6:
       	 	r.ZLcPeriodic07 = NewUnionNullFloat()

		return
	
	
        
	case 7:
       	 	r.ZLcPeriodic08 = NewUnionNullFloat()

		return
	
	
        
	case 8:
       	 	r.ZLcPeriodic09 = NewUnionNullFloat()

		return
	
	
        
	case 9:
       	 	r.ZLcPeriodic10 = NewUnionNullFloat()

		return
	
	
        
	case 10:
       	 	r.ZLcPeriodic11 = NewUnionNullFloat()

		return
	
	
        
	case 11:
       	 	r.ZLcPeriodic12 = NewUnionNullFloat()

		return
	
	
        
	case 12:
       	 	r.ZLcPeriodic13 = NewUnionNullFloat()

		return
	
	
        
	case 13:
       	 	r.ZLcPeriodic14 = NewUnionNullFloat()

		return
	
	
        
	case 14:
       	 	r.ZLcPeriodic15 = NewUnionNullFloat()

		return
	
	
        
	case 15:
       	 	r.ZLcPeriodic16 = NewUnionNullFloat()

		return
	
	
        
	case 16:
       	 	r.ZLcPeriodic17 = NewUnionNullFloat()

		return
	
	
        
	case 17:
       	 	r.ZLcPeriodic18 = NewUnionNullFloat()

		return
	
	
        
	case 18:
       	 	r.ZLcPeriodic19 = NewUnionNullFloat()

		return
	
	
        
	case 19:
       	 	r.ZLcPeriodic20 = NewUnionNullFloat()

		return
	
	
        
	case 20:
       	 	r.ZLcPeriodic21 = NewUnionNullFloat()

		return
	
	
        
	case 21:
       	 	r.ZLcPeriodic22 = NewUnionNullFloat()

		return
	
	
        
	case 22:
       	 	r.ZLcPeriodic23 = NewUnionNullFloat()

		return
	
	
        
	case 23:
       	 	r.ZLcPeriodic24 = NewUnionNullFloat()

		return
	
	
        
	case 24:
       	 	r.ZLcPeriodic25 = NewUnionNullFloat()

		return
	
	
        
	case 25:
       	 	r.ZLcPeriodic26 = NewUnionNullFloat()

		return
	
	
        
	case 26:
       	 	r.ZLcPeriodic27 = NewUnionNullFloat()

		return
	
	
        
	case 27:
       	 	r.ZLcPeriodic28 = NewUnionNullFloat()

		return
	
	
        
	case 28:
       	 	r.ZLcPeriodic29 = NewUnionNullFloat()

		return
	
	
        
	case 29:
       	 	r.ZLcPeriodic30 = NewUnionNullFloat()

		return
	
	
        
	case 30:
       	 	r.ZLcPeriodic31 = NewUnionNullFloat()

		return
	
	
        
	case 31:
       	 	r.ZLcPeriodic32 = NewUnionNullFloat()

		return
	
	
	}
	panic("Unknown field index")
}

func (_ *ZLcPeriodic) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *ZLcPeriodic) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *ZLcPeriodic) Finalize() { }


func (_ *ZLcPeriodic) AvroCRC64Fingerprint() []byte {
  return []byte(ZLcPeriodicAvroCRC64Fingerprint)
}

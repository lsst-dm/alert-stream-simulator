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


type YLcPeriodic struct {

	
	
		YLcPeriodic01 *UnionNullFloat
	

	
	
		YLcPeriodic02 *UnionNullFloat
	

	
	
		YLcPeriodic03 *UnionNullFloat
	

	
	
		YLcPeriodic04 *UnionNullFloat
	

	
	
		YLcPeriodic05 *UnionNullFloat
	

	
	
		YLcPeriodic06 *UnionNullFloat
	

	
	
		YLcPeriodic07 *UnionNullFloat
	

	
	
		YLcPeriodic08 *UnionNullFloat
	

	
	
		YLcPeriodic09 *UnionNullFloat
	

	
	
		YLcPeriodic10 *UnionNullFloat
	

	
	
		YLcPeriodic11 *UnionNullFloat
	

	
	
		YLcPeriodic12 *UnionNullFloat
	

	
	
		YLcPeriodic13 *UnionNullFloat
	

	
	
		YLcPeriodic14 *UnionNullFloat
	

	
	
		YLcPeriodic15 *UnionNullFloat
	

	
	
		YLcPeriodic16 *UnionNullFloat
	

	
	
		YLcPeriodic17 *UnionNullFloat
	

	
	
		YLcPeriodic18 *UnionNullFloat
	

	
	
		YLcPeriodic19 *UnionNullFloat
	

	
	
		YLcPeriodic20 *UnionNullFloat
	

	
	
		YLcPeriodic21 *UnionNullFloat
	

	
	
		YLcPeriodic22 *UnionNullFloat
	

	
	
		YLcPeriodic23 *UnionNullFloat
	

	
	
		YLcPeriodic24 *UnionNullFloat
	

	
	
		YLcPeriodic25 *UnionNullFloat
	

	
	
		YLcPeriodic26 *UnionNullFloat
	

	
	
		YLcPeriodic27 *UnionNullFloat
	

	
	
		YLcPeriodic28 *UnionNullFloat
	

	
	
		YLcPeriodic29 *UnionNullFloat
	

	
	
		YLcPeriodic30 *UnionNullFloat
	

	
	
		YLcPeriodic31 *UnionNullFloat
	

	
	
		YLcPeriodic32 *UnionNullFloat
	

}

const YLcPeriodicAvroCRC64Fingerprint = "\xbc?\x8a㣗Q\x04"

func NewYLcPeriodic() (*YLcPeriodic) {
	return &YLcPeriodic{}
}

func DeserializeYLcPeriodic(r io.Reader) (*YLcPeriodic, error) {
	t := NewYLcPeriodic()
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

func DeserializeYLcPeriodicFromSchema(r io.Reader, schema string) (*YLcPeriodic, error) {
	t := NewYLcPeriodic()

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

func writeYLcPeriodic(r *YLcPeriodic, w io.Writer) error {
	var err error
	
	err = writeUnionNullFloat( r.YLcPeriodic01, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.YLcPeriodic02, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.YLcPeriodic03, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.YLcPeriodic04, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.YLcPeriodic05, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.YLcPeriodic06, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.YLcPeriodic07, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.YLcPeriodic08, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.YLcPeriodic09, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.YLcPeriodic10, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.YLcPeriodic11, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.YLcPeriodic12, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.YLcPeriodic13, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.YLcPeriodic14, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.YLcPeriodic15, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.YLcPeriodic16, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.YLcPeriodic17, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.YLcPeriodic18, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.YLcPeriodic19, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.YLcPeriodic20, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.YLcPeriodic21, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.YLcPeriodic22, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.YLcPeriodic23, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.YLcPeriodic24, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.YLcPeriodic25, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.YLcPeriodic26, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.YLcPeriodic27, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.YLcPeriodic28, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.YLcPeriodic29, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.YLcPeriodic30, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.YLcPeriodic31, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.YLcPeriodic32, w)
	if err != nil {
		return err
	}
	
	return err
}

func (r *YLcPeriodic) Serialize(w io.Writer) error {
	return writeYLcPeriodic(r, w)
}

func (r *YLcPeriodic) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"yLcPeriodic01\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"yLcPeriodic02\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"yLcPeriodic03\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"yLcPeriodic04\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"yLcPeriodic05\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"yLcPeriodic06\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"yLcPeriodic07\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"yLcPeriodic08\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"yLcPeriodic09\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"yLcPeriodic10\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"yLcPeriodic11\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"yLcPeriodic12\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"yLcPeriodic13\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"yLcPeriodic14\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"yLcPeriodic15\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"yLcPeriodic16\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"yLcPeriodic17\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"yLcPeriodic18\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"yLcPeriodic19\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"yLcPeriodic20\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"yLcPeriodic21\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"yLcPeriodic22\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"yLcPeriodic23\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"yLcPeriodic24\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"yLcPeriodic25\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"yLcPeriodic26\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"yLcPeriodic27\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"yLcPeriodic28\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"yLcPeriodic29\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"yLcPeriodic30\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"yLcPeriodic31\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"yLcPeriodic32\",\"type\":[\"null\",\"float\"]}],\"name\":\"lsst.yLcPeriodic\",\"type\":\"record\"}"
}

func (r *YLcPeriodic) SchemaName() string {
	return "lsst.yLcPeriodic"
}

func (_ *YLcPeriodic) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *YLcPeriodic) SetInt(v int32) { panic("Unsupported operation") }
func (_ *YLcPeriodic) SetLong(v int64) { panic("Unsupported operation") }
func (_ *YLcPeriodic) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *YLcPeriodic) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *YLcPeriodic) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *YLcPeriodic) SetString(v string) { panic("Unsupported operation") }
func (_ *YLcPeriodic) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *YLcPeriodic) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
			r.YLcPeriodic01 = NewUnionNullFloat()

		
		
			return r.YLcPeriodic01
		
	
	case 1:
		
			r.YLcPeriodic02 = NewUnionNullFloat()

		
		
			return r.YLcPeriodic02
		
	
	case 2:
		
			r.YLcPeriodic03 = NewUnionNullFloat()

		
		
			return r.YLcPeriodic03
		
	
	case 3:
		
			r.YLcPeriodic04 = NewUnionNullFloat()

		
		
			return r.YLcPeriodic04
		
	
	case 4:
		
			r.YLcPeriodic05 = NewUnionNullFloat()

		
		
			return r.YLcPeriodic05
		
	
	case 5:
		
			r.YLcPeriodic06 = NewUnionNullFloat()

		
		
			return r.YLcPeriodic06
		
	
	case 6:
		
			r.YLcPeriodic07 = NewUnionNullFloat()

		
		
			return r.YLcPeriodic07
		
	
	case 7:
		
			r.YLcPeriodic08 = NewUnionNullFloat()

		
		
			return r.YLcPeriodic08
		
	
	case 8:
		
			r.YLcPeriodic09 = NewUnionNullFloat()

		
		
			return r.YLcPeriodic09
		
	
	case 9:
		
			r.YLcPeriodic10 = NewUnionNullFloat()

		
		
			return r.YLcPeriodic10
		
	
	case 10:
		
			r.YLcPeriodic11 = NewUnionNullFloat()

		
		
			return r.YLcPeriodic11
		
	
	case 11:
		
			r.YLcPeriodic12 = NewUnionNullFloat()

		
		
			return r.YLcPeriodic12
		
	
	case 12:
		
			r.YLcPeriodic13 = NewUnionNullFloat()

		
		
			return r.YLcPeriodic13
		
	
	case 13:
		
			r.YLcPeriodic14 = NewUnionNullFloat()

		
		
			return r.YLcPeriodic14
		
	
	case 14:
		
			r.YLcPeriodic15 = NewUnionNullFloat()

		
		
			return r.YLcPeriodic15
		
	
	case 15:
		
			r.YLcPeriodic16 = NewUnionNullFloat()

		
		
			return r.YLcPeriodic16
		
	
	case 16:
		
			r.YLcPeriodic17 = NewUnionNullFloat()

		
		
			return r.YLcPeriodic17
		
	
	case 17:
		
			r.YLcPeriodic18 = NewUnionNullFloat()

		
		
			return r.YLcPeriodic18
		
	
	case 18:
		
			r.YLcPeriodic19 = NewUnionNullFloat()

		
		
			return r.YLcPeriodic19
		
	
	case 19:
		
			r.YLcPeriodic20 = NewUnionNullFloat()

		
		
			return r.YLcPeriodic20
		
	
	case 20:
		
			r.YLcPeriodic21 = NewUnionNullFloat()

		
		
			return r.YLcPeriodic21
		
	
	case 21:
		
			r.YLcPeriodic22 = NewUnionNullFloat()

		
		
			return r.YLcPeriodic22
		
	
	case 22:
		
			r.YLcPeriodic23 = NewUnionNullFloat()

		
		
			return r.YLcPeriodic23
		
	
	case 23:
		
			r.YLcPeriodic24 = NewUnionNullFloat()

		
		
			return r.YLcPeriodic24
		
	
	case 24:
		
			r.YLcPeriodic25 = NewUnionNullFloat()

		
		
			return r.YLcPeriodic25
		
	
	case 25:
		
			r.YLcPeriodic26 = NewUnionNullFloat()

		
		
			return r.YLcPeriodic26
		
	
	case 26:
		
			r.YLcPeriodic27 = NewUnionNullFloat()

		
		
			return r.YLcPeriodic27
		
	
	case 27:
		
			r.YLcPeriodic28 = NewUnionNullFloat()

		
		
			return r.YLcPeriodic28
		
	
	case 28:
		
			r.YLcPeriodic29 = NewUnionNullFloat()

		
		
			return r.YLcPeriodic29
		
	
	case 29:
		
			r.YLcPeriodic30 = NewUnionNullFloat()

		
		
			return r.YLcPeriodic30
		
	
	case 30:
		
			r.YLcPeriodic31 = NewUnionNullFloat()

		
		
			return r.YLcPeriodic31
		
	
	case 31:
		
			r.YLcPeriodic32 = NewUnionNullFloat()

		
		
			return r.YLcPeriodic32
		
	
	}
	panic("Unknown field index")
}

func (r *YLcPeriodic) SetDefault(i int) {
	switch (i) {
	
        
	case 0:
       	 	r.YLcPeriodic01 = NewUnionNullFloat()

		return
	
	
        
	case 1:
       	 	r.YLcPeriodic02 = NewUnionNullFloat()

		return
	
	
        
	case 2:
       	 	r.YLcPeriodic03 = NewUnionNullFloat()

		return
	
	
        
	case 3:
       	 	r.YLcPeriodic04 = NewUnionNullFloat()

		return
	
	
        
	case 4:
       	 	r.YLcPeriodic05 = NewUnionNullFloat()

		return
	
	
        
	case 5:
       	 	r.YLcPeriodic06 = NewUnionNullFloat()

		return
	
	
        
	case 6:
       	 	r.YLcPeriodic07 = NewUnionNullFloat()

		return
	
	
        
	case 7:
       	 	r.YLcPeriodic08 = NewUnionNullFloat()

		return
	
	
        
	case 8:
       	 	r.YLcPeriodic09 = NewUnionNullFloat()

		return
	
	
        
	case 9:
       	 	r.YLcPeriodic10 = NewUnionNullFloat()

		return
	
	
        
	case 10:
       	 	r.YLcPeriodic11 = NewUnionNullFloat()

		return
	
	
        
	case 11:
       	 	r.YLcPeriodic12 = NewUnionNullFloat()

		return
	
	
        
	case 12:
       	 	r.YLcPeriodic13 = NewUnionNullFloat()

		return
	
	
        
	case 13:
       	 	r.YLcPeriodic14 = NewUnionNullFloat()

		return
	
	
        
	case 14:
       	 	r.YLcPeriodic15 = NewUnionNullFloat()

		return
	
	
        
	case 15:
       	 	r.YLcPeriodic16 = NewUnionNullFloat()

		return
	
	
        
	case 16:
       	 	r.YLcPeriodic17 = NewUnionNullFloat()

		return
	
	
        
	case 17:
       	 	r.YLcPeriodic18 = NewUnionNullFloat()

		return
	
	
        
	case 18:
       	 	r.YLcPeriodic19 = NewUnionNullFloat()

		return
	
	
        
	case 19:
       	 	r.YLcPeriodic20 = NewUnionNullFloat()

		return
	
	
        
	case 20:
       	 	r.YLcPeriodic21 = NewUnionNullFloat()

		return
	
	
        
	case 21:
       	 	r.YLcPeriodic22 = NewUnionNullFloat()

		return
	
	
        
	case 22:
       	 	r.YLcPeriodic23 = NewUnionNullFloat()

		return
	
	
        
	case 23:
       	 	r.YLcPeriodic24 = NewUnionNullFloat()

		return
	
	
        
	case 24:
       	 	r.YLcPeriodic25 = NewUnionNullFloat()

		return
	
	
        
	case 25:
       	 	r.YLcPeriodic26 = NewUnionNullFloat()

		return
	
	
        
	case 26:
       	 	r.YLcPeriodic27 = NewUnionNullFloat()

		return
	
	
        
	case 27:
       	 	r.YLcPeriodic28 = NewUnionNullFloat()

		return
	
	
        
	case 28:
       	 	r.YLcPeriodic29 = NewUnionNullFloat()

		return
	
	
        
	case 29:
       	 	r.YLcPeriodic30 = NewUnionNullFloat()

		return
	
	
        
	case 30:
       	 	r.YLcPeriodic31 = NewUnionNullFloat()

		return
	
	
        
	case 31:
       	 	r.YLcPeriodic32 = NewUnionNullFloat()

		return
	
	
	}
	panic("Unknown field index")
}

func (_ *YLcPeriodic) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *YLcPeriodic) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *YLcPeriodic) Finalize() { }


func (_ *YLcPeriodic) AvroCRC64Fingerprint() []byte {
  return []byte(YLcPeriodicAvroCRC64Fingerprint)
}

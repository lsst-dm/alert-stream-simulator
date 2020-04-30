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


type SsObject struct {

	
	
		SsObjectId int64
	

	
	
		Oe *UnionNullOe
	

	
	
		Oe_Cov *UnionNullOe_Cov
	

	
	
		Arc *UnionNullFloat
	

	
	
		OrbFitLnL *UnionNullFloat
	

	
	
		OrbFitChi2 *UnionNullFloat
	

	
	
		OrbFitNdata *UnionNullInt
	

	
	
		MOID1 *UnionNullFloat
	

	
	
		MOID2 *UnionNullFloat
	

	
	
		MoidLon1 *UnionNullDouble
	

	
	
		MoidLon2 *UnionNullDouble
	

	
	
		UH *UnionNullFloat
	

	
	
		UHErr *UnionNullFloat
	

	
	
		UG1 *UnionNullFloat
	

	
	
		UG1Err *UnionNullFloat
	

	
	
		UG2 *UnionNullFloat
	

	
	
		UG2Err *UnionNullFloat
	

	
	
		GH *UnionNullFloat
	

	
	
		GHErr *UnionNullFloat
	

	
	
		GG1 *UnionNullFloat
	

	
	
		GG1Err *UnionNullFloat
	

	
	
		GG2 *UnionNullFloat
	

	
	
		GG2Err *UnionNullFloat
	

	
	
		RH *UnionNullFloat
	

	
	
		RHErr *UnionNullFloat
	

	
	
		RG1 *UnionNullFloat
	

	
	
		RG1Err *UnionNullFloat
	

	
	
		RG2 *UnionNullFloat
	

	
	
		RG2Err *UnionNullFloat
	

	
	
		IH *UnionNullFloat
	

	
	
		IHErr *UnionNullFloat
	

	
	
		IG1 *UnionNullFloat
	

	
	
		IG1Err *UnionNullFloat
	

	
	
		IG2 *UnionNullFloat
	

	
	
		IG2Err *UnionNullFloat
	

	
	
		ZH *UnionNullFloat
	

	
	
		ZHErr *UnionNullFloat
	

	
	
		ZG1 *UnionNullFloat
	

	
	
		ZG1Err *UnionNullFloat
	

	
	
		ZG2 *UnionNullFloat
	

	
	
		ZG2Err *UnionNullFloat
	

	
	
		YH *UnionNullFloat
	

	
	
		YHErr *UnionNullFloat
	

	
	
		YG1 *UnionNullFloat
	

	
	
		YG1Err *UnionNullFloat
	

	
	
		YG2 *UnionNullFloat
	

	
	
		YG2Err *UnionNullFloat
	

	
	
		Flags int64
	

}

const SsObjectAvroCRC64Fingerprint = "\xd3\f\xa2\x8aD\xde}\xca"

func NewSsObject() (*SsObject) {
	return &SsObject{}
}

func DeserializeSsObject(r io.Reader) (*SsObject, error) {
	t := NewSsObject()
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

func DeserializeSsObjectFromSchema(r io.Reader, schema string) (*SsObject, error) {
	t := NewSsObject()

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

func writeSsObject(r *SsObject, w io.Writer) error {
	var err error
	
	err = vm.WriteLong( r.SsObjectId, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullOe( r.Oe, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullOe_Cov( r.Oe_Cov, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.Arc, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.OrbFitLnL, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.OrbFitChi2, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullInt( r.OrbFitNdata, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.MOID1, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.MOID2, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullDouble( r.MoidLon1, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullDouble( r.MoidLon2, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.UH, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.UHErr, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.UG1, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.UG1Err, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.UG2, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.UG2Err, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.GH, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.GHErr, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.GG1, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.GG1Err, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.GG2, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.GG2Err, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.RH, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.RHErr, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.RG1, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.RG1Err, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.RG2, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.RG2Err, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.IH, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.IHErr, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.IG1, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.IG1Err, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.IG2, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.IG2Err, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.ZH, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.ZHErr, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.ZG1, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.ZG1Err, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.ZG2, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.ZG2Err, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.YH, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.YHErr, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.YG1, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.YG1Err, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.YG2, w)
	if err != nil {
		return err
	}
	
	err = writeUnionNullFloat( r.YG2Err, w)
	if err != nil {
		return err
	}
	
	err = vm.WriteLong( r.Flags, w)
	if err != nil {
		return err
	}
	
	return err
}

func (r *SsObject) Serialize(w io.Writer) error {
	return writeSsObject(r, w)
}

func (r *SsObject) Schema() string {
	return "{\"fields\":[{\"name\":\"ssObjectId\",\"type\":\"long\"},{\"default\":null,\"name\":\"oe\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"q\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"e\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"i\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"lan\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"aop\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"M\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"epoch\",\"type\":[\"null\",\"double\"]}],\"name\":\"oe\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"oe_Cov\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"qSigma\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"eSigma\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"iSigma\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"lanSigma\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"aopSigma\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"MSigma\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"epochSigma\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"q_e_Cov\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"q_i_Cov\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"q_lan_Cov\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"q_aop_Cov\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"q_M_Cov\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"q_epoch_Cov\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"e_i_Cov\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"e_lan_Cov\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"e_aop_Cov\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"e_M_Cov\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"e_epoch_Cov\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"i_lan_Cov\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"i_aop_Cov\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"i_M_Cov\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"i_epoch_Cov\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"lan_aop_Cov\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"lan_M_Cov\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"lan_epoch_Cov\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"aop_M_Cov\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"aop_epoch_Cov\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"M_epoch_Cov\",\"type\":[\"null\",\"double\"]}],\"name\":\"oe_Cov\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"arc\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"orbFitLnL\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"orbFitChi2\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"orbFitNdata\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"MOID1\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"MOID2\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"moidLon1\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"moidLon2\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"uH\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"uHErr\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"uG1\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"uG1Err\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"uG2\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"uG2Err\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"gH\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"gHErr\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"gG1\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"gG1Err\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"gG2\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"gG2Err\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"rH\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"rHErr\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"rG1\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"rG1Err\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"rG2\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"rG2Err\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"iH\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"iHErr\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"iG1\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"iG1Err\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"iG2\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"iG2Err\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"zH\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"zHErr\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"zG1\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"zG1Err\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"zG2\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"zG2Err\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"yH\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"yHErr\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"yG1\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"yG1Err\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"yG2\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"yG2Err\",\"type\":[\"null\",\"float\"]},{\"name\":\"flags\",\"type\":\"long\"}],\"name\":\"lsst.ssObject\",\"type\":\"record\"}"
}

func (r *SsObject) SchemaName() string {
	return "lsst.ssObject"
}

func (_ *SsObject) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *SsObject) SetInt(v int32) { panic("Unsupported operation") }
func (_ *SsObject) SetLong(v int64) { panic("Unsupported operation") }
func (_ *SsObject) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *SsObject) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *SsObject) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *SsObject) SetString(v string) { panic("Unsupported operation") }
func (_ *SsObject) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *SsObject) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
		
			return (*types.Long)(&r.SsObjectId)
		
	
	case 1:
		
			r.Oe = NewUnionNullOe()

		
		
			return r.Oe
		
	
	case 2:
		
			r.Oe_Cov = NewUnionNullOe_Cov()

		
		
			return r.Oe_Cov
		
	
	case 3:
		
			r.Arc = NewUnionNullFloat()

		
		
			return r.Arc
		
	
	case 4:
		
			r.OrbFitLnL = NewUnionNullFloat()

		
		
			return r.OrbFitLnL
		
	
	case 5:
		
			r.OrbFitChi2 = NewUnionNullFloat()

		
		
			return r.OrbFitChi2
		
	
	case 6:
		
			r.OrbFitNdata = NewUnionNullInt()

		
		
			return r.OrbFitNdata
		
	
	case 7:
		
			r.MOID1 = NewUnionNullFloat()

		
		
			return r.MOID1
		
	
	case 8:
		
			r.MOID2 = NewUnionNullFloat()

		
		
			return r.MOID2
		
	
	case 9:
		
			r.MoidLon1 = NewUnionNullDouble()

		
		
			return r.MoidLon1
		
	
	case 10:
		
			r.MoidLon2 = NewUnionNullDouble()

		
		
			return r.MoidLon2
		
	
	case 11:
		
			r.UH = NewUnionNullFloat()

		
		
			return r.UH
		
	
	case 12:
		
			r.UHErr = NewUnionNullFloat()

		
		
			return r.UHErr
		
	
	case 13:
		
			r.UG1 = NewUnionNullFloat()

		
		
			return r.UG1
		
	
	case 14:
		
			r.UG1Err = NewUnionNullFloat()

		
		
			return r.UG1Err
		
	
	case 15:
		
			r.UG2 = NewUnionNullFloat()

		
		
			return r.UG2
		
	
	case 16:
		
			r.UG2Err = NewUnionNullFloat()

		
		
			return r.UG2Err
		
	
	case 17:
		
			r.GH = NewUnionNullFloat()

		
		
			return r.GH
		
	
	case 18:
		
			r.GHErr = NewUnionNullFloat()

		
		
			return r.GHErr
		
	
	case 19:
		
			r.GG1 = NewUnionNullFloat()

		
		
			return r.GG1
		
	
	case 20:
		
			r.GG1Err = NewUnionNullFloat()

		
		
			return r.GG1Err
		
	
	case 21:
		
			r.GG2 = NewUnionNullFloat()

		
		
			return r.GG2
		
	
	case 22:
		
			r.GG2Err = NewUnionNullFloat()

		
		
			return r.GG2Err
		
	
	case 23:
		
			r.RH = NewUnionNullFloat()

		
		
			return r.RH
		
	
	case 24:
		
			r.RHErr = NewUnionNullFloat()

		
		
			return r.RHErr
		
	
	case 25:
		
			r.RG1 = NewUnionNullFloat()

		
		
			return r.RG1
		
	
	case 26:
		
			r.RG1Err = NewUnionNullFloat()

		
		
			return r.RG1Err
		
	
	case 27:
		
			r.RG2 = NewUnionNullFloat()

		
		
			return r.RG2
		
	
	case 28:
		
			r.RG2Err = NewUnionNullFloat()

		
		
			return r.RG2Err
		
	
	case 29:
		
			r.IH = NewUnionNullFloat()

		
		
			return r.IH
		
	
	case 30:
		
			r.IHErr = NewUnionNullFloat()

		
		
			return r.IHErr
		
	
	case 31:
		
			r.IG1 = NewUnionNullFloat()

		
		
			return r.IG1
		
	
	case 32:
		
			r.IG1Err = NewUnionNullFloat()

		
		
			return r.IG1Err
		
	
	case 33:
		
			r.IG2 = NewUnionNullFloat()

		
		
			return r.IG2
		
	
	case 34:
		
			r.IG2Err = NewUnionNullFloat()

		
		
			return r.IG2Err
		
	
	case 35:
		
			r.ZH = NewUnionNullFloat()

		
		
			return r.ZH
		
	
	case 36:
		
			r.ZHErr = NewUnionNullFloat()

		
		
			return r.ZHErr
		
	
	case 37:
		
			r.ZG1 = NewUnionNullFloat()

		
		
			return r.ZG1
		
	
	case 38:
		
			r.ZG1Err = NewUnionNullFloat()

		
		
			return r.ZG1Err
		
	
	case 39:
		
			r.ZG2 = NewUnionNullFloat()

		
		
			return r.ZG2
		
	
	case 40:
		
			r.ZG2Err = NewUnionNullFloat()

		
		
			return r.ZG2Err
		
	
	case 41:
		
			r.YH = NewUnionNullFloat()

		
		
			return r.YH
		
	
	case 42:
		
			r.YHErr = NewUnionNullFloat()

		
		
			return r.YHErr
		
	
	case 43:
		
			r.YG1 = NewUnionNullFloat()

		
		
			return r.YG1
		
	
	case 44:
		
			r.YG1Err = NewUnionNullFloat()

		
		
			return r.YG1Err
		
	
	case 45:
		
			r.YG2 = NewUnionNullFloat()

		
		
			return r.YG2
		
	
	case 46:
		
			r.YG2Err = NewUnionNullFloat()

		
		
			return r.YG2Err
		
	
	case 47:
		
		
			return (*types.Long)(&r.Flags)
		
	
	}
	panic("Unknown field index")
}

func (r *SsObject) SetDefault(i int) {
	switch (i) {
	
        
	
        
	case 1:
       	 	r.Oe = NewUnionNullOe()

		return
	
	
        
	case 2:
       	 	r.Oe_Cov = NewUnionNullOe_Cov()

		return
	
	
        
	case 3:
       	 	r.Arc = NewUnionNullFloat()

		return
	
	
        
	case 4:
       	 	r.OrbFitLnL = NewUnionNullFloat()

		return
	
	
        
	case 5:
       	 	r.OrbFitChi2 = NewUnionNullFloat()

		return
	
	
        
	case 6:
       	 	r.OrbFitNdata = NewUnionNullInt()

		return
	
	
        
	case 7:
       	 	r.MOID1 = NewUnionNullFloat()

		return
	
	
        
	case 8:
       	 	r.MOID2 = NewUnionNullFloat()

		return
	
	
        
	case 9:
       	 	r.MoidLon1 = NewUnionNullDouble()

		return
	
	
        
	case 10:
       	 	r.MoidLon2 = NewUnionNullDouble()

		return
	
	
        
	case 11:
       	 	r.UH = NewUnionNullFloat()

		return
	
	
        
	case 12:
       	 	r.UHErr = NewUnionNullFloat()

		return
	
	
        
	case 13:
       	 	r.UG1 = NewUnionNullFloat()

		return
	
	
        
	case 14:
       	 	r.UG1Err = NewUnionNullFloat()

		return
	
	
        
	case 15:
       	 	r.UG2 = NewUnionNullFloat()

		return
	
	
        
	case 16:
       	 	r.UG2Err = NewUnionNullFloat()

		return
	
	
        
	case 17:
       	 	r.GH = NewUnionNullFloat()

		return
	
	
        
	case 18:
       	 	r.GHErr = NewUnionNullFloat()

		return
	
	
        
	case 19:
       	 	r.GG1 = NewUnionNullFloat()

		return
	
	
        
	case 20:
       	 	r.GG1Err = NewUnionNullFloat()

		return
	
	
        
	case 21:
       	 	r.GG2 = NewUnionNullFloat()

		return
	
	
        
	case 22:
       	 	r.GG2Err = NewUnionNullFloat()

		return
	
	
        
	case 23:
       	 	r.RH = NewUnionNullFloat()

		return
	
	
        
	case 24:
       	 	r.RHErr = NewUnionNullFloat()

		return
	
	
        
	case 25:
       	 	r.RG1 = NewUnionNullFloat()

		return
	
	
        
	case 26:
       	 	r.RG1Err = NewUnionNullFloat()

		return
	
	
        
	case 27:
       	 	r.RG2 = NewUnionNullFloat()

		return
	
	
        
	case 28:
       	 	r.RG2Err = NewUnionNullFloat()

		return
	
	
        
	case 29:
       	 	r.IH = NewUnionNullFloat()

		return
	
	
        
	case 30:
       	 	r.IHErr = NewUnionNullFloat()

		return
	
	
        
	case 31:
       	 	r.IG1 = NewUnionNullFloat()

		return
	
	
        
	case 32:
       	 	r.IG1Err = NewUnionNullFloat()

		return
	
	
        
	case 33:
       	 	r.IG2 = NewUnionNullFloat()

		return
	
	
        
	case 34:
       	 	r.IG2Err = NewUnionNullFloat()

		return
	
	
        
	case 35:
       	 	r.ZH = NewUnionNullFloat()

		return
	
	
        
	case 36:
       	 	r.ZHErr = NewUnionNullFloat()

		return
	
	
        
	case 37:
       	 	r.ZG1 = NewUnionNullFloat()

		return
	
	
        
	case 38:
       	 	r.ZG1Err = NewUnionNullFloat()

		return
	
	
        
	case 39:
       	 	r.ZG2 = NewUnionNullFloat()

		return
	
	
        
	case 40:
       	 	r.ZG2Err = NewUnionNullFloat()

		return
	
	
        
	case 41:
       	 	r.YH = NewUnionNullFloat()

		return
	
	
        
	case 42:
       	 	r.YHErr = NewUnionNullFloat()

		return
	
	
        
	case 43:
       	 	r.YG1 = NewUnionNullFloat()

		return
	
	
        
	case 44:
       	 	r.YG1Err = NewUnionNullFloat()

		return
	
	
        
	case 45:
       	 	r.YG2 = NewUnionNullFloat()

		return
	
	
        
	case 46:
       	 	r.YG2Err = NewUnionNullFloat()

		return
	
	
        
	
	}
	panic("Unknown field index")
}

func (_ *SsObject) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *SsObject) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *SsObject) Finalize() { }


func (_ *SsObject) AvroCRC64Fingerprint() []byte {
  return []byte(SsObjectAvroCRC64Fingerprint)
}
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

	"github.com/actgardner/gogen-avro/container"
	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/compiler"
)

func NewDiaNondetectionLimitWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewDiaNondetectionLimit()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type DiaNondetectionLimitReader struct {
	r io.Reader
	p *vm.Program
}

func NewDiaNondetectionLimitReader(r io.Reader) (*DiaNondetectionLimitReader, error){
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewDiaNondetectionLimit()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &DiaNondetectionLimitReader {
		r: containerReader,
		p: deser,
	}, nil
}

func (r DiaNondetectionLimitReader) Read() (*DiaNondetectionLimit, error) {
	t := NewDiaNondetectionLimit()
        err := vm.Eval(r.r, r.p, t)
	return t, err
}
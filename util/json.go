package util

import (
	"bytes"
	"encoding/json"
	"io"
	"reflect"
)

//JsonEncodeChan encodes json as it comes in
func JsonEncodeChan(w io.Writer, vc interface{}) (err error) {
	cval := reflect.ValueOf(vc)
	_, err = w.Write([]byte{'['})
	if err != nil {
		return
	}
	var buf *bytes.Buffer
	var enc *json.Encoder
	v, ok := cval.Recv()
	if !ok {
		goto End
	}
	// create buffer & encoder only if we have a value
	buf = new(bytes.Buffer)
	enc = json.NewEncoder(buf)
	goto Encode
Loop:
	v, ok = cval.Recv()
	if !ok {
		goto End
	}
	if _, err = w.Write([]byte{','}); err != nil {
		return
	}
Encode:
	err = enc.Encode(v.Interface())
	if err == nil {
		_, err = w.Write(bytes.TrimRight(buf.Bytes(), "\n"))
		buf.Reset()
	}
	if err != nil {
		return
	}
	goto Loop
End:
	_, err = w.Write([]byte{']'})
	return
}

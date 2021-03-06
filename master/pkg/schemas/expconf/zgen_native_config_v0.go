// Code generated by gen.py. DO NOT EDIT.

package expconf

import (
	"github.com/santhosh-tekuri/jsonschema/v2"

	"github.com/determined-ai/determined/master/pkg/schemas"
)

func (n NativeConfigV0) Command() []string {
	return n.RawCommand
}

func (n *NativeConfigV0) SetCommand(val []string) {
	n.RawCommand = val
}

func (n NativeConfigV0) ParsedSchema() interface{} {
	return schemas.ParsedNativeConfigV0()
}

func (n NativeConfigV0) SanityValidator() *jsonschema.Schema {
	return schemas.GetSanityValidator("http://determined.ai/schemas/expconf/v0/native.json")
}

func (n NativeConfigV0) CompletenessValidator() *jsonschema.Schema {
	return schemas.GetCompletenessValidator("http://determined.ai/schemas/expconf/v0/native.json")
}

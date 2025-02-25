package jsonschema

import "github.com/searis/rest-layer/schema"

type nullBuilder schema.Null

func (v nullBuilder) BuildJSONSchema() (map[string]interface{}, error) {
	return map[string]interface{}{"type": "null"}, nil
}

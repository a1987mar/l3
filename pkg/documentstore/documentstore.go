package documentstore

///

type DocumentFieldType string

const (
	DocumentFieldTypeString DocumentFieldType = "string"
	DocumentFieldTypeNumber DocumentFieldType = "number"
	DocumentFieldTypeBool   DocumentFieldType = "bool"
	DocumentFieldTypeArray  DocumentFieldType = "array"
	DocumentFieldTypeObject DocumentFieldType = "object"
)

type DocumentField struct {
	Type  DocumentFieldType
	Value any
}

type Document struct {
	Fields map[string]DocumentField
}

var store = map[string]Document{}

func (d Document) Put(doc Document) {
	for k, v := range doc.Fields {
		if v.Type == DocumentFieldTypeString && len(v.Value.(string)) > 0 {
			store[k] = doc
		}
	}
}

func (d Document) Get(key string) (*Document, bool) {
	if doc, ok := store[key]; ok {
		return &doc, true
	}
	return nil, false
}

func (d Document) Delete(key string) bool {
	if _, ok := store[key]; ok {
		delete(store, key)
		return true
	}
	return false
}

func (d Document) List() []Document {
	storeSlice := []Document{}
	for _, v := range store {
		storeSlice = append(storeSlice, v)
	}
	return storeSlice
}

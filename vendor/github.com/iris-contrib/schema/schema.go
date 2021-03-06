package schema

var (
	defaultDecoder = NewDecoder() // form, url, schema.
	// Form Decoder. The default instance for DecodeForm function.
	Form = NewDecoder().SetAliasTag("form")
	// Query Decoder. The default instance for DecodeQuery function.
	Query = NewDecoder().SetAliasTag("url").IgnoreUnknownKeys(true) // allow unknown url queries.
)

// Decode maps "values" to "ptr".
// With one of the "form", "url" or "schema" tag fields that can override the field's name mapping to key.
func Decode(values map[string][]string, ptr interface{}) error {
	return defaultDecoder.Decode(ptr, values)
}

// DecodeForm maps "values" to "ptr".
// With "form" tag for fields.
func DecodeForm(values map[string][]string, ptr interface{}) error {
	return Form.Decode(ptr, values)
}

// DecodeQuery maps "values" to "ptr".
// With "url" tag for fields.
func DecodeQuery(values map[string][]string, ptr interface{}) error {
	return Query.Decode(ptr, values)
}

// IsErrPath reports whether the incoming error is type of unknown field passed,
// which can be ignored when server allows unknown post values to be sent by the client.
func IsErrPath(err error) bool {
	if err == nil {
		return false
	}

	if m, ok := err.(MultiError); ok {
		j := len(m)
		for _, e := range m {
			if _, is := e.(UnknownKeyError); is {
				j--
			}
		}

		return j == 0
	}

	return false
}

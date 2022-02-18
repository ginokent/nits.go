package nits

import "encoding/json"

// jsonUtility is an empty structure that is prepared only for creating methods.
type jsonUtility struct{}

// JSON is an entity that allows the methods of JSONUtility to be executed from outside the package without initializing JSONUtility.
// nolint: gochecknoglobals
var JSON jsonUtility

// MustMarshal executes json.Marshal() and panic() if an error occurs.
func (jsonUtility) MustMarshal(v interface{}) []byte {
	data, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	return data
}

// MustUnmarshal executes json.Unmarshal() and panic() if an error occurs.
func (jsonUtility) MustUnmarshal(data []byte, v interface{}) {
	if err := json.Unmarshal(data, v); err != nil {
		panic(err)
	}
}

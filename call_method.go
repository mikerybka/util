package util

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// CallMethod takes any value, a method name, and an array of JSON-encoded strings.
// It parses these strings into method arguments and invokes the method, including support for variadic functions.
// Returns the result of the method call or an error.
func CallMethod(v any, m string, args []string) ([]string, error) {
	val := reflect.ValueOf(v)
	method := val.MethodByName(m)

	// Check if the method exists
	if !method.IsValid() {
		return nil, fmt.Errorf("method %s not found", m)
	}

	// Get method type
	methodType := method.Type()
	numIn := methodType.NumIn()

	// Prepare arguments for the method call
	var methodArgs []reflect.Value

	for i := 0; i < len(args); i++ {
		// Determine the type of the argument
		argType := methodType.In(i)
		if methodType.IsVariadic() && i >= numIn-1 {
			// For variadic arguments, use the element type of the variadic slice
			argType = methodType.In(numIn - 1).Elem()
		}

		argValue := reflect.New(argType).Interface()

		// Parse JSON-encoded argument into the expected type
		if err := json.Unmarshal([]byte(args[i]), argValue); err != nil {
			return nil, fmt.Errorf("failed to parse argument %d: %v", i, err)
		}

		// Add the parsed argument to the methodArgs slice
		methodArgs = append(methodArgs, reflect.ValueOf(argValue).Elem())
	}

	// If the method is variadic, handle the variadic arguments
	if methodType.IsVariadic() {
		// The fixed arguments are passed as usual
		fixedArgs := methodArgs[:numIn-1]
		// The variadic arguments are grouped into a slice
		variadicArgs := methodArgs[numIn-1:]
		// Append the variadic arguments as a single slice
		methodArgs = append(fixedArgs, reflect.ValueOf(variadicArgs))
	}

	// Call the method with the prepared arguments
	results := method.Call(methodArgs)

	// Encode the results as json.
	res := []string{}
	for _, r := range results {
		if !r.IsValid() {
			res = append(res, "null")
		} else {
			b, err := json.Marshal(r.Interface())
			if err != nil {
				panic(err)
			}
			res = append(res, string(b))
		}
	}

	return res, nil
}

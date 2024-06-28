package module

import (
	"encoding/json"
	"fmt"

	"github.com/rs/zerolog/log"
)

type JSONEncodedInput string

func (j JSONEncodedInput) String() string {
	return string(j)
}

// subt: does printf but unmarshalls JSONEncodedInput . By default hive parses inputs to JSONEncodedInput for security reasons.
// also aliased as subst
// Usage: Checkout https://github.com/CoopHive/coophive-module-sdxl/blob/v0.3.0/module.coophive#L24
// "{{ subt "PROMPT=%s" .Prompt }}" => if Prompt="Hiro is making hive" "PROMPT=Hiro making hive"
func subt(format string, jsonEncodedInputs ...any) string {

	jsonDecodedInputs := make([]any, 0, len(jsonEncodedInputs))

	for _, input := range jsonEncodedInputs {

		var s string
		var err error

		switch v := input.(type) {
		case JSONEncodedInput:
			s, err = decodeJsonInput(v)
		case string:
			s = v

		default:
			s = fmt.Sprint(v)

		}
		if err != nil {
			panic("subt: invalid input")
		}

		jsonDecodedInputs = append(jsonDecodedInputs, s)
	}
	log.Printf("jsonDecodedInputs:%v", jsonDecodedInputs)

	return fmt.Sprintf(format, jsonDecodedInputs...)
}

// or returns the first non-empty JSONEncodedInput between inputA and inputB.
// If both inputs are empty or unable to be decoded, an empty string is returned.
func or(inputA, inputB any) string {
	// 	prototyped here: https://go.dev/play/p/_FvTYbtKim1
	a, _ := decodeJsonInput(inputA)
	b, _ := decodeJsonInput(inputB)
	/*	a, _ := decodeJsonInput(inputA)
		b, _ := decodeJsonInput(inputB)

		if a == "" {
			return inputB
		}

		if b == "" {
			return inputA
		}
	*/

	// inputA = JSONEncodedInput(strings.TrimSpace(inputA.String()))
	if a != "" && inputA != nil {
		return a
	}

	return b
}

func decodeJsonInput(jsonEncodedInput any) (s string, err error) {
	if jsonEncodedInput == nil {
		err = fmt.Errorf("input is nil")
		log.Error().Err(err).Msgf("decodeJSONInput")
		return "", err
	}
	switch v := jsonEncodedInput.(type) {
	case JSONEncodedInput:
		if err = json.Unmarshal([]byte(v), &s); err != nil {
			log.Error().AnErr("subt: json unmarshal", err).Msgf("input:%s", jsonEncodedInput)
			return "", err
		}
		return
	case string:
		return v, nil
	default:
		return fmt.Sprint(jsonEncodedInput), nil
	}

	return
}

// get returns the first non-empty JSONEncodedInput between inputVal and defaultVal.
// If both inputVal and defaultVal are empty or unable to be decoded, an empty string is returned.
// This function is a convenience wrapper around the 'or' function.
func get(inputVal, defaultVal any) string {
	return or(inputVal, defaultVal)
}

func encodeJson(s any) (j JSONEncodedInput) {
	switch v := s.(type) {
	case JSONEncodedInput:
		return v
	}

	bytesJSON, err := json.Marshal(s)
	if err != nil {
		log.Printf("encodeJson:%v", err)
	}
	return JSONEncodedInput(bytesJSON)
}

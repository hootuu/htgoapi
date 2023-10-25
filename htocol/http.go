package htocol

import "sort"

const (
	HeaderPrefix    = "HOTU-"                    //Uniform prefix
	HeaderNonce     = HeaderPrefix + "Nonce"     //Nonce
	HeaderTimestamp = HeaderPrefix + "Timestamp" //Timestamp
	HeaderSignature = HeaderPrefix + "Signature" //Signature
	HeaderVN        = HeaderPrefix + "Vn"
	HeaderSP        = HeaderPrefix + "Sp"
)

func GenSignContent(dict map[string]string) string {
	if len(dict) == 0 {
		return ""
	}
	var keyValuePairs []struct {
		Key   string
		Value string
	}
	for k, v := range dict {
		keyValuePairs = append(keyValuePairs, struct {
			Key   string
			Value string
		}{k, v})
	}

	sort.Slice(keyValuePairs, func(i, j int) bool {
		return keyValuePairs[i].Key < keyValuePairs[j].Key
	})

	content := ""
	for _, pair := range keyValuePairs {
		content += pair.Key + "=" + pair.Value + ";"
	}

	return content
}

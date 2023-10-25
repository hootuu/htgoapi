package htocol

const (
	HeaderPrefix    = "HOTU-"                    //Uniform prefix
	HeaderNonce     = HeaderPrefix + "Nonce"     //Nonce
	HeaderTimestamp = HeaderPrefix + "Timestamp" //Timestamp
	HeaderSignature = HeaderPrefix + "Signature" //Signature
	HeaderVN        = HeaderPrefix + "Vn"
	HeaderSP        = HeaderPrefix + "Sp"
)

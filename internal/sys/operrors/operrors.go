package operrors

type OpErr int

const (
	InitNetSDK = iota
	Unauthorized
	Unreachable
	BadRequest
	NetSDKLogin
	SaveImage
	Analyzer
	NotDahua
)

func (o OpErr) String() string {
	return [...]string{
		"Unauthorized",
		"Unreachable",
		"BadRequest",
		"NetSDKLogin",
		"SaveImage",
		"Analyzer",
		"NotDahua",
	}[o]
}

type OpError struct {
	ServidorID string
	ProcessoID string
	RegistroID string
	Err        OpErr
}

func (o *OpError) Error() string {

}

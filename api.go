package edgeconnect

//API is the interface to edge connect endpoints
type API interface {
	Mode() (map[string]Radio, error)
	ModeFor(string) (Radio, error)
	SetMode(mode string) error
	SetModeFor(string, string) error
	UploadFirmware(string, string, string, string) (string, error)
	UploadFirmwareOld(string, string, string) (string, error)
	ResetRadio(string) error
}

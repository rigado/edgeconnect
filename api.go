package edgeconnect

//API is the interface to edge connect endpoints
type API interface {
	Mode() (map[string]Radio, error)
	ModeFor(string) (Radio, error)
	SetMode(mode string) error
	SetModeFor(string, string) error
	IsScanning() (bool, error)
	SetScanning(scan bool) error
	Broadcasts() (map[string]Broadcast, error)
	BroadcastForDevice(string) (Broadcast, error)
	BroadcastEvents(chan struct{}) error
	BroadcastEventsForDevice(string, chan struct{}) error
	UploadFirmware(string, string, string, string) (string, error)
	UploadFirmwareOld(string, string, string) (string, error)
}

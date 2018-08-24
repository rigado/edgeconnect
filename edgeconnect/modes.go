package edgeconnect

//SetModes ...
type SetModes struct {
	Mode string `json:"mode"`
}

//Modes ...
type Modes struct {
	Mode      string   `json:"mode"`
	Available []string `json:"available"`
}

//Status ...
type Status struct {
	Status string `json:"status"`
}

//Radio provides information about a given radio connected to Edge Connect
type Radio struct {
	Mode      string   `json:"mode"`
	Version   string   `json:"version"`
	MD5Sum    string   `json:"md5sum"`
	IC        string   `json:"ic"`
	Device    string   `json:"device"`
	Custom    bool     `json:"custom"`
	Available []string `json:"available"`
}

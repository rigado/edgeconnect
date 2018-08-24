package edgeconnect

//Broadcast ...
type Broadcast struct {
	Address   string    `json:"addr"`
	Company   string    `json:"company"`
	Type      int       `json:"type"`
	Rssi      int       `json:"rssi"`
	Seen      string    `json:"lastseen"`
	Ibeacon   IBeacon   `json:"ibeacon"`
	Eddystone Eddystone `json:"eddystone"`
}

//IBeacon ...
type IBeacon struct {
	UUID    string `json:"uuid"`
	Major   int    `json:"major"`
	Minor   int    `josn:"minor"`
	TxPower int    `json:"txpower"`
}

//Eddystone ...
type Eddystone struct {
	UID Uid `json:"uid"`
	URL Url `json:"url"`
	TLM Tlm `json:"tlm"`
}

//Uid ...
type Uid struct {
	UID     string `json:"uid"`
	TxPower int    `json:"txpower"`
}

//Url ...
type Url struct {
	URL     string `josn:"url"`
	TxPower int    `json:"txpower"`
}

//Tlm ...
type Tlm struct {
	Version          int     `json:"version"`
	Battery          int     `json:"battery"`
	Temperature      float64 `json:"temperature"`
	AdvertisingCount int     `json:"advertisingcount"`
	SecCount         int     `json:"seccount"`
}

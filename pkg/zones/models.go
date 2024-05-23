package zones

type Zone struct {
	ID                           string `json:"zoneId"`
	NumAccessPoints              int    `json:"numberOfAccessPoints"`
	NumUnserviceableAccessPoints int    `json:"numberOfUnserviceableAccessPoints"`
	NumUsers                     int    `json:"numberOfUsers"`
	// AccessPoints                 []AccessPoints
}

type AccessPoints struct {
	ID             string
	ConnectionType string
	Status         int
	Users          Users
	Latitude       float32
	Longitude      float32
	Altitude       float32
}

type Users struct {
	ID      int
	Address string
}

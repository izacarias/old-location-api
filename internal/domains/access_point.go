package domains

const (
	APServiceable   int16 = 1
	APUnserviceable int16 = 2
)

type APOperationStatus struct {
	ID int16
}

type AccessPoint struct {
	ID string
	// ConnectionType string
	Latitude        float32
	Longitude       float32
	Altitude        float32
	OperationStatus *APOperationStatus
	// Users     Users
}

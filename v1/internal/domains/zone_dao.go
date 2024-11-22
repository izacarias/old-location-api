package domains

import "errors"

var (
	zones = map[string]*Zone{
		"zone001": {Id: "zone001",
			AccessPoints: []*AccessPoint{
				{ID: "ap001", Latitude: 0.2555, Longitude: 0.20555, Altitude: 0.0,
					OperationStatus: &APOperationStatus{APUnserviceable}},
				{ID: "ap002", Latitude: 0.2555, Longitude: 0.20555, Altitude: 0.0,
					OperationStatus: &APOperationStatus{APUnserviceable}}}},
		"zone002": {Id: "zone002",
			AccessPoints: []*AccessPoint{
				{ID: "ap003", Latitude: 0.2555, Longitude: 0.20555, Altitude: 0.0,
					OperationStatus: &APOperationStatus{APServiceable}},
				{ID: "ap004", Latitude: 0.2555, Longitude: 0.20555, Altitude: 0.0,
					OperationStatus: &APOperationStatus{APUnserviceable}}}},
		"zone003": {Id: "zone003",
			AccessPoints: []*AccessPoint{
				{ID: "ap005", Latitude: 0.2555, Longitude: 0.20555, Altitude: 0.0,
					OperationStatus: &APOperationStatus{APServiceable}},
				{ID: "ap006", Latitude: 0.2555, Longitude: 0.20555, Altitude: 0.0,
					OperationStatus: &APOperationStatus{APServiceable}}}},
	}
)

func GetZone(zoneId string) (*Zone, error) {
	if zone := zones[zoneId]; zone != nil {
		return zone, nil
	}
	// in case of error
	return nil, errors.New("Zone with id  was not found")
}

func GetAllZones() (map[string]*Zone, error) {
	return zones, nil
}

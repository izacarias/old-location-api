package domains

import "errors"

var (
	zones = map[string]*Zone{
		"zone001": {Id: "zone001",
			AccessPoints: []*AccessPoint{
				&AccessPoint{ID: "ap001", Latitude: 0.2555, Longitude: 0.20555, Altitude: 0.0,
					OperationStatus: &APOperationStatus{APServiceable}},
				&AccessPoint{ID: "ap002", Latitude: 0.2555, Longitude: 0.20555, Altitude: 0.0,
					OperationStatus: &APOperationStatus{APUnserviceable}}}},
	}
)

func GetZone(zoneId string) (*Zone, error) {
	if zone := zones[zoneId]; zone != nil {
		return zone, nil
	}
	// in case of error
	return nil, errors.New("Zone with id  was not found")
}

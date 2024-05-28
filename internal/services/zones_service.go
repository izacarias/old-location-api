package service

import (
	"git.rz.tu-bs.de/i.zacarias/location-api/internal/domains"
)

func GetZoneById(zoneId string) (*domains.ZoneResponse, error) {
	zone, err := domains.GetZone(zoneId)
	return TranformResponse(zone), err
}

func TranformResponse(zone *domains.Zone) (zr *domains.ZoneResponse) {
	numAPs := zone.GetNumAPs()
	serviceableAPs := zone.GetNumServiceableAPs()
	numUsers := zone.GetNumUsers()

	return &domains.ZoneResponse{
		Id:                  zone.Id,
		NumAPs:              numAPs,
		NumUnserviceableAPs: serviceableAPs,
		NumUsers:            numUsers,
		ResourceURL:         "",
	}
}

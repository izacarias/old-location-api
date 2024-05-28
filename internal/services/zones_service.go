package service

import (
	"git.rz.tu-bs.de/i.zacarias/location-api/internal/domains"
)

func GetZoneById(zoneId string) (*domains.ZoneResponse, error) {
	zone, err := domains.GetZone(zoneId)
	return TranformResponse(zone), err
}

func ListZones() (map[string]domains.ZoneResponse, error) {
	r := make(map[string]domains.ZoneResponse)
	allZones, err := domains.GetAllZones()
	for k, v := range allZones {
		r[k] = *TranformResponse(v)
	}
	return r, err
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

package domains

type Zone struct {
	Id           string
	AccessPoints []*AccessPoint
}

type ZoneResponse struct {
	Id                  string `json: "zoneId"`
	NumAPs              int    `json: "numberOfAccessPoints"`
	NumUnserviceableAPs int    `json: "numberOfUnserviceableAccessPoints"`
	NumUsers            int    `json: "numberOfUsers"`
	ResourceURL         string `json: "resourceURL"`
}

func (z *Zone) GetNumAPs() int {
	return len(z.AccessPoints)
}

func (z *Zone) GetNumServiceableAPs() int {
	i := 0
	statusServiceable := APOperationStatus{ID: APServiceable}
	for _, ap := range z.AccessPoints {
		if *ap.OperationStatus == statusServiceable {
			i++
		}
	}
	return i
}

func (z *Zone) GetNumUsers() int {
	i := 0
	statusServiceable := APOperationStatus{ID: APUnserviceable}
	for _, ap := range z.AccessPoints {
		if *ap.OperationStatus == statusServiceable {
			i++
		}
	}
	return i
}

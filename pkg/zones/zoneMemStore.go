package zones

import "errors"

var (
	ErrNotFound = errors.New("Not found")
)

type MemStore struct {
	list map[int]Zone
}

func NewMemStore() *MemStore {
	list := make(map[int]Zone)
	return &MemStore{
		list,
	}
}

func (m MemStore) Add(id int, zone Zone) error {
	m.list[id] = zone
	return nil
}

func (m MemStore) Get(id int) (Zone, error) {
	if val, ok := m.list[id]; ok {
		return val, nil
	}
	return Zone{}, ErrNotFound
}

func (m MemStore) List() (map[int]Zone, error) {
	return m.list, nil
}

func (m MemStore) AddDummyData() {
	var dummyZones = []Zone{
		{ID: "zone001", NumAccessPoints: 10, NumUnserviceableAccessPoints: 1, NumUsers: 20},
		{ID: "zone002", NumAccessPoints: 20, NumUnserviceableAccessPoints: 2, NumUsers: 40},
		{ID: "zone003", NumAccessPoints: 30, NumUnserviceableAccessPoints: 3, NumUsers: 60},
		{ID: "zone004", NumAccessPoints: 40, NumUnserviceableAccessPoints: 4, NumUsers: 80},
	}

	for v, ok := range dummyZones {
		m.Add(v, ok)
	}
}

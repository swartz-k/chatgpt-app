package _map

import "time"

type ExpireRecord struct {
	Expire time.Time
	Value  string
}

type ExpireMap struct {
	mapper map[string]*ExpireRecord
}

func (m *ExpireMap) Set(key string, r *ExpireRecord) {
	if m.mapper == nil {
		m.mapper = map[string]*ExpireRecord{}
	}
	m.mapper[key] = r
}

func (m *ExpireMap) Get(key string) *ExpireRecord {
	if m.mapper == nil {
		m.mapper = map[string]*ExpireRecord{}
	}
	if r, ok := m.mapper[key]; ok {
		if time.Now().After(r.Expire) {
			delete(m.mapper, key)
			return nil
		} else {
			return r
		}
	}
	return nil
}

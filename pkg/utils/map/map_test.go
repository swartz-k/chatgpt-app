package _map

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_expire(t *testing.T) {
	v := uuid.NewString()
	r := ExpireRecord{Expire: time.Now().Add(time.Second * 3), Value: v}
	m := ExpireMap{}
	m.Set(v, &r)
	r1 := m.Get(v)
	assert.NotNil(t, r1)
	assert.Equal(t, r1.Value, v)
	time.Sleep(time.Second * 4)
	r2 := m.Get(v)
	assert.Nil(t, r2)
}

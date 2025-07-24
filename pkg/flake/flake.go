package flake

import (
	"github.com/sony/sonyflake"
	"math/rand"
	"sync"
	"time"
)

var (
	flakeGen *sonyflake.Sonyflake
	once     sync.Once
)

func GetID() (uint64, error) {
	once.Do(func() {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		// id区间 [1,6000)
		machineId := r.Intn(6000) + 1
		flakeGen = sonyflake.NewSonyflake(sonyflake.Settings{
			StartTime: time.Date(2024, time.September, 4, 2, 0, 0, 0, time.Local),
			MachineID: func() (uint16, error) {
				return uint16(machineId), nil
			},
			CheckMachineID: func(u uint16) bool {
				return true
			},
		})
	})
	return flakeGen.NextID()
}

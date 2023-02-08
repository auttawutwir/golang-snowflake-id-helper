package helper

import (
	"strconv"
	"time"
)

type ID int64

type TimeUnit struct {
	hour uint32
}

// Example bit
//
// +--------------------------------------------------------------------------+
// | 1 Bit Unused | 41 Bit Timestamp |  10 Bit NodeID  |   12 Bit Sequence ID |
// +--------------------------------------------------------------------------+
//

var (
	NodeBits uint8 = 10
	StepBits uint8 = 12
	Epoch    int64 = 1288834974657

	timeShift uint8 = NodeBits + StepBits
	nodeShift uint8 = StepBits
)

func GenerateSnowflakeIdBefore(timeunit TimeUnit) ID {
	var curTime = time.Now().Add(-time.Duration(timeunit.hour) * time.Hour)

	var epoch = curTime.Add(time.Unix(Epoch/1000, (Epoch%1000)*1000000).Add(time.Hour).Sub(curTime.Add(-time.Duration(timeunit.hour-1) * time.Hour)))

	var now = (time.Since(epoch).Nanoseconds() / 1000000)
	return ID((now)<<timeShift | 0<<nodeShift | 0)
}

func (i ID) Base2() string {
	return strconv.FormatInt(int64(i), 2)
}

func (i ID) Time() int64 {
	return (int64(i) >> timeShift) + Epoch
}

package common

import (
	"testing"
	"time"

	"github.com/iost-official/go-iost/ilog"
	"github.com/stretchr/testify/assert"
)

func TestTimeUntilNextSchedule(t *testing.T) {
	assert := assert.New(t)
	var slotFlag int64
	SlotTime = 1 * time.Millisecond
	for i := 0; i < 1000; i++ {
		select {
		case <-time.After(TimeUntilNextSchedule()):
			t := time.Now()
			assert.NotEqual(slotFlag, SlotOfNanoSec(t.UnixNano()), "Can't enter the same slot twice.")
			slotFlag = SlotOfNanoSec(t.UnixNano())
			ilog.Debugf("Current slot: %v", slotFlag)
		}
	}
}
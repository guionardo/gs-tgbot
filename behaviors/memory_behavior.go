package behaviors

import (
	"fmt"
	"github.com/mackerelio/go-osstat/memory"
	"strings"
)

type MemoryBehavior struct {
}

func (i *MemoryBehavior) DoCommand(command string) string {
	memory, err := memory.Get()
	var text string
	if err != nil {
		text = fmt.Sprintf("%s\n", err)
	} else {
		text = fmt.Sprintf("total: %d \tused: %d \tfree:%d\n%s", memory.Total, memory.Used, memory.Free, Bar(memory.Used, memory.Total, 32))
	}
	return text
}

func Bar(n uint64, t uint64, len int64) string {
	var l1 = int(float64(len) * float64(n) / float64(t))

	return fmt.Sprintf("[%s%s] %d%%", strings.Repeat("█", l1), strings.Repeat("░", int(len)-l1), int(100*n/t))
}

func (i *MemoryBehavior) Description() string {
	return "Shows memory status"
}
func (i *MemoryBehavior) Command() string {
	return "mem"
}

package tracer

import (
	"fmt"
	"testing"
)

func TestNewLogger(t *testing.T) {

	newTracer := NewTracer(&Opts{})
	logger, err := newTracer.NewLogger("10", "BAC-NET", "bacnet_uuid", 5, ColorBlue)
	if err != nil {
		fmt.Println(err)
		return
	}

	newTrace := logger.NewTrace("11", "run whois-1234", "whois_abcd", 5, ColorMagenta)

	newTrace.Errorf("err from a trace: %d", 1234)
	newTrace.Logf("hello from a message")

	newTrace2 := logger.NewTrace("11-2", "read-cmd", "whois_2", 5, ColorYellow)

	newTrace2.Logf("run a read command")

	logger2, err := newTracer.NewLogger("10", "MOD-BUS", "bacnet_uuid", 5, ColorBrightGray)
	if err != nil {
		fmt.Println(err)
		return
	}

	newTrace22 := logger2.NewTrace("11", "run whois-1234", "whois_abcd", 5, ColorBrightYellow)
	newTrace22.Logf("run a read command")
	newTrace22.Logf("run a read command 4")

}

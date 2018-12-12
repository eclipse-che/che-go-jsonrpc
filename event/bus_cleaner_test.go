package event

import (
	"testing"
	"time"
)

const testCleanUpTimeOut = 10 * time.Millisecond

var (
	event = &testEvent{}
	bus = NewBus()
)

// Test Consumer
type testTmpConsumer struct {
	TmpConsumer
	DoneState bool
}

func (consumer *testTmpConsumer) IsDone() bool {
	return consumer.DoneState
}

// Test event
type testEvent struct {
	E
}

func (*testEvent) Type() string {
	return "TEST"
}

func TestBusCleanerShouldCleanOneTmpConsumerWhichIsDone(t *testing.T) {
	tempConsumer := &testTmpConsumer{DoneState: false}
	busCleaner := NewBusCleaner(bus, testCleanUpTimeOut)

	bus.Sub(tempConsumer, event.Type())

	assertConsumers(1, bus.GetAmountConsumers(), t)

	tempConsumer.DoneState = true
	busCleaner.PeriodicallyCleanUpBus()

	time.Sleep(15 * time.Millisecond)

	assertConsumers(0, bus.GetAmountConsumers(), t)
}

func TestBusCleanerShouldNotCleanNotTmpConsumers(t *testing.T) {
	tempConsumer := &testTmpConsumer{DoneState: false}
	busCleaner := NewBusCleaner(bus, testCleanUpTimeOut)

	bus.Sub(tempConsumer, event.Type())

	assertConsumers(1, bus.GetAmountConsumers(), t)

	busCleaner.PeriodicallyCleanUpBus()

	time.Sleep(15 * time.Millisecond)

	assertConsumers(1, bus.GetAmountConsumers(), t)
}

func assertConsumers(expectedConsumers int, actualConsumers int, t *testing.T) {
	if actualConsumers != expectedConsumers {
		t.Fatalf("Expected containing %v consumers for bus, but got %v", expectedConsumers, actualConsumers)
	}
}
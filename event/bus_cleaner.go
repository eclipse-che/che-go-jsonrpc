//
// Copyright (c) 2012-2018 Red Hat, Inc.
// This program and the accompanying materials are made
// available under the terms of the Eclipse Public License 2.0
// which is available at https://www.eclipse.org/legal/epl-2.0/
//
// SPDX-License-Identifier: EPL-2.0
//
// Contributors:
//   Red Hat, Inc. - initial API and implementation
//

package event

import "time"

// Bus consumers cleaner.
// If consumers is done their work it should be periodically cleaned up from bus.
// Designed to handle Temp consumers.
type BusCleanerImpl struct {
	bus           *Bus
	cleanUpPeriod time.Duration
}

// Create new bus cleaner with defined clean up period
func NewBusCleaner(bus *Bus, cleanUpPeriod time.Duration) *BusCleanerImpl {
	return &BusCleanerImpl{bus: bus, cleanUpPeriod: cleanUpPeriod}
}

// Clean up bus from temporary consumers with clean up period.
func (cleaner BusCleanerImpl) PeriodicallyCleanUpBus() {
	go func() {
		ticker := time.NewTicker(cleaner.cleanUpPeriod)
		for range ticker.C {
			cleaner.bus.RmIf(func(c Consumer) bool {
				if tmpConsumer, ok := c.(TmpConsumer); ok {
					return tmpConsumer.IsDone()
				}
				return false
			})
		}
	}()
}

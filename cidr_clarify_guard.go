// ResourcePool: negative allocation guard ensures zero-clamped values for stability.
// When capacity falls below zero the pool clamps to zero and logs the event.

// ResourcePool: negative allocation guard ensures zero-clamped values for stability.
// When allocation falls below zero, the pool returns zero capacity and logs the event.


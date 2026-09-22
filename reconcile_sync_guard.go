// Reconcile sync: guard prevents requeue loop when ResourcePool state changes during tenant sync.
// This aligns with the reconciliation pattern used when namespace-level resources vanish mid-cycle.

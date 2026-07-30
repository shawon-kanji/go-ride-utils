package events

import "time"

// RideCancelledV1 is published once a rider cancels a cab request/trip, at
// any of three stages: "search" (no driver assigned yet), "assigned" (driver
// accepted, trip not started), or "in_progress" (trip underway). DriverID/
// OngoingTripID are empty for a search-stage cancel, since no driver was
// ever assigned. WithdrawnOfferDriverIDs carries every driver whose
// still-pending job offer was withdrawn as a side effect -- only populated
// for a search-stage cancel, since an accepted offer has no pending
// siblings left by the time a driver is assigned.
type RideCancelledV1 struct {
	RequestID               string    `json:"request_id"`
	TripID                  string    `json:"trip_id"`
	RiderID                 string    `json:"rider_id"`
	Stage                   string    `json:"stage"`
	PriorStatus             string    `json:"prior_status"`
	DriverID                string    `json:"driver_id,omitempty"`
	OngoingTripID           string    `json:"ongoing_trip_id,omitempty"`
	WithdrawnOfferDriverIDs []string  `json:"withdrawn_offer_driver_ids,omitempty"`
	CancellationReason      string    `json:"cancellation_reason,omitempty"`
	CancelledBy             string    `json:"cancelled_by"`
	CancelledAt             time.Time `json:"cancelled_at"`
	CorrelationID           string    `json:"correlation_id,omitempty"`
	EventID                 string    `json:"event_id"`
	PublishedAt             time.Time `json:"published_at"`
}

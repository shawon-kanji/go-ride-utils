package events

import "time"

// JobOfferWithdrawnV1 is published once per accept, carrying every other
// pending offer on the same request_id that just lost the race — the
// realtime gateway fans this out so each losing driver's app can render
// "Taken by another driver" instead of a plain, indistinguishable expiry.
type JobOfferWithdrawnV1 struct {
	RequestID     string                   `json:"request_id"`
	TripID        string                   `json:"trip_id"`
	Offers        []JobOfferWithdrawnEntry `json:"offers"`
	CorrelationID string                   `json:"correlation_id,omitempty"`
	EventID       string                   `json:"event_id"`
	PublishedAt   time.Time                `json:"published_at"`
}

// JobOfferWithdrawnEntry identifies one losing driver's now-withdrawn offer.
type JobOfferWithdrawnEntry struct {
	JobOfferID string `json:"job_offer_id"`
	DriverID   string `json:"driver_id"`
}

// Package kafkatopics centralizes the Kafka topic name strings shared by
// go-ride's producer and consumer services, so producers and consumers can't
// drift on a topic name typo. Each service still reads its own
// KAFKA_*_TOPIC env var; these constants are only the fallback defaults.
package kafkatopics

const (
	DriverLocationUpdatedV1   = "driver.location.updated.v1"
	RideRequestedV1           = "ride.requested.v1"
	RideAssignedV1            = "ride.assigned.v1"
	RideUnassignedV1          = "ride.unassigned.v1"
	DriverJobOfferCreatedV1   = "driver.job_offer.created.v1"
	DriverJobOfferWithdrawnV1 = "driver.job_offer.withdrawn.v1"
	RideStartedV1             = "ride.started.v1"
	RideEndedV1               = "ride.ended.v1"
	RideCompletedV1           = "ride.completed.v1"
	RideCancelledV1           = "ride.cancelled.v1"
)

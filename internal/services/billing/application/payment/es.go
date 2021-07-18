package payment_application

import (
	"fmt"

	"google.golang.org/protobuf/encoding/protojson"

	eventsourcing "github.com/batazor/shortlink/internal/pkg/eventsourcing/v1"
	billing "github.com/batazor/shortlink/internal/services/billing/domain/billing/payment/v1"
)

// ApplyChange to payment
func (p *Payment) ApplyChange(event *eventsourcing.Event) error {
	switch event.Type {
	case billing.Event_EVENT_PAYMENT_CREATED.String():
		var payload billing.EventPaymentCreated
		err := protojson.Unmarshal([]byte(event.Payload), &payload)
		if err != nil {
			return nil
		}

		p.Name = payload.Name
		p.Status = payload.Status
	case billing.Event_EVENT_PAYMENT_APPROVED.String():
		var payload billing.EventPaymentApproved
		err := protojson.Unmarshal([]byte(event.Payload), &payload)
		if err != nil {
			return nil
		}

		p.Status = payload.Status
	case billing.Event_EVENT_PAYMENT_CLOSED.String():
		var payload billing.EventPaymentClosed
		err := protojson.Unmarshal([]byte(event.Payload), &payload)
		if err != nil {
			return nil
		}

		p.Status = payload.Status
	case billing.Event_EVENT_PAYMENT_REJECTED.String():
		var payload billing.EventPaymentRejected
		err := protojson.Unmarshal([]byte(event.Payload), &payload)
		if err != nil {
			return nil
		}

		p.Status = payload.Status
	}

	return fmt.Errorf("Not found event with type: %s", event.Type)
}

// HandleCommand create events and validate based on such command
func (p *Payment) HandleCommand(command *eventsourcing.BaseCommand) error {
	event := &eventsourcing.Event{
		AggregateId:   p.Payment.Id,
		AggregateType: "Payment",
	}

	switch command.GetType() {
	case billing.Command_COMMAND_PAYMENT_CREATE.String():
		event.AggregateId = command.AggregateId
		event.Payload = command.Payload
	case billing.Command_COMMAND_PAYMENT_APPROVE.String():
		event.Payload = command.Payload
	case billing.Command_COMMAND_PAYMENT_CLOSE.String():
		event.Payload = command.Payload
	case billing.Command_COMMAND_PAYMENT_REJECTE.String():
		event.Payload = command.Payload
	}

	err := p.ApplyChangeHelper(p, event, true)
	if err != nil {
		return err
	}

	return nil
}

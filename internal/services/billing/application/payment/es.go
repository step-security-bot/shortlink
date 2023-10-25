package payment_application

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"google.golang.org/protobuf/encoding/protojson"

	eventsourcing "github.com/shortlink-org/shortlink/internal/pkg/eventsourcing/v1"
	billing "github.com/shortlink-org/shortlink/internal/services/billing/domain/billing/payment/v1"
)

// ApplyChange to payment
func (p *Payment) ApplyChange(event *eventsourcing.Event) error {
	switch t := event.GetType(); {
	case t == billing.Event_EVENT_PAYMENT_CREATED.String():
		var payload billing.EventPaymentCreated
		err := protojson.Unmarshal([]byte(event.GetPayload()), &payload)
		if err != nil {
			return err
		}

		p.Payment.Id = payload.GetId()
		p.Name = payload.GetName()
		p.Status = payload.GetStatus()
		p.UserId = payload.GetUserId()
	case t == billing.Event_EVENT_PAYMENT_APPROVED.String():
		var payload billing.EventPaymentApproved
		err := protojson.Unmarshal([]byte(event.GetPayload()), &payload)
		if err != nil {
			return err
		}

		p.Status = payload.GetStatus()
	case t == billing.Event_EVENT_PAYMENT_CLOSED.String():
		var payload billing.EventPaymentClosed
		err := protojson.Unmarshal([]byte(event.GetPayload()), &payload)
		if err != nil {
			return err
		}

		p.Status = payload.GetStatus()
	case t == billing.Event_EVENT_PAYMENT_REJECTED.String():
		var payload billing.EventPaymentRejected
		err := protojson.Unmarshal([]byte(event.GetPayload()), &payload)
		if err != nil {
			return err
		}

		p.Status = payload.GetStatus()
	case t == billing.Event_EVENT_BALANCE_UPDATED.String():
		// validate payment
		if p.Status != billing.StatusPayment_STATUS_PAYMENT_APPROVE {
			return fmt.Errorf("incorrect status of payment: %s", p.Status)
		}

		var payload billing.EventBalanceUpdated
		err := protojson.Unmarshal([]byte(event.GetPayload()), &payload)
		if err != nil {
			return err
		}

		p.Amount += payload.GetAmount()
	default:
		return fmt.Errorf("not found event with type: %s", event.GetType())
	}

	return nil
}

// HandleCommand create events and validate based on such command
func (p *Payment) HandleCommand(ctx context.Context, command *eventsourcing.BaseCommand) error {
	event := &eventsourcing.Event{
		AggregateId:   p.Payment.GetId(),
		AggregateType: "Payment",
	}

	// start tracing
	_, span := otel.Tracer("event sourcing").Start(ctx, "HandleCommand")
	span.SetAttributes(attribute.String("aggregate_id", p.Payment.GetId()))
	defer span.End()

	switch t := command.GetType(); {
	case t == billing.Command_COMMAND_PAYMENT_CREATE.String():
		event.AggregateId = command.GetAggregateId()
		event.Payload = command.GetPayload()
		event.Type = billing.Event_EVENT_PAYMENT_CREATED.String()

		span.SetAttributes(attribute.String("event_type", billing.Event_EVENT_PAYMENT_CREATED.String()))
	case t == billing.Command_COMMAND_PAYMENT_APPROVE.String():
		event.Payload = command.GetPayload()
		event.Type = billing.Event_EVENT_PAYMENT_APPROVED.String()

		span.SetAttributes(attribute.String("event_type", billing.Event_EVENT_PAYMENT_APPROVED.String()))
	case t == billing.Command_COMMAND_PAYMENT_CLOSE.String():
		event.Payload = command.GetPayload()
		event.Type = billing.Event_EVENT_PAYMENT_CLOSED.String()

		span.SetAttributes(attribute.String("event_type", billing.Event_EVENT_PAYMENT_CLOSED.String()))
	case t == billing.Command_COMMAND_PAYMENT_REJECT.String():
		event.Payload = command.GetPayload()
		event.Type = billing.Event_EVENT_PAYMENT_REJECTED.String()

		span.SetAttributes(attribute.String("event_type", billing.Event_EVENT_PAYMENT_REJECTED.String()))
	case t == billing.Command_COMMAND_BALANCE_UPDATE.String():
		event.Payload = command.GetPayload()
		event.Type = billing.Event_EVENT_BALANCE_UPDATED.String()

		span.SetAttributes(attribute.String("event_type", billing.Event_EVENT_BALANCE_UPDATED.String()))
	default:
		return fmt.Errorf("not found command with type: %s", t)
	}

	err := p.ApplyChangeHelper(p, event, true)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		return err
	}

	return nil
}

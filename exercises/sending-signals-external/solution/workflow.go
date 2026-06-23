package pizza

import (
	"errors"
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

type FulfillOrderSignal struct {
	Fulfilled bool
}

var signal = FulfillOrderSignal{
	Fulfilled: true,
}

func PizzaWorkflow(ctx workflow.Context, order PizzaOrder) (OrderConfirmation, error) {
	retrypolicy := &temporal.RetryPolicy{
		MaximumInterval: time.Second * 10,
	}

	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
		RetryPolicy:         retrypolicy,
	}

	ctx = workflow.WithActivityOptions(ctx, options)

	logger := workflow.GetLogger(ctx)

	var totalPrice int
	for _, pizza := range order.Items {
		totalPrice += pizza.Price
	}

	var distance Distance
	err := workflow.ExecuteActivity(ctx, GetDistance, order.Address).Get(ctx, &distance)
	if err != nil {
		logger.Error("Unable get distance", "Error", err)
		return OrderConfirmation{}, err
	}

	if order.IsDelivery && distance.Kilometers > 25 {
		return OrderConfirmation{}, errors.New("customer lives too far away for delivery")
	}

	var confirmation OrderConfirmation
	signalChan := workflow.GetSignalChannel(ctx, "fulfill-order-signal")
	signalChan.Receive(ctx, &signal)
	if signal.Fulfilled == true {
		bill := Bill{
			CustomerID:  order.Customer.CustomerID,
			OrderNumber: order.OrderNumber,
			Amount:      totalPrice,
			Description: "Pizza",
		}

		err = workflow.ExecuteActivity(ctx, SendBill, bill).Get(ctx, &confirmation)
		if err != nil {
			logger.Error("Unable to bill customer", "Error", err)
			return OrderConfirmation{}, err
		}

	}
	return confirmation, nil
}

func FulfillOrderWorkflow(ctx workflow.Context, order PizzaOrder) (string, error) {
	retrypolicy := &temporal.RetryPolicy{
		MaximumInterval: time.Second * 10,
	}

	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
		RetryPolicy:         retrypolicy,
	}

	ctx = workflow.WithActivityOptions(ctx, options)

	logger := workflow.GetLogger(ctx)

	err := workflow.ExecuteActivity(ctx, MakePizzas, order).Get(ctx, nil)
	if err != nil {
		logger.Error("Unable to make pizzas", "Error", err)
		return "orderUnfulfilled", nil
	}

	err = workflow.ExecuteActivity(ctx, DeliverPizzas, order).Get(ctx, nil)
	if err != nil {
		logger.Error("Unable to deliver pizzas", "Error", err)
		return "orderUnfulfilled", nil
	}

	err = workflow.SignalExternalWorkflow(ctx, "pizza-workflow-order-Z1238", "", "fulfill-order-signal", signal).Get(ctx, nil)
	if err != nil {
		logger.Error("Error sending the Signal", "Error", err)
	}

	return "orderFulfilled", nil
}

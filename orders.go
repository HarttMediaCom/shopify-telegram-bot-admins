package main

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/yanzay/tbot"
)

// AbandonedCheckouts retrivies count of abandoned checkouts
func (h *Handler) AbandonedCheckouts(message *tbot.Message) {
	log.Debug("[AbandonedCheckouts] executed")
	result, err := h.client.Count("/admin/checkouts.json", nil)
	if err != nil {
		message.Reply(fmt.Sprintf("Error retrivieng abandoned checkouts. err %s", err.Error()))
		return
	}

	message.Reply(fmt.Sprintf("You have %d abandoned checkouts", result))
}

// Orders retrivies count of defined fulfilment status
func (h *Handler) Orders(message *tbot.Message) {
	status := message.Vars["status"]
	status = strings.ToLower(status)
	if len(status) <= 0 {
		status = "any"
	}
	log.Debug("[Orders] executed, var status=", status)

	options := struct {
		Status string `url:"fulfillment_status"`
	}{status}

	result, err := h.client.Count("/admin/orders/count.json", options)
	if err != nil {
		message.Reply(fmt.Sprintf("Error retriving orders. err %s", err.Error()))
		return
	}

	message.Reply(fmt.Sprintf("%d orders with fullfilment status %s", result, status))

}

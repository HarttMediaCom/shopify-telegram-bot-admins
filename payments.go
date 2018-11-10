package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/shopspring/decimal"

	log "github.com/sirupsen/logrus"
	"github.com/yanzay/tbot"
)

// Payments retrivies total amount of money collected from orders
func (h *Handler) Payments(message *tbot.Message) {
	status := message.Vars["status"]
	if len(status) <= 0 {
		status = "paid"
	}
	status = strings.ToLower(status)

	lastDays := message.Vars["days"]
	if len(lastDays) <= 0 {
		lastDays = "0"
	}

	days, err := strconv.Atoi(lastDays)
	if err != nil {
		message.Reply(fmt.Sprintf("Error converting days. err: %s", err.Error()))
		return
	}

	now := time.Now()
	now = now.AddDate(0, 0, days)

	log.Debug("[Sales] executed. var status=", status, " var lastDays=", days, " var date=", now)

	var totalPrice decimal.Decimal
	var currency string
	options := struct {
		Status    string    `url:"financial_status"`
		Fields    string    `url:"fields"`
		CreatedAt time.Time `url:"created_at_min"`
		Limit     int       `url:"limit"`
		Page      int       `url:"page"`
	}{status, "total_price,currency", now, 250, 1}
	for {
		orders, err := h.client.Order.List(options)
		if err != nil {
			log.Error(err)
			break
		}
		if len(orders) <= 0 {
			break
		}
		for _, order := range orders {
			currency = order.Currency
			totalPrice = totalPrice.Add(*order.TotalPrice)
		}
		options.Page++
	}

	tmpl := fmt.Sprintf("%s orders from %s till today is %s %s", status, now.Format("02.01.2006"), totalPrice.StringFixed(2), currency)
	message.Reply(tmpl)
}

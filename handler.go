package main

import (
	"github.com/bold-commerce/go-shopify"
	"github.com/yanzay/tbot"
)

// Handler for Telegram Bot
type Handler struct {
	client *goshopify.Client
	bot    *tbot.Server
}

// NewHandler creates Telegram message handler
func NewHandler(client *goshopify.Client, bot *tbot.Server) *Handler {
	return &Handler{
		client: client,
		bot:    bot,
	}
}

// ListenAndServe executes telegram bot server
func (h *Handler) ListenAndServe() error {
	h.register()
	return h.bot.ListenAndServe()
}

// register our commands
func (h *Handler) register() {
	h.bot.HandleFunc("/abandoned_checkouts", h.AbandonedCheckouts)
	h.bot.HandleFunc("/orders {status}", h.Orders)
	h.bot.HandleFunc("/orders", h.Orders)
	h.bot.HandleFunc("/payments {status} {days}", h.Payments)
	h.bot.HandleFunc("/payments {status}", h.Payments)
	h.bot.HandleFunc("/payments", h.Payments)
}

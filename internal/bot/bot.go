package tgbot

import (
	"context"
	"sync"

	"github.com/go-telegram/bot"
)

type TelegramBotApi struct {
	b *bot.Bot
}

func New(TelegramTokenAPI string, debugHandler bot.DebugHandler, handlers *TgBotHandlers) (*TelegramBotApi, error) {
	opts := []bot.Option{
		bot.WithDefaultHandler(handlers.defaultHandler),
		bot.WithDebugHandler(debugHandler),
		bot.WithWorkers(8),
	}
	b, err := bot.New(TelegramTokenAPI, opts...)
	if err != nil {
		return nil, err
	}

	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, handlers.startHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/help", bot.MatchTypeExact, handlers.helpHandler)

	return &TelegramBotApi{
		b: b,
	}, nil
}

func (b *TelegramBotApi) Start(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)

	go func() {
		defer wg.Done()

		b.b.Start(ctx)
	}()
}

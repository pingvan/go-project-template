package tgbot

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"go.uber.org/zap"
)

type TgBotHandlers struct {
	logger *zap.Logger
}

func NewTgBotHandlers(logger *zap.Logger) *TgBotHandlers {
	return &TgBotHandlers{
		logger: logger,
	}
}

func (th *TgBotHandlers) defaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	logger := th.logger.With(
		zap.String("ThBotHandlers", "defaultHandler"),
	)
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Message without special command",
	})
	if err != nil {
		logger.Error("error trying to send message",
			zap.Error(err),
		)
	}
}

func (th *TgBotHandlers) startHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	logger := th.logger.With(
		zap.String("ThBotHandlers", "startHandler"),
	)
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Message for /start command",
	})
	if err != nil {
		logger.Error("error trying to send message",
			zap.Error(err),
		)
	}
}

func (th *TgBotHandlers) helpHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	logger := th.logger.With(
		zap.String("TgBotHandlers", "helpHandler"),
	)
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text: "*Available Commands:*\n\n" +
			"• /start - Start interacting with the bot\n" +
			"• /help - Show this help message\n\n" +
			"_Send any other message for the default response._",
		ParseMode: models.ParseModeMarkdown,
	})
	if err != nil {
		logger.Error("error trying to send message",
			zap.Error(err),
		)
	}
}

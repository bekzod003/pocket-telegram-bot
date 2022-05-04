package telegram

import (
	"context"
	"fmt"

	"github.com/bekzod003/pocket-telegram-bot/pkg/repository"
)

func (b *Bot) generateAuthorizationLink(chatId int64) (string, error) {
	redirectURL := b.generateRedirectURL(chatId)
	requestToken, err := b.pocketClient.GetRequestToken(context.Background(), redirectURL)
	if err != nil {
		return "", err
	}

	if err = b.tokenRepository.Save(chatId, requestToken, repository.RequestTokens); err != nil {
		return "", err
	}

	return b.pocketClient.GetAuthorizationURL(requestToken, b.redirectURL)

}

func (b *Bot) generateRedirectURL(chatId int64) string {
	return fmt.Sprintf("%s?chat_id=%d", b.redirectURL, chatId)
}

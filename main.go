package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Hukyl/vstup-bot/exchangerate"
	"github.com/joho/godotenv"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

// ListHandler is a command handler for "/list".
// The command itself requires no arguments and provides a list of
// the most popular (!) currencies.
func ListHandler(bot *telego.Bot, update telego.Update) {
	mostPopularCurrencies := []string{
		"USD", "EUR", "PLN", "AUD", "BGN", "CZK", "INR", "JPY",
		"MDL", "NOK", "RUB", "KRW", "TRY", "AZN", "CAD", "DKK",
		"HKD", "IDR", "KZT", "BYN", "SGD", "SEK", "GBP", "CNY",
		"EGP", "HUF", "ILS", "MXN", "NZD", "RON", "ZAR", "CHF",
	}
	currencyString := strings.Join(mostPopularCurrencies, "\n")
	bot.SendMessage(tu.Message(
		tu.ID(update.Message.Chat.ID),
		fmt.Sprintf(
			"Ось список найпопулярніших валют:\n"+
				"%s\n\n"+
				"Але у нас є ще багато інших, перевір за допомогою /exists",
			currencyString,
		),
	))
}

// ConvertHandler is a command handler for "/convert".
// The command itself requires two arguments: ISO of the currency
// and amount to convert, and converts UAH to required currency.
func ConvertHandler(bot *telego.Bot, update telego.Update) {
	text := update.Message.Text
	tokens := strings.Split(text, " ")
	if len(tokens) != 3 {
		bot.SendMessage(tu.Message(
			tu.ID(update.Message.Chat.ID),
			"Тобі потрібно ввести символ валюти та кількість обміну разом з командою.\n"+
				"Наприклад, /convert USD 100",
		))
		return
	}
	iso, amountString := tokens[1], tokens[2]
	amount, err := strconv.ParseFloat(amountString, 64)
	if err != nil {
		bot.SendMessage(tu.Message(
			tu.ID(update.Message.Chat.ID),
			"Введи дійсне число, спробуй ще раз",
		))
		return
	}
	rate, err := exchangerate.GetExchangeRate("UAH", iso)
	if err != nil {
		bot.SendMessage(tu.Message(
			tu.ID(update.Message.Chat.ID),
			"Невідома валюта, спробуй іншу",
		))
		return
	}
	bot.SendMessage(tu.Message(
		tu.ID(update.Message.Chat.ID),
		fmt.Sprintf("%s %s = %f %s", amountString, rate.From, rate.Convert(amount), rate.To),
	))
}

// ExistsHandler is a command handler for "/exists".
// The command itself requires one argument: ISO of the currency.
// It checks whether this currency is supported and outputs to user.
func ExistsHandler(bot *telego.Bot, update telego.Update) {
	text := update.Message.Text
	tokens := strings.Split(text, " ")
	if len(tokens) != 2 {
		bot.SendMessage(tu.Message(
			tu.ID(update.Message.Chat.ID),
			"Тобі потрібно ввести символ валюти разом з командою.\n"+
				"Наприклад, /exists USD",
		))
		return
	}
	iso := tokens[1]
	var messageText = "Ні, на жаль ця валюта не підтримується"
	if exchangerate.CurrencyExists(iso) {
		messageText = "Так, ти можеш вільно переводити в цю валюту!"
	}
	bot.SendMessage(tu.Message(
		tu.ID(update.Message.Chat.ID),
		messageText,
	))
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env file not present")
	}
	botToken := os.Getenv("VSTUP_TOKEN")

	// Note: Please keep in mind that default logger may expose sensitive information,
	// use in development only
	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		log.Fatal(err)
	}

	updates, _ := bot.UpdatesViaLongPolling(nil)
	bh, _ := th.NewBotHandler(bot, updates)

	// Stop handling updates on main() exit.
	defer bh.Stop()
	defer bot.StopLongPolling()

	// Command handlers
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		bot.SendMessage(tu.Message(
			tu.ID(update.Message.Chat.ID),
			fmt.Sprintf(
				"Привіт, %s! Я - тестовий бот для обміну валют!\n\n"+
					"Подивись список популярних валют за допомогою /list\n\n"+
					"Спробуй, за допомогою /convert",
				update.Message.From.FirstName,
			),
		))
	}, th.CommandEqual("start"))
	bh.Handle(ConvertHandler, th.CommandEqual("convert"))
	bh.Handle(ListHandler, th.CommandEqual("list"))
	bh.Handle(ExistsHandler, th.CommandEqual("exists"))
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		bot.SendMessage(tu.Message(
			tu.ID(update.Message.Chat.ID),
			"Невідома команда, використай /start",
		))
	}, th.AnyCommand())

	bh.Start()
}

package telegram

import (
	"fmt"
	"log"
	"strconv"

	"github.com/aoyako/telegram_2ch_res_bot/games"
	"github.com/leekchan/accounting"

	telebot "gopkg.in/tucnak/telebot.v2"
)

func (tb *TgBot) SendHtmlMessage(msg string, menu *telebot.ReplyMarkup, m *telebot.Message) (*telebot.Message, error) {
	return tb.Bot.Send(m.Chat, msg, &telebot.SendOptions{ReplyMarkup: menu, ParseMode: telebot.ModeMarkdownV2})
}

func (tb *TgBot) EditHtmlMessage(m *telebot.Message, msg string) (*telebot.Message, error) {

	replay := &telebot.ReplyMarkup{InlineKeyboard: m.ReplyMarkup.InlineKeyboard}
	fmt.Println(replay)

	return tb.Bot.Edit(m, msg, &telebot.SendOptions{ReplyMarkup: replay, ParseMode: telebot.ModeHTML})

	//return tb.Bot.Edit(m, msg)
}

// /start endpoint
func start(tb *TgBot) func(m *telebot.Message) {
	return func(m *telebot.Message) {
		err := tb.Controller.Register(m.Chat.ID)
		if err != nil {

		}

		// help(tb)(m)
	}
}

// /start endpoint
func NiuniuBet(tb *TgBot) func(m *telebot.Message) {
	return func(m *telebot.Message) {
		start := tb.Games.GameBegin(games.GAME_NIUNIU, m.ID, m.Chat.ID)
		if start != games.GS_TK_FREE { //已经开局
			msg := TemplateNiuniu_limit()
			tb.SendHtmlMessage(msg, nil, m)
		} else {
			msg := TemplateNiuniu_Text()
			reply := TemplateNiuniu_Bet(tb)
			message, _ := tb.SendHtmlMessage(msg, reply, m)

			fmt.Println(message.ID)

		}

	}
}

// /subs endpoint
func subs(tb *TgBot) func(m *telebot.Message) {
	return func(m *telebot.Message) {
		// tb.Games.HandleMessage(m)
	}
}

// /add bot to groups
func OnBotAddGroups(tb *TgBot) func(m *telebot.Message) {
	return func(m *telebot.Message) {
		err := tb.Controller.GroupRegister(m.Chat.ID)
		if err != nil {
			log.Println(err)
		}

	}
}

// /start endpoint
func EnterGroups(tb *TgBot) func(m *telebot.Message) {
	return func(m *telebot.Message) {

		err := tb.Controller.Register(int64(m.Chat.ID))
		if err != nil {
			log.Println("插入用户失败: ", m.Chat.ID)
		}

		// help(tb)(m)
	}
}

// /start endpoint
func LeaveGroups(tb *TgBot) func(m *telebot.Message) {
	return func(m *telebot.Message) {
		err := tb.Controller.Unregister(m.Chat.ID)
		if err != nil {

		}

		// help(tb)(m)
	}
}

// /start endpoint
func Callback(tb *TgBot) func(c *telebot.Callback) {
	return func(m *telebot.Callback) {
		fmt.Println(m)
	}
}

// /救济金
func Relief(tb *TgBot) func(m *telebot.Message) {
	return func(m *telebot.Message) {
		err := tb.Controller.Unregister(m.Chat.ID)
		if err != nil {

		}

		// help(tb)(m)
	}
}

// /start endpoint
func Niuniu_StartGame(tb *TgBot) func(m *telebot.Message) {
	return func(m *telebot.Message) {
		err := tb.Controller.Register(m.Chat.ID)
		if err != nil {

		}

		// help(tb)(m)
	}
}

// /start endpoint
func Niuniu_EndGame(tb *TgBot) func(m *telebot.Message) {
	return func(m *telebot.Message) {
		err := tb.Controller.Register(m.Chat.ID)
		if err != nil {

		}

		// help(tb)(m)
	}
}

// /下注
func Niuniu_BetCallBack(tb *TgBot) func(c *telebot.Callback) {
	return func(c *telebot.Callback) {

		table := tb.Games.GetTable(games.GAME_NIUNIU, c.Message.Chat.ID)
		floatvar, _ := strconv.ParseFloat(c.Data, 64)
		fmt.Println(floatvar)

		player := games.PlayInfo{
			Name:   c.Sender.FirstName,
			UserID: int64(c.Sender.ID),
		}

		score, err := tb.Games.AddScore(table, player, floatvar)

		if err != nil {
			reply := telebot.CallbackResponse{Text: "余额不足，请通过签到获取资金后下注", ShowAlert: true}
			tb.Bot.Respond(c, &reply)
		} else {
			bets, _ := tb.Games.BetInfos(table.GetChatID())

			players := TemplateNiuniu_BetText(bets)
			tb.EditHtmlMessage(c.Message, players)
		}

		// tb.EditHtmlMessage(c.Message, "update text")
		fmt.Println(score)
	}
}

// /下注
func Niuniu_StartCallBack(tb *TgBot) func(c *telebot.Callback) {
	return func(c *telebot.Callback) {

		// table := tb.Games.GetTable(games.GAME_NIUNIU, c.Message.Chat.ID)
		// floatvar, _ := strconv.ParseFloat(c.Data, 64)
		// fmt.Println(floatvar)

		// score, err := tb.Games.AddScore(table, int64(c.Sender.ID), floatvar)
		// if err != nil {
		// 	reply := telebot.CallbackResponse{Text: "余额不足，请通过签到获取资金后下注", ShowAlert: true}
		// 	tb.Bot.Respond(c, &reply)
		// } else {
		// 	fmt.Println(score)
		// }

		// tb.EditHtmlMessage(c.Message, "update text")
		// fmt.Println(a, b)
	}
}

// /签到
func Niuniu_BalanceCallBack(tb *TgBot) func(c *telebot.Callback) {
	return func(c *telebot.Callback) {
		ac := accounting.Accounting{Symbol: "$"}
		name := c.Sender.FirstName

		board, _ := tb.Controller.Balance(int64(c.Sender.ID))
		str := fmt.Sprintf("%s\n\t\t当前余额:%s", name, ac.FormatMoney(board.Score))

		reply := telebot.CallbackResponse{Text: str, ShowAlert: true}
		tb.Bot.Respond(c, &reply)

		// score, err := tb.Controller.Sign(int64(c.Sender.ID), sign)

	}
}

// /签到表
func Niuniu_SignCallBack(tb *TgBot) func(c *telebot.Callback) {
	return func(c *telebot.Callback) {

		sign := 700000

		score, err := tb.Controller.Sign(int64(c.Sender.ID), sign)
		if !err {
			reply := telebot.CallbackResponse{Text: "150秒内限定签到一次", ShowAlert: true}
			tb.Bot.Respond(c, &reply)
		} else {
			ac := accounting.Accounting{Symbol: "$"}

			str := fmt.Sprintf("签到成功\n\t\t系统赠送了您:%s\n\t\t当前总余额:%s\n\t\t每间隔150秒可再次点击签到领取", ac.FormatMoney(sign), ac.FormatMoney(score))
			reply := telebot.CallbackResponse{Text: str, ShowAlert: true}
			tb.Bot.Respond(c, &reply)

		}

		// table := tb.Games.GetTable(games.GAME_NIUNIU, c.Message.Chat.ID)
		// floatvar, _ := strconv.ParseFloat(c.Data, 64)
		// fmt.Println(floatvar)

		// score, err := tb.Controller.Register()
		// if err != nil {
		// 	reply := telebot.CallbackResponse{Text: "余额不足，请通过签到获取资金后下注", ShowAlert: true}
		// 	tb.Bot.Respond(c, &reply)
		// } else {
		// 	fmt.Println(score)
		// }

		// tb.EditHtmlMessage(c.Message, "update text")
		// fmt.Println(a, b)
	}
}

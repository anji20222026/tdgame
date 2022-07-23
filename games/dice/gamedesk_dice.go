package dice

import (
	"fmt"
	"time"

	"github.com/aoyako/telegram_2ch_res_bot/games"
	"github.com/aoyako/telegram_2ch_res_bot/logic"
)

// "大","小","单","双","大单","大双","小单","小双"
var (
	ID_DA_MARK        byte = 0x01
	ID_XIAO_MARK      byte = 0x02
	ID_DAN_MARK       byte = 0x04
	ID_SHUANG_MARK    byte = 0x08
	ID_DADAN_SHU      byte = 0x10
	ID_DASHUANG_SHU   byte = 0x11
	ID_XIAODAN_SHU    byte = 0x12
	ID_XIAOSHUANG_SHU byte = 0x13
)

//骰子
type Dice struct {
	games.GameDesk
}

func (g *Dice) AddScore(player games.PlayInfo, score float64) (int64, error) {
	return g.GameDesk.AddScore(player, score)

}
func (g *Dice) Bet(userid int64, area int) (bool, error) {
	return g.GameDesk.Bet(userid, area)
}

func (g *Dice) EndGame() error {

	g.UnInitTable()
	g.GameStation = games.GS_TK_FREE

	return nil
}

//下注信息
func (g *Dice) GetBetInfos() (bets []logic.Bets, err error) {
	return g.GameDesk.GetBetInfos()

}

//获取ID
func (g *Dice) GetPeriodID() string {
	t1 := time.Now().Year()
	t2 := time.Now().Month()
	t3 := time.Now().Day()
	date := fmt.Sprintf("%d%02d%02d", t1, t2, t3)
	fmt.Println(date)
	isexist, _, err := g.Rdb.GetValue(date)
	fmt.Println(isexist, err)

	return ""
}

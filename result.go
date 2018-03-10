package iksm

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
	"strings"
)

type gear struct {
	Id    string `json:"id"`
	Brand struct {
		Id    string `json:"id,omitempty"`
		Name  string `json:"name,omitempty"`
		Image string `json:"image,omitempty"`
	} `json:"brand"`
	Image  string `json:"image"`
	Rarity int    `json:"rarity"`
	Name   string `json:"name"`
}

type gearskill struct {
	Main struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Image string `json:"image"`
	} `json:"main"`
	Sub []struct {
		Id    string `json:"id,omitempty"`
		Name  string `json:"name,omitempty"`
		Image string `json:"image,omitempty"`
	} `json:"subs"`
}

type weapon struct {
	Name  string `json:"name"`
	Image string `json:"image"`
	Id    string `json:"id"`
	Sub   struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"sub"`
	Special struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"special"`
}

type player struct {
	PaintPoint   int `json:"game_paint_point"`
	DeathCount   int `json:"death_count"`
	KillCount    int `json:"kill_count"`
	AssistCount  int `json:"assist_count"`
	SpecialCount int `json:"special_count"`
	Info         struct {
		Head         gear      `json:"head"`
		Clothes      gear      `json:"clothes"`
		Shoes        gear      `json:"shoes"`
		HeadSkill    gearskill `json:"head_skills"`
		ClothesSkill gearskill `json:"clothes_skills"`
		ShoesSkill   gearskill `json:"shoes_skills"`
		NickName     string    `json:"nickname"`
		Rank         int       `json:"player_rank"`
		Weapon       weapon    `json:"weapon"`
	} `json:"player"`
}

type result struct {
	// id of the battle
	BattleNumber string `json:"battle_number"`
	Rule         struct {
		Name string `json:"name"`
	} `json:"rule"`
	Stage struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Image string `json:"image"`
	} `json:"stage"`
	GameMode struct {
		Key  string `json:"key"`
		Name string `json:"name"`
	} `json:"game_mode"`
	// victory or defeat
	Result struct {
		Result string `json:"key"`
	} `json:"my_team_result"`
	PlayerRank int `json:"player_rank"`
	// My score
	Player player `json:"player_result"`
	// My team members' score
	MyTeamMembers []player `json:"my_team_members"`
	// Other team member's score
	OtherTeamMembers []player `json:"other_team_members"`
}

func (data *result) SaveToDB(dbname string, overwrite bool) error {
	db, err := sql.Open("sqlite3", dbname)
	if err != nil {
		return err
	}
	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS "RESULTS" (
			"ID" INTEGER PRIMARY KEY,
			"RESULT" INTEGER,
			"RULE" VARCHAR(50),
			"TEAMMEMBER1_NAME" VARCHAR(30),
			"TEAMMEMBER1_RANK" INTEGER,
			"TEAMMEMBER1_PAINTPOINT" INTEGER,
			"TEAMMEMBER1_DEATHCOUNT" INTEGER,
			"TEAMMEMBER1_KILLCOUNT" INTEGER,
			"TEAMMEMBER1_ASSISTCOUNT" INTEGER,
			"TEAMMEMBER1_SPECIALCOUNT" INTEGER,
			"TEAMMEMBER1_WEAPON_NAME" VARCHAR(30),
			"TEAMMEMBER1_WEAPON_SUB" VARCHAR(30),
			"TEAMMEMBER1_WEAPON_SPECIAL" VARCHAR(30),
			"TEAMMEMBER1_HEAD_NAME" VARCHAR(30),
			"TEAMMEMBER1_HEAD_MAIN" VARCHAR(30),
			"TEAMMEMBER1_HEAD_SUB1" VARCHAR(30),
			"TEAMMEMBER1_HEAD_SUB2" VARCHAR(30),
			"TEAMMEMBER1_HEAD_SUB3" VARCHAR(30),
			"TEAMMEMBER1_SHOES_NAME" VARCHAR(30),
			"TEAMMEMBER1_SHOES_MAIN" VARCHAR(30),
			"TEAMMEMBER1_SHOES_SUB1" VARCHAR(30),
			"TEAMMEMBER1_SHOES_SUB2" VARCHAR(30),
			"TEAMMEMBER1_SHOES_SUB3" VARCHAR(30),
			"TEAMMEMBER1_CLOTHES_NAME" VARCHAR(30),
			"TEAMMEMBER1_CLOTHES_MAIN" VARCHAR(30),
			"TEAMMEMBER1_CLOTHES_SUB1" VARCHAR(30),
			"TEAMMEMBER1_CLOTHES_SUB2" VARCHAR(30),
			"TEAMMEMBER1_CLOTHES_SUB3" VARCHAR(30),
			"TEAMMEMBER2_NAME" VARCHAR(30),
			"TEAMMEMBER2_RANK" INTEGER,
			"TEAMMEMBER2_PAINTPOINT" INTEGER,
			"TEAMMEMBER2_DEATHCOUNT" INTEGER,
			"TEAMMEMBER2_KILLCOUNT" INTEGER,
			"TEAMMEMBER2_ASSISTCOUNT" INTEGER,
			"TEAMMEMBER2_SPECIALCOUNT" INTEGER,
			"TEAMMEMBER2_WEAPON_NAME" VARCHAR(30),
			"TEAMMEMBER2_WEAPON_SUB" VARCHAR(30),
			"TEAMMEMBER2_WEAPON_SPECIAL" VARCHAR(30),
			"TEAMMEMBER2_HEAD_NAME" VARCHAR(30),
			"TEAMMEMBER2_HEAD_MAIN" VARCHAR(30),
			"TEAMMEMBER2_HEAD_SUB1" VARCHAR(30),
			"TEAMMEMBER2_HEAD_SUB2" VARCHAR(30),
			"TEAMMEMBER2_HEAD_SUB3" VARCHAR(30),
			"TEAMMEMBER2_SHOES_NAME" VARCHAR(30),
			"TEAMMEMBER2_SHOES_MAIN" VARCHAR(30),
			"TEAMMEMBER2_SHOES_SUB1" VARCHAR(30),
			"TEAMMEMBER2_SHOES_SUB2" VARCHAR(30),
			"TEAMMEMBER2_SHOES_SUB3" VARCHAR(30),
			"TEAMMEMBER2_CLOTHES_NAME" VARCHAR(30),
			"TEAMMEMBER2_CLOTHES_MAIN" VARCHAR(30),
			"TEAMMEMBER2_CLOTHES_SUB1" VARCHAR(30),
			"TEAMMEMBER2_CLOTHES_SUB2" VARCHAR(30),
			"TEAMMEMBER2_CLOTHES_SUB3" VARCHAR(30),
			"TEAMMEMBER3_NAME" VARCHAR(30),
			"TEAMMEMBER3_RANK" INTEGER,
			"TEAMMEMBER3_PAINTPOINT" INTEGER,
			"TEAMMEMBER3_DEATHCOUNT" INTEGER,
			"TEAMMEMBER3_KILLCOUNT" INTEGER,
			"TEAMMEMBER3_ASSISTCOUNT" INTEGER,
			"TEAMMEMBER3_SPECIALCOUNT" INTEGER,
			"TEAMMEMBER3_WEAPON_NAME" VARCHAR(30),
			"TEAMMEMBER3_WEAPON_SUB" VARCHAR(30),
			"TEAMMEMBER3_WEAPON_SPECIAL" VARCHAR(30),
			"TEAMMEMBER3_HEAD_NAME" VARCHAR(30),
			"TEAMMEMBER3_HEAD_MAIN" VARCHAR(30),
			"TEAMMEMBER3_HEAD_SUB1" VARCHAR(30),
			"TEAMMEMBER3_HEAD_SUB2" VARCHAR(30),
			"TEAMMEMBER3_HEAD_SUB3" VARCHAR(30),
			"TEAMMEMBER3_SHOES_NAME" VARCHAR(30),
			"TEAMMEMBER3_SHOES_MAIN" VARCHAR(30),
			"TEAMMEMBER3_SHOES_SUB1" VARCHAR(30),
			"TEAMMEMBER3_SHOES_SUB2" VARCHAR(30),
			"TEAMMEMBER3_SHOES_SUB3" VARCHAR(30),
			"TEAMMEMBER3_CLOTHES_NAME" VARCHAR(30),
			"TEAMMEMBER3_CLOTHES_MAIN" VARCHAR(30),
			"TEAMMEMBER3_CLOTHES_SUB1" VARCHAR(30),
			"TEAMMEMBER3_CLOTHES_SUB2" VARCHAR(30),
			"TEAMMEMBER3_CLOTHES_SUB3" VARCHAR(30),
			"TEAMMEMBER4_NAME" VARCHAR(30),
			"TEAMMEMBER4_RANK" INTEGER,
			"TEAMMEMBER4_PAINTPOINT" INTEGER,
			"TEAMMEMBER4_DEATHCOUNT" INTEGER,
			"TEAMMEMBER4_KILLCOUNT" INTEGER,
			"TEAMMEMBER4_ASSISTCOUNT" INTEGER,
			"TEAMMEMBER4_SPECIALCOUNT" INTEGER,
			"TEAMMEMBER4_WEAPON_NAME" VARCHAR(30),
			"TEAMMEMBER4_WEAPON_SUB" VARCHAR(30),
			"TEAMMEMBER4_WEAPON_SPECIAL" VARCHAR(30),
			"TEAMMEMBER4_HEAD_NAME" VARCHAR(30),
			"TEAMMEMBER4_HEAD_MAIN" VARCHAR(30),
			"TEAMMEMBER4_HEAD_SUB1" VARCHAR(30),
			"TEAMMEMBER4_HEAD_SUB2" VARCHAR(30),
			"TEAMMEMBER4_HEAD_SUB3" VARCHAR(30),
			"TEAMMEMBER4_SHOES_NAME" VARCHAR(30),
			"TEAMMEMBER4_SHOES_MAIN" VARCHAR(30),
			"TEAMMEMBER4_SHOES_SUB1" VARCHAR(30),
			"TEAMMEMBER4_SHOES_SUB2" VARCHAR(30),
			"TEAMMEMBER4_SHOES_SUB3" VARCHAR(30),
			"TEAMMEMBER4_CLOTHES_NAME" VARCHAR(30),
			"TEAMMEMBER4_CLOTHES_MAIN" VARCHAR(30),
			"TEAMMEMBER4_CLOTHES_SUB1" VARCHAR(30),
			"TEAMMEMBER4_CLOTHES_SUB2" VARCHAR(30),
			"TEAMMEMBER4_CLOTHES_SUB3" VARCHAR(30),
			"OTHERMEMBER1_NAME" VARCHAR(30),
			"OTHERMEMBER1_RANK" INTEGER,
			"OTHERMEMBER1_PAINTPOINT" INTEGER,
			"OTHERMEMBER1_DEATHCOUNT" INTEGER,
			"OTHERMEMBER1_KILLCOUNT" INTEGER,
			"OTHERMEMBER1_ASSISTCOUNT" INTEGER,
			"OTHERMEMBER1_SPECIALCOUNT" INTEGER,
			"OTHERMEMBER1_WEAPON_NAME" VARCHAR(30),
			"OTHERMEMBER1_WEAPON_SUB" VARCHAR(30),
			"OTHERMEMBER1_WEAPON_SPECIAL" VARCHAR(30),
			"OTHERMEMBER1_HEAD_NAME" VARCHAR(30),
			"OTHERMEMBER1_HEAD_MAIN" VARCHAR(30),
			"OTHERMEMBER1_HEAD_SUB1" VARCHAR(30),
			"OTHERMEMBER1_HEAD_SUB2" VARCHAR(30),
			"OTHERMEMBER1_HEAD_SUB3" VARCHAR(30),
			"OTHERMEMBER1_SHOES_NAME" VARCHAR(30),
			"OTHERMEMBER1_SHOES_MAIN" VARCHAR(30),
			"OTHERMEMBER1_SHOES_SUB1" VARCHAR(30),
			"OTHERMEMBER1_SHOES_SUB2" VARCHAR(30),
			"OTHERMEMBER1_SHOES_SUB3" VARCHAR(30),
			"OTHERMEMBER1_CLOTHES_NAME" VARCHAR(30),
			"OTHERMEMBER1_CLOTHES_MAIN" VARCHAR(30),
			"OTHERMEMBER1_CLOTHES_SUB1" VARCHAR(30),
			"OTHERMEMBER1_CLOTHES_SUB2" VARCHAR(30),
			"OTHERMEMBER1_CLOTHES_SUB3" VARCHAR(30),
			"OTHERMEMBER2_NAME" VARCHAR(30),
			"OTHERMEMBER2_RANK" INTEGER,
			"OTHERMEMBER2_PAINTPOINT" INTEGER,
			"OTHERMEMBER2_DEATHCOUNT" INTEGER,
			"OTHERMEMBER2_KILLCOUNT" INTEGER,
			"OTHERMEMBER2_ASSISTCOUNT" INTEGER,
			"OTHERMEMBER2_SPECIALCOUNT" INTEGER,
			"OTHERMEMBER2_WEAPON_NAME" VARCHAR(30),
			"OTHERMEMBER2_WEAPON_SUB" VARCHAR(30),
			"OTHERMEMBER2_WEAPON_SPECIAL" VARCHAR(30),
			"OTHERMEMBER2_HEAD_NAME" VARCHAR(30),
			"OTHERMEMBER2_HEAD_MAIN" VARCHAR(30),
			"OTHERMEMBER2_HEAD_SUB1" VARCHAR(30),
			"OTHERMEMBER2_HEAD_SUB2" VARCHAR(30),
			"OTHERMEMBER2_HEAD_SUB3" VARCHAR(30),
			"OTHERMEMBER2_SHOES_NAME" VARCHAR(30),
			"OTHERMEMBER2_SHOES_MAIN" VARCHAR(30),
			"OTHERMEMBER2_SHOES_SUB1" VARCHAR(30),
			"OTHERMEMBER2_SHOES_SUB2" VARCHAR(30),
			"OTHERMEMBER2_SHOES_SUB3" VARCHAR(30),
			"OTHERMEMBER2_CLOTHES_NAME" VARCHAR(30),
			"OTHERMEMBER2_CLOTHES_MAIN" VARCHAR(30),
			"OTHERMEMBER2_CLOTHES_SUB1" VARCHAR(30),
			"OTHERMEMBER2_CLOTHES_SUB2" VARCHAR(30),
			"OTHERMEMBER2_CLOTHES_SUB3" VARCHAR(30),
			"OTHERMEMBER3_NAME" VARCHAR(30),
			"OTHERMEMBER3_RANK" INTEGER,
			"OTHERMEMBER3_PAINTPOINT" INTEGER,
			"OTHERMEMBER3_DEATHCOUNT" INTEGER,
			"OTHERMEMBER3_KILLCOUNT" INTEGER,
			"OTHERMEMBER3_ASSISTCOUNT" INTEGER,
			"OTHERMEMBER3_SPECIALCOUNT" INTEGER,
			"OTHERMEMBER3_WEAPON_NAME" VARCHAR(30),
			"OTHERMEMBER3_WEAPON_SUB" VARCHAR(30),
			"OTHERMEMBER3_WEAPON_SPECIAL" VARCHAR(30),
			"OTHERMEMBER3_HEAD_NAME" VARCHAR(30),
			"OTHERMEMBER3_HEAD_MAIN" VARCHAR(30),
			"OTHERMEMBER3_HEAD_SUB1" VARCHAR(30),
			"OTHERMEMBER3_HEAD_SUB2" VARCHAR(30),
			"OTHERMEMBER3_HEAD_SUB3" VARCHAR(30),
			"OTHERMEMBER3_SHOES_NAME" VARCHAR(30),
			"OTHERMEMBER3_SHOES_MAIN" VARCHAR(30),
			"OTHERMEMBER3_SHOES_SUB1" VARCHAR(30),
			"OTHERMEMBER3_SHOES_SUB2" VARCHAR(30),
			"OTHERMEMBER3_SHOES_SUB3" VARCHAR(30),
			"OTHERMEMBER3_CLOTHES_NAME" VARCHAR(30),
			"OTHERMEMBER3_CLOTHES_MAIN" VARCHAR(30),
			"OTHERMEMBER3_CLOTHES_SUB1" VARCHAR(30),
			"OTHERMEMBER3_CLOTHES_SUB2" VARCHAR(30),
			"OTHERMEMBER3_CLOTHES_SUB3" VARCHAR(30),
			"OTHERMEMBER4_NAME" VARCHAR(30),
			"OTHERMEMBER4_RANK" INTEGER,
			"OTHERMEMBER4_PAINTPOINT" INTEGER,
			"OTHERMEMBER4_DEATHCOUNT" INTEGER,
			"OTHERMEMBER4_KILLCOUNT" INTEGER,
			"OTHERMEMBER4_ASSISTCOUNT" INTEGER,
			"OTHERMEMBER4_SPECIALCOUNT" INTEGER,
			"OTHERMEMBER4_WEAPON_NAME" VARCHAR(30),
			"OTHERMEMBER4_WEAPON_SUB" VARCHAR(30),
			"OTHERMEMBER4_WEAPON_SPECIAL" VARCHAR(30),
			"OTHERMEMBER4_HEAD_NAME" VARCHAR(30),
			"OTHERMEMBER4_HEAD_MAIN" VARCHAR(30),
			"OTHERMEMBER4_HEAD_SUB1" VARCHAR(30),
			"OTHERMEMBER4_HEAD_SUB2" VARCHAR(30),
			"OTHERMEMBER4_HEAD_SUB3" VARCHAR(30),
			"OTHERMEMBER4_SHOES_NAME" VARCHAR(30),
			"OTHERMEMBER4_SHOES_MAIN" VARCHAR(30),
			"OTHERMEMBER4_SHOES_SUB1" VARCHAR(30),
			"OTHERMEMBER4_SHOES_SUB2" VARCHAR(30),
			"OTHERMEMBER4_SHOES_SUB3" VARCHAR(30),
			"OTHERMEMBER4_CLOTHES_NAME" VARCHAR(30),
			"OTHERMEMBER4_CLOTHES_MAIN" VARCHAR(30),
			"OTHERMEMBER4_CLOTHES_SUB1" VARCHAR(30),
			"OTHERMEMBER4_CLOTHES_SUB2" VARCHAR(30),
			"OTHERMEMBER4_CLOTHES_SUB3" VARCHAR(30)
		)`)
	if err != nil {
		return err
	}
	id, err := strconv.Atoi(data.BattleNumber)
	if err != nil {
		return err
	}
	result := 0
	if data.Result.Result == `victory` {
		result = 1
	}
	if overwrite {
		_, err = db.Exec(
			`INSERT OR REPLACE INTO RESULTS (ID, RESULT, RULE) VALUES (?,?,?)`,
			id, result, data.Rule.Name,
		)
	} else {
		_, err = db.Exec(
			`INSERT INTO RESULTS (ID, RESULT, RULE) VALUES (?,?,?)`,
			id, result, data.Rule.Name,
		)
	}
	if err != nil {
		return err
	}
	err = writePlayer(db, data.Player, `TEAMMEMBER1`, id)
	if err != nil {
		db.Exec(`DELETE FROM RESULTS WHERE ID=?`, id)
		return err
	}
	for i, player := range data.MyTeamMembers {
		err = writePlayer(db, player, fmt.Sprintf(`TEAMMEMBER%d`, i+2), id)
		if err != nil {
			db.Exec(`DELETE FROM RESULTS WHERE ID=?`, id)
			return err
		}
	}
	for i, player := range data.OtherTeamMembers {
		err = writePlayer(db, player, fmt.Sprintf(`OTHERMEMBER%d`, i+1), id)
		if err != nil {
			db.Exec(`DELETE FROM RESULTS WHERE ID=?`, id)
			return err
		}
	}
	fmt.Printf("Successfully write the battle %d\n", id)
	return nil
}

func writePlayer(db *sql.DB, p player, name string, id int) error {
	template := `(<NAME>_NAME, <NAME>_RANK, <NAME>_PAINTPOINT, <NAME>_DEATHCOUNT, <NAME>_KILLCOUNT, <NAME>_ASSISTCOUNT, <NAME>_SPECIALCOUNT, <NAME>_WEAPON_NAME, <NAME>_WEAPON_SUB, <NAME>_WEAPON_SPECIAL, <NAME>_HEAD_NAME, <NAME>_HEAD_MAIN, <NAME>_HEAD_SUB1, <NAME>_HEAD_SUB2, <NAME>_HEAD_SUB3, <NAME>_SHOES_NAME, <NAME>_SHOES_MAIN, <NAME>_SHOES_SUB1, <NAME>_SHOES_SUB2, <NAME>_SHOES_SUB3, <NAME>_CLOTHES_NAME, <NAME>_CLOTHES_MAIN, <NAME>_CLOTHES_SUB1, <NAME>_CLOTHES_SUB2, <NAME>_CLOTHES_SUB3)`
	cnt := len(strings.Split(template, `,`))
	qs := `(` + strings.Repeat(`?,`, cnt)[:cnt*2-1] + `)`
	keys := strings.Replace(template, `<NAME>`, name, -1)
	_, err := db.Exec(
		`UPDATE RESULTS set `+keys+` = `+qs+` WHERE ID=?`,
		p.Info.NickName,
		p.Info.Rank,
		p.PaintPoint,
		p.DeathCount,
		p.KillCount,
		p.AssistCount,
		p.SpecialCount,
		p.Info.Weapon.Name,
		p.Info.Weapon.Sub.Name,
		p.Info.Weapon.Special.Name,
		p.Info.Head.Name,
		p.Info.HeadSkill.Main.Name,
		p.Info.HeadSkill.Sub[0].Name,
		p.Info.HeadSkill.Sub[1].Name,
		p.Info.HeadSkill.Sub[2].Name,
		p.Info.Shoes.Name,
		p.Info.ShoesSkill.Main.Name,
		p.Info.ShoesSkill.Sub[0].Name,
		p.Info.ShoesSkill.Sub[1].Name,
		p.Info.ShoesSkill.Sub[2].Name,
		p.Info.Clothes.Name,
		p.Info.ClothesSkill.Main.Name,
		p.Info.ClothesSkill.Sub[0].Name,
		p.Info.ClothesSkill.Sub[1].Name,
		p.Info.ClothesSkill.Sub[2].Name,
		id,
	)
	return err
}
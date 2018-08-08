package main

// import (
// 	"ibgame/logs"
// 	"ibgame/models/manage_model"
// 	"ibgame/models/mysql"
// )

// func main() {
// 	pi := make([]manage_model.PlayerInfo, 0)
// 	pe := make([]manage_model.PlayerExtro, 0)
// 	extroMap := make(map[int64]manage_model.PlayerExtro)
// 	pidMap := make(map[int64]manage_model.PlayerInfo)
// 	pids := make([]int64, 0)
// 	engine, e := mysql.GetEngine()
// 	if e != nil {
// 		logs.Info.Println("db err", e)
// 	}
// 	if e1 := engine.Find(&pi); e1 != nil {
// 		logs.Info.Println("db err1", e1)
// 	}
// 	if e2 := engine.Find(&pe); e2 != nil {
// 		logs.Info.Println("db err2", e2)
// 	}
// 	for _, v := range pi {
// 		pids = append(pids, v.PlayerId)
// 		pidMap[v.PlayerId] = manage_model.PlayerInfo{Name: v.Name, NickName: v.NickName, Position: v.Position, SecondPosition: v.SecondPosition, Type: v.Type, Score: v.Score, Rebound: v.Rebound, Assist: v.Assist, Steal: v.Steal, Cap: v.Cap, AppearNum: v.AppearNum}
// 	}
// 	for _, pe := range pe {
// 		extroMap[pe.PlayerId] = manage_model.PlayerExtro{OffensiveRebound: pe.OffensiveRebound, DefensiveRebound: pe.DefensiveRebound, ThreeAttempt: pe.ThreeAttempt, ThreeMade: pe.ThreeMade, InsideAttempt: pe.InsideAttempt, InsideMade: pe.InsideMade, TurnOff: pe.TurnOff, ShotAttempt: pe.ShotAttempt, ShotMade: pe.ShotMade, Faul: pe.Faul}
// 	}
// }

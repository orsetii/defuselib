package parse

import (
	"fmt"
	"os"

	"github.com/golang/geo/r3"
	demo "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs"
	common "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/common"
	events "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/events"
	"github.com/orsetii/defuse/cmd"
)

type nadePath struct {
	wep  common.EquipmentType
	path []r3.Vector
	team common.Team
}

type nadeMap map[int64]*nadePath

type teamData struct {
	ClanName string
	Players  []*common.Player
}

// @TODO Add more fields and data here.
// @TODO Make a constructor ton extract all applicable data possible before actual parsing of demo starts.
//SerData is a struct containing the serialized data extracted from a demo file.
type SerData struct {
	NadePaths nadeMap
	Teams     [2]teamData // Teams contains the name of the team, as well as data on the players.
	Header    common.DemoHeader
}

// ParseDemo the function to extract and serialize all data from a demo file.
func ParseDemo(f *os.File, verbose bool) (err error) {
	// @TODO Will need to add buffered parsing at some point.
	// Likely of 2MB or more.
	defer f.Close() // @TODO Remove after changing from os.Open method to a buffered reader.

	// Create new parser to parse file pointed to by args.
	p := demo.NewParser(f)
	defer p.Close()

	// Create new demo data struct from constructor.
	demoInfo, err := NewSerData(p)
	if err != nil {
		return err
	}

	// Event Handler to register and store nade paths into a map.
	p.RegisterEventHandler(func(e events.GrenadeProjectileDestroy) {
		id := e.Projectile.UniqueID()

		// Sometimes the thrower is nil, in that case we want the team to be unassigned (which is the default value)
		var team common.Team
		if e.Projectile.Thrower != nil {
			team = common.Team(e.Projectile.Thrower.Team)
		}

		if demoInfo.NadePaths[id] == nil {
			demoInfo.NadePaths[id] = &nadePath{
				wep:  e.Projectile.WeaponInstance.Type,
				team: team,
			}
		}

		demoInfo.NadePaths[id].path = e.Projectile.Trajectory
	})

	fmt.Printf("\n")
	// @TODO Create custom parsing loop here, print metadata & progress as it goes. remember carriage returns
	for frame := 0; frame < demoInfo.Header.PlaybackTicks; frame++ {
		prog := p.Progress() * 100
		cmd.PrintInfo(fmt.Sprintf("%f%% Done\r", prog))
		cont, err := p.ParseNextFrame()
		if !cont {
			break
		}
		if err != nil {
			return err
		}
	}

	// // Parse demo to its end.
	// // Check for errors and ret if any found.
	// err = p.ParseToEnd()
	// if err != nil {
	// 	return err
	// }
	// We have to assign some data here as not possible to get pre-parsing.
	demoInfo.Teams[0].ClanName = p.GameState().TeamCounterTerrorists().ClanName()
	demoInfo.Teams[1].ClanName = p.GameState().TeamTerrorists().ClanName()
	demoInfo.Teams[0].Players = p.GameState().TeamCounterTerrorists().Members()
	demoInfo.Teams[1].Players = p.GameState().TeamTerrorists().Members()
	fmt.Println(demoInfo.Teams[0].Players[0].SteamID64)
	fmt.Printf("%v", demoInfo.Teams)

	return nil

}

func NewSerData(p demo.Parser) (demInfo SerData, err error) {

	// Parse header and store in our struct
	demInfo.Header, err = p.ParseHeader()
	if err != nil {
		return demInfo, err
	}

	// Alloc memory for nadepath map.
	// Alloc a reasonable number(cant set max due to OTs) 2 nades thrown per round * 10 players * 30 rounds = 600 elements this is NOT capcity, just intial mem alloc
	demInfo.NadePaths = make(nadeMap, 600)

	return demInfo, nil

}

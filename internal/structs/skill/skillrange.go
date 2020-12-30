package skill

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/nivek706/d2skillcalc/pkg/fileutil"
)

type SkillRange struct {
	skillName   string
	skillFile   *fileutil.File
	missileFile *fileutil.File
	startLevel  int
	endLevel    int
	skillArray  []Skill
}

func NewSkillRange(skillName string, skillFile *fileutil.File, missileFile *fileutil.File, startLevel int, endLevel int) *SkillRange {
	skillRange := &SkillRange{
		skillName:   skillName,
		skillFile:   skillFile,
		missileFile: missileFile,
		startLevel:  startLevel,
		endLevel:    endLevel}

	skillArray := make([]Skill, (endLevel-startLevel)+1)
	for i := range skillArray {
		tmpSkill := NewSkill(skillName, skillFile, missileFile, startLevel+i)
		tmpSkill.PopulateMissileDamage()
		tmpSkill.populateElementalDamage()
		skillArray[i] = *tmpSkill
	}
	skillRange.skillArray = skillArray
	return skillRange
}

func (sr SkillRange) PrintSkillTable() {

	leveloffset := len(sr.skillArray)

	skillinfo := make(map[string][]interface{})

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	//set levels
	skillinfo["level"] = make([]interface{}, leveloffset+1)
	skillinfo["level"][0] = "Level"
	for i := 1; i <= leveloffset; i++ {
		skillinfo["level"][i] = float64(i + (sr.startLevel - 1))
	}
	t.AppendHeader(skillinfo["level"])

	//get skill mana costs
	skillinfo["manacost"] = make([]interface{}, leveloffset+1)
	skillinfo["manacost"][0] = "Mana Cost"
	for i := 1; i <= leveloffset; i++ {
		skillinfo["manacost"][i] = fmt.Sprintf("%.1f", sr.skillArray[i-1].manacost)
	}
	t.AppendRow(skillinfo["manacost"])

	// get skill damage information
	skillinfo["eledmg"] = make([]interface{}, leveloffset+1)
	skillinfo["eledmg"][0] = fmt.Sprintf("%s Dmg", sr.skillArray[0].elementalDamage.DmgType)
	for i := 1; i <= leveloffset; i++ {
		eleDmg := sr.skillArray[i-1].elementalDamage
		skillinfo["eledmg"][i] = fmt.Sprintf("%.1f - %.1f", eleDmg.Min, eleDmg.Max)
	}
	t.AppendRow(skillinfo["eledmg"])

	for i, missile := range sr.skillArray[0].missileDamage {
		missileIndex := fmt.Sprintf("missile%d", i)
		skillinfo[fmt.Sprintf(missileIndex)] = make([]interface{}, 1)
		skillinfo[missileIndex][0] = missile.DmgType
	}

	// get missile damage
	// fmt.Printf("len(sr.skillArray): %v\n", len(sr.skillArray))
	for j := 0; j < len(sr.skillArray); j++ {
		missileDmg := sr.skillArray[j].missileDamage
		for i := 0; i < len(missileDmg); i++ {
			missileIndex := fmt.Sprintf("missile%d", i)
			// skillinfo[missileIndex][i] = fmt.Sprintf("%.0f - %.0f", missileDmg[j][i], missileDmg[j][i+1])
			skillinfo[missileIndex] = append(skillinfo[missileIndex], fmt.Sprintf("%.1f - %.1f", missileDmg[i].Min, missileDmg[i].Max))
		}
	}

	for i := range sr.skillArray[0].missileDamage {
		missileIndex := fmt.Sprintf("missile%d", i)
		t.AppendRow(skillinfo[missileIndex])
	}

	t.Render()
}

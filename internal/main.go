package main

import (
	"os"
	"fmt"
	"math"
	"bufio"
	"strconv"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"

	"github.com/nivek706/d2skillcalc/pkg/fileutil"
	"github.com/nivek706/d2skillcalc/pkg/index/skills"

)

func getSkillRecord(file *fileutil.File, skillname string) []string {
	var skill []string
	for row := range file.Rows {
		if (file.Rows[row][skills.Skill] == skillname) {
			skill = file.Rows[row]
		}
	}
	return skill
}

func printSkillInfo(file *fileutil.File, skillname string, skilllevel int) {
	skill := getSkillRecord(file, skillname)

	if skill != nil {
		fmt.Println("Skill: " + skill[skills.Skill])
		fmt.Println("Id: " + skill[skills.Id])
		basemana, _ := strconv.ParseFloat(skill[skills.Mana], 64)
		startmana, _ := strconv.ParseFloat(skill[skills.StartMana], 64)
		lvlmana, _ := strconv.ParseFloat(skill[skills.LvlMana], 64)
		manashift, _ := strconv.ParseFloat(skill[skills.ManaShift], 64)
		minmana, _ := strconv.ParseFloat(skill[skills.MinMana], 64)
		manacost := calcManaCost(skilllevel, basemana, startmana, lvlmana, manashift, minmana)
		fmt.Printf("Calculated mana cost: %f\n", manacost)


		// get damage information
	}
}

func printSkillTable(file *fileutil.File, skillname string, startlevel int, maxlevel int) {
	leveloffset := maxlevel-startlevel+1
	skillrecord := getSkillRecord(file, skillname)

	if skillrecord == nil {
		fmt.Println("Could not find skill name: " + skillname)
		return
	}

/* 	
	fmt.Println("HitShift: " + skillrecord[skills.HitShift])
	fmt.Println("SrcDam: " + skillrecord[skills.SrcDam])
	fmt.Println("MinDam: " + skillrecord[skills.MinDam])
	fmt.Println("MaxDam: " + skillrecord[skills.MaxDam])
	fmt.Printf("MinLevDam1-5: %s, %s, %s, %s, %s\n", 
		skillrecord[skills.MinLevDam1], 
		skillrecord[skills.MinLevDam2],
		skillrecord[skills.MinLevDam3], 
		skillrecord[skills.MinLevDam4], 
		skillrecord[skills.MinLevDam5])
	fmt.Printf("MaxLevDam1-5: %s, %s, %s, %s, %s\n", 
		skillrecord[skills.MaxLevDam1], 
		skillrecord[skills.MaxLevDam2],
		skillrecord[skills.MaxLevDam3], 
		skillrecord[skills.MaxLevDam4], 
		skillrecord[skills.MaxLevDam5])
	fmt.Println("DmgSymPerCalc: " + skillrecord[skills.DmgSymPerCalc])
	fmt.Println("EType: " + skillrecord[skills.EType])
	fmt.Println("EMin: " + skillrecord[skills.EMin])
	fmt.Println("EMax: " + skillrecord[skills.EMax])
	fmt.Printf("EMinLev1-5: %s, %s, %s, %s, %s\n", 
		skillrecord[skills.EMinLev1], 
		skillrecord[skills.EMinLev2],
		skillrecord[skills.EMinLev3], 
		skillrecord[skills.EMinLev4], 
		skillrecord[skills.EMinLev5])
	fmt.Printf("EMaxLev1-5: %s, %s, %s, %s, %s\n", 
		skillrecord[skills.EMaxLev1], 
		skillrecord[skills.EMaxLev2],
		skillrecord[skills.EMaxLev3], 
		skillrecord[skills.EMaxLev4], 
		skillrecord[skills.EMaxLev5])
	fmt.Println("EDmgSymPerCalc: " + skillrecord[skills.EDmgSymPerCalc])
	fmt.Println("ELen: " + skillrecord[skills.ELen])
	fmt.Printf("ELevLen1-3: %s, %s, %s\n", skillrecord[skills.ELevLen1], skillrecord[skills.ELevLen2], skillrecord[skills.ELevLen3])
	fmt.Println("ELenSymPerCalc: " + skillrecord[skills.ELenSymPerCalc])
	fmt.Println("par8? : " + skillrecord[skills.Param8])
 */
	skillinfo := make(map[string][]interface{})

	//set levels
	skillinfo["level"] = make([]interface{}, leveloffset+1)
	skillinfo["level"][0] = "Level"
	for i := 1; i <= leveloffset; i++ {
		skillinfo["level"][i] = float64(i+(startlevel-1))
	}

	//get skill mana costs
	skillinfo["manacost"] = make([]interface{}, leveloffset+1)
	skillinfo["manacost"][0] = "Mana Cost"
	for i := 1; i <= leveloffset; i++ {
		skillinfo["manacost"][i] = fmt.Sprintf("%.1f", getSkillManaCost(skillrecord, int(i+(startlevel-1))))
	}

	// get skill damage information (this is going to be difficult)
	skillinfo["eledmg"] = make([]interface{}, leveloffset+1)
	skillinfo["eledmg"][0] = "Ele Dmg"
	for i := 1; i <= leveloffset; i++ {
		eleDmg := getSkillDamageValues(skillrecord, int(i+(startlevel-1)))
		skillinfo["eledmg"][i] = fmt.Sprintf("%.0f - %.0f", eleDmg[0], eleDmg[1])
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(skillinfo["level"])
	t.AppendRow(skillinfo["manacost"])
	t.AppendRow(skillinfo["eledmg"])
	t.Render()
}

func getSkillDamageValues(skillRecord []string, sLvl int) []float64 {
	// return value indices
	// 0 - calculated minimum damage
	// 1 - calculated max damage
	damageValues := make([]float64, 2)

	//calculate elemental damage
	//damage[type] = ((base_damage*srcdamage/128)+(skill_damage*effectiveshift*(100+synergy)/100))*(100+enhancer_stat)/100

	//LevDmg1 sLvl 2-8
	//LevDmg2 sLvl 9-16
	//LevDmg3 sLvl 17-22
	//LevDmg4 sLvl 23-28
	//LevDmg5 sLvl 29+
	hitShift, _ := strconv.ParseFloat(skillRecord[skills.HitShift], 64)

	//get min ele damage stats
	eMin, _ := strconv.ParseFloat(skillRecord[skills.EMin], 64)
	eMinLev1, _ := strconv.ParseFloat(skillRecord[skills.EMinLev1], 64)
	eMinLev2, _ := strconv.ParseFloat(skillRecord[skills.EMinLev2], 64)
	eMinLev3, _ := strconv.ParseFloat(skillRecord[skills.EMinLev3], 64)
	eMinLev4, _ := strconv.ParseFloat(skillRecord[skills.EMinLev4], 64)
	eMinLev5, _ := strconv.ParseFloat(skillRecord[skills.EMinLev5], 64)

	minEleDmg := calculateDamage(sLvl, hitShift, eMin, eMinLev1, eMinLev2, eMinLev3, eMinLev4, eMinLev5)


	//get max ele damage stats
	eMax, _ := strconv.ParseFloat(skillRecord[skills.EMax], 64)
	eMaxLev1, _ := strconv.ParseFloat(skillRecord[skills.EMaxLev1], 64)
	eMaxLev2, _ := strconv.ParseFloat(skillRecord[skills.EMaxLev2], 64)
	eMaxLev3, _ := strconv.ParseFloat(skillRecord[skills.EMaxLev3], 64)
	eMaxLev4, _ := strconv.ParseFloat(skillRecord[skills.EMaxLev4], 64)
	eMaxLev5, _ := strconv.ParseFloat(skillRecord[skills.EMaxLev5], 64)

	maxEleDmg := calculateDamage(sLvl, hitShift, eMax, eMaxLev1, eMaxLev2, eMaxLev3, eMaxLev4, eMaxLev5)

	damageValues[0] = minEleDmg
	damageValues[1] = maxEleDmg
	return damageValues
}

func calculateDamage(sLvl int, hitShift float64, baseDmg float64, levDmg1 float64, levDmg2 float64, levDmg3 float64, levDmg4 float64, levDmg5 float64) float64 {
	var calcDmg float64
	calcDmg = baseDmg
	//get added damage sLvl 2-8
	if sLvl > 8 {
		calcDmg = calcDmg + (levDmg1 * 7)
	} else if sLvl > 1 {
		calcDmg = calcDmg + (levDmg1 * float64(sLvl - 1))
	}
	//get added damage sLvl 9-16
	if sLvl > 16 {
		calcDmg = calcDmg + (levDmg2 * 8)
	} else if sLvl > 8 {
		calcDmg = calcDmg + (levDmg2 * float64(sLvl - 8))
	}
	//get added damage sLvl 17-22
	if sLvl > 22 {
		calcDmg = calcDmg + (levDmg3 * 6)
	} else if sLvl > 16 {
		calcDmg = calcDmg + (levDmg3 * float64(sLvl - 16))
	}
	//get added damage sLvl 23-28
	if sLvl > 28 {
		calcDmg = calcDmg + (levDmg4 * 6)
	} else if sLvl > 22 {
		calcDmg = calcDmg + (levDmg4 * float64(sLvl - 22))
	}
	//get added damage sLvl 29+
	if sLvl > 29 {
		calcDmg = calcDmg + (levDmg5 * float64(sLvl -28))
	}

	//add in hitshift operator
	effectiveShift := math.Exp2(float64(hitShift))/256

	calcDmg = calcDmg * effectiveShift

	//add in synergy damage??

	return calcDmg
}

func getSkillManaCost(skillRecord []string, sLvl int) float64 {
	basemana, _ := strconv.ParseFloat(skillRecord[skills.Mana], 64)
	startmana, _ := strconv.ParseFloat(skillRecord[skills.StartMana], 64)
	lvlmana, _ := strconv.ParseFloat(skillRecord[skills.LvlMana], 64)
	manashift, _ := strconv.ParseFloat(skillRecord[skills.ManaShift], 64)
	minmana, _ := strconv.ParseFloat(skillRecord[skills.MinMana], 64)
	manacost := calcManaCost(sLvl, basemana, startmana, lvlmana, manashift, minmana)
	return manacost
}

func calcManaCost(sLvl int, basemana float64, startmana float64, lvlmana float64, manashift float64, minmana float64) float64 {
	// manacost = max((mana+lvlmana*(sLvl-1))*effectiveshift,minmana);

	effectiveshift := math.Exp2(float64(manashift))/256


	calcmana := ((basemana + (lvlmana * float64(sLvl-1))) * effectiveshift) + startmana

	if (calcmana < minmana) {
		calcmana = minmana
	}

	return calcmana
}

func main() {

	//read the Skills.txt file, load into memory
	skillfile := fileutil.ReadFile("../assets/113c-data/Skills.txt")

	startInputLoop(skillfile)
}

func startInputLoop(skillfile *fileutil.File) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter skill name:")
		text, _ := reader.ReadString('\n')
		text = strings.TrimRight(text, "\r\n")
		if (text == "exit") { break }
		printSkillTable(skillfile, text, 1, 10)
		printSkillTable(skillfile, text, 11, 20)
		printSkillTable(skillfile, text, 21, 30)
		printSkillTable(skillfile, text, 31, 40)
	}
}
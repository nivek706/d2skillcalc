package skill

import (
	"math"
	"strconv"

	"github.com/nivek706/d2skillcalc/internal/common"
	"github.com/nivek706/d2skillcalc/internal/structs/damage"
	"github.com/nivek706/d2skillcalc/pkg/fileutil"
	"github.com/nivek706/d2skillcalc/pkg/index/eletypelookup"
	"github.com/nivek706/d2skillcalc/pkg/index/missiles"
	"github.com/nivek706/d2skillcalc/pkg/index/skills"
)

type Skill struct {
	name            string
	skillFile       *fileutil.File
	missileFile     *fileutil.File
	level           int
	manacost        float64
	physicalDamage  damage.Damage
	elementalDamage damage.Damage
	missileDamage   []damage.Damage
}

func NewSkill(name string, skillFile *fileutil.File, missileFile *fileutil.File, level int) *Skill {
	s := &Skill{name: name, skillFile: skillFile, missileFile: missileFile, level: level}
	s.manacost = s.getManaCost()
	return s
}

func (skill Skill) populateSkillDamage() {

}

func (skill Skill) populatePhysicalDamage() {

}

func (skill *Skill) populateElementalDamage() {
	skill.elementalDamage = skill.getSkillEleDamageValues()

}

func (skill *Skill) PopulateMissileDamage() {
	skillRecord := skill.getSkillRecord()
	// fmt.Println(skillRecord)
	missileDamageArray := skill.getSkillMissileDamage(skillRecord)
	skill.missileDamage = missileDamageArray
	// for i := 0; i < len(missileDamageArray); i++ {
	// 	// tempMissileDmg := make(Damage)
	// 	fmt.Printf("missileDamageArray[%d]: %v\n", i, missileDamageArray[i])
	// }

}

func (skill Skill) getSkillRecord() []string {
	var skillRecord []string
	for row := range skill.skillFile.Rows {
		if skill.skillFile.Rows[row][skills.Skill] == skill.name {
			skillRecord = skill.skillFile.Rows[row]
		}
	}
	return skillRecord
}

func (skill Skill) getMissileRecord(missileName string) []string {
	var missile []string
	for row := range skill.missileFile.Rows {
		if skill.missileFile.Rows[row][missiles.Missile] == missileName {
			missile = skill.missileFile.Rows[row]
		}
	}
	return missile
}

func (skill Skill) getManaCost() float64 {
	skillRecord := skill.getSkillRecord()
	basemana, _ := strconv.ParseFloat(skillRecord[skills.Mana], 64)
	startmana, _ := strconv.ParseFloat(skillRecord[skills.StartMana], 64)
	lvlmana, _ := strconv.ParseFloat(skillRecord[skills.LvlMana], 64)
	manashift, _ := strconv.ParseFloat(skillRecord[skills.ManaShift], 64)
	minmana, _ := strconv.ParseFloat(skillRecord[skills.MinMana], 64)
	manacost := calcManaCost(skill.level, basemana, startmana, lvlmana, manashift, minmana)
	return manacost
}

func (skill Skill) getSkillEleDamageValues() damage.Damage {
	skillRecord := skill.getSkillRecord()

	// return value indices
	// 0 - calculated minimum damage
	// 1 - calculated max damage
	// damageValues := make([]float64, 2)

	//calculate elemental damage
	//damage[type] = ((base_damage*srcdamage/128)+(skill_damage*effectiveshift*(100+synergy)/100))*(100+enhancer_stat)/100

	//LevDmg1 sLvl 2-8
	//LevDmg2 sLvl 9-16
	//LevDmg3 sLvl 17-22
	//LevDmg4 sLvl 23-28
	//LevDmg5 sLvl 29+
	hitShift, _ := strconv.ParseFloat(skillRecord[skills.HitShift], 64)

	/*
		Another shot at the logic for calculating missile damage
		If the skillrow contains anything in srvmissile*
			For each value of srvmissile*
				Calculate that missile damage (may need some hardcode stuff here, ex: firewallmaker vs firewall)
				May need to use the damage columns from the missiles.txt, otherwise inherit parent skills.txt damage columns
		Else
			Assume that the skill has no missile, but if it has damage columns, then calculate that damage

	*/

	//get min ele damage stats
	eMin, _ := strconv.ParseFloat(skillRecord[skills.EMin], 64)
	eMinLev1, _ := strconv.ParseFloat(skillRecord[skills.EMinLev1], 64)
	eMinLev2, _ := strconv.ParseFloat(skillRecord[skills.EMinLev2], 64)
	eMinLev3, _ := strconv.ParseFloat(skillRecord[skills.EMinLev3], 64)
	eMinLev4, _ := strconv.ParseFloat(skillRecord[skills.EMinLev4], 64)
	eMinLev5, _ := strconv.ParseFloat(skillRecord[skills.EMinLev5], 64)

	minEleDmg := calculateDamage(skill.level, hitShift, eMin, eMinLev1, eMinLev2, eMinLev3, eMinLev4, eMinLev5)

	//get max ele damage stats
	eMax, _ := strconv.ParseFloat(skillRecord[skills.EMax], 64)
	eMaxLev1, _ := strconv.ParseFloat(skillRecord[skills.EMaxLev1], 64)
	eMaxLev2, _ := strconv.ParseFloat(skillRecord[skills.EMaxLev2], 64)
	eMaxLev3, _ := strconv.ParseFloat(skillRecord[skills.EMaxLev3], 64)
	eMaxLev4, _ := strconv.ParseFloat(skillRecord[skills.EMaxLev4], 64)
	eMaxLev5, _ := strconv.ParseFloat(skillRecord[skills.EMaxLev5], 64)

	maxEleDmg := calculateDamage(skill.level, hitShift, eMax, eMaxLev1, eMaxLev2, eMaxLev3, eMaxLev4, eMaxLev5)

	//TODO: factor in missile damage functions/calculations
	//reference info: https://d2mods.info/forum/viewtopic.php?f=122&t=29595
	/*
		 * Trying to plan out the logic
		 * Traverse backward from srvmissilec to srvmissile value
		 * Once (if) a non-empty value is found, set missileFunc to that value
		 * If missileFunc is non-empty, apply any damage calculation
		 ** How to handle skills with two damage instances? aka, Meteor
				 - When finding the missilefunc, travel into the details of it in Missiles.txt
				 - Find and submissiles in Missiles.txt
				 - If there is damage info in Missiles.txt (EMin, MinELev1, etc), then calculate that as a new line of damage
				 	- example: Fire Wall has no EMin, etc in Missiles.txt, so there is only one damage instance
				 	- Meteor has EMin and other) damage info, so there are two damage instances, one from Skills.txt and Missiles.txt
		 *
	*/

	eLen, _ := strconv.ParseFloat(skillRecord[skills.ELen], 64)
	eLevLen1, _ := strconv.ParseFloat(skillRecord[skills.ELevLen1], 64)
	eLevLen2, _ := strconv.ParseFloat(skillRecord[skills.ELevLen2], 64)
	eLevLen3, _ := strconv.ParseFloat(skillRecord[skills.ELevLen3], 64)

	length := calcLength(skill.level, eLen, eLevLen1, eLevLen2, eLevLen3)

	var missileFunc string = ""
	if skillRecord[skills.SrvMissileC] != "" {
		missileFunc = skillRecord[skills.SrvMissileC]
	} else if skillRecord[skills.SrvMissileB] != "" {
		missileFunc = skillRecord[skills.SrvMissileB]
	} else if skillRecord[skills.SrvMissileA] != "" {
		missileFunc = skillRecord[skills.SrvMissileA]
	} else if skillRecord[skills.SrvMissile] != "" {
		missileFunc = skillRecord[skills.SrvMissile]
	}

	if missileFunc != "" {
		//apply the missileFunc damage calculation/transformation
		minEleDmg = calculateMissileFuncDamage(missileFunc, minEleDmg, length)
		maxEleDmg = calculateMissileFuncDamage(missileFunc, maxEleDmg, length)
	}

	damageValues := damage.Damage{DmgType: eletypelookup.EType[skillRecord[skills.EType]], Min: minEleDmg, Max: maxEleDmg}

	return damageValues
}

func (s Skill) getSkillMissileDamage(skillRecord []string) []damage.Damage {
	// fmt.Println("Entered getSkillMissileDamageValues")
	//returns a 2D array of all missile damage values for a skill
	missileDamageValues := make([]damage.Damage, 0)

	//srvmissile
	if skillRecord[skills.SrvMissile] != "" {
		missileDmg := calculateMissileDamage(skillRecord[skills.SrvMissile], s.level, s)
		for j := 0; j < len(missileDmg); j++ {
			missileDamageValues = append(missileDamageValues, missileDmg[j])
		}
	}
	//srvmissilea
	if skillRecord[skills.SrvMissileA] != "" {
		missileDmg := calculateMissileDamage(skillRecord[skills.SrvMissileA], s.level, s)
		for j := 0; j < len(missileDmg); j++ {
			missileDamageValues = append(missileDamageValues, missileDmg[j])
		}
	}

	//srvmissileb
	if skillRecord[skills.SrvMissileB] != "" {
		missileDmg := calculateMissileDamage(skillRecord[skills.SrvMissileB], s.level, s)
		for j := 0; j < len(missileDmg); j++ {
			missileDamageValues =
				append(missileDamageValues,
					missileDmg[j])
			// fmt.Printf("missileDamageValues: %v\n", missileDamageValues)
			// fmt.Println(missileDamageValues)
			// fmt.Printf("missileDamageValues: %v\n", missileDamageValues)
		}
	}

	//srvmissilec
	if skillRecord[skills.SrvMissileC] != "" {
		missileDmg := calculateMissileDamage(skillRecord[skills.SrvMissileC], s.level, s)
		for j := 0; j < len(missileDmg); j++ {
			missileDamageValues =
				append(missileDamageValues,
					missileDmg[j])
			// fmt.Println(missileDamageValues)
			// fmt.Printf("missileDamageValues: %v\n", missileDamageValues)
		}

	}

	// TODO: remove duplicates from the Missile damage array
	missileDamageValues = common.RemoveDuplicateDamageRows(missileDamageValues)

	return missileDamageValues
}
func (skill Skill) getSkillMissileDamageValues(skillRecord []string, missileFile *fileutil.File, skillFile *fileutil.File, startlevel int, maxlevel int) [][]damage.Damage {
	// fmt.Println("Entered getSkillMissileDamageValues")
	//returns a 2D array of all missile damage values for a skill
	missileDamageValues := make([][]damage.Damage, 0)

	leveloffset := maxlevel - startlevel + 1
	damageValuesIndexOffset := 0

	//srvmissile
	if skillRecord[skills.SrvMissile] != "" {
		// fmt.Println("srvmissile")
		for i := 1; i <= leveloffset; i++ {
			missileDmg := calculateMissileDamage(skillRecord[skills.SrvMissile], int(i+(startlevel-1)), skill)
			// tempArray := make([][]float64, len(missileDmg))
			// missileDamageValues = append(missileDamageValues, tempArray...)
			// for i := 0; i < len(missileDmg); i++ {
			// 	missileDamageValues[i] = append(missileDamageValues[i], missileDmg[i])

			// }
			// fmt.Printf("len(missileDamageValues): %d\n", len(missileDamageValues))
			// if len(missileDamageValues) < len(missileDmg) {
			// 	tempArray := make([][]float64, (len(missileDmg) - len(missileDamageValues)))
			// 	missileDamageValues = append(missileDamageValues, tempArray...)
			// 	fmt.Printf("missileDamageValues length: %d\n", len(missileDamageValues))
			// }
			for j := 0; j < len(missileDmg); j++ {
				// on the first iteration, add new row(s)
				if i == 1 {
					tempArray := make([][]damage.Damage, 1)
					missileDamageValues = append(missileDamageValues, tempArray...)
				}
				missileDamageValues[j] = append(missileDamageValues[j],
					missileDmg[j])
			}
			// fmt.Println(missileDamageValues)
		}

	}

	//srvmissilea
	if skillRecord[skills.SrvMissileA] != "" {
		// fmt.Println("srvmissilea")
		for i := 1; i <= leveloffset; i++ {
			// fmt.Printf("i: %d\n", i)
			missileDmg := calculateMissileDamage(skillRecord[skills.SrvMissileA], int(i+(startlevel-1)), skill)
			// fmt.Printf("len(missileDamageValues): %d\n", len(missileDamageValues))
			// fmt.Printf("len(missileDmg): %d\n", len(missileDmg))
			// fmt.Printf("missileDmg: %v\n", missileDmg)

			if i == 1 {
				damageValuesIndexOffset = len(missileDamageValues)
			}

			for j := 0; j < len(missileDmg); j++ {
				// on the first iteration, add new row(s)
				if i == 1 {
					tempArray := make([][]damage.Damage, 1)
					missileDamageValues = append(missileDamageValues, tempArray...)
					// fmt.Printf("missileDamageValues after tempArray: %v\n", missileDamageValues)
				}
				// fmt.Printf("missileDmg[%d]: %v\n", i, missileDmg[i])
				// fmt.Printf("j: %d\n", j)
				// fmt.Printf("damageValuesIndexOffset: %d\n", damageValuesIndexOffset)
				// fmt.Printf("missileDamageValues: %v\n", missileDamageValues)
				// fmt.Printf("missileDmg: %v\n", missileDmg)
				// missileDamageValues[damageValuesIndexOffset+j] = append(missileDamageValues[damageValuesIndexOffset+j], missileDmg[j]...)
				missileDamageValues[damageValuesIndexOffset+j] =
					append(missileDamageValues[damageValuesIndexOffset+j],
						missileDmg[j])
				// fmt.Printf("missileDamageValues: %v\n", missileDamageValues)
			}
		}

	}

	//srvmissileb
	if skillRecord[skills.SrvMissileB] != "" {
		// fmt.Println("srvmissileb")
		for i := 1; i <= leveloffset; i++ {
			missileDmg := calculateMissileDamage(skillRecord[skills.SrvMissileB], int(i+(startlevel-1)), skill)
			// fmt.Println(missileDmg)

			if i == 1 {
				damageValuesIndexOffset = len(missileDamageValues)
			}

			for j := 0; j < len(missileDmg); j++ {
				// on the first iteration, add new row(s)
				if i == 1 {
					tempArray := make([][]damage.Damage, 1)
					missileDamageValues = append(missileDamageValues, tempArray...)
				}
				// fmt.Printf("missileDmg[%d]: %v\n", i, missileDmg[i])
				// missileDamageValues[j] = append(missileDamageValues[j], missileDmg[j]...)
				// missileDamageValues[damageValuesIndexOffset+j] = append(missileDamageValues[damageValuesIndexOffset+j], missileDmg[j]...)
				missileDamageValues[damageValuesIndexOffset+j] =
					append(missileDamageValues[damageValuesIndexOffset+j],
						missileDmg[j])
				// fmt.Printf("missileDamageValues: %v\n", missileDamageValues)
				// fmt.Println(missileDamageValues)
				// fmt.Printf("missileDamageValues: %v\n", missileDamageValues)
			}
		}
	}

	//srvmissilec
	if skillRecord[skills.SrvMissileC] != "" {
		// fmt.Println("srvmissilec")
		for i := 1; i <= leveloffset; i++ {
			missileDmg := calculateMissileDamage(skillRecord[skills.SrvMissileC], int(i+(startlevel-1)), skill)
			// fmt.Println(missileDmg)
			if i == 1 {
				damageValuesIndexOffset = len(missileDamageValues)
			}

			for j := 0; j < len(missileDmg); j++ {
				// on the first iteration, add new row(s)
				if i == 1 {
					tempArray := make([][]damage.Damage, 1)
					missileDamageValues = append(missileDamageValues, tempArray...)
				}
				// fmt.Printf("missileDmg[%d]: %v\n", i, missileDmg[i])
				// missileDamageValues[damageValuesIndexOffset+j] = append(missileDamageValues[damageValuesIndexOffset+j], missileDmg[j]...)
				missileDamageValues[damageValuesIndexOffset+j] =
					append(missileDamageValues[damageValuesIndexOffset+j],
						missileDmg[j])
				// fmt.Println(missileDamageValues)
				// fmt.Printf("missileDamageValues: %v\n", missileDamageValues)
			}

		}

	}
	// TODO: remove duplicates from the Missile damage array
	// missileDamageValues = common.RemoveDuplicateDamageRows(missileDamageValues)

	return missileDamageValues
}

func calculateDamage(sLvl int, hitShift float64, baseDmg float64, levDmg1 float64, levDmg2 float64, levDmg3 float64, levDmg4 float64, levDmg5 float64) float64 {
	var calcDmg float64
	calcDmg = baseDmg

	//get added damage sLvl 2-8
	if sLvl >= 2 && sLvl <= 8 {
		calcDmg = calcDmg + (levDmg1 * float64(sLvl-1))
	} else if sLvl > 8 {
		calcDmg = calcDmg + (levDmg1 * 7)
	}
	//get added damage sLvl 9-16
	if sLvl >= 9 && sLvl <= 16 {
		calcDmg = calcDmg + (levDmg2 * float64(sLvl-8))
	} else if sLvl > 16 {
		calcDmg = calcDmg + (levDmg2 * 8)
	}
	//get added damage sLvl 17-22
	if sLvl >= 17 && sLvl <= 22 {
		calcDmg = calcDmg + (levDmg3 * float64(sLvl-16))
	} else if sLvl > 22 {
		calcDmg = calcDmg + (levDmg3 * 6)
	}
	//get added damage sLvl 23-28
	if sLvl >= 23 && sLvl <= 28 {
		calcDmg = calcDmg + (levDmg4 * float64(sLvl-22))
	} else if sLvl > 28 {
		calcDmg = calcDmg + (levDmg4 * 6)
		//get added damage sLvl 29+
		calcDmg = calcDmg + (levDmg5 * float64(sLvl-28))
	}

	//add in hitshift operator
	effectiveShift := math.Exp2(float64(hitShift)) / 256

	calcDmg = calcDmg * effectiveShift

	//add in synergy damage??

	return calcDmg
}

func calcLength(sLvl int, baseLength float64, levLen1 float64, levLen2 float64, levLen3 float64) float64 {
	var calcLen float64
	calcLen = baseLength
	//get added length sLvl 2-8
	if sLvl >= 2 && sLvl <= 8 {
		calcLen = calcLen + (levLen1 * float64(sLvl-1))
	} else if sLvl > 8 {
		calcLen = calcLen + (levLen1 * 7)
	}
	//get added length sLvl 9-16
	if sLvl >= 8 && sLvl <= 16 {
		calcLen = calcLen + (levLen2 * float64(sLvl-8))
	} else if sLvl > 16 {
		calcLen = calcLen + (levLen2 * 8)
		//get added length sLvl 17+
		calcLen = calcLen + (levLen3 * float64(sLvl-16))
	}
	return calcLen
}

func calcManaCost(sLvl int, basemana float64, startmana float64, lvlmana float64, manashift float64, minmana float64) float64 {
	// manacost = max((mana+lvlmana*(sLvl-1))*effectiveshift,minmana);

	effectiveshift := math.Exp2(float64(manashift)) / 256

	calcmana := ((basemana + (lvlmana * float64(sLvl-1))) * effectiveshift) + startmana

	if calcmana < minmana {
		calcmana = minmana
	}

	return calcmana
}

func calculateMissileDamage(
	missileName string,
	sLvl int,
	skill Skill) []damage.Damage {
	/*
		Another shot at the logic for calculating missile damage
		If the skillrow contains anything in srvmissile*
			For each value of srvmissile*
				Calculate that missile damage (may need some hardcode stuff here, ex: firewallmaker vs firewall)
				May need to use the damage columns from the missiles.txt, otherwise inherit parent skills.txt damage columns
		Else
			Assume that the skill has no missile, but if it has damage columns, then calculate that damage

	*/

	/*
		NEW INFO::::

		Skill: If you enter the ID Pointer ('the name') of a skill here this missile will retrieve all of it's damage/hit data from this skill.
		The columns being read from skills.txt instead of missiles.txt are the following: ResultFlags, HitFlags, HitShift, HitClass,
		SrcDamage (SrcDam in skills.txt!), MinDam, MinLevDam1-5, MaxDam, MaxLevDam1-5, DmgSymPerCalc, EType, EMin, EMinLev1-5, EMax, EMaxLev1-5,
		EDmgSymPerCalc, ELen, ELenLev1-3, ELenSymPerCalc.
	*/

	// fmt.Println("Entered calculateMissileDamage")

	missileDamageSlice := make([]damage.Damage, 0)

	missileRecord := skill.getMissileRecord(missileName)

	if missileRecord[missiles.Skill] != "" {
		missileSkill := NewSkill(missileRecord[missiles.Skill], skill.skillFile, skill.missileFile, sLvl)

		// then, calculate the damage for this missile
		tmpSkillDamage := missileSkill.getSkillEleDamageValues()
		missileDamageSlice = append(missileDamageSlice, damage.Damage{missileSkill.name, tmpSkillDamage.Min, tmpSkillDamage.Max})

	} else if missileRecord[missiles.EMin] != "" {
		eMin, _ := strconv.ParseFloat(missileRecord[missiles.EMin], 64)
		eMinLev1, _ := strconv.ParseFloat(missileRecord[missiles.MinELev1], 64)
		eMinLev2, _ := strconv.ParseFloat(missileRecord[missiles.MinELev2], 64)
		eMinLev3, _ := strconv.ParseFloat(missileRecord[missiles.MinELev3], 64)
		eMinLev4, _ := strconv.ParseFloat(missileRecord[missiles.MinELev4], 64)
		eMinLev5, _ := strconv.ParseFloat(missileRecord[missiles.MinELev5], 64)

		eMax, _ := strconv.ParseFloat(missileRecord[missiles.Emax], 64)
		eMaxLev1, _ := strconv.ParseFloat(missileRecord[missiles.MaxELev1], 64)
		eMaxLev2, _ := strconv.ParseFloat(missileRecord[missiles.MaxELev2], 64)
		eMaxLev3, _ := strconv.ParseFloat(missileRecord[missiles.MaxELev3], 64)
		eMaxLev4, _ := strconv.ParseFloat(missileRecord[missiles.MaxELev4], 64)
		eMaxLev5, _ := strconv.ParseFloat(missileRecord[missiles.MaxELev5], 64)

		hitShift, _ := strconv.ParseFloat(missileRecord[missiles.HitShift], 64)

		eMissileDamageMin := calculateDamage(sLvl, hitShift, eMin, eMinLev1, eMinLev2, eMinLev3, eMinLev4, eMinLev5)
		eMissileDamageMax := calculateDamage(sLvl, hitShift, eMax, eMaxLev1, eMaxLev2, eMaxLev3, eMaxLev4, eMaxLev5)

		eLen, _ := strconv.ParseFloat(missileRecord[missiles.ELen], 64)
		eLevLen1, _ := strconv.ParseFloat(missileRecord[missiles.ELevLen1], 64)
		eLevLen2, _ := strconv.ParseFloat(missileRecord[missiles.ELevLen2], 64)
		eLevLen3, _ := strconv.ParseFloat(missileRecord[missiles.ELevLen3], 64)

		length := calcLength(sLvl, eLen, eLevLen1, eLevLen2, eLevLen3)

		eMissileDamageMin = calculateMissileFuncDamage(missileName, eMissileDamageMin, length)
		eMissileDamageMax = calculateMissileFuncDamage(missileName, eMissileDamageMax, length)

		missileDamageSlice = append(missileDamageSlice, damage.Damage{missileName, eMissileDamageMin, eMissileDamageMax})

	}

	//if there is a SubMissile of any type, calculate that and append
	//ExplosionMissile
	if missileRecord[missiles.ExplosionMissile] != "" {
		// fmt.Println("getting ExplosionMissile")
		missileDamageSlice = append(missileDamageSlice,
			calculateMissileDamage(missileRecord[missiles.ExplosionMissile],
				sLvl,
				skill)...)
	}

	//SubMissile1
	if missileRecord[missiles.SubMissile1] != "" {
		// fmt.Println("getting SubMissile1")
		missileDamageSlice = append(missileDamageSlice,
			calculateMissileDamage(missileRecord[missiles.SubMissile1],
				sLvl,
				skill)...)
	}

	//SubMissile2
	if missileRecord[missiles.SubMissile2] != "" {
		// fmt.Println("getting SubMissile2")
		missileDamageSlice = append(missileDamageSlice,
			calculateMissileDamage(missileRecord[missiles.SubMissile2],
				sLvl,
				skill)...)
	}

	//SubMissile3
	if missileRecord[missiles.SubMissile3] != "" {
		// fmt.Println("getting SubMissile3")
		missileDamageSlice = append(missileDamageSlice,
			calculateMissileDamage(missileRecord[missiles.SubMissile3],
				sLvl,
				skill)...)
	}

	//HitSubMissile1
	if missileRecord[missiles.HitSubMissile1] != "" {
		// fmt.Println("getting HitSubMissile1")
		missileDamageSlice = append(missileDamageSlice,
			calculateMissileDamage(missileRecord[missiles.HitSubMissile1],
				sLvl,
				skill)...)
	}

	//HitSubMissile2
	if missileRecord[missiles.HitSubMissile2] != "" {
		// fmt.Println("getting HitSubMissile2")
		missileDamageSlice = append(missileDamageSlice,
			calculateMissileDamage(missileRecord[missiles.HitSubMissile2],
				sLvl,
				skill)...)
	}

	//HitSubMissile3
	if missileRecord[missiles.HitSubMissile3] != "" {
		// fmt.Println("getting HitSubMissile3")
		missileDamageSlice = append(missileDamageSlice,
			calculateMissileDamage(missileRecord[missiles.HitSubMissile3],
				sLvl,
				skill)...)
	}

	//HitSubMissile4
	if missileRecord[missiles.HitSubMissile4] != "" {
		// fmt.Println("getting HitSubMissile4")
		missileDamageSlice = append(missileDamageSlice,
			calculateMissileDamage(missileRecord[missiles.HitSubMissile4],
				sLvl,
				skill)...)
	}

	return missileDamageSlice
}

func calculateMissileFuncDamage(missileFunc string, damage float64, length float64) float64 {
	var returnDmg float64
	// fmt.Printf("missileFunc: %s\n", missileFunc)
	switch missileFunc {
	case "firewall":
		returnDmg = damage * 25 * 3
	case "meteorfire":
		returnDmg = damage * 25 * 3
	case "firestormmaker":
		returnDmg = damage * 24 * 3
	case "arcticblast1":
		returnDmg = damage * 25
	case "poisonnova":
		returnDmg = damage * length
	case "poisonexplosioncloud":
		returnDmg = damage * length
	default:
		returnDmg = damage
	}

	return returnDmg

}

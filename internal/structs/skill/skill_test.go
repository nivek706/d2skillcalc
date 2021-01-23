package skill

import (
	"fmt"
	"math"
	"testing"

	"github.com/nivek706/d2skillcalc/configs"
	"github.com/nivek706/d2skillcalc/internal/structs/damage"
	"github.com/nivek706/d2skillcalc/pkg/fileutil"
)

func TestSorcColdSkills(t *testing.T) {
	c, err := configs.LoadConfig(".")
	if err != nil {
		fmt.Println("fatal")
	}
	skillFile := fileutil.ReadFile(fmt.Sprintf("%sSkills.txt", c.TxtDirPath))
	missileFile := fileutil.ReadFile(fmt.Sprintf("%sMissiles.txt", c.TxtDirPath))

	/* Ice Bolt */
	icebolt1 := NewSkill("Ice Bolt", skillFile, missileFile, 1)
	icebolt1.PopulateSkillDamage()
	compareDamageWithExpected(t, icebolt1.elementalDamage, 3, 5)
	compareDamageWithExpected(t, icebolt1.physicalDamage, 0.0, 0.0)
	iceboltmissile1, err := icebolt1.GetMissileDamageByName("Ice Bolt")
	compareDamageWithExpected(t, *iceboltmissile1, 3, 5)

	icebolt10 := NewSkill("Ice Bolt", skillFile, missileFile, 10)
	icebolt10.PopulateSkillDamage()
	compareDamageWithExpected(t, icebolt10.elementalDamage, 14, 20)
	compareDamageWithExpected(t, icebolt10.physicalDamage, 0.0, 0.0)
	iceboltmissile10, err := icebolt10.GetMissileDamageByName("Ice Bolt")
	compareDamageWithExpected(t, *iceboltmissile10, 14, 20)

	icebolt20 := NewSkill("Ice Bolt", skillFile, missileFile, 20)
	icebolt20.PopulateSkillDamage()
	compareDamageWithExpected(t, icebolt20.elementalDamage, 38, 49)
	compareDamageWithExpected(t, icebolt20.physicalDamage, 0.0, 0.0)
	iceboltmissile20, err := icebolt20.GetMissileDamageByName("Ice Bolt")
	compareDamageWithExpected(t, *iceboltmissile20, 38, 49)

	/* Frozen Armor */

	/* Frost Nova */
	frostnova1 := NewSkill("Frost Nova", skillFile, missileFile, 1)
	frostnova1.PopulateSkillDamage()
	compareDamageWithExpected(t, frostnova1.elementalDamage, 2, 4)
	compareDamageWithExpected(t, frostnova1.physicalDamage, 0.0, 0.0)
	frostnovamissile1, err := frostnova1.GetMissileDamageByName("Frost Nova")
	compareDamageWithExpected(t, *frostnovamissile1, 2, 4)

	frostnova10 := NewSkill("Frost Nova", skillFile, missileFile, 10)
	frostnova10.PopulateSkillDamage()
	compareDamageWithExpected(t, frostnova10.elementalDamage, 22, 28)
	compareDamageWithExpected(t, frostnova10.physicalDamage, 0.0, 0.0)
	frostnovamissile10, err := frostnova10.GetMissileDamageByName("Frost Nova")
	compareDamageWithExpected(t, *frostnovamissile10, 22, 28)

	frostnova20 := NewSkill("Frost Nova", skillFile, missileFile, 20)
	frostnova20.PopulateSkillDamage()
	compareDamageWithExpected(t, frostnova20.elementalDamage, 56, 67)
	compareDamageWithExpected(t, frostnova20.physicalDamage, 0.0, 0.0)
	frostnovamissile20, err := frostnova20.GetMissileDamageByName("Frost Nova")
	compareDamageWithExpected(t, *frostnovamissile20, 56, 67)

	/* Ice Blast */
	iceblast1 := NewSkill("Ice Blast", skillFile, missileFile, 1)
	iceblast1.PopulateSkillDamage()
	compareDamageWithExpected(t, iceblast1.elementalDamage, 8, 12)
	compareDamageWithExpected(t, iceblast1.physicalDamage, 0.0, 0.0)
	iceblastmissile1, err := iceblast1.GetMissileDamageByName("Ice Blast")
	compareDamageWithExpected(t, *iceblastmissile1, 8, 12)

	iceblast10 := NewSkill("Ice Blast", skillFile, missileFile, 10)
	iceblast10.PopulateSkillDamage()
	compareDamageWithExpected(t, iceblast10.elementalDamage, 85, 93)
	compareDamageWithExpected(t, iceblast10.physicalDamage, 0.0, 0.0)
	iceblastmissile10, err := iceblast10.GetMissileDamageByName("Ice Blast")
	compareDamageWithExpected(t, *iceblastmissile10, 85, 93)

	iceblast20 := NewSkill("Ice Blast", skillFile, missileFile, 20)
	iceblast20.PopulateSkillDamage()
	compareDamageWithExpected(t, iceblast20.elementalDamage, 253, 266)
	compareDamageWithExpected(t, iceblast20.physicalDamage, 0.0, 0.0)
	iceblastmissile20, err := iceblast20.GetMissileDamageByName("Ice Blast")
	compareDamageWithExpected(t, *iceblastmissile20, 253, 266)

	/* Shiver Armor */

	/* Glacial Spike */
	glacialspike1 := NewSkill("Glacial Spike", skillFile, missileFile, 1)
	glacialspike1.PopulateSkillDamage()
	compareDamageWithExpected(t, glacialspike1.elementalDamage, 16, 24)
	compareDamageWithExpected(t, glacialspike1.physicalDamage, 0.0, 0.0)
	glacialspikemissile1, err := glacialspike1.GetMissileDamageByName("Glacial Spike")
	compareDamageWithExpected(t, *glacialspikemissile1, 16, 24)

	glacialspike10 := NewSkill("Glacial Spike", skillFile, missileFile, 10)
	glacialspike10.PopulateSkillDamage()
	compareDamageWithExpected(t, glacialspike10.elementalDamage, 91, 103)
	compareDamageWithExpected(t, glacialspike10.physicalDamage, 0.0, 0.0)
	glacialspikemissile10, err := glacialspike10.GetMissileDamageByName("Glacial Spike")
	compareDamageWithExpected(t, *glacialspikemissile10, 91, 103)

	glacialspike20 := NewSkill("Glacial Spike", skillFile, missileFile, 20)
	glacialspike20.PopulateSkillDamage()
	compareDamageWithExpected(t, glacialspike20.elementalDamage, 225, 242)
	compareDamageWithExpected(t, glacialspike20.physicalDamage, 0.0, 0.0)
	glacialspikemissile20, err := glacialspike20.GetMissileDamageByName("Glacial Spike")
	compareDamageWithExpected(t, *glacialspikemissile20, 225, 242)

	/* Blizzard */
	blizzard1 := NewSkill("Blizzard", skillFile, missileFile, 1)
	blizzard1.PopulateSkillDamage()
	compareDamageWithExpected(t, blizzard1.elementalDamage, 45, 75)
	compareDamageWithExpected(t, blizzard1.physicalDamage, 0.0, 0.0)
	blizzardmissile1, err := blizzard1.GetMissileDamageByName("Blizzard")
	compareDamageWithExpected(t, *blizzardmissile1, 45, 75)

	blizzard10 := NewSkill("Blizzard", skillFile, missileFile, 10)
	blizzard10.PopulateSkillDamage()
	compareDamageWithExpected(t, blizzard10.elementalDamage, 210, 249)
	compareDamageWithExpected(t, blizzard10.physicalDamage, 0.0, 0.0)
	blizzardmissile10, err := blizzard10.GetMissileDamageByName("Blizzard")
	compareDamageWithExpected(t, *blizzardmissile10, 210, 249)

	blizzard20 := NewSkill("Blizzard", skillFile, missileFile, 20)
	blizzard20.PopulateSkillDamage()
	compareDamageWithExpected(t, blizzard20.elementalDamage, 570, 619)
	compareDamageWithExpected(t, blizzard20.physicalDamage, 0.0, 0.0)
	blizzardmissile20, err := blizzard20.GetMissileDamageByName("Blizzard")
	compareDamageWithExpected(t, *blizzardmissile20, 570, 619)

	/* Chilling Armor */

	/* Frozen Orb */
	frozenorb1 := NewSkill("Frozen Orb", skillFile, missileFile, 1)
	frozenorb1.PopulateSkillDamage()
	compareDamageWithExpected(t, frozenorb1.elementalDamage, 40, 45)
	compareDamageWithExpected(t, frozenorb1.physicalDamage, 0.0, 0.0)
	frozenorbmissile1, err := frozenorb1.GetMissileDamageByName("Frozen Orb")
	compareDamageWithExpected(t, *frozenorbmissile1, 40, 45)

	frozenorb10 := NewSkill("Frozen Orb", skillFile, missileFile, 10)
	frozenorb10.PopulateSkillDamage()
	compareDamageWithExpected(t, frozenorb10.elementalDamage, 134, 143)
	compareDamageWithExpected(t, frozenorb10.physicalDamage, 0.0, 0.0)
	frozenorbmissile10, err := frozenorb10.GetMissileDamageByName("Frozen Orb")
	compareDamageWithExpected(t, *frozenorbmissile10, 134, 143)

	frozenorb20 := NewSkill("Frozen Orb", skillFile, missileFile, 20)
	frozenorb20.PopulateSkillDamage()
	compareDamageWithExpected(t, frozenorb20.elementalDamage, 262, 276)
	compareDamageWithExpected(t, frozenorb20.physicalDamage, 0.0, 0.0)
	frozenorbmissile20, err := frozenorb20.GetMissileDamageByName("Frozen Orb")
	compareDamageWithExpected(t, *frozenorbmissile20, 262, 276)

	/* Cold Mastery */

}

func TestSorcLightningSkills(t *testing.T) {
	c, err := configs.LoadConfig(".")
	if err != nil {
		fmt.Println("fatal")
	}
	skillFile := fileutil.ReadFile(fmt.Sprintf("%sSkills.txt", c.TxtDirPath))
	missileFile := fileutil.ReadFile(fmt.Sprintf("%sMissiles.txt", c.TxtDirPath))

	/* Charged Bolt */
	chargedbolt1 := NewSkill("Charged Bolt", skillFile, missileFile, 1)
	chargedbolt1.PopulateSkillDamage()
	compareDamageWithExpected(t, chargedbolt1.elementalDamage, 2, 4)
	compareDamageWithExpected(t, chargedbolt1.physicalDamage, 0.0, 0.0)
	chargedboltmissile1, err := chargedbolt1.GetMissileDamageByName("Charged Bolt")
	compareDamageWithExpected(t, *chargedboltmissile1, 2, 4)

	chargedbolt10 := NewSkill("Charged Bolt", skillFile, missileFile, 10)
	chargedbolt10.PopulateSkillDamage()
	compareDamageWithExpected(t, chargedbolt10.elementalDamage, 6, 8)
	compareDamageWithExpected(t, chargedbolt10.physicalDamage, 0.0, 0.0)
	chargedboltmissile10, err := chargedbolt10.GetMissileDamageByName("Charged Bolt")
	compareDamageWithExpected(t, *chargedboltmissile10, 6, 8)

	chargedbolt20 := NewSkill("Charged Bolt", skillFile, missileFile, 20)
	chargedbolt20.PopulateSkillDamage()
	compareDamageWithExpected(t, chargedbolt20.elementalDamage, 13, 15)
	compareDamageWithExpected(t, chargedbolt20.physicalDamage, 0.0, 0.0)
	chargedboltmissile20, err := chargedbolt20.GetMissileDamageByName("Charged Bolt")
	compareDamageWithExpected(t, *chargedboltmissile20, 13, 15)

	/* Telekinesis */
	telekinesis1 := NewSkill("Telekinesis", skillFile, missileFile, 1)
	telekinesis1.PopulateSkillDamage()
	compareDamageWithExpected(t, telekinesis1.elementalDamage, 1, 2)
	compareDamageWithExpected(t, telekinesis1.physicalDamage, 0.0, 0.0)

	telekinesis10 := NewSkill("Telekinesis", skillFile, missileFile, 10)
	telekinesis10.PopulateSkillDamage()
	compareDamageWithExpected(t, telekinesis10.elementalDamage, 10, 11)
	compareDamageWithExpected(t, telekinesis10.physicalDamage, 0.0, 0.0)

	telekinesis20 := NewSkill("Telekinesis", skillFile, missileFile, 20)
	telekinesis20.PopulateSkillDamage()
	compareDamageWithExpected(t, telekinesis20.elementalDamage, 20, 21)
	compareDamageWithExpected(t, telekinesis20.physicalDamage, 0.0, 0.0)

	/* Static Field */

	/* Lightning */
	lightning1 := NewSkill("Lightning", skillFile, missileFile, 1)
	lightning1.PopulateSkillDamage()
	compareDamageWithExpected(t, lightning1.elementalDamage, 1, 40)
	compareDamageWithExpected(t, lightning1.physicalDamage, 0.0, 0.0)
	lightningmissile1, err := lightning1.GetMissileDamageByName("Lightning")
	compareDamageWithExpected(t, *lightningmissile1, 1, 40)

	lightning10 := NewSkill("Lightning", skillFile, missileFile, 10)
	lightning10.PopulateSkillDamage()
	compareDamageWithExpected(t, lightning10.elementalDamage, 1, 120)
	compareDamageWithExpected(t, lightning10.physicalDamage, 0.0, 0.0)
	lightningmissile10, err := lightning10.GetMissileDamageByName("Lightning")
	compareDamageWithExpected(t, *lightningmissile10, 1, 120)

	lightning20 := NewSkill("Lightning", skillFile, missileFile, 20)
	lightning20.PopulateSkillDamage()
	compareDamageWithExpected(t, lightning20.elementalDamage, 1, 272)
	compareDamageWithExpected(t, lightning20.physicalDamage, 0.0, 0.0)
	lightningmissile20, err := lightning20.GetMissileDamageByName("Lightning")
	compareDamageWithExpected(t, *lightningmissile20, 1, 272)

	/* Nova */
	nova1 := NewSkill("Nova", skillFile, missileFile, 1)
	nova1.PopulateSkillDamage()
	compareDamageWithExpected(t, nova1.elementalDamage, 1, 20)
	compareDamageWithExpected(t, nova1.physicalDamage, 0.0, 0.0)
	novamissile1, err := nova1.GetMissileDamageByName("Nova")
	compareDamageWithExpected(t, *novamissile1, 1, 20)

	nova10 := NewSkill("Nova", skillFile, missileFile, 10)
	nova10.PopulateSkillDamage()
	compareDamageWithExpected(t, nova10.elementalDamage, 57, 94)
	compareDamageWithExpected(t, nova10.physicalDamage, 0.0, 0.0)
	novamissile10, err := nova10.GetMissileDamageByName("Nova")
	compareDamageWithExpected(t, *novamissile10, 57, 94)

	nova20 := NewSkill("Nova", skillFile, missileFile, 20)
	nova20.PopulateSkillDamage()
	compareDamageWithExpected(t, nova20.elementalDamage, 131, 188)
	compareDamageWithExpected(t, nova20.physicalDamage, 0.0, 0.0)
	novamissile20, err := nova20.GetMissileDamageByName("Nova")
	compareDamageWithExpected(t, *novamissile20, 131, 188)

	/* Chain Lightning */
	chainlightning1 := NewSkill("Chain Lightning", skillFile, missileFile, 1)
	chainlightning1.PopulateSkillDamage()
	compareDamageWithExpected(t, chainlightning1.elementalDamage, 1, 40)
	compareDamageWithExpected(t, chainlightning1.physicalDamage, 0.0, 0.0)
	chainlightningmissile1, err := chainlightning1.GetMissileDamageByName("Chain Lightning")
	compareDamageWithExpected(t, *chainlightningmissile1, 1, 40)

	chainlightning10 := NewSkill("Chain Lightning", skillFile, missileFile, 10)
	chainlightning10.PopulateSkillDamage()
	compareDamageWithExpected(t, chainlightning10.elementalDamage, 1, 143)
	compareDamageWithExpected(t, chainlightning10.physicalDamage, 0.0, 0.0)
	chainlightningmissile10, err := chainlightning10.GetMissileDamageByName("Chain Lightning")
	compareDamageWithExpected(t, *chainlightningmissile10, 1, 143)

	chainlightning20 := NewSkill("Chain Lightning", skillFile, missileFile, 20)
	chainlightning20.PopulateSkillDamage()
	compareDamageWithExpected(t, chainlightning20.elementalDamage, 1, 281)
	compareDamageWithExpected(t, chainlightning20.physicalDamage, 0.0, 0.0)
	chainlightningmissile20, err := chainlightning20.GetMissileDamageByName("Chain Lightning")
	compareDamageWithExpected(t, *chainlightningmissile20, 1, 281)

	/* Teleport */

	/* Thunder Storm */
	thunderstorm1 := NewSkill("Thunder Storm", skillFile, missileFile, 1)
	thunderstorm1.PopulateSkillDamage()
	compareDamageWithExpected(t, thunderstorm1.elementalDamage, 1, 100)
	compareDamageWithExpected(t, thunderstorm1.physicalDamage, 0.0, 0.0)
	thunderstormmissile1, err := thunderstorm1.GetMissileDamageByName("Thunder Storm")
	compareDamageWithExpected(t, *thunderstormmissile1, 1, 100)

	thunderstorm10 := NewSkill("Thunder Storm", skillFile, missileFile, 10)
	thunderstorm10.PopulateSkillDamage()
	compareDamageWithExpected(t, thunderstorm10.elementalDamage, 91, 190)
	compareDamageWithExpected(t, thunderstorm10.physicalDamage, 0.0, 0.0)
	thunderstormmissile10, err := thunderstorm10.GetMissileDamageByName("Thunder Storm")
	compareDamageWithExpected(t, *thunderstormmissile10, 91, 190)

	thunderstorm20 := NewSkill("Thunder Storm", skillFile, missileFile, 20)
	thunderstorm20.PopulateSkillDamage()
	compareDamageWithExpected(t, thunderstorm20.elementalDamage, 195, 294)
	compareDamageWithExpected(t, thunderstorm20.physicalDamage, 0.0, 0.0)
	thunderstormmissile20, err := thunderstorm20.GetMissileDamageByName("Thunder Storm")
	compareDamageWithExpected(t, *thunderstormmissile20, 195, 294)

	/* Energy Shield */

	/* Lightning Mastery */

}

func TestSorcFireSkills(t *testing.T) {
	c, err := configs.LoadConfig(".")
	if err != nil {
		fmt.Println("fatal")
	}
	skillFile := fileutil.ReadFile(fmt.Sprintf("%sSkills.txt", c.TxtDirPath))
	missileFile := fileutil.ReadFile(fmt.Sprintf("%sMissiles.txt", c.TxtDirPath))

	/* Fire Bolt */
	firebolt1 := NewSkill("Fire Bolt", skillFile, missileFile, 1)
	firebolt1.PopulateSkillDamage()
	compareDamageWithExpected(t, firebolt1.elementalDamage, 3, 6)
	compareDamageWithExpected(t, firebolt1.physicalDamage, 0.0, 0.0)
	fireboltmissile1, err := firebolt1.GetMissileDamageByName("Fire Bolt")
	// I don't like(or fully understand) the pointer business here - try to learn if there is a better way to do all of this
	compareDamageWithExpected(t, *fireboltmissile1, 3, 6)

	firebolt10 := NewSkill("Fire Bolt", skillFile, missileFile, 10)
	firebolt10.PopulateSkillDamage()
	compareDamageWithExpected(t, firebolt10.elementalDamage, 17, 22)
	compareDamageWithExpected(t, firebolt10.physicalDamage, 0.0, 0.0)
	fireboltmissile10, err := firebolt10.GetMissileDamageByName("Fire Bolt")
	compareDamageWithExpected(t, *fireboltmissile10, 17, 22)

	firebolt20 := NewSkill("Fire Bolt", skillFile, missileFile, 20)
	firebolt20.PopulateSkillDamage()
	compareDamageWithExpected(t, firebolt20.elementalDamage, 45, 60)
	compareDamageWithExpected(t, firebolt20.physicalDamage, 0.0, 0.0)
	fireboltmissile20, err := firebolt20.GetMissileDamageByName("Fire Bolt")
	compareDamageWithExpected(t, *fireboltmissile20, 45, 60)

	/* Warmth */

	/* Inferno */
	inferno1 := NewSkill("Inferno", skillFile, missileFile, 1)
	inferno1.PopulateSkillDamage()
	compareDamageWithExpected(t, inferno1.elementalDamage, 12, 25)
	compareDamageWithExpected(t, inferno1.physicalDamage, 0.0, 0.0)
	infernomissile1, err := inferno1.GetMissileDamageByName("Inferno")
	compareDamageWithExpected(t, *infernomissile1, 12, 25)

	inferno10 := NewSkill("Inferno", skillFile, missileFile, 10)
	inferno10.PopulateSkillDamage()
	compareDamageWithExpected(t, inferno10.elementalDamage, 98, 111)
	compareDamageWithExpected(t, inferno10.physicalDamage, 0.0, 0.0)
	infernomissile10, err := inferno10.GetMissileDamageByName("Inferno")
	compareDamageWithExpected(t, *infernomissile10, 98, 111)

	inferno20 := NewSkill("Inferno", skillFile, missileFile, 20)
	inferno20.PopulateSkillDamage()
	compareDamageWithExpected(t, inferno20.elementalDamage, 203, 220)
	compareDamageWithExpected(t, inferno20.physicalDamage, 0.0, 0.0)
	infernomissile20, err := inferno20.GetMissileDamageByName("Inferno")
	compareDamageWithExpected(t, *infernomissile20, 203, 220)

	/* Blaze */
	blaze1 := NewSkill("Blaze", skillFile, missileFile, 1)
	blaze1.PopulateSkillDamage()
	compareDamageWithExpected(t, blaze1.elementalDamage, 18, 37)
	compareDamageWithExpected(t, blaze1.physicalDamage, 0.0, 0.0)
	blazemissile1, err := blaze1.GetMissileDamageByName("Blaze")
	compareDamageWithExpected(t, *blazemissile1, 18, 37)

	blaze10 := NewSkill("Blaze", skillFile, missileFile, 10)
	blaze10.PopulateSkillDamage()
	compareDamageWithExpected(t, blaze10.elementalDamage, 112, 131)
	compareDamageWithExpected(t, blaze10.physicalDamage, 0.0, 0.0)
	blazemissile10, err := blaze10.GetMissileDamageByName("Blaze")
	compareDamageWithExpected(t, *blazemissile10, 112, 131)

	blaze20 := NewSkill("Blaze", skillFile, missileFile, 20)
	blaze20.PopulateSkillDamage()
	compareDamageWithExpected(t, blaze20.elementalDamage, 271, 290)
	compareDamageWithExpected(t, blaze20.physicalDamage, 0.0, 0.0)
	blazemissile20, err := blaze20.GetMissileDamageByName("Blaze")
	compareDamageWithExpected(t, *blazemissile20, 271, 290)

	/* Fireball */
	fireball1 := NewSkill("Fire Ball", skillFile, missileFile, 1)
	fireball1.PopulateSkillDamage()
	compareDamageWithExpected(t, fireball1.elementalDamage, 6, 14)
	compareDamageWithExpected(t, fireball1.physicalDamage, 0.0, 0.0)
	fireballmissile1, err := fireball1.GetMissileDamageByName("Fire Ball")
	compareDamageWithExpected(t, *fireballmissile1, 6, 14)

	fireball10 := NewSkill("Fire Ball", skillFile, missileFile, 10)
	fireball10.PopulateSkillDamage()
	compareDamageWithExpected(t, fireball10.elementalDamage, 74, 91)
	compareDamageWithExpected(t, fireball10.physicalDamage, 0.0, 0.0)
	fireballmissile10, err := fireball10.GetMissileDamageByName("Fire Ball")
	compareDamageWithExpected(t, *fireballmissile10, 74, 91)

	fireball20 := NewSkill("Fire Ball", skillFile, missileFile, 20)
	fireball20.PopulateSkillDamage()
	compareDamageWithExpected(t, fireball20.elementalDamage, 199, 226)
	compareDamageWithExpected(t, fireball20.physicalDamage, 0.0, 0.0)
	fireballmissile20, err := fireball20.GetMissileDamageByName("Fire Ball")
	compareDamageWithExpected(t, *fireballmissile20, 199, 226)

	/* Fire Wall */
	firewall1 := NewSkill("Fire Wall", skillFile, missileFile, 1)
	firewall1.PopulateSkillDamage()
	compareDamageWithExpected(t, firewall1.elementalDamage, 70, 93)
	compareDamageWithExpected(t, firewall1.physicalDamage, 0.0, 0.0)
	firewallmissile1, err := firewall1.GetMissileDamageByName("Fire Wall")
	compareDamageWithExpected(t, *firewallmissile1, 70, 93)

	firewall10 := NewSkill("Fire Wall", skillFile, missileFile, 10)
	firewall10.PopulateSkillDamage()
	compareDamageWithExpected(t, firewall10.elementalDamage, 496, 520)
	compareDamageWithExpected(t, firewall10.physicalDamage, 0.0, 0.0)
	firewallmissile10, err := firewall10.GetMissileDamageByName("Fire Wall")
	compareDamageWithExpected(t, *firewallmissile10, 496, 520)

	firewall20 := NewSkill("Fire Wall", skillFile, missileFile, 20)
	firewall20.PopulateSkillDamage()
	compareDamageWithExpected(t, firewall20.elementalDamage, 1284, 1307)
	compareDamageWithExpected(t, firewall20.physicalDamage, 0.0, 0.0)
	firewallmissile20, err := firewall20.GetMissileDamageByName("Fire Wall")
	compareDamageWithExpected(t, *firewallmissile20, 1284, 1307)

	/* Enchant */

	/* Meteor */
	meteor1 := NewSkill("Meteor", skillFile, missileFile, 1)
	meteor1.PopulateSkillDamage()
	compareDamageWithExpected(t, meteor1.elementalDamage, 80.0, 100.0)
	compareDamageWithExpected(t, meteor1.physicalDamage, 0.0, 0.0)
	meteorfire1, err := meteor1.GetMissileDamageByName("meteorfire")
	compareDamageWithExpected(t, *meteorfire1, 35.0, 58.0)

	meteor10 := NewSkill("Meteor", skillFile, missileFile, 10)
	meteor10.PopulateSkillDamage()
	compareDamageWithExpected(t, meteor10.elementalDamage, 319.0, 357.0)
	compareDamageWithExpected(t, meteor10.physicalDamage, 0.0, 0.0)
	meteorfire10, _ := meteor10.GetMissileDamageByName("meteorfire")
	compareDamageWithExpected(t, *meteorfire10, 124.0, 147.0)

	meteor20 := NewSkill("Meteor", skillFile, missileFile, 20)
	meteor20.PopulateSkillDamage()
	compareDamageWithExpected(t, meteor20.elementalDamage, 869.0, 927.0)
	compareDamageWithExpected(t, meteor20.physicalDamage, 0.0, 0.0)
	meteorfire20, _ := meteor20.GetMissileDamageByName("meteorfire")
	compareDamageWithExpected(t, *meteorfire20, 250.0, 274.0)

	/* Fire Mastery */

	/* Hydra */
	hydra1 := NewSkill("Hydra", skillFile, missileFile, 1)
	hydra1.PopulateSkillDamage()
	compareDamageWithExpected(t, hydra1.elementalDamage, 14, 19)
	compareDamageWithExpected(t, hydra1.physicalDamage, 0.0, 0.0)

	hydra10 := NewSkill("Hydra", skillFile, missileFile, 10)
	hydra10.PopulateSkillDamage()
	compareDamageWithExpected(t, hydra10.elementalDamage, 67, 82)
	compareDamageWithExpected(t, hydra10.physicalDamage, 0.0, 0.0)

	hydra20 := NewSkill("Hydra", skillFile, missileFile, 20)
	hydra20.PopulateSkillDamage()
	compareDamageWithExpected(t, hydra20.elementalDamage, 150, 175)
	compareDamageWithExpected(t, hydra20.physicalDamage, 0.0, 0.0)
}

func fuzzyFloatEquals(f1 float64, f2 float64) bool {
	// this function might be needed to do comparison due to the rounding on Arreat Summit
	// (ex: 3.5 might be rounded to 3, or 4.125 rounded to 4.1, etc)
	if f1 == f2 {
		// if they're exactly the same, return true
		return true
	}

	tolerance := 1.0
	if math.Abs(f2-f1) < tolerance {
		return true
	}

	return false
}

func compareDamageWithExpected(t *testing.T, dmg damage.Damage, expectedMinDmg float64, expectedMaxDmg float64) {

	if !fuzzyFloatEquals(dmg.Min, expectedMinDmg) {
		t.Log(fmt.Sprintf("%s min damage expected %f, got: ", dmg.Name, expectedMinDmg), dmg.Min)
		t.Fail()
	}

	if !fuzzyFloatEquals(dmg.Max, expectedMaxDmg) {
		t.Log(fmt.Sprintf("%s max damage expected %f, got: ", dmg.Name, expectedMaxDmg), dmg.Max)
		t.Fail()
	}

}

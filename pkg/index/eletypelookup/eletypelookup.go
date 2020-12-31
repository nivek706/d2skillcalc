package eletypelookup

//EType - Lookup structure for translating EType in Skills.txt
var EType = map[string]string{
	"fire": "Fire",
	"cold": "Cold",
	"ltng": "Lightning",
	"pois": "Poison",
	"mag":  "Magic",
	"stun": "Stun",
	"phys": "Physical"}

func GetType(dmgType string) string {
	switch dmgType {
	case "fire":
		return "Fire"
	case "cold":
		return "Cold"
	case "ltng":
		return "Lightning"
	case "pois":
		return "Poison"
	case "mag":
		return "Magic"
	case "stun":
		return "Stun"
	case "phys":
		return "Physical"
	default:
		return "Physical"
	}
}

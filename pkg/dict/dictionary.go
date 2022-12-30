package dict

var UnitMap = map[string]string{
	"0": "zero",
	"1": "jeden",
	"2": "dwa",
	"3": "trzy",
	"4": "cztery",
	"5": "pięć",
	"6": "sześć",
	"7": "siedem",
	"8": "osiem",
	"9": "dziewięć",
}

var TenthsMap = map[string]string{
	"10": "dziesięć",
	"11": "jedenaście",
	"12": "dwanaście",
	"13": "trzynaście",
	"14": "czternaście",
	"15": "piętnaście",
	"16": "szesnaście",
	"17": "siedemnaście",
	"18": "osiemnaście",
	"19": "dziewiętnaście",
}

var UpperTenthsMap = map[string]string{
	"2": "dwadzieścia",
	"3": "trzydzieści",
	"4": "czterdzieści",
	"5": "pięćdziesiąt",
	"6": "sześćdziesiąt",
	"7": "siedemdziesiąt",
	"8": "osiemdziesiąt",
	"9": "dziewięćdziesiąt",
}

var HundredsMap = map[string]string{
	"1": "sto",
	"2": "dwieście",
	"3": "trzysta",
	"4": "czterysta",
	"5": "pięćset",
	"6": "sześćset",
	"7": "siedemset",
	"8": "osiemset",
	"9": "dziewięćset",
}

var Singular = []string{" ", "tysiąc ", "milion ", "miliard ", "bilion ", "biliard ", "trylion ", "tryliard ", "kwadrylion ", "kwadryliard ", "kwintylion ", "kwintyliard "}
var PluralSimple = []string{" ", "tysiące ", "miliony ", "miliardy ", "biliony ", "biliardy ", "tryliony ", "tryliardy ", "kwadryliony ", "kwadryliardy ", "kwintyliony ", "kwintyliardy "}
var PluralUpper = []string{" ", "tysięcy ", "milionów ", "miliardów ", "bilionów ", "biliardów ", "trylionów ", "tryliardów ", "kwadrylionów ", "kwadryliardów ", "kwintylionów ", "kwintyliardów "}

var Relation = map[string][]string{
	"jeden ":  Singular,
	"dwa ":    PluralSimple,
	"trzy ":   PluralSimple,
	"cztery ": PluralSimple,
}

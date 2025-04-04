package lang

var Dictionary = map[string]ProtoWord{
	"birch": {
		atr: true,
		syllables: []ProtoSyllable{
			{Consonant{}, Front, Back, 0},
			{V, Central, 0, 0},
		},
	},
	"eat": {
		atr: false,
		syllables: []ProtoSyllable{
			{Consonant{}, Front, 0, 0},
			{T, Central, 0, 0},
		},
	},
}

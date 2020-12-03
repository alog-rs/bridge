package mocks

import rs3pb "github.com/alog-rs/proto/rs3"

var combatLevel int32 = 138
var questInfo = &rs3pb.QuestData{
	Completed:  116,
	Started:    10,
	NotStarted: 170, // lol noob
}
var skills = []*rs3pb.SkillData{
	{
		Skill:        rs3pb.Skill_ARCHAEOLOGY,
		Rank:         8103,
		Level:        120,
		VirtualLevel: 124,
		Xp:           156613628,
	},
	{
		Skill:        rs3pb.Skill_INVENTION,
		Rank:         36175,
		Level:        120,
		VirtualLevel: 133,
		Xp:           124540713,
	},
	{
		Skill:        rs3pb.Skill_FARMING,
		Rank:         53681,
		Level:        109,
		VirtualLevel: 109,
		Xp:           36577510,
	},
	{
		Skill:        rs3pb.Skill_CONSTITUTION,
		Rank:         134891,
		Level:        99,
		VirtualLevel: 107,
		Xp:           30706448,
	},
	{
		Skill:        rs3pb.Skill_STRENGTH,
		Rank:         77253,
		Level:        99,
		VirtualLevel: 105,
		Xp:           25693275,
	},
	{
		Skill:        rs3pb.Skill_ATTACK,
		Rank:         79314,
		Level:        99,
		VirtualLevel: 105,
		Xp:           25290491,
	},
	{
		Skill:        rs3pb.Skill_DEFENCE,
		Rank:         100769,
		Level:        99,
		VirtualLevel: 105,
		Xp:           25246413,
	},
	{
		Skill:        rs3pb.Skill_HERBLORE,
		Rank:         78877,
		Level:        102,
		VirtualLevel: 102,
		Xp:           19050147,
	},
	{
		Skill:        rs3pb.Skill_RANGED,
		Rank:         138600,
		Level:        99,
		VirtualLevel: 102,
		Xp:           17869777,
	},
	{
		Skill:        rs3pb.Skill_FIREMAKING,
		Rank:         72256,
		Level:        99,
		VirtualLevel: 101,
		Xp:           17021945,
	},
	{
		Skill:        rs3pb.Skill_SLAYER,
		Rank:         135028,
		Level:        101,
		VirtualLevel: 101,
		Xp:           16287051,
	},
	{
		Skill:        rs3pb.Skill_CRAFTING,
		Rank:         68467,
		Level:        99,
		VirtualLevel: 100,
		Xp:           15812588,
	},
	{
		Skill:        rs3pb.Skill_SMITHING,
		Rank:         79274,
		Level:        99,
		VirtualLevel: 100,
		Xp:           15352629,
	},
	{
		Skill:        rs3pb.Skill_PRAYER,
		Rank:         96495,
		Level:        99,
		VirtualLevel: 100,
		Xp:           14762610,
	},
	{
		Skill:        rs3pb.Skill_WOODCUTTING,
		Rank:         111023,
		Level:        99,
		VirtualLevel: 100,
		Xp:           14559209,
	},
	{
		Skill:        rs3pb.Skill_MAGIC,
		Rank:         193334,
		Level:        99,
		VirtualLevel: 100,
		Xp:           14503526,
	},
	{
		Skill:        rs3pb.Skill_AGILITY,
		Rank:         92535,
		Level:        99,
		VirtualLevel: 99,
		Xp:           14186053,
	},
	{
		Skill:        rs3pb.Skill_SUMMONING,
		Rank:         90819,
		Level:        99,
		VirtualLevel: 99,
		Xp:           14081303,
	},
	{
		Skill:        rs3pb.Skill_DIVINATION,
		Rank:         94631,
		Level:        99,
		VirtualLevel: 99,
		Xp:           13872498,
	},
	{
		Skill:        rs3pb.Skill_MINING,
		Rank:         159554,
		Level:        99,
		VirtualLevel: 99,
		Xp:           13470262,
	},
	{
		Skill:        rs3pb.Skill_FLETCHING,
		Rank:         128678,
		Level:        99,
		VirtualLevel: 99,
		Xp:           13470186,
	},
	{
		Skill:        rs3pb.Skill_RUNECRAFTING,
		Rank:         102418,
		Level:        99,
		VirtualLevel: 99,
		Xp:           13347283,
	},
	{
		Skill:        rs3pb.Skill_CONSTRUCTION,
		Rank:         114642,
		Level:        99,
		VirtualLevel: 99,
		Xp:           13340157,
	},
	{
		Skill:        rs3pb.Skill_FISHING,
		Rank:         165458,
		Level:        99,
		VirtualLevel: 99,
		Xp:           13289198,
	},
	{
		Skill:        rs3pb.Skill_THIEVING,
		Rank:         139467,
		Level:        99,
		VirtualLevel: 99,
		Xp:           13281980,
	},
	{
		Skill:        rs3pb.Skill_COOKING,
		Rank:         180059,
		Level:        99,
		VirtualLevel: 99,
		Xp:           13280626,
	},
	{
		Skill:        rs3pb.Skill_HUNTER,
		Rank:         144349,
		Level:        99,
		VirtualLevel: 99,
		Xp:           13046679,
	},
	{
		Skill:        rs3pb.Skill_DUNGEONEERING,
		Rank:         320425,
		Level:        79, // yikes
		VirtualLevel: 79,
		Xp:           1881689,
	},
}
var activity = []*rs3pb.PlayerActivityItem{
	{
		Timestamp: 1605558120,
		Long:      "I now have at least 124000000 experience points in the Invention skill.",
		Short:     "124000000XP in Invention",
	},
	{
		Timestamp: 1605515520,
		Long:      "I now have at least 156000000 experience points in the Archaeology skill.",
		Short:     "156000000XP in Archaeology",
	},
	{
		Timestamp: 1605514080,
		Long:      "I now have at least 154000000 experience points in the Archaeology skill.",
		Short:     "154000000XP in Archaeology",
	},
	{
		Timestamp: 1605513840,
		Long:      "I now have at least 152000000 experience points in the Archaeology skill.",
		Short:     "152000000XP in Archaeology",
	},
	{
		Timestamp: 1605513480,
		Long:      "I now have at least 150000000 experience points in the Archaeology skill.",
		Short:     "150000000XP in Archaeology",
	},
	{
		Timestamp: 1605347940,
		Long:      "I now have at least 122000000 experience points in the Invention skill.",
		Short:     "122000000XP in Invention",
	},
	{
		Timestamp: 1605342060,
		Long:      "I now have at least 148000000 experience points in the Archaeology skill.",
		Short:     "148000000XP in Archaeology",
	},
	{
		Timestamp: 1605339660,
		Long:      "I now have at least 36000000 experience points in the Farming skill.",
		Short:     "36000000XP in Farming",
	},
	{
		Timestamp: 1605259260,
		Long:      "I now have at least 146000000 experience points in the Archaeology skill.",
		Short:     "146000000XP in Archaeology",
	},
	{
		Timestamp: 1605258900,
		Long:      "I now have at least 120000000 experience points in the Invention skill.",
		Short:     "120000000XP in Invention",
	},
	{
		Timestamp: 1605228540,
		Long:      "I now have at least 144000000 experience points in the Archaeology skill.",
		Short:     "144000000XP in Archaeology",
	},
	{
		Timestamp: 1605225360,
		Long:      "I now have at least 118000000 experience points in the Invention skill.",
		Short:     "118000000XP in Invention",
	},
	{
		Timestamp: 1605213540,
		Long:      "I now have at least 142000000 experience points in the Archaeology skill.",
		Short:     "142000000XP in Archaeology",
	},
	{
		Timestamp: 1605173880,
		Long:      "I now have at least 116000000 experience points in the Invention skill.",
		Short:     "116000000XP in Invention",
	},
	{
		Timestamp: 1605172500,
		Long:      "I now have at least 140000000 experience points in the Archaeology skill.",
		Short:     "140000000XP in Archaeology",
	},
	{
		Timestamp: 1605172500,
		Long:      "I levelled my  Farming skill, I am now level 109.",
		Short:     "Levelled up Farming.",
	},
	{
		Timestamp: 1605160560,
		Long:      "I now have at least 114000000 experience points in the Invention skill.",
		Short:     "114000000XP in Invention",
	},
	{
		Timestamp: 1605160260,
		Long:      "I now have at least 138000000 experience points in the Archaeology skill.",
		Short:     "138000000XP in Archaeology",
	},
	{
		Timestamp: 1605088200,
		Long:      "I now have at least 136000000 experience points in the Archaeology skill.",
		Short:     "136000000XP in Archaeology",
	},
	{
		Timestamp: 1605086700,
		Long:      "I now have at least 112000000 experience points in the Invention skill.",
		Short:     "112000000XP in Invention",
	},
}

// RuneMetricsPlayerProfile mocks a valid PlayerProfile which includes all data returned from
// RuneMetrics
var RuneMetricsPlayerProfile = &rs3pb.PlayerProfile{
	Name:        "Uss",
	Rank:        61678,
	TotalLevel:  2809,
	TotalXp:     720435874,
	CombatLevel: &combatLevel,
	QuestInfo:   questInfo,
	Skills:      skills,
	Activity:    activity,
}

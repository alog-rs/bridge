package mocks

// RuneMetricsProfileResponseNotFound mocks a requested profile that was not found
var RuneMetricsProfileResponseNotFound = []byte(`
	{
		"error":"NO_PROFILE",
		"loggedIn":"false"
	}
`)

// RuneMetricsProfileResponsePrivate mocks a requested profile that is private
var RuneMetricsProfileResponsePrivate = []byte(`
	{
		"error": "PROFILE_PRIVATE",
		"loggedIn": "false"
	}
`)

// RuneMetricsProfileResponseSuccess mocks a successful profile request
var RuneMetricsProfileResponseSuccess = []byte(`
	{
		"magic":14762610,
		"questsstarted":10,
		"totalskill":2809,
		"questscomplete":116,
		"questsnotstarted":170,
		"totalxp":720435874,
		"ranged":30706448,
		"activities":[
			{
				"date":"16-Nov-2020 20:22",
				"details":"I now have at least 124000000 experience points in the Invention skill.",
				"text":"124000000XP in Invention"
			},
			{
				"date":"16-Nov-2020 08:32",
				"details":"I now have at least 156000000 experience points in the Archaeology skill.",
				"text":"156000000XP in Archaeology"
			},
			{
				"date":"16-Nov-2020 08:08",
				"details":"I now have at least 154000000 experience points in the Archaeology skill.",
				"text":"154000000XP in Archaeology"
			},
			{
				"date":"16-Nov-2020 08:04",
				"details":"I now have at least 152000000 experience points in the Archaeology skill.",
				"text":"152000000XP in Archaeology"
			},
			{
				"date":"16-Nov-2020 07:58",
				"details":"I now have at least 150000000 experience points in the Archaeology skill.",
				"text":"150000000XP in Archaeology"
			},
			{
				"date":"14-Nov-2020 09:59",
				"details":"I now have at least 122000000 experience points in the Invention skill.",
				"text":"122000000XP in Invention"
			},
			{
				"date":"14-Nov-2020 08:21",
				"details":"I now have at least 148000000 experience points in the Archaeology skill.",
				"text":"148000000XP in Archaeology"
			},
			{
				"date":"14-Nov-2020 07:41",
				"details":"I now have at least 36000000 experience points in the Farming skill.",
				"text":"36000000XP in Farming"
			},
			{
				"date":"13-Nov-2020 09:21",
				"details":"I now have at least 146000000 experience points in the Archaeology skill.",
				"text":"146000000XP in Archaeology"
			},
			{
				"date":"13-Nov-2020 09:15",
				"details":"I now have at least 120000000 experience points in the Invention skill.",
				"text":"120000000XP in Invention"
			},
			{
				"date":"13-Nov-2020 00:49",
				"details":"I now have at least 144000000 experience points in the Archaeology skill.",
				"text":"144000000XP in Archaeology"
			},
			{
				"date":"12-Nov-2020 23:56",
				"details":"I now have at least 118000000 experience points in the Invention skill.",
				"text":"118000000XP in Invention"
			},
			{
				"date":"12-Nov-2020 20:39",
				"details":"I now have at least 142000000 experience points in the Archaeology skill.",
				"text":"142000000XP in Archaeology"
			},
			{
				"date":"12-Nov-2020 09:38",
				"details":"I now have at least 116000000 experience points in the Invention skill.",
				"text":"116000000XP in Invention"
			},
			{
				"date":"12-Nov-2020 09:15",
				"details":"I now have at least 140000000 experience points in the Archaeology skill.",
				"text":"140000000XP in Archaeology"
			},
			{
				"date":"12-Nov-2020 09:15",
				"details":"I levelled my  Farming skill, I am now level 109.",
				"text":"Levelled up Farming."
			},
			{
				"date":"12-Nov-2020 05:56",
				"details":"I now have at least 114000000 experience points in the Invention skill.",
				"text":"114000000XP in Invention"
			},
			{
				"date":"12-Nov-2020 05:51",
				"details":"I now have at least 138000000 experience points in the Archaeology skill.",
				"text":"138000000XP in Archaeology"
			},
			{
				"date":"11-Nov-2020 09:50",
				"details":"I now have at least 136000000 experience points in the Archaeology skill.",
				"text":"136000000XP in Archaeology"
			},
			{
				"date":"11-Nov-2020 09:25",
				"details":"I now have at least 112000000 experience points in the Invention skill.",
				"text":"112000000XP in Invention"
			}
		],
		"skillvalues":[
			{
				"level":120,
				"xp":1566136286,
				"rank":8103,
				"id":27
			},
			{
				"level":120,
				"xp":1245407138,
				"rank":36175,
				"id":26
			},
			{
				"level":109,
				"xp":365775109,
				"rank":53681,
				"id":19
			},
			{
				"level":99,
				"xp":307064485,
				"rank":134891,
				"id":3
			},
			{
				"level":99,
				"xp":256932754,
				"rank":77253,
				"id":2
			},
			{
				"level":99,
				"xp":252904913,
				"rank":79314,
				"id":0
			},
			{
				"level":99,
				"xp":252464135,
				"rank":100769,
				"id":1
			},
			{
				"level":102,
				"xp":190501476,
				"rank":78877,
				"id":15
			},
			{
				"level":99,
				"xp":178697775,
				"rank":138600,
				"id":4
			},
			{
				"level":99,
				"xp":170219459,
				"rank":72256,
				"id":11
			},
			{
				"level":101,
				"xp":162870518,
				"rank":135028,
				"id":18
			},
			{
				"level":99,
				"xp":158125882,
				"rank":68467,
				"id":12
			},
			{
				"level":99,
				"xp":153526293,
				"rank":79274,
				"id":13
			},
			{
				"level":99,
				"xp":147626103,
				"rank":96495,
				"id":5
			},
			{
				"level":99,
				"xp":145592098,
				"rank":111023,
				"id":8
			},
			{
				"level":99,
				"xp":145035260,
				"rank":193334,
				"id":6
			},
			{
				"level":99,
				"xp":141860531,
				"rank":92535,
				"id":16
			},
			{
				"level":99,
				"xp":140813033,
				"rank":90819,
				"id":23
			},
			{
				"level":99,
				"xp":138724982,
				"rank":94631,
				"id":25
			},
			{
				"level":99,
				"xp":134702629,
				"rank":159554,
				"id":14
			},
			{
				"level":99,
				"xp":134701869,
				"rank":128678,
				"id":9
			},
			{
				"level":99,
				"xp":133472834,
				"rank":102418,
				"id":20
			},
			{
				"level":99,
				"xp":133401578,
				"rank":114642,
				"id":22
			},
			{
				"level":99,
				"xp":132891981,
				"rank":165458,
				"id":10
			},
			{
				"level":99,
				"xp":132819800,
				"rank":139467,
				"id":17
			},
			{
				"level":99,
				"xp":132806265,
				"rank":180059,
				"id":7
			},
			{
				"level":99,
				"xp":130466790,
				"rank":144349,
				"id":21
			},
			{
				"level":79,
				"xp":18816892,
				"rank":320425,
				"id":24
			}
		],
		"name":"Uss",
		"rank":"61,678",
		"melee":372841143,
		"combatlevel":138,
		"loggedIn":"false"
	}
`)

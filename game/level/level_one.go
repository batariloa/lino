package level

type LevelOne struct {
}

func NewLevelOne() *LevelOne {

	one := LevelOne{}
	one.GenerateMap()

	return &one
}

func (*LevelOne) GenerateMap() {

	LightOn := true

	LevelStatus = &Status{
		LightOn: &LightOn,
	}

	LevelMap = &[][]int{

		{
			63, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40,
			63, 40, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 40, 40,
			63, 62, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 63, 40,
			63, 62, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 63, 40,

			63, 62, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 63, 40,
			63, 62, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 63, 40,
			63, 62, 42, 42, 42, 42, 42, 44, 42, 42, 42, 42, 42, 42, 42, 42, 63, 40,
			63, 62, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 63, 40,

			63, 62, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 63, 40,
			63, 62, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 45, 43, 63, 40,
			63, 62, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 63, 40,
			63, 62, 43, 43, 43, 45, 43, 43, 66, 41, 41, 41, 41, 41, 41, 41, 40, 40,

			63, 62, 43, 43, 43, 45, 43, 43, 63, 40, 40, 40, 40, 40, 40, 40, 40, 40,
		},

		{
			63, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40,
			63, 40, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 46, 47, 48, 60, 40, 40,
			63, 62, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 49, 50, 51, 42, 63, 40,
			63, 62, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 52, 53, 54, 43, 63, 40,

			63, 62, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 63, 40,
			63, 62, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 63, 40,
			63, 62, 42, 42, 42, 42, 42, 44, 42, 42, 42, 42, 42, 42, 42, 42, 63, 40,
			63, 62, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 63, 40,
		},

		{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,

			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,

			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 67, 0, 0, 0, 0, 0, 0, 0, 0, 0,

			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		},
	}

	Interactables = []int{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,

		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,

		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0,

		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}

	Triggers = []int{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,

		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,

		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,

		0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}

	MaxLevelWidth = 18 * 20
	MaxLevelHeight = 13 * 20
}

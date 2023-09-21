package util

func SwitchLightOn(level *[][]int) {

	size := len((*level)[0])
	newLayer := make([]int, size)

	for i := range (*level)[0] {
		//dark shade
		newLayer[i] = 399
	}

	*level = append(*level, newLayer)
}

func SwitchLightOff(level *[][]int) {

	if len(*level) > 0 {
		*level = (*level)[:len(*level)-1]
	}
}

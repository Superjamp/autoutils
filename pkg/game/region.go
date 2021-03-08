package game

type Region int

const (
	NA Region = 13
	AS Region = 4
	EU Region = 6
	IP Region = 8
)

func (r Region) ToString() string {
	switch r {
	case NA:
		return "North America"
	case EU:
		return "Europe"
	case AS:
		return "Asia"
	case IP:
		return "Impostor"
	}
	return "Unknown"
}

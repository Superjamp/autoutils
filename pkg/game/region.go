package game

type Region int

const (
	NA Region = 0
	AS Region = 1
	EU Region = 2
	NA2 Region = 13
	AS2 Region = 4
	EU2 Region = 6
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
	case NA2:
		return "North America"
	case EU2:
		return "Europe"
	case AS2:
		return "Asia"
	case IP:
		return "Impostor"
	}
	return "Unknown"
}

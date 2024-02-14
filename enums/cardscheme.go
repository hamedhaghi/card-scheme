package enums

type CardScheme string

const (
	AmericanExpress CardScheme = "American Express"
	JCB             CardScheme = "JCB"
	Maestro         CardScheme = "Maestro"
	Visa            CardScheme = "Visa"
	MasterCard      CardScheme = "Master Card"
)

func (cs CardScheme) String() string {
	switch cs {
	case AmericanExpress:
		return string(AmericanExpress)
	case JCB:
		return string(JCB)
	case Maestro:
		return string(Maestro)
	case Visa:
		return string(Visa)
	case MasterCard:
		return string(MasterCard)
	}
	return ""
}

package weightconv

import "fmt"

type Kilogram float64
type Pound float64

func KtoP(k Kilogram) Pound { return Pound(k * 0.45359237) }

func PtoK(p Pound) Kilogram { return Kilogram(p / 0.45359237) }

func (k Kilogram) String() string { return fmt.Sprintf("%gKg", k) }

func (p Pound) String() string { return fmt.Sprintf("%glb", p) }

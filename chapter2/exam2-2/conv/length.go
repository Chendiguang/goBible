//!+

// lenconv performs Kilometer and meter

package conv

import (
	"fmt"
)

type Meter float64
type Kilometer float64

func (m Meter) String() string     { return fmt.Sprintf("%gm", m) }
func (k Kilometer) String() string { return fmt.Sprintf("%gkm", k) }

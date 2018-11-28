//!+

// weight converts grams and kilograms

package conv

import "fmt"

type Grams float64
type Kilograms float64

func (g Grams) String() string      { return fmt.Sprintf("%gg", g) }
func (kg Kilograms) String() string { return fmt.Sprintf("%gkg", kg) }

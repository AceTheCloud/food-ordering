package enums

// Status - Custom type to hold value for status ranging from 1-6
type Status int

// Declare related constants for each weekday starting with index 1
const (
	Ordering  Status = iota + 1 // EnumIndex = 1
	Placed                      // EnumIndex = 2
	Preparing                   // EnumIndex = 3
	Shipping                    // EnumIndex = 4
	Delivered                   // EnumIndex = 5
	Cancelled                   // EnumIndex = 6
)

// String - Creating common behavior - give the type a String function
func (s Status) String() string {
	return [...]string{"Ordering", "Placed", "Preparing", "Shipping", "Delivered", "Cancelled"}[s-1]
}

// EnumIndex - Creating common behavior - give the type a EnumIndex function
func (s Status) EnumIndex() int {
	return int(s)
}

// Example Use
// var status = Ordering
// fmt.Println(status)             // Print : Ordering
// fmt.Println(status.String())    // Print : Ordering
// fmt.Println(status.EnumIndex()) // Print : 1

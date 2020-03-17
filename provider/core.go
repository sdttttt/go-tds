package provider

// Service Info
type Service struct {
	*Address
}

// Address is Physical address
type Address struct {
	IP   string
	Port string
}

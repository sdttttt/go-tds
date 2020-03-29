package tds

type Onion struct {
	middleware []func()
}

package controller

type KeyStates struct {
	PrevEKeyState bool
}

func NewKeyStates() *KeyStates {
	return &KeyStates{}
}

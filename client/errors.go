package client

// 404
type MapNotFound struct {
}

// 486
type ActionInProgress struct {
}

// 490
type CharacterAtDestinationError struct {
}

// 497
type CharacterInvFull struct {
}

// 498
type CharacterNotFound struct {
}

// 499
type CharacterInCooldown struct {
}

// 598
type MonsterNotFound struct {
}

func (m MapNotFound) Error() string {
	return ""
}

func (a ActionInProgress) Error() string {
	return ""
}

func (e CharacterAtDestinationError) Error() string {
	return ""
}

func (i CharacterInvFull) Error() string {
	return ""
}

func (c CharacterNotFound) Error() string {
	return ""
}

func (i CharacterInCooldown) Error() string {
	return ""
}

func (m MonsterNotFound) Error() string {
	return ""
}

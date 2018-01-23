package fsmalexa

type Status struct {
	Code string `json:"code"`
}

type ResolvedValue struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type SlotValue struct {
	Value ResolvedValue `json:"value"`
}

type Authority struct {
	Authority string      `json:"authority"`
	Status    Status      `json:"status"`
	Values    []SlotValue `json:"values"`
}

type Resolutions struct {
	ResolutionsPerAuthority []Authority `json:"resolutionsPerAuthority"`
}

type SlotObject struct {
	Name               string      `json:"name"`
	Value              string      `json:"value"`
	ConfirmationStatus string      `json:"confirmationStatus"`
	Resolutions        Resolutions `json:"resolutions"`
}

type Intent struct {
	Name               string                `json:"name"`
	ConfirmationStatus string                `json:"confirmationStatus"`
	Slots              map[string]SlotObject `json:"slots"`
}

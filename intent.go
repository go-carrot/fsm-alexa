package fsmalexa

type Status struct {
	Code string `json:"code,omitempty"`
}

type ResolvedValue struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type SlotValue struct {
	Value ResolvedValue `json:"value,omitempty"`
}

type Authority struct {
	Authority string      `json:"authority,omitempty"`
	Status    Status      `json:"status,omitempty"`
	Values    []SlotValue `json:"values,omitempty"`
}

type Resolutions struct {
	ResolutionsPerAuthority []Authority `json:"resolutionsPerAuthority,omitempty"`
}

type SlotObject struct {
	Name               string      `json:"name,omitempty"`
	Value              string      `json:"value,omitempty"`
	ConfirmationStatus string      `json:"confirmationStatus,omitempty"`
	Resolutions        Resolutions `json:"resolutions,omitempty"`
}

type Intent struct {
	Name               string                `json:"name,omitempty"`
	ConfirmationStatus string                `json:"confirmationStatus,omitempty"`
	Slots              map[string]SlotObject `json:"slots,omitempty"`
}

package request

type Area struct {
	Level int64  `json:"-" binding:"oneof=0 1 2 3 4" validate:"oneof=0 1 2 3 4"`
	ID    int64  `json:"id,omitempty" binding:"gte=0" validate:"gte=0"`
	Key   string `json:"key,omitempty"`
}

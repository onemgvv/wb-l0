package domain

type Order struct {
	UID  string `json:"id" mapstructure:"id" db:"id"`
	Data string `json:"data" mapstructure:"data" db:"data"`
}

type OrderJSON map[string]any

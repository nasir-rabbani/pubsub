package models

// MsgData - The model to Map the message received from MQ
type MsgData struct {
	Offers []Offers `json:"offers"`
}

// Offers - Model to hold Offers object
type Offers struct {
	Hotel    Hotel    `json:"hotel"`
	Room     Room     `json:"room"`
	RatePlan RatePlan `json:"rate_plan"`
}

// Hotel - Model to hold Hotel Object
type Hotel struct {
	HotelID   string  `json:"hotel_id" gorm:"type:varchar(20);primaryKey"`
	Name      string  `json:"name" gorm:"type:varchar(200)"`
	Country   string  `json:"country" gorm:"type:varchar(100)"`
	Address   string  `json:"address" gorm:"type:varchar(200)"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Telephone string  `json:"telephone"`
	// Amenities   []string `json:"amenities"`
	Description string     `json:"description" gorm:"type:varchar(500)"`
	RoomCount   int        `json:"room_count" gorm:"type:int(8)"`
	Currency    string     `json:"currency" gorm:"type:varchar(50)"`
	Room        []Room     `gorm:"foreignkey:HotelID;references:HotelID"`
	RatePlan    []RatePlan `json:"rate_plan" gorm:"foreignkey:HotelID;references:HotelID"`
}

// Room - Model to hold Room Object
type Room struct {
	HotelID     string `json:"hotel_id" gorm:"type:varchar(20)"`
	RoomID      string `json:"room_id" gorm:"type:varchar(20)"`
	Description string `json:"description" gorm:"type:varchar(500)"`
	Name        string `json:"name"`
	// Capacity    Capacity `json:"capacity"`
}

// RatePlan - Model to hold RatePlan Object
type RatePlan struct {
	HotelID    string `json:"hotel_id" gorm:"type:varchar(20)"`
	RatePlanID string `json:"rate_plan_id" gorm:"type:varchar(20)"`
	// CancellationPolicy []CancellationPolicy `json:"cancellation_policy"`
	Name string `json:"name" gorm:"type:varchar(20)"`
	// OtherConditions    []string      `json:"other_conditions"`
	MealPlan string `json:"meal_plan" gorm:"type:varchar(200)"`
}

// Capacity -
type Capacity struct {
	MaxAdults     int `json:"max_adults"`
	ExtraChildren int `json:"extra_children"`
}

// CancellationPolicy -
type CancellationPolicy struct {
	Type              string `json:"type"`
	ExpiresDaysBefore int    `json:"expires_days_before"`
}

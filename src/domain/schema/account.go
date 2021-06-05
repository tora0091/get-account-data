package schema

type Account struct {
	PROPERTY_ID int64  `json:"property_id"`
	PLAN_ID     int64  `json:"plan_id"`
	CATEGORY_ID int64  `json:"category_id"`
	PRICE       int64  `json:"price"`
	NAME        string `json:"name"`
	START_DATE  string `json:"start_date"`
	END_DATE    string `json:"end_date"`
	SECTION     int64  `json:"section"`
}

type Accounts []Account

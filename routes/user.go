package routes



type User struct{
	//this is not the model User,see this as the serializer
	ID uint `json:"id" gorm:"primaryKey"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}
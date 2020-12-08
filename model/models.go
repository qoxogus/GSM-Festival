package model

//User struct
type User struct {
	ID        string `json:"id" bson:"id"`
	Username  string `json:"username" bson:"username"`
	Classnum  int    `json:"classnum" bson:"classnum"`
	Email     string `json:"email" bson:"email"`
	Password  string `json:"password" bson:"password"`
	Token     string `json:"token"`
	IsTeacher bool   `json:"is_teacher" bson:"is_teacher"`
}

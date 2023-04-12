package web

type CreateBookRequest struct {
	Title       string `validate:"required,min=1,max=255" json:"title"`
	Description string `validate:"required" json:"description"`
	Author      string `validate:"required" json:"author"`
	PublishDate string `validate:"required" json:"publish_date"`
}

type UpdateBookRequest struct {
	Id          int64  `validate:"required" json:"id"`
	Title       string `validate:"required,min=1,max=255" json:"title"`
	Description string `validate:"required" json:"description"`
	Author      string `validate:"required" json:"author"`
	PublishDate string `validate:"required" json:"publish_date"`
}

type CreateAdminRequest struct {
	Username string `validate:"required,min=1,max=50" json:"username"`
	Password string `validate:"required,min=1,max=255" json:"password"`
}

type LoginAdminRequest struct {
	Username string `validate:"required,min=1,max=50" json:"username"`
	Password string `validate:"required,min=1,max=255" json:"password"`
}

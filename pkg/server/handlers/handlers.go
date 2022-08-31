package handlers

type Handler struct {
	Session []string
}

func NewHandler() *Handler {
	return &Handler{
		Session: []string{},
	}
}
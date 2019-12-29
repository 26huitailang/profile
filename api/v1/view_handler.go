package v1

type ViewHandler struct {
	store Store
}

type Store interface {
	IDeviceManager
}

func NewViewHandler(store Store) *ViewHandler {
	return &ViewHandler{
		store: store,
	}
}

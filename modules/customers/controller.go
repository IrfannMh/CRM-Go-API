package customers

type ControllerCustomer struct {
	useCase *UseCase
}
type CreateResponse struct {
	Message string               `json:"message"`
	Data    CustomerItemResponse `json:"data"`
}
type CustomerItemResponse struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
}

type ReadCustomerResponse struct {
	Data []CustomerItemResponse `json:"data"`
}

func NewController(useCase *UseCase) *ControllerCustomer {
	return &ControllerCustomer{
		useCase: useCase,
	}
}

func (c ControllerCustomer) Create(req *CreateRequest) (*CreateResponse, error) {
	customer := Customer{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Avatar:    req.Avatar,
	}
	err := c.useCase.Create(&customer)
	if err != nil {
		return nil, err
	}

	res := &CreateResponse{
		Message: "Success",
		Data: CustomerItemResponse{
			ID:        customer.ID,
			FirstName: customer.FirstName,
			LastName:  customer.LastName,
			Email:     customer.Email,
			Avatar:    customer.Avatar,
		},
	}

	return res, nil
}

func (c ControllerCustomer) GetAll(page int) (*ReadCustomerResponse, error) {
	customers, err := c.useCase.GetAll(page)
	if err != nil {
		return nil, err
	}
	result := &ReadCustomerResponse{}

	for _, v := range customers {
		item := CustomerItemResponse{
			ID:        v.ID,
			FirstName: v.FirstName,
			LastName:  v.LastName,
			Email:     v.Email,
			Avatar:    v.Avatar,
		}
		result.Data = append(result.Data, item)
	}
	return result, nil
}

func (c ControllerCustomer) Delete(customer *Customer) error {
	return c.useCase.Delete(customer)
}

func (c ControllerCustomer) FindByID(id string) (*Customer, error) {
	var customer Customer
	err := c.useCase.FindById(&customer, id)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

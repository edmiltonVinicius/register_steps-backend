package contracts

type ContractError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

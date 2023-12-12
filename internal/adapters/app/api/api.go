package api

import "github.com/matheus-alpe/go-hex-structure-example/internal/ports"

type Adapter struct {
	db    ports.DbPort
	arith ports.ArithmeticPort
}

func NewAdapter(db ports.DbPort, arith ports.ArithmeticPort) *Adapter {
	return &Adapter{db: db, arith: arith}
}

func (api *Adapter) GetAddition(a, b int32) (int32, error) {
	answer, err := api.arith.Addition(a, b)
	if err != nil {
		return 0, err
	}

	err = api.db.AddToHistory(answer, "addition")
	if err != nil {
		return 0, err
	}

	return answer, nil
}

func (api *Adapter) GetSubtraction(a, b int32) (int32, error) {
	answer, err := api.arith.Subtraction(a, b)
	if err != nil {
		return 0, err
	}

	err = api.db.AddToHistory(answer, "subtraction")
	if err != nil {
		return 0, err
	}

	return answer, nil
}

func (api *Adapter) GetMultiplication(a, b int32) (int32, error) {
	answer, err := api.arith.Multiplication(a, b)
	if err != nil {
		return 0, err
	}
	err = api.db.AddToHistory(answer, "multiplication")
	if err != nil {
		return 0, err
	}

	return answer, nil
}

func (api *Adapter) GetDivision(a, b int32) (int32, error) {
	answer, err := api.arith.Division(a, b)
	if err != nil {
		return 0, err
	}

	err = api.db.AddToHistory(answer, "division")
	if err != nil {
		return 0, err
	}

	return answer, nil
}

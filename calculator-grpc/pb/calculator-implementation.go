package pb

import "context"

type CalculatorServer struct {
	UnimplementedCalculatorServiceServer
}

func (s *CalculatorServer) Add(ctx context.Context, payload *CalculatorPayload) (*CalculatorResult, error) {
	result := payload.X + payload.Y
	return &CalculatorResult{Result: result}, nil
}

func (s *CalculatorServer) Sub(ctx context.Context, payload *CalculatorPayload) (*CalculatorResult, error) {
	result := payload.X - payload.Y
	return &CalculatorResult{Result: result}, nil
}

func (s *CalculatorServer) Multiply(ctx context.Context, payload *CalculatorPayload) (*CalculatorResult, error) {
	result := payload.X * payload.Y
	return &CalculatorResult{Result: result}, nil
}

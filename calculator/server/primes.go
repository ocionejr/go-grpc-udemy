package main

import (
	pb "github.com/ocionejr/go-grpc-udemy/calculator/proto"
)

func (s *Server) Primes(in *pb.PrimeRequest, stream pb.CalculatorService_PrimesServer) error {
	k := int64(2)
	N := in.Number

	for  N > 1 {
		if N % k == 0 { 
				stream.Send(&pb.PrimeResponse{
					Result: k,
				})   
				N = N / k   
		} else {
				k++
		}
	}

	return nil
}
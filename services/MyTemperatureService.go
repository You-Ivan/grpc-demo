package services

import (
	client_stream "grpc-tutorial/api/proto/client-stream"
	"io"
	"log"
)

type MyTemperatureService struct {
	client_stream.UnimplementedTemperatureServiceServer
}

func (m MyTemperatureService) GetAvgTemp(stream client_stream.TemperatureService_GetAvgTempServer) error {
	sum := 0.0
	cnt := 0

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		sum += resp.Degree
		cnt++
		log.Printf("Time: %v, Avg Temp: %v\n", resp.Time, sum/float64(cnt))
	}
	return nil
}

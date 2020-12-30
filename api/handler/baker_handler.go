package handler

import (
	"context"
	"math/rand"
	"sync"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"pancake.maker/gen/api"
)

func init() {
	rand.Seed(time.Now().UnixNano()) // Initialize seeds that affect pancake score
}

type BakerHandler struct {
	report *report
}

// report records the number of pancakes baked for each menu
type report struct {
	sync.Mutex // Safe for multiple chefs to bake at the same time
	data       map[api.Pancake_Menu]int
}

func NewBakerHandler() *BakerHandler {
	return &BakerHandler{
		report: &report{
			data: make(map[api.Pancake_Menu]int),
		},
	}
}

// Bake a pancake of the specified menu and return as response
func (h *BakerHandler) Bake(_ context.Context, req *api.BakeRequest) (*api.BakeResponse, error) {
	// validation
	if req.Menu == api.Pancake_UNKNOWN || req.Menu > api.Pancake_SPICY_CURRY {
		return nil, status.Errorf(codes.InvalidArgument, "select a pancake from the menu")
	}

	// bake a pancake and update the report
	now := time.Now()
	panCake := &api.Pancake{
		ChefName:       "Martin", // alone
		Menu:           req.Menu,
		TechnicalScore: rand.Float32(),
		CreateTime: &timestamp.Timestamp{ // Timestamp defined by .proto file is converted to timestamp.Timestamp type
			Seconds: now.Unix(),
			Nanos:   int32(now.Nanosecond()),
		},
	}

	h.report.Lock()
	h.report.data[req.Menu] = h.report.data[req.Menu] + 1
	h.report.Unlock()

	return &api.BakeResponse{Pancake: panCake}, nil
}

// Report the number of pancakes baked
func (h *BakerHandler) Report(_ context.Context, req *api.ReportRequest) (*api.ReportResponse, error) {
	bakeCounts := make([]*api.Report_BakeCount, 0)

	h.report.Lock()
	for menu, count := range h.report.data {
		bakeCounts = append(bakeCounts, &api.Report_BakeCount{
			Menu:  menu,
			Count: int32(count),
		})
	}
	h.report.Unlock()

	report := &api.Report{
		BakeCounts: bakeCounts,
	}

	return &api.ReportResponse{Report: report}, nil
}

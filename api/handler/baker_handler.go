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
	rand.Seed(time.Now().UnixNano()) // パンケーキの仕上がりに影響するseedを初期化
}

type BakerHandler struct {
	report *report
}

// メニューごとの、焼いたパンケーキの枚数を記録するレポート
type report struct {
	sync.Mutex // 複数人が同時に焼いても大丈夫なように
	data       map[api.Pancake_Menu]int
}

func NewBakerHandler() *BakerHandler {
	return &BakerHandler{
		report: &report{
			data: make(map[api.Pancake_Menu]int),
		},
	}
}

// 指定されたメニューのパンケーキを焼いて、レスポンスとして返す
func (h *BakerHandler) Bake(_ context.Context, req *api.BakeRequest) (*api.BakeResponse, error) {
	// validation
	if req.Menu == api.Pancake_UNKNOWN || req.Menu > api.Pancake_SPICY_CURRY {
		return nil, status.Errorf(codes.InvalidArgument, "select a pancake from the menu")
	}

	// パンケーキを焼いて、レポートのそのメニューの焼いた数を更新する
	now := time.Now()
	panCake := &api.Pancake{
		ChefName:       "Martin", // 現在はワンオペ
		Menu:           req.Menu,
		TechnicalScore: rand.Float32(),
		CreateTime: &timestamp.Timestamp{ // .protoで定義したTimestampはtimestamp.Timestamp型に変換される
			Seconds: now.Unix(),
			Nanos:   int32(now.Nanosecond()),
		},
	}

	h.report.Lock()
	h.report.data[req.Menu] = h.report.data[req.Menu] + 1
	h.report.Unlock()

	return &api.BakeResponse{Pancake: panCake}, nil
}

// 焼いたパンケーキの数を報告する
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

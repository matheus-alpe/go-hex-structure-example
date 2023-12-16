package rpc

import (
	"log"
	"net"
	"os"
	"testing"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"

	"github.com/matheus-alpe/go-hex-structure-example/internal/adapters/app/api"
	"github.com/matheus-alpe/go-hex-structure-example/internal/adapters/core/arithmetic"
	"github.com/matheus-alpe/go-hex-structure-example/internal/adapters/framework/left/grpc/pb"
	"github.com/matheus-alpe/go-hex-structure-example/internal/adapters/framework/right/db"
	"github.com/matheus-alpe/go-hex-structure-example/internal/ports"
	"github.com/stretchr/testify/require"
)

const bufSize = 1024 * 1024

var listener *bufconn.Listener

func init() {
	listener = bufconn.Listen(bufSize)
	grpcServer := grpc.NewServer()

	// ports
    dbDriver := os.Getenv("DB_DRIVER")
    dbDSN := os.Getenv("DB_DATA_SOURCE_NAME")

    var dbAdapter ports.DbPort = db.NewAdapter(dbDriver, dbDSN)
    var core ports.ArithmeticPort = arithmetic.NewAdapter()
    var appAdapter ports.APIPort = api.NewAdapter(dbAdapter, core)
    var grpcAdapter ports.GRPCPort = NewAdapter(appAdapter)

	pb.RegisterArithmeticServiceServer(grpcServer, grpcAdapter)
	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("test server start error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return listener.Dial()
}

func getGRPCConnection(ctx context.Context, t *testing.T) *grpc.ClientConn {
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("failed to dial bufnet: %v", err)
	}

	return conn
}

func TestGetAddition(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConnection(ctx, t)
	defer conn.Close()

	client := pb.NewArithmeticServiceClient(conn)
	params := &pb.OperationParameters{
		A: 1,
		B: 1,
	}

	answer, err := client.GetAddition(ctx, params)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	require.Equal(t, answer.Value, int32(2))
}

func TestGetSubtraction(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConnection(ctx, t)
	defer conn.Close()

	client := pb.NewArithmeticServiceClient(conn)
	params := &pb.OperationParameters{
		A: 1,
		B: 1,
	}

	answer, err := client.GetSubtraction(ctx, params)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	require.Equal(t, answer.Value, int32(0))
}

func TestGetMultiplication(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConnection(ctx, t)
	defer conn.Close()

	client := pb.NewArithmeticServiceClient(conn)
	params := &pb.OperationParameters{
		A: 1,
		B: 1,
	}

	answer, err := client.GetMultiplication(ctx, params)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	require.Equal(t, answer.Value, int32(1))
}

func TestGetDivision(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConnection(ctx, t)
	defer conn.Close()

	client := pb.NewArithmeticServiceClient(conn)
	params := &pb.OperationParameters{
		A: 1,
		B: 1,
	}

	answer, err := client.GetDivision(ctx, params)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	require.Equal(t, answer.Value, int32(1))
}

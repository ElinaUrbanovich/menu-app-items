package items

import (
	"context"
	"log"
	"net"

	"github.com/ElinaUrbanovich/menu-app-items/pkg/items/pb"
	"github.com/jackc/pgx/v4"
	"google.golang.org/grpc"
)

const port string = ":10091"

type ItemServiceServer struct {
	Conn *pgx.Conn
	pb.UnimplementedItemServiceServer
}

func NewCategoryServer() *ItemServiceServer {
	return &ItemServiceServer{}
}

func (server *ItemServiceServer) Run() error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterItemServiceServer(s, server)
	log.Printf("server listening at %v", lis.Addr())

	return s.Serve(lis)
}

func (server *ItemServiceServer) CreateNewCategory(ctx context.Context, in *pb.NewCategory) (*pb.Category, error) {

	log.Printf("Received: %v", in.GetName())

	createdCategory := &pb.Category{Name: in.GetName()}
	tx, err := server.Conn.Begin(context.Background())
	if err != nil {
		log.Fatalf("conn.Begin failed: %v", err)
	}

	_, err = tx.Exec(context.Background(), "insert into categories(name) values ($1)",
		createdCategory.Name)
	if err != nil {
		log.Fatalf("tx.Exec failed: %v", err)
	}
	tx.Commit(context.Background())
	return createdCategory, nil
}

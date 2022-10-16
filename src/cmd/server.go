package cmd

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
	"os"
	"os/signal"
	"social-work_notification_service/src/config"
	notification_transport "social-work_notification_service/src/features/notification/transport"
	notificationpb "social-work_notification_service/src/pb/notification"
	"sync"
)

func server(config config.IConfig) *cobra.Command {
	return &cobra.Command{
		Use: "server",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			var wg sync.WaitGroup
			wg.Add(2)
			go runGrpcServer(ctx, config, &wg)

			go runGrpcGateway(ctx, config, &wg)
			wg.Wait()
			zap.S().Info("shutdown application")
		},
	}
}

func runGrpcServer(ctx context.Context, config config.IConfig, wg *sync.WaitGroup) {
	server := grpc.NewServer()
	reflection.Register(server)

	notificationTransport := notification_transport.New()

	notificationpb.RegisterNotificationServiceServer(server, notificationTransport)

	lis, err := net.Listen("tcp", config.GetGrpcAddress())
	if err != nil {
		zap.S().Fatalln(err)
	}

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	signal.Notify(sigint, os.Kill)

	go func() {
		zap.S().Info(fmt.Sprintf("grpc server is running at %s", config.GetGrpcAddress()))
		err = server.Serve(lis)
		if err != nil {
			zap.S().Fatalln(err)
		}

	}()

	<-sigint
	server.GracefulStop()
	wg.Done()
	zap.S().Info("shutdown grpc server")

}

func runGrpcGateway(ctx context.Context, config config.IConfig, wg *sync.WaitGroup) {
	conn, err := grpc.DialContext(ctx, config.GetGrpcAddress(), grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		zap.S().Fatalln(err)
	}

	gw := runtime.NewServeMux()
	err = notificationpb.RegisterNotificationServiceHandler(ctx, gw, conn)
	if err != nil {
		zap.S().Fatalln(err)
	}

	gwServer := &http.Server{
		Addr:    ":21000",
		Handler: gw,
	}

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	signal.Notify(sigint, os.Kill)

	go func() {
		zap.S().Info(fmt.Sprintf("grpc gateway server is running at %s", "21000"))
		err = gwServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			zap.S().Fatalln(err)
		}
	}()

	<-sigint
	gwServer.Shutdown(ctx)
	wg.Done()
	zap.S().Info("shutdown grpc gateway server")
}

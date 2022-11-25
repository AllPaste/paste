package interceptor

import (
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var customFunc grpc_recovery.RecoveryHandlerFunc

func customRecoveryHandler() []grpc_recovery.Option {
	// Define customfunc to handle panic
	customFunc = func(p interface{}) (err error) {
		return status.Errorf(codes.Unknown, "panic triggered: %+v", p)
	}
	// Shared options for the logger, with a custom gRPC code to log level function.
	return []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(customFunc),
	}
}

var RecoveryUnaryServerInterceptor = grpc_recovery.UnaryServerInterceptor(
	customRecoveryHandler()...)

var RecoveryStreamUnaryServerInterceptor = grpc_recovery.StreamServerInterceptor(
	customRecoveryHandler()...)

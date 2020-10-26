package utils

import (
	"errors"

	"google.golang.org/grpc/status"

	"github.com/zerops-io/zcli/src/zeropsApiProtocol"
	"github.com/zerops-io/zcli/src/zeropsVpnProtocol"
)

func HandleGrpcApiError(
	response interface {
		GetError() *zeropsApiProtocol.Error
	},
	err error,
) error {
	if err != nil {
		if s, ok := status.FromError(err); ok {
			return errors.New(s.Message())
		}
		return err
	}
	if response.GetError().GetCode() != zeropsApiProtocol.ErrorCode_NO_ERROR {
		return errors.New(response.GetError().GetMessage())
	}

	return nil
}

func HandleVpnApiError(
	response interface {
		GetError() *zeropsVpnProtocol.Error
	},
	err error,
) error {
	if err != nil {
		if s, ok := status.FromError(err); ok {
			return errors.New(s.Message())
		}
		return err
	}
	if response.GetError().GetCode() != zeropsVpnProtocol.ErrorCode_NO_ERROR {
		return errors.New(response.GetError().GetMessage())
	}

	return nil
}

func HandleDaemonError(
	err error,
) error {
	if err != nil {
		if s, ok := status.FromError(err); ok {
			return errors.New(s.Message())
		}
		return err
	}

	return nil
}

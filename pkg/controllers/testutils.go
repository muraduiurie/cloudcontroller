package controllers

import (
	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

func testLogger() logr.Logger {
	logger := zap.New()
	ctrl.SetLogger(logger)
	return ctrl.Log.WithName("test")
}

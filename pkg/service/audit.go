package service

import (
	"context"
	"github.com/Onelvay/halyklife-test-task/pkg/domain"
	"go.mongodb.org/mongo-driver/mongo"
	_ "go.mongodb.org/mongo-driver/mongo"
)

type AuditServer struct {
	db *mongo.Collection
}

func NewAuditServer(db *mongo.Collection) *AuditServer {
	return &AuditServer{db}
}

func (a *AuditServer) Log(ctx context.Context, log domain.Log) error {
	_, err := a.db.InsertOne(ctx, log)
	return err
}

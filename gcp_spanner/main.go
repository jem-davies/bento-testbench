package main

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/spanner"
	database "cloud.google.com/go/spanner/admin/database/apiv1"
	"cloud.google.com/go/spanner/admin/database/apiv1/databasepb"
	instance "cloud.google.com/go/spanner/admin/instance/apiv1"
	"cloud.google.com/go/spanner/admin/instance/apiv1/instancepb"
	"github.com/anicoll/screamer/pkg/interceptor"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	os.Setenv("SPANNER_EMULATOR_HOST", "localhost:9010")

	// create instance

	instanceAdminClient, err := instance.NewInstanceAdminClient(ctx)
	if err != nil {
		panic(err)
	}
	defer instanceAdminClient.Close()

	op, err := instanceAdminClient.CreateInstance(ctx, &instancepb.CreateInstanceRequest{
		Parent:     "projects/project-id",
		InstanceId: "instance-id",
		Instance: &instancepb.Instance{
			Config:          "projects/model/instanceConfigs/regional-us-central1",
			DisplayName:     "instance-id",
			ProcessingUnits: 100,
		},
	})
	if err != nil {
		panic(err)
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Name)

	// create database

	databaseAdminClient, err := database.NewDatabaseAdminClient(ctx)
	if err != nil {
		panic(err)
	}
	defer databaseAdminClient.Close()

	op2, err := databaseAdminClient.CreateDatabase(ctx, &databasepb.CreateDatabaseRequest{
		Parent:          resp.Name,
		CreateStatement: fmt.Sprintf("CREATE DATABASE `%s`", "stream-database"),
	})
	if err != nil {
		panic(err)
	}

	resp2, err := op2.Wait(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp2)

	// create table and changeStream

	op3, err := databaseAdminClient.UpdateDatabaseDdl(ctx, &databasepb.UpdateDatabaseDdlRequest{
		Database: "projects/project-id/instances/instance-id/databases/stream-database",
		Statements: []string{
			fmt.Sprintf(`CREATE TABLE %s (
				Int64           INT64,
				String          STRING(MAX),
			) PRIMARY KEY (Int64)`, "test"),
			fmt.Sprintf(`CREATE CHANGE STREAM %s FOR %s`, "Stream", "test"),
		},
	})
	if err != nil {
		panic(err)
	}

	if err := op3.Wait(ctx); err != nil {
		panic(err)
	}

	// send data to the table
	spannerInterceptor := interceptor.NewQueueInterceptor(100)
	client, err := spanner.NewClient(ctx, "projects/project-id/instances/instance-id/databases/stream-database", option.WithGRPCDialOption(grpc.WithChainUnaryInterceptor(spannerInterceptor.UnaryInterceptor)))

	for {
		client.Apply(ctx, []*spanner.Mutation{
			spanner.InsertOrUpdate("test", []string{"Int64", "String"}, []interface{}{23, "Alice"}),
		})
	}
}

package main

import (
	"log"

	"github.com/azybk/music-catalog/internal/configs"
	membershipsHandler "github.com/azybk/music-catalog/internal/handler/memberships"
	"github.com/azybk/music-catalog/internal/models/memberships"
	membershipsRepo "github.com/azybk/music-catalog/internal/repository/memberships"
	membershipsSvc "github.com/azybk/music-catalog/internal/service/memberships"
	"github.com/azybk/music-catalog/pkg/internalsql"
	"github.com/gin-gonic/gin"
)

func main() {
	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolder([]string{"./internal/configs"}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)
	if err != nil {
		log.Fatal("Gagal inisiasi Config", err)
	}

	cfg = configs.Get()

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatalf("gagal konek ke database, err: %+v", err)
	}

	db.AutoMigrate(&memberships.User{})
	r := gin.Default()

	membershipsRepo := membershipsRepo.NewRepository(db)
	membershipsSvc := membershipsSvc.NewService(cfg, membershipsRepo)
	membershipsHandler := membershipsHandler.NewHandler(r, membershipsSvc)

	membershipsHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}

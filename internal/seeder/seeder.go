package seeder

import (
	"context"
	"errors"
	"os"
	"pleasurelove/internal/constanta"
	"pleasurelove/internal/models"
	"pleasurelove/internal/utils"
	"pleasurelove/pkg/logger"
	"time"

	"gorm.io/gorm"
)

func SeedSuperAdmin(ctx context.Context, db *gorm.DB) error {
	if db == nil {
		err := errors.New("database connection is nil")
		logger.Error(ctx, "Database connection is nil", err)
		return err
	}

	// Periksa apakah role superadmin sudah ada
	var role models.Roles
	if err := db.Where("name = ?", "superadmin").First(&role).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Buat role superadmin
			role = models.Roles{
				Name: "superadmin",
				Code: constanta.RoleCodeSuperAdmin,
			}
			if err := db.Create(&role).Error; err != nil {
				logger.Error(ctx, "Failed to create superadmin role", err)
				return err
			}
		} else {
			logger.Error(ctx, "Failed to query superadmin role", err)
			return err
		}
	}

	// Ambil email dan password dari environment variables
	superAdminEmail := os.Getenv("SUPERADMIN_EMAIL")
	superAdminPassword := os.Getenv("SUPERADMIN_PASSWORD")
	if superAdminEmail == "" || superAdminPassword == "" {
		err := errors.New("SUPERADMIN_EMAIL or SUPERADMIN_PASSWORD environment variable is not set")
		logger.Error(ctx, "SUPERADMIN_EMAIL or SUPERADMIN_PASSWORD environment variable is not set", err)
		return err
	}

	// Periksa apakah user superadmin sudah ada
	var user models.User
	if err := db.Where("email = ?", superAdminEmail).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Hash password
			hashedPassword, err := utils.HashPassword(superAdminPassword)
			if err != nil {
				logger.Error(ctx, "Failed to hash password", err)
				return err
			}

			// Buat user superadmin
			user = models.User{
				Username:  "superadmin",
				Email:     superAdminEmail,
				Password:  hashedPassword,
				RoleID:    role.ID,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
			if err := db.Create(&user).Error; err != nil {
				logger.Error(ctx, "Failed to create superadmin user", err)
				return err
			}
		} else {
			logger.Error(ctx, "Failed to query superadmin user", err)
			return err
		}
	}

	logger.Info(ctx, "Superadmin seeding completed successfully", nil)
	return nil
}

package usecase

import (
	"context"
	"pleasurelove/internal/constanta"

	"gorm.io/gorm"
)

func processWithTx(ctx context.Context, db *gorm.DB, fn func(ctx context.Context) error) error {
    tx := db.Begin()
    if tx.Error != nil {
        return tx.Error
    }

    // Menyimpan database awal
    originalDB := db

    // Membuat context baru dengan transaksi
    ctx = context.WithValue(ctx, constanta.Tx, tx)

    defer func() {
        if r := recover(); r != nil {
            tx.Rollback()
            panic(r)
        }

        // Reset repository ke database utama
        ctx = context.WithValue(ctx, constanta.Tx, originalDB)
    }()

    if err := fn(ctx); err != nil {
        tx.Rollback()
        return err
    }

    return tx.Commit().Error
}


// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package analytic

import (
	"context"
	"github.com/authgear/authgear-server/pkg/lib/analytic"
	"github.com/authgear/authgear-server/pkg/lib/config"
	"github.com/authgear/authgear-server/pkg/lib/infra/db"
	"github.com/authgear/authgear-server/pkg/lib/infra/db/appdb"
	"github.com/authgear/authgear-server/pkg/lib/infra/db/auditdb"
	"github.com/authgear/authgear-server/pkg/lib/infra/db/globaldb"
	"github.com/authgear/authgear-server/pkg/lib/infra/redis"
	"github.com/authgear/authgear-server/pkg/lib/infra/redis/analyticredis"
	"github.com/authgear/authgear-server/pkg/util/clock"
	"github.com/authgear/authgear-server/pkg/util/periodical"
)

// Injectors from wire.go:

func NewUserWeeklyReport(ctx context.Context, pool *db.Pool, databaseCredentials *config.DatabaseCredentials) *analytic.UserWeeklyReport {
	globalDatabaseCredentialsEnvironmentConfig := NewGlobalDatabaseCredentials(databaseCredentials)
	databaseEnvironmentConfig := config.NewDefaultDatabaseEnvironmentConfig()
	factory := NewLoggerFactory()
	handle := globaldb.NewHandle(ctx, pool, globalDatabaseCredentialsEnvironmentConfig, databaseEnvironmentConfig, factory)
	sqlBuilder := globaldb.NewSQLBuilder(globalDatabaseCredentialsEnvironmentConfig)
	sqlExecutor := globaldb.NewSQLExecutor(ctx, handle)
	globalDBStore := &analytic.GlobalDBStore{
		SQLBuilder:  sqlBuilder,
		SQLExecutor: sqlExecutor,
	}
	databaseConfig := NewDatabaseConfig()
	appdbHandle := appdb.NewHandle(ctx, pool, databaseConfig, databaseCredentials, factory)
	appdbSQLBuilder := appdb.NewSQLBuilder(databaseCredentials)
	appdbSQLExecutor := appdb.NewSQLExecutor(ctx, appdbHandle)
	appDBStore := &analytic.AppDBStore{
		SQLBuilder:  appdbSQLBuilder,
		SQLExecutor: appdbSQLExecutor,
	}
	userWeeklyReport := &analytic.UserWeeklyReport{
		GlobalHandle:  handle,
		GlobalDBStore: globalDBStore,
		AppDBHandle:   appdbHandle,
		AppDBStore:    appDBStore,
	}
	return userWeeklyReport
}

func NewProjectWeeklyReport(ctx context.Context, pool *db.Pool, databaseCredentials *config.DatabaseCredentials, auditDatabaseCredentials *config.AuditDatabaseCredentials) *analytic.ProjectWeeklyReport {
	globalDatabaseCredentialsEnvironmentConfig := NewGlobalDatabaseCredentials(databaseCredentials)
	databaseEnvironmentConfig := config.NewDefaultDatabaseEnvironmentConfig()
	factory := NewLoggerFactory()
	handle := globaldb.NewHandle(ctx, pool, globalDatabaseCredentialsEnvironmentConfig, databaseEnvironmentConfig, factory)
	sqlBuilder := globaldb.NewSQLBuilder(globalDatabaseCredentialsEnvironmentConfig)
	sqlExecutor := globaldb.NewSQLExecutor(ctx, handle)
	globalDBStore := &analytic.GlobalDBStore{
		SQLBuilder:  sqlBuilder,
		SQLExecutor: sqlExecutor,
	}
	databaseConfig := NewDatabaseConfig()
	appdbHandle := appdb.NewHandle(ctx, pool, databaseConfig, databaseCredentials, factory)
	appdbSQLBuilder := appdb.NewSQLBuilder(databaseCredentials)
	appdbSQLExecutor := appdb.NewSQLExecutor(ctx, appdbHandle)
	appDBStore := &analytic.AppDBStore{
		SQLBuilder:  appdbSQLBuilder,
		SQLExecutor: appdbSQLExecutor,
	}
	readHandle := auditdb.NewReadHandle(ctx, pool, databaseConfig, auditDatabaseCredentials, factory)
	auditdbSQLBuilder := auditdb.NewSQLBuilder(auditDatabaseCredentials)
	readSQLExecutor := auditdb.NewReadSQLExecutor(ctx, readHandle)
	auditDBReadStore := &analytic.AuditDBReadStore{
		SQLBuilder:  auditdbSQLBuilder,
		SQLExecutor: readSQLExecutor,
	}
	projectWeeklyReport := &analytic.ProjectWeeklyReport{
		GlobalHandle:  handle,
		GlobalDBStore: globalDBStore,
		AppDBHandle:   appdbHandle,
		AppDBStore:    appDBStore,
		AuditDBHandle: readHandle,
		AuditDBStore:  auditDBReadStore,
	}
	return projectWeeklyReport
}

func NewProjectMonthlyReport(ctx context.Context, pool *db.Pool, databaseCredentials *config.DatabaseCredentials, auditDatabaseCredentials *config.AuditDatabaseCredentials) *analytic.ProjectMonthlyReport {
	globalDatabaseCredentialsEnvironmentConfig := NewGlobalDatabaseCredentials(databaseCredentials)
	databaseEnvironmentConfig := config.NewDefaultDatabaseEnvironmentConfig()
	factory := NewLoggerFactory()
	handle := globaldb.NewHandle(ctx, pool, globalDatabaseCredentialsEnvironmentConfig, databaseEnvironmentConfig, factory)
	sqlBuilder := globaldb.NewSQLBuilder(globalDatabaseCredentialsEnvironmentConfig)
	sqlExecutor := globaldb.NewSQLExecutor(ctx, handle)
	globalDBStore := &analytic.GlobalDBStore{
		SQLBuilder:  sqlBuilder,
		SQLExecutor: sqlExecutor,
	}
	databaseConfig := NewDatabaseConfig()
	readHandle := auditdb.NewReadHandle(ctx, pool, databaseConfig, auditDatabaseCredentials, factory)
	auditdbSQLBuilder := auditdb.NewSQLBuilder(auditDatabaseCredentials)
	readSQLExecutor := auditdb.NewReadSQLExecutor(ctx, readHandle)
	auditDBReadStore := &analytic.AuditDBReadStore{
		SQLBuilder:  auditdbSQLBuilder,
		SQLExecutor: readSQLExecutor,
	}
	projectMonthlyReport := &analytic.ProjectMonthlyReport{
		GlobalHandle:  handle,
		GlobalDBStore: globalDBStore,
		AuditDBHandle: readHandle,
		AuditDBStore:  auditDBReadStore,
	}
	return projectMonthlyReport
}

func NewCountCollector(ctx context.Context, pool *db.Pool, databaseCredentials *config.DatabaseCredentials, auditDatabaseCredentials *config.AuditDatabaseCredentials, redisPool *redis.Pool, credentials *config.AnalyticRedisCredentials) *analytic.CountCollector {
	globalDatabaseCredentialsEnvironmentConfig := NewGlobalDatabaseCredentials(databaseCredentials)
	databaseEnvironmentConfig := config.NewDefaultDatabaseEnvironmentConfig()
	factory := NewLoggerFactory()
	handle := globaldb.NewHandle(ctx, pool, globalDatabaseCredentialsEnvironmentConfig, databaseEnvironmentConfig, factory)
	sqlBuilder := globaldb.NewSQLBuilder(globalDatabaseCredentialsEnvironmentConfig)
	sqlExecutor := globaldb.NewSQLExecutor(ctx, handle)
	globalDBStore := &analytic.GlobalDBStore{
		SQLBuilder:  sqlBuilder,
		SQLExecutor: sqlExecutor,
	}
	databaseConfig := NewDatabaseConfig()
	appdbHandle := appdb.NewHandle(ctx, pool, databaseConfig, databaseCredentials, factory)
	appdbSQLBuilder := appdb.NewSQLBuilder(databaseCredentials)
	appdbSQLExecutor := appdb.NewSQLExecutor(ctx, appdbHandle)
	appDBStore := &analytic.AppDBStore{
		SQLBuilder:  appdbSQLBuilder,
		SQLExecutor: appdbSQLExecutor,
	}
	readHandle := auditdb.NewReadHandle(ctx, pool, databaseConfig, auditDatabaseCredentials, factory)
	auditdbSQLBuilder := auditdb.NewSQLBuilder(auditDatabaseCredentials)
	readSQLExecutor := auditdb.NewReadSQLExecutor(ctx, readHandle)
	auditDBReadStore := &analytic.AuditDBReadStore{
		SQLBuilder:  auditdbSQLBuilder,
		SQLExecutor: readSQLExecutor,
	}
	writeHandle := auditdb.NewWriteHandle(ctx, pool, databaseConfig, auditDatabaseCredentials, factory)
	writeSQLExecutor := auditdb.NewWriteSQLExecutor(ctx, writeHandle)
	auditDBWriteStore := &analytic.AuditDBWriteStore{
		SQLBuilder:  auditdbSQLBuilder,
		SQLExecutor: writeSQLExecutor,
	}
	redisConfig := NewRedisConfig()
	analyticredisHandle := analyticredis.NewHandle(redisPool, redisConfig, credentials, factory)
	readStoreRedis := &analytic.ReadStoreRedis{
		Context: ctx,
		Redis:   analyticredisHandle,
	}
	countCollector := &analytic.CountCollector{
		GlobalHandle:       handle,
		GlobalDBStore:      globalDBStore,
		AppDBHandle:        appdbHandle,
		AppDBStore:         appDBStore,
		AuditDBReadHandle:  readHandle,
		AuditDBReadStore:   auditDBReadStore,
		AuditDBWriteHandle: writeHandle,
		AuditDBWriteStore:  auditDBWriteStore,
		CounterStore:       readStoreRedis,
	}
	return countCollector
}

func NewPeriodicalArgumentParser() *periodical.ArgumentParser {
	clock := _wireSystemClockValue
	argumentParser := &periodical.ArgumentParser{
		Clock: clock,
	}
	return argumentParser
}

var (
	_wireSystemClockValue = clock.NewSystemClock()
)

// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package accountdeletion

import (
	"context"
	"github.com/authgear/authgear-server/pkg/lib/config"
	"github.com/authgear/authgear-server/pkg/lib/infra/db"
	"github.com/authgear/authgear-server/pkg/lib/infra/db/globaldb"
	"github.com/authgear/authgear-server/pkg/util/backgroundjob"
	"github.com/authgear/authgear-server/pkg/util/clock"
	"github.com/authgear/authgear-server/pkg/util/log"
)

// Injectors from wire.go:

func newRunnable(context2 context.Context, pool *db.Pool, globalDBCredentials *config.GlobalDatabaseCredentialsEnvironmentConfig, databaseCfg *config.DatabaseEnvironmentConfig, logFactory *log.Factory, clock2 clock.Clock, appContextResolver AppContextResolver, userServiceFactory UserServiceFactory) backgroundjob.Runnable {
	handle := globaldb.NewHandle(context2, pool, globalDBCredentials, databaseCfg, logFactory)
	sqlBuilder := globaldb.NewSQLBuilder(globalDBCredentials)
	sqlExecutor := globaldb.NewSQLExecutor(context2, handle)
	store := &Store{
		Handle:      handle,
		SQLBuilder:  sqlBuilder,
		SQLExecutor: sqlExecutor,
		Clock:       clock2,
	}
	runnableLogger := NewRunnableLogger(logFactory)
	runnable := &Runnable{
		Context:            context2,
		Store:              store,
		AppContextResolver: appContextResolver,
		UserServiceFactory: userServiceFactory,
		Logger:             runnableLogger,
	}
	return runnable
}

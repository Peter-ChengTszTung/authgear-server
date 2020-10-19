package portal

import (
	"github.com/google/wire"

	adminauthz "github.com/authgear/authgear-server/pkg/lib/admin/authz"
	"github.com/authgear/authgear-server/pkg/portal/db"
	"github.com/authgear/authgear-server/pkg/portal/deps"
	"github.com/authgear/authgear-server/pkg/portal/endpoint"
	"github.com/authgear/authgear-server/pkg/portal/graphql"
	"github.com/authgear/authgear-server/pkg/portal/loader"
	"github.com/authgear/authgear-server/pkg/portal/service"
	"github.com/authgear/authgear-server/pkg/portal/task"
	"github.com/authgear/authgear-server/pkg/portal/template"
	"github.com/authgear/authgear-server/pkg/portal/transport"
	"github.com/authgear/authgear-server/pkg/util/clock"
)

var DependencySet = wire.NewSet(
	deps.DependencySet,
	deps.TaskDependencySet,

	service.DependencySet,
	adminauthz.DependencySet,
	clock.DependencySet,
	db.DependencySet,
	template.DependencySet,
	endpoint.DependencySet,

	wire.Bind(new(service.AuthzAdder), new(*adminauthz.Adder)),
	wire.Bind(new(service.CollaboratorServiceTaskQueue), new(*task.InProcessQueue)),
	wire.Bind(new(service.CollaboratorServiceEndpointsProvider), new(*endpoint.EndpointsProvider)),

	loader.DependencySet,
	wire.Bind(new(loader.AppService), new(*service.AppService)),
	wire.Bind(new(loader.DomainService), new(*service.DomainService)),
	wire.Bind(new(loader.CollaboratorService), new(*service.CollaboratorService)),
	wire.Bind(new(loader.AuthzService), new(*service.AuthzService)),

	graphql.DependencySet,
	wire.Bind(new(graphql.ViewerLoader), new(*loader.ViewerLoader)),
	wire.Bind(new(graphql.AppLoader), new(*loader.AppLoader)),
	wire.Bind(new(graphql.DomainLoader), new(*loader.DomainLoader)),
	wire.Bind(new(graphql.CollaboratorLoader), new(*loader.CollaboratorLoader)),

	transport.DependencySet,
	wire.Bind(new(transport.AdminAPIService), new(*service.AdminAPIService)),
	wire.Bind(new(transport.AdminAPIAuthzService), new(*service.AuthzService)),
)

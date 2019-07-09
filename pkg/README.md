# Backend code structure

The backend is using mysql as storage. 

## Central components of the backend

| folder | description |
| ------- | ----------- |
| /pkg/http | Http Handlers and routing. |
| /pkg/infra | Packages in infra should be packages that are used in multiple places without knowing anything about the domain. E.g. Logs, Metrics, Traces. |
| /pkg/models | This is where we keep our domain model. This package should not depend on any package outside standard library.  |
| /pkg/registry | Service management package. |
| /pkg/services | Packages in services are responsible for peristing domain objects and manage the relationship between domain objects. |
| /pkg/setting | Anything related to global configuration should be dealt with in this package. |
| /pkg/storage | Where are database calls resides. |

## Testing

Tests uses standard library and `testify/assert`.

## Services/Repositories

Services should be self-contained and only talk to other parts using the repositories that have been made available through service registry. All services should register themselves to the `registry` package in an init function. Only registration should be done in the init function. Init functions should be avoided as much as possible.

When app starts all init functions within the services will be called and register themselves.
# User

### Available Operations

* [GetCollaborations](#getcollaborations) - Collaborations
* [GetCurrentUser](#getcurrentuser) - User Information
* [GetUser](#getuser) - User Information

## GetCollaborations

Provides the set of organizations of which a user is a member or a collaborator.

The set of organizations that a user can collaborate on is composed of:

* Organizations that the current user belongs to across VCS types (e.g. BitBucket, GitHub)
* The parent organization of repository that the user can collaborate on, but is not necessarily a member of
* The organization of the current user's account

### Example Usage

```go
package main

import(
	"context"
	"log"
	"CircleCi"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            APIKeyHeader: sdk.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.User.GetCollaborations(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.Collaborations != nil {
        // handle response
    }
}
```

## GetCurrentUser

Provides information about the user that is currently signed in.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"CircleCi"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            APIKeyHeader: sdk.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.User.GetCurrentUser(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.User != nil {
        // handle response
    }
}
```

## GetUser

Provides information about the user with the given ID.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"CircleCi"
	"CircleCi/pkg/models/operations"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            APIKeyHeader: sdk.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.User.GetUser(ctx, operations.GetUserRequest{
        ID: "ce6c5561-46c3-4e25-8fb0-08c42e141aac",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.User != nil {
        // handle response
    }
}
```

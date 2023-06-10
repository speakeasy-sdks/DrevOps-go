# Schedule

### Available Operations

* [CreateSchedule](#createschedule) - Create a schedule
* [DeleteScheduleByID](#deleteschedulebyid) - Delete a schedule
* [GetScheduleByID](#getschedulebyid) - Get a schedule
* [ListSchedulesForProject](#listschedulesforproject) - Get all schedules
* [UpdateSchedule](#updateschedule) - Update a schedule

## CreateSchedule

Creates a schedule and returns the created schedule.

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
    res, err := s.Schedule.CreateSchedule(ctx, operations.CreateScheduleRequest{
        RequestBody: &operations.CreateScheduleCreateScheduleParameters{
            AttributionActor: operations.CreateScheduleCreateScheduleParametersAttributionActorCurrent,
            Description: sdk.String("soluta"),
            Name: "Ted Kling",
            Parameters: map[string]interface{}{
                "necessitatibus": "distinctio",
                "asperiores": "nihil",
                "ipsum": "voluptate",
            },
            Timetable: operations.CreateScheduleCreateScheduleParametersTimetable2{
                DaysOfMonth: []int64{
                    263322,
                    137220,
                    20651,
                    229219,
                },
                DaysOfWeek: []CreateScheduleCreateScheduleParametersTimetable2DaysOfWeek{
                    operations.CreateScheduleCreateScheduleParametersTimetable2DaysOfWeekFri,
                    operations.CreateScheduleCreateScheduleParametersTimetable2DaysOfWeekSun,
                    operations.CreateScheduleCreateScheduleParametersTimetable2DaysOfWeekFri,
                    operations.CreateScheduleCreateScheduleParametersTimetable2DaysOfWeekSun,
                },
                HoursOfDay: []int64{
                    588317,
                    324683,
                    831049,
                },
                Months: []CreateScheduleCreateScheduleParametersTimetable2Months{
                    operations.CreateScheduleCreateScheduleParametersTimetable2MonthsApr,
                    operations.CreateScheduleCreateScheduleParametersTimetable2MonthsMar,
                    operations.CreateScheduleCreateScheduleParametersTimetable2MonthsJan,
                },
                PerHour: 311860,
            },
        },
        ProjectSlug: "tempora",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Schedule != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.CreateScheduleRequest](../../models/operations/createschedulerequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.CreateScheduleResponse](../../models/operations/createscheduleresponse.md), error**


## DeleteScheduleByID

Deletes the schedule by id.

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
    res, err := s.Schedule.DeleteScheduleByID(ctx, operations.DeleteScheduleByIDRequest{
        ScheduleID: "6ce2af7a-73cf-43be-853f-870b326b5a73",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.MessageResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.DeleteScheduleByIDRequest](../../models/operations/deleteschedulebyidrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.DeleteScheduleByIDResponse](../../models/operations/deleteschedulebyidresponse.md), error**


## GetScheduleByID

Get a schedule by id.

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
    res, err := s.Schedule.GetScheduleByID(ctx, operations.GetScheduleByIDRequest{
        ScheduleID: "429cdb1a-8422-4bb6-b9d2-322715bf0cbb",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Schedule != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetScheduleByIDRequest](../../models/operations/getschedulebyidrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.GetScheduleByIDResponse](../../models/operations/getschedulebyidresponse.md), error**


## ListSchedulesForProject

Returns all schedules for this project.

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
    res, err := s.Schedule.ListSchedulesForProject(ctx, operations.ListSchedulesForProjectRequest{
        PageToken: sdk.String("et"),
        ProjectSlug: "saepe",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListSchedulesForProject200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.ListSchedulesForProjectRequest](../../models/operations/listschedulesforprojectrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.ListSchedulesForProjectResponse](../../models/operations/listschedulesforprojectresponse.md), error**


## UpdateSchedule

Updates a schedule and returns the updated schedule.

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
    res, err := s.Schedule.UpdateSchedule(ctx, operations.UpdateScheduleRequest{
        RequestBody: &operations.UpdateScheduleUpdateScheduleParameters{
            AttributionActor: operations.UpdateScheduleUpdateScheduleParametersAttributionActorCurrent.ToPointer(),
            Description: sdk.String("ipsum"),
            Name: sdk.String("Gayle Lueilwitz"),
            Parameters: map[string]interface{}{
                "delectus": "dolorem",
            },
            Timetable: &operations.UpdateScheduleUpdateScheduleParametersTimetable{
                DaysOfMonth: []int64{
                    286915,
                    240829,
                },
                DaysOfWeek: []UpdateScheduleUpdateScheduleParametersTimetableDaysOfWeek{
                    operations.UpdateScheduleUpdateScheduleParametersTimetableDaysOfWeekTue,
                    operations.UpdateScheduleUpdateScheduleParametersTimetableDaysOfWeekTue,
                    operations.UpdateScheduleUpdateScheduleParametersTimetableDaysOfWeekTue,
                },
                HoursOfDay: []int64{
                    929530,
                    9240,
                    669917,
                },
                Months: []UpdateScheduleUpdateScheduleParametersTimetableMonths{
                    operations.UpdateScheduleUpdateScheduleParametersTimetableMonthsAug,
                    operations.UpdateScheduleUpdateScheduleParametersTimetableMonthsJul,
                    operations.UpdateScheduleUpdateScheduleParametersTimetableMonthsJun,
                    operations.UpdateScheduleUpdateScheduleParametersTimetableMonthsSep,
                },
                PerHour: sdk.Int64(586410),
            },
        },
        ScheduleID: "21879fce-953f-473e-b7fb-c7abd74dd39c",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Schedule != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.UpdateScheduleRequest](../../models/operations/updateschedulerequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.UpdateScheduleResponse](../../models/operations/updatescheduleresponse.md), error**


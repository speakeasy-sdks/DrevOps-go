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
            APIKeyHeader: sdk.String("YOUR_API_KEY_HERE"),
        }),
    )

    ctx := context.Background()
    res, err := s.Schedule.CreateSchedule(ctx, operations.CreateScheduleRequest{
        RequestBody: &operations.CreateScheduleCreateScheduleParameters{
            AttributionActor: operations.CreateScheduleCreateScheduleParametersAttributionActorEnumCurrent,
            Description: sdk.String("similique"),
            Name: "Cristina Hahn",
            Parameters: map[string]interface{}{
                "officiis": "qui",
                "dolorum": "a",
                "esse": "harum",
                "iusto": "ipsum",
            },
            Timetable: operations.CreateScheduleCreateScheduleParametersTimetable2{
                DaysOfMonth: []int64{
                    229442,
                    730856,
                    880298,
                    253941,
                },
                DaysOfWeek: []CreateScheduleCreateScheduleParametersTimetable2DaysOfWeekEnum{
                    operations.CreateScheduleCreateScheduleParametersTimetable2DaysOfWeekEnumSat,
                    operations.CreateScheduleCreateScheduleParametersTimetable2DaysOfWeekEnumFri,
                },
                HoursOfDay: []int64{
                    471752,
                    25662,
                    711584,
                },
                Months: []CreateScheduleCreateScheduleParametersTimetable2MonthsEnum{
                    operations.CreateScheduleCreateScheduleParametersTimetable2MonthsEnumNov,
                },
                PerHour: 424685,
            },
        },
        ProjectSlug: "libero",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Schedule != nil {
        // handle response
    }
}
```

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
            APIKeyHeader: sdk.String("YOUR_API_KEY_HERE"),
        }),
    )

    ctx := context.Background()
    res, err := s.Schedule.DeleteScheduleByID(ctx, operations.DeleteScheduleByIDRequest{
        ScheduleID: "5a73429c-db1a-4842-abb6-79d2322715bf",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.MessageResponse != nil {
        // handle response
    }
}
```

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
            APIKeyHeader: sdk.String("YOUR_API_KEY_HERE"),
        }),
    )

    ctx := context.Background()
    res, err := s.Schedule.GetScheduleByID(ctx, operations.GetScheduleByIDRequest{
        ScheduleID: "0cbb1e31-b8b9-40f3-843a-1108e0adcf4b",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Schedule != nil {
        // handle response
    }
}
```

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
            APIKeyHeader: sdk.String("YOUR_API_KEY_HERE"),
        }),
    )

    ctx := context.Background()
    res, err := s.Schedule.ListSchedulesForProject(ctx, operations.ListSchedulesForProjectRequest{
        PageToken: sdk.String("cupiditate"),
        ProjectSlug: "qui",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListSchedulesForProject200ApplicationJSONObject != nil {
        // handle response
    }
}
```

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
            APIKeyHeader: sdk.String("YOUR_API_KEY_HERE"),
        }),
    )

    ctx := context.Background()
    res, err := s.Schedule.UpdateSchedule(ctx, operations.UpdateScheduleRequest{
        RequestBody: &operations.UpdateScheduleUpdateScheduleParameters{
            AttributionActor: operations.UpdateScheduleUpdateScheduleParametersAttributionActorEnumCurrent.ToPointer(),
            Description: sdk.String("quae"),
            Name: sdk.String("Darren McClure"),
            Parameters: map[string]interface{}{
                "omnis": "quis",
                "ipsum": "delectus",
                "voluptate": "consectetur",
                "vero": "tenetur",
            },
            Timetable: &operations.UpdateScheduleUpdateScheduleParametersTimetable{
                DaysOfMonth: []int64{
                    941378,
                    715561,
                },
                DaysOfWeek: []UpdateScheduleUpdateScheduleParametersTimetableDaysOfWeekEnum{
                    operations.UpdateScheduleUpdateScheduleParametersTimetableDaysOfWeekEnumMon,
                    operations.UpdateScheduleUpdateScheduleParametersTimetableDaysOfWeekEnumThu,
                    operations.UpdateScheduleUpdateScheduleParametersTimetableDaysOfWeekEnumThu,
                    operations.UpdateScheduleUpdateScheduleParametersTimetableDaysOfWeekEnumFri,
                },
                HoursOfDay: []int64{
                    293020,
                    844550,
                },
                Months: []UpdateScheduleUpdateScheduleParametersTimetableMonthsEnum{
                    operations.UpdateScheduleUpdateScheduleParametersTimetableMonthsEnumDec,
                    operations.UpdateScheduleUpdateScheduleParametersTimetableMonthsEnumApr,
                    operations.UpdateScheduleUpdateScheduleParametersTimetableMonthsEnumAug,
                    operations.UpdateScheduleUpdateScheduleParametersTimetableMonthsEnumMar,
                },
                PerHour: sdk.Int64(974259),
            },
        },
        ScheduleID: "5d2cff7c-70a4-4562-ad43-6813f16d9f5f",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Schedule != nil {
        // handle response
    }
}
```

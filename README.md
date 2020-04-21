solar-staff-sdk-go
------------------

Golang SDK to [solar-staff API](https://help.solar-staff.com/ru/collections/681023-api-documentation#api-reference)

## Usage example

```go
package main

import (
    "log"

    "github.com/spacetab-io/solar-staff-sdk-go/config"
    "github.com/spacetab-io/solar-staff-sdk-go/pkg/repo/rest"
    "github.com/spacetab-io/solar-staff-sdk-go/pkg/service"
)

func main() {
    cfg := &config.SolarStaffConfig{
        URL:      "https://api.solar-staff.com/v1/",
        ClientID: "xxx",
        Salt:     "yyy",
    }
    taskID := uint64(123123123)
    
    r := rest.NewRepo(cfg)
    s := service.NewService(r)
    
    b, err := s.Balance()
    if err != nil {
        log.Fatal(err)
    }
    
    log.Printf("balance: %v", b)
    wl, err := s.WorkerList()
    if err != nil {
        log.Fatal(err)
    }
    
    log.Printf("worker list: %v", wl)
    
    t, err := s.TaskByID(taskID)
    if err != nil {
        log.Fatal(err)
    }
    
    log.Printf("task: %v", t)
}
```


solar-staff-sdk-go
------------------

Golang SDK to [solar-staff API](https://help.solar-staff.com/ru/collections/681023-api-documentation#api-reference)

## Methods implementation status

Info:
* [x] [balance](https://help.solar-staff.com/ru/articles/1601247-balance)
* [ ] [get_exchange_rate](https://help.solar-staff.com/ru/articles/1601278-get_exchange_rate)

Workers:
* [ ] [card_verify](https://help.solar-staff.com/ru/articles/1601252-card_verify)
* [ ] [card_status](https://help.solar-staff.com/ru/articles/1601259-card_status)
* [ ] [card_remove](https://help.solar-staff.com/ru/articles/1601240-card_remove)
* [ ] [cards_list](https://help.solar-staff.com/ru/articles/1601226-cards_list)

Payment:
* [x] [payout](https://help.solar-staff.com/ru/articles/1601285-payout)
* [ ] [payouts](https://help.solar-staff.com/ru/articles/3347018-payouts)
* [ ] [payout_epayments, payout_paypal, payout_webmoney, payout_qiwi](https://help.solar-staff.com/ru/articles/3345987-payout_epayments-payout_paypal-payout_webmoney-payout_qiwi)

Task:
* [x] [search](https://help.solar-staff.com/ru/articles/1601302-search)
* [ ] [change_status](https://help.solar-staff.com/ru/articles/1601309-change_status)

Transaction:
* [ ] [transactions](https://help.solar-staff.com/ru/articles/1601310-transactions)
* [ ] [transaction_status](https://help.solar-staff.com/ru/articles/1601317-transaction_status)

Worker:
* [ ] [worker_create](https://help.solar-staff.com/ru/articles/1601322-worker_create)
* [ ] [worker_update](https://help.solar-staff.com/ru/articles/2824899-worker_update)
* [ ] [worker_invite](https://help.solar-staff.com/ru/articles/2501349-worker_invite)
* [ ] [worker_find](https://help.solar-staff.com/ru/articles/1601324-worker_find)
* [ ] [worker_remove](https://help.solar-staff.com/ru/articles/1601349-worker_remove)
* [ ] [worker_balance](https://help.solar-staff.com/ru/articles/1601350-worker_balance)
* [ ] [worker_transactions](https://help.solar-staff.com/ru/articles/1601362-worker_transactions)
* [x] [worker_list](https://help.solar-staff.com/ru/articles/1601374-workers_list)
* [ ] [worker_terminal](https://help.solar-staff.com/ru/articles/2860181-worker_terminal)

## Usage example

```go
package main

import (
    "log"
    
    "github.com/spacetab-io/solar-staff-sdk-go/config"
    "github.com/spacetab-io/solar-staff-sdk-go/pkg/contracts"
    "github.com/spacetab-io/solar-staff-sdk-go/pkg/models"
    "github.com/spacetab-io/solar-staff-sdk-go/pkg/repo/rest"
    "github.com/spacetab-io/solar-staff-sdk-go/pkg/service"
)

func main() {
    cfg := &config.SolarStaffConfig{
        URL:      "https://demo-api.solar-staff.com/v1/",
        ClientID: "123",
        Salt:     "xxxxxxxxxxxxxxxxxx",
    }
    
    r := rest.NewRepo(cfg)
    s := service.NewService(r)
    
    b, err := s.Balance()
    if err != nil {
        log.Fatal(err)
    }
    
    log.Printf("balance: %+v", b)
    
    wl, err := s.WorkerList()
    if err != nil {
        log.Fatal(err)
    }
    
    log.Printf("worker list: %+v", wl)
    
    workerID := uint64(100001212)
    
    task := contracts.PaymentRequest{
        WorkerID:        workerID,
        CustomerTaskID:  "SOME_PROJECT-105",
        Currency:        models.CurrencyCode("RUB"),
        Amount:          202.50,
        TaskTitle:       "Testing some functional",
        TaskDescription: "1. Go to main page\n2. See button\n3. Push the button\n4. Get the result",
        TaskCategoryID:  1435,
        ProjectName:     "SOME PROJECT",
    }
    p, err := s.PaymentOnTask(task)
    if err != nil {
        log.Fatal(err)
    }
    
    log.Printf("payment data: %+v", p)
    
    t1, err := s.TaskByCustomerTaskID(task.CustomerTaskID)
    if err != nil {
        log.Fatal(err)
    }
    
    log.Printf("task by customer task id: %+v", t1)
    
    t2, err := s.TaskByID(t1.ID)
    if err != nil {
        log.Fatal(err)
    }
    
    log.Printf("task by task id: %+v", t2)
}
```


package models

/**
1  - Не подтверждена - Not confirmed
2 - В работе - In progress
3 - Ждет приемки - Pending acceptance
4 - Ждет оплату - Pending payment
5 - Завершена - Completed
6 - Отклонена исполнителем - Declined by contractor
8 - Отклонена заказчиком - Declined by customer
10 - Ошибка при выплате - Payment declined
11 - Ждет подтверждения отмены от фрилансера - Pending cancellation confirmation
*/
type TaskStatus struct {
	ID uint64
	Title
}

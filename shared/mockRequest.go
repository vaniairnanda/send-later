package shared

import (
	"errors"
	"fmt"
)

func MockBusinessServiceGetUserLogin(login string) error {
	if login == "WRONG_LOGIN" {
		return errors.New("you do not have access to complete this action")
	}
	return nil
}


func PublishEventDisbursementApply(item interface{}) {
	fmt.Printf("success push data into kafka for disbursement: %v", item)
	// message should be consumed by itself
	// loop each disbursement item
	// proceed to call TransactionService and others

}
func PublishEventApprovalExpired(item interface{}) {
	fmt.Printf("success push data into kafka for expired approval: %v", item)
}

func PublishEventApprovalReminder(item interface{}) {
	fmt.Printf("success push data into kafka for reminder approval: %v", item)
}
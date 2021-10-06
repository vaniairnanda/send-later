# send-later

Send-later is a feature that allows user to send batch disbursement at a certain scheduled date using CronJob. 
The feature will also send email to remind user to approve the batch disbursement on D-1 of scheduled date 
and mark it as expired if left unapproved after the scheduled date has passed. 
If the expiry took place, the user will have options to Schedule for a later date or Disburse immediately. 

## Getting Started

1. Copy `.env.example` to `.env` and configure the env
   
   ```
   cp .env.example .env
   ```

2. Create table 
   ```
   CREATE TABLE batch_disbursements (
      	id serial NOT NULL,
      	reference varchar NULL,
      	scheduled_date date NULL,
      	country_code varchar NOT NULL,
      	client_id int4 NOT NULL,
      	is_send_later bool NULL,
      	approved_at timestamp NULL,
      	total_uploaded_amount int8 NULL,
      	total_uploaded_count int8 NULL,
      	status varchar NOT NULL,
      	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
      	updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
      	reminder_sent_at timestamp(0) NULL,
      	CONSTRAINT batch_disbursements_pkey PRIMARY KEY (id)
      );

   CREATE TABLE disbursements (
   	id serial NOT NULL,
   	batch_disbursement_id int4 NOT NULL,
   	amount int8 NOT NULL,
   	bank_code varchar NOT NULL,
   	bank_account_name varchar NOT NULL,
   	bank_account_number varchar NOT NULL,
   	external_id varchar NOT NULL,
   	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
   	updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
   	description varchar NULL,
   	CONSTRAINT disbursements_pkey PRIMARY KEY (id)
   );
   ```

3. Run `go run .`


## Run tests

1. Run `go test -v ./...`


## Code Structure
Code structure is sorted by application flow from inbound to repository of each domain
- server.go
- cronjob.go
- api
    - [domain-name]
        - handler
        - repository
- app
        - [base-function]
- shared
    - [utility-func]
    - [external-request-func]
- mocks
    - api
        - [domain-name]
            - handler
            - repository
- config
    - kafka
    - [connection-func]       
- environment
    - [environment-func]      
- model 
    - constant
    - [data-access-object]
    - [data-transfer-object]
    - enum
        
## APIs 

1. Create Scheduled Batch Disbursement


sample curl
```
curl --location --request POST 'http://localhost:8080/batch_disbursements' \
   --header 'USER_LOGIN_KEY: random' \
   --header 'Content-Type: application/json' \
   --data-raw '{
       "reference": "demo_trial",
       "country_code": "Asia/Jakarta",
       "client_id": 5,
       "scheduled_date": "2021-10-07",
       "is_send_later": true,
       "disbursements": [
         {
           "external_id": "demo_123_1",
           "bank_code": "BCA",
           "bank_account_name": "Stanley Nguyen",
           "bank_account_number": "12345678",
           "description": "Reimbursement for pair of shoes 2222",
           "amount": 20000
           }
       ]
     }
   '
```

2. Approve Scheduled Disbursement and Reschedule/Disburse Immediately


sample curl

```
curl --location --request PATCH 'http://localhost:8080/batch_disbursements/approve/:batchId' \
--header 'Content-Type: application/json' \
--data-raw '{
    "is_instant_disbursement": false,
    "new_scheduled_date": "2021-10-07"
}'
```

# Jobs

1. Schedule Batch Disbursement

setup env 

```
SCHEDULED_BATCH_DISBURSEMENT=* * * * *
SCHEDULED_DISBURSEMENT_TIME=07:00
```
       
2. Send Email Approval Reminder 

setup env 

```
SEND_APPROVAL_REMINDER=* * * * *
```       

3. Mark Expired and Send Notification


setup env 

```
MARK_APPROVAL_EXPIRED=* * * * *
```
# send-later



## How to Run

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

   CREATE TABLE public.disbursements (
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
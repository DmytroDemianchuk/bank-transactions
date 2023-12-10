CREATE TABLE transactions
(
    id                  serial       not null unique,
	transaction_id      integer      not null default 0,
	request_id          integer      not null default 0,
	terminal_id         integer      not null default 0,
	partner_object_id   integer      not null default 0,
	amount_total        real         not null default 0.0,
	amount_original     real         not null default 0.0,
	commission_ps       real         not null default 0.0,
	commission_client   real         not null default 0.0,
	commission_provider real         not null default 0.0,
	date_input          timestamptz  not null default CURRENT_TIMESTAMP,
	date_post           timestamptz  not null default CURRENT_TIMESTAMP,
	status              varchar(255) not null default '',
	payment_type        varchar(255) not null default '',
	payment_number      varchar(255) not null default '',
	service_id          integer      not null default 0,
	service             varchar(255) not null default '',
	payee_id            integer      not null default 0,
	payee_name          varchar(255) not null default '',
	payee_bnank_mfo     integer      not null default 0,
	payee_bnank_account varchar(255) not null default '',
	payment_narrative   varchar(255) not null default ''
);

CREATE TABLE sourceCSV
(
    id                  serial       not null unique,
	transaction_id      integer      not null default 0,
	request_id          integer      not null default 0,
	terminal_id         integer      not null default 0,
	partner_object_id   integer      not null default 0,
	amount_total        real         not null default 0.0,
	amount_original     real         not null default 0.0,
	commission_ps       real         not null default 0.0,
	commission_client   real         not null default 0.0,
	commission_provider real         not null default 0.0,
	date_input          timestamptz  not null default CURRENT_TIMESTAMP,
	date_post           timestamptz  not null default CURRENT_TIMESTAMP,
	status              varchar(255) not null default '',
	payment_type        varchar(255) not null default '',
	payment_number      varchar(255) not null default '',
	service_id          integer      not null default 0,
	service             varchar(255) not null default '',
	payee_id            integer      not null default 0,
	payee_name          varchar(255) not null default '',
	payee_bnank_mfo     integer      not null default 0,
	payee_bnank_account varchar(255) not null default '',
	payment_narrative   varchar(255) not null default ''
);

INSERT INTO sourceCSV (transaction_id, request_id, terminal_id, partner_object_id, amount_total, amount_original, commission_ps, commission_client, commission_provider, date_input, date_post, status, payment_type, payment_number, service_id, service, payee_id, payee_name, payee_bnank_mfo, payee_bnank_account, payment_narrative) 
VALUES (1,10020,3506,1111,1.00,1.00,0.00,0.00,0.00,					'2023-08-12 11:25:27',			'2023-08-12 14:25:27',		'accepted','cash',					'PS16698205',13980,'Replenishment of cards',14232155,'Bank of America',254751, 'USD713451373919523', 'Transfer of funds according to the contract for the provision of services B12/77771 dated September 19.2007');
INSERT INTO sourceCSV (transaction_id, request_id, terminal_id, partner_object_id, amount_total, amount_original, commission_ps, commission_client, commission_provider, date_input, date_post, status, payment_type, payment_number, service_id, service, payee_id, payee_name, payee_bnank_mfo, payee_bnank_account, payment_narrative) 
VALUES (2,10030,3507,1111,1.00,1.00,0.00,0.00,0.00,					'2023-08-12 12:36:52',			'2023-08-12 15:36:53',		'accepted','cash',					'PS16698215',13990,'Replenishment of cards',14332255,'Goldman Sachs',255752, 'USD713461333619513', 'Transfer of funds according to the contract for the provision of services B12/77771 dated September 19.2007');
INSERT INTO sourceCSV (transaction_id, request_id, terminal_id, partner_object_id, amount_total, amount_original, commission_ps, commission_client, commission_provider, date_input, date_post, status, payment_type, payment_number, service_id, service, payee_id, payee_name, payee_bnank_mfo, payee_bnank_account, payment_narrative) 
VALUES (3,10040,3508,1111,3.00,3.00,0.00,0.00,-0.01,				'2023-08-17 9:53:43',			'2023-08-17 12:53:44',		'accepted','cash',					'PS16698225',14000,'Replenishment of cards',14432355,'Citigroup',256753, 'USD713471293319503', 'Transfer of funds according to the contract for the provision of services B12/77771 dated September 19.2007');
INSERT INTO sourceCSV (transaction_id, request_id, terminal_id, partner_object_id, amount_total, amount_original, commission_ps, commission_client, commission_provider, date_input, date_post, status, payment_type, payment_number, service_id, service, payee_id, payee_name, payee_bnank_mfo, payee_bnank_account, payment_narrative) 
VALUES (4,10050,3509,1111,115.00,115.00,0.08,0.00,-0.23,			'2023-08-17 11:22:54',			'2023-08-17 14:22:54',		'accepted','cash',				'PS16698235',14010,'Replenishment of cards',14532455,'Bank of America',257754, 'USD713481253019493', 'Transfer of funds according to the contract for the provision of services B12/77771 dated September 19.2007');
INSERT INTO sourceCSV (transaction_id, request_id, terminal_id, partner_object_id, amount_total, amount_original, commission_ps, commission_client, commission_provider, date_input, date_post, status, payment_type, payment_number, service_id, service, payee_id, payee_name, payee_bnank_mfo, payee_bnank_account, payment_narrative) 
VALUES (5,10060,3510,1111,343.00,343.00,0.24,0.00,-0.69,			'2023-08-23 8:48:32',			'2023-08-23 11:48:33',		'accepted','cash',				'PS16698245',14020,'Replenishment of cards',14632555,'Goldman Sachs',258755, 'USD713491212719483', 'Transfer of funds according to the contract for the provision of services B12/77771 dated September 19.2007');
INSERT INTO sourceCSV (transaction_id, request_id, terminal_id, partner_object_id, amount_total, amount_original, commission_ps, commission_client, commission_provider, date_input, date_post, status, payment_type, payment_number, service_id, service, payee_id, payee_name, payee_bnank_mfo, payee_bnank_account, payment_narrative) 
VALUES (6,10070,3511,1111,1799.00,1799.00,1.26,0.00,-3.60,			'2023-08-23 8:48:47',			'2023-08-23 11:48:47',		'accepted','cash',				'PS16698255',14030,'Replenishment of cards',14732655,'Citigroup',259756, 'USD713501172419473', 'Transfer of funds according to the contract for the provision of services B12/77771 dated September 19.2007');
INSERT INTO sourceCSV (transaction_id, request_id, terminal_id, partner_object_id, amount_total, amount_original, commission_ps, commission_client, commission_provider, date_input, date_post, status, payment_type, payment_number, service_id, service, payee_id, payee_name, payee_bnank_mfo, payee_bnank_account, payment_narrative) 
VALUES (7,10080,3512,1111,2499.00,2499.00,1.75,0.00,-5.00,			'2023-08-23 8:49:01',			'2023-08-23 11:49:02',		'accepted','cash',				'PS16698265',14040,'Replenishment of cards',14832755,'Goldman Sachs',260757, 'USD713511132119463', 'Transfer of funds according to the contract for the provision of services B12/77771 dated September 19.2007');
INSERT INTO sourceCSV (transaction_id, request_id, terminal_id, partner_object_id, amount_total, amount_original, commission_ps, commission_client, commission_provider, date_input, date_post, status, payment_type, payment_number, service_id, service, payee_id, payee_name, payee_bnank_mfo, payee_bnank_account, payment_narrative) 
VALUES (8,10090,3513,1111,99.00,99.00,0.07,0.00,-0.20,				'2023-08-23 9:00:30',			'2023-08-23 12:00:30',		'accepted','card',					'PS16698275',14050,'Replenishment of cards',14932855,'Goldman Sachs',261758, 'USD713521091819453', 'Transfer of funds according to the contract for the provision of services B12/77771 dated September 19.2007');
INSERT INTO sourceCSV (transaction_id, request_id, terminal_id, partner_object_id, amount_total, amount_original, commission_ps, commission_client, commission_provider, date_input, date_post, status, payment_type, payment_number, service_id, service, payee_id, payee_name, payee_bnank_mfo, payee_bnank_account, payment_narrative) 
VALUES (9,10100,3514,1111,1349.00,1349.00,0.94,0.00,-2.70,			'2023-08-23 9:00:44',			'2023-08-23 12:00:45',		'accepted','cash',				'PS16698285',14060,'Replenishment of cards',15032955,'Goldman Sachs',262759, 'USD713531051519443', 'Transfer of funds according to the contract for the provision of services B12/77771 dated September 19.2007');
INSERT INTO sourceCSV (transaction_id, request_id, terminal_id, partner_object_id, amount_total, amount_original, commission_ps, commission_client, commission_provider, date_input, date_post, status, payment_type, payment_number, service_id, service, payee_id, payee_name, payee_bnank_mfo, payee_bnank_account, payment_narrative) 
VALUES (10,10110,3515,1111,119.00,119.00,0.08,1.00,-0.24,			'2023-08-23 9:03:37',			'2023-08-23 12:03:38',		'accepted','cash',				'PS16698295',14070,'Replenishment of cards',15133055,'Bank of America',263760, 'USD713541011219433', 'Transfer of funds according to the contract for the provision of services B12/77771 dated September 19.2007');
INSERT INTO sourceCSV (transaction_id, request_id, terminal_id, partner_object_id, amount_total, amount_original, commission_ps, commission_client, commission_provider, date_input, date_post, status, payment_type, payment_number, service_id, service, payee_id, payee_name, payee_bnank_mfo, payee_bnank_account, payment_narrative) 
VALUES (11,10120,3516,1111,119.00,119.00,0.08,0.00,-0.24,			'2023-08-23 9:03:52',			'2023-08-23 12:03:52',		'accepted','cash',				'PS16698305',14080,'Replenishment of cards',15233155,'U.S. Bancorp',264761, 'USD713550970919423', 'Transfer of funds according to the contract for the provision of services B12/77771 dated September 19.2007');
INSERT INTO sourceCSV (transaction_id, request_id, terminal_id, partner_object_id, amount_total, amount_original, commission_ps, commission_client, commission_provider, date_input, date_post, status, payment_type, payment_number, service_id, service, payee_id, payee_name, payee_bnank_mfo, payee_bnank_account, payment_narrative) 
VALUES (12,10130,3517,1111,119.00,119.00,0.08,0.00,-0.24,			'2023-08-23 9:04:06',			'2023-08-23 12:04:07',		'accepted','cash',				'PS16698315',14090,'Replenishment of cards',15333255,'U.S. Bancorp',265762, 'USD713560930619413', 'Transfer of funds according to the contract for the provision of services B12/77771 dated September 19.2007');
INSERT INTO sourceCSV (transaction_id, request_id, terminal_id, partner_object_id, amount_total, amount_original, commission_ps, commission_client, commission_provider, date_input, date_post, status, payment_type, payment_number, service_id, service, payee_id, payee_name, payee_bnank_mfo, payee_bnank_account, payment_narrative) 
VALUES (13,10140,3518,1111,913.00,913.00,0.64,0.00,-1.83,			'2023-08-23 9:04:20',			'2023-08-23 12:04:21',		'accepted','cash',				'PS16698325',14100,'Replenishment of cards',15433355,'Citigroup',266763, 'USD713570890319403', 'Transfer of funds according to the contract for the provision of services B12/77771 dated September 19.2007');
INSERT INTO sourceCSV (transaction_id, request_id, terminal_id, partner_object_id, amount_total, amount_original, commission_ps, commission_client, commission_provider, date_input, date_post, status, payment_type, payment_number, service_id, service, payee_id, payee_name, payee_bnank_mfo, payee_bnank_account, payment_narrative) 
VALUES (14,10150,3519,1111,1349.00,1349.00,0.94,0.00,-2.70,			'2023-08-23 9:02:25',			'2023-08-23 12:02:26',		'accepted','cash',			'PS16698335',14110,'Replenishment of cards',15533455,'Citigroup',267764, 'USD713580850019393', 'Transfer of funds according to the contract for the provision of services B12/77771 dated September 19.2007');
INSERT INTO sourceCSV (transaction_id, request_id, terminal_id, partner_object_id, amount_total, amount_original, commission_ps, commission_client, commission_provider, date_input, date_post, status, payment_type, payment_number, service_id, service, payee_id, payee_name, payee_bnank_mfo, payee_bnank_account, payment_narrative) 
VALUES (15,10160,3520,1111,770.00,770.00,0.54,0.00,-1.54,			'2023-08-23 8:59:41',			'2023-08-23 11:59:42',		'accepted','cash',				'PS16698345',14120,'Replenishment of cards',15633555,'U.S. Bancorp',268765, 'USD713590809719383', 'Transfer of funds according to the contract for the provision of services B12/77771 dated September 19.2007');
INSERT INTO sourceCSV (transaction_id, request_id, terminal_id, partner_object_id, amount_total, amount_original, commission_ps, commission_client, commission_provider, date_input, date_post, status, payment_type, payment_number, service_id, service, payee_id, payee_name, payee_bnank_mfo, payee_bnank_account, payment_narrative) 
VALUES (16,10170,3521,1111,120.00,120.00,0.08,0.00,-0.24,			'2023-08-23 8:59:55',			'2023-08-23 11:59:56',		'accepted','cash',				'PS16698355',14130,'Replenishment of cards',15733655,'Goldman Sachs',269766, 'USD713600769419373', 'Transfer of funds according to the contract for the provision of services B12/77771 dated September 19.2007');
INSERT INTO sourceCSV (transaction_id, request_id, terminal_id, partner_object_id, amount_total, amount_original, commission_ps, commission_client, commission_provider, date_input, date_post, status, payment_type, payment_number, service_id, service, payee_id, payee_name, payee_bnank_mfo, payee_bnank_account, payment_narrative) 
VALUES (17,10180,3522,1111,120.00,120.00,0.08,0.00,-0.24,			'2023-08-23 8:58:02',			'2023-08-23 11:58:02',		'accepted','cash',				'PS16698365',14140,'Replenishment of cards',15833755,'Bank of America',270767, 'USD713610729119363', 'Transfer of funds according to the contract for the provision of services B12/77771 dated September 19.2007');
INSERT INTO sourceCSV (transaction_id, request_id, terminal_id, partner_object_id, amount_total, amount_original, commission_ps, commission_client, commission_provider, date_input, date_post, status, payment_type, payment_number, service_id, service, payee_id, payee_name, payee_bnank_mfo, payee_bnank_account, payment_narrative) 
VALUES (18,10190,3523,1111,120.00,120.00,0.08,0.00,-0.24,			'2023-08-23 8:58:16',			'2023-08-23 11:58:16',		'accepted','cash',				'PS16698375',14150,'Replenishment of cards',15933855,'U.S. Bancorp',271768, 'USD713620688819353', 'Transfer of funds according to the contract for the provision of services B12/77771 dated September 19.2007');
INSERT INTO sourceCSV (transaction_id, request_id, terminal_id, partner_object_id, amount_total, amount_original, commission_ps, commission_client, commission_provider, date_input, date_post, status, payment_type, payment_number, service_id, service, payee_id, payee_name, payee_bnank_mfo, payee_bnank_account, payment_narrative) 
VALUES (19,10200,3524,1111,120.00,120.00,0.08,0.00,-0.24,			'2023-08-23 8:58:30',			'2023-08-23 11:58:30',		'accepted','cash',				'PS16698385',14160,'Replenishment of cards',16033955,'Citigroup',272769, 'USD713630648519343', 'Transfer of funds according to the contract for the provision of services B12/77771 dated September 19.2007');

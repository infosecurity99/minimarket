
/*branch*/
CREATE TABLE branch (
  id UUID PRIMARY KEY,
  name VARCHAR(30),
  address VARCHAR(30)
);

/*sale*/
CREATE TABLE sale (
  id UUID PRIMARY KEY,
  branch_id UUID REFERENCES branch(id),
  shopassistant_id UUID REFERENCES staff(id),
  cashier_id UUID REFERENCES staff(id),
  payment_type payment_type_enum,
  price INT,
  status_type  status_enum,
  clientname VARCHAR(30)
);
create type payment_type_enum as enum ( 'shopassistant', 'cashier');

/*transaction*/
CREATE TABLE transaction (
  id UUID PRIMARY KEY,
  sale_id UUID REFERENCES sale(id),
  staff_id UUID REFERENCES staff(id),
  transaction tarnsaction_type_enum,
  sourcetype source_type_enum,
  amount INT,
  description VARCHAR(30)
);

create type tarnsaction_type_enum as enum ( );

create type source_type_enum as enum ( 'bonus'  ,'sales');

/*staff*/
CREATE TABLE staff (
  id UUID PRIMARY KEY,
  branch_id UUID REFERENCES branch(id),
  tarif_id UUID REFERENCES staff_tarif(id),
  type_stuff type_stuf_enum,
  name VARCHAR(30),
  balance VARCHAR(30),
  age INT,
  birthdate INT,
  login VARCHAR(30),
  password VARCHAR(30)
);

create type type_stuf_enum as enum ( 'withdraw'  ,'topup');

/*stafftarif*/
CREATE TABLE staff_tarif (
  id UUID PRIMARY KEY,
  name VARCHAR(30),
  tarif_type tarif_type_enum,
  amount_for_cashe INT,
  amount_for_card INT
);
create type tarif_type_enum as enum ( 'percent'  ,'fixed');
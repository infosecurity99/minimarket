/*branch*/
CREATE TABLE branch (
id UUID PRIMARY KEY,
name VARCHAR(30),
address VARCHAR(30),
create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TYPE status_enum AS ENUM ('In process', 'Success', 'Cancel');
/*sale*/
CREATE TABLE sale (
 id UUID PRIMARY KEY,
 branch_id UUID REFERENCES branch(id),
 shopassistant_id UUID REFERENCES staff(id),
 cashier_id UUID REFERENCES staff(id),
 payment_type payment_type_enum,
 price INT,
 status_type  status_enum,
 clientname VARCHAR(30),
 create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
create type payment_type_enum as enum ('card', 'cashe');

/*transaction*/
CREATE TABLE transaction (
id UUID PRIMARY KEY,
sale_id UUID REFERENCES sale(id),
staff_id UUID REFERENCES staff(id),
transaction tarnsaction_type_enum,
sourcetype source_type_enum,
amount INT,
description VARCHAR(30),
create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

create type tarnsaction_type_enum as enum ('withdraw'  ,'topup' );

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
password VARCHAR(30),
create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

create type type_stuf_enum as enum ( 'shopassistant', 'cashier');

/*stafftarif*/
CREATE TABLE staff_tarif (
id UUID PRIMARY KEY,
name VARCHAR(30),
tarif_type tarif_type_enum,
amount_for_cashe INT,
amount_for_card INT,
create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
create type tarif_type_enum as enum ( 'percent'  ,'fixed');








/*storage*/
CREATE TABLE storage (
 id UUID PRIMARY KEY,
 product_id UUID REFERENCES product(id),
 branch_id UUID REFERENCES branch(id),
 count INT
);

/*product*/
CREATE TABLE product (
 id UUID PRIMARY KEY,
 name VARCHAR(30),
 price INT,
 barcode SERIAL,
 category_id UUID REFERENCES category(id)
);

-- Barcode'i otomatik olarak oluşturacak bir fonksiyon
CREATE OR REPLACE FUNCTION generate_barcode()
 RETURNS TRIGGER AS $$
BEGIN
 NEW.barcode = LPAD(NEW.id::TEXT, 5, '0');
 RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Fonksiyonu tetikleyici olarak ekleyin
CREATE TRIGGER product_before_insert
 BEFORE INSERT ON product
 FOR EACH ROW EXECUTE FUNCTION generate_barcode();

/*category*/
CREATE TABLE category (
  id UUID PRIMARY KEY,
  name VARCHAR(30),
  parent_id UUID
);

/*storage transaction*/
CREATE TABLE storage_transaction (
 id UUID PRIMARY KEY,
 branch_id UUID REFERENCES branch(id),
 staff_id UUID REFERENCES staff(id),
 product_id UUID REFERENCES product(id),
 transaction_type storage_transaction_type_enum,
 price INT,
 quantity INT
);


CREATE TYPE storage_transaction_type_enum AS ENUM ('minus', 'plus');
/*basket*/

CREATE TABLE basket (
sale_id UUID REFERENCES sale(id),
product_id UUID REFERENCES product(id),
quantity INT,
price INT
);
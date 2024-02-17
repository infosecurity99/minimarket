-- ENUM tipleri
CREATE TYPE payment_type_enum AS ENUM ('card', 'cash');
CREATE TYPE status_enum AS ENUM ('InProgress', 'Success', 'Cancel');
CREATE TYPE transaction_type_enum AS ENUM ('withdraw', 'topup');
CREATE TYPE source_type_enum AS ENUM ('bonus', 'sales');
CREATE TYPE tarif_type_enum AS ENUM ('percent', 'fixed');
CREATE TYPE type_stuf_enum AS ENUM ('ShopAssistant', 'Cashier');
CREATE TYPE storage_transaction_type_enum AS ENUM ('minus', 'plus');

-- Branch tablosu
CREATE TABLE branch (
                        id UUID PRIMARY KEY,
                        name VARCHAR(255),
                        address VARCHAR(255)
);

-- Staff Tarif tablosu
CREATE TABLE staff_tarif (
                             id UUID PRIMARY KEY,
                             name VARCHAR(255),
                             tarif_type tarif_type_enum,
                             amount_for_cash INT,
                             amount_for_card INT

);

-- Staff tablosu
CREATE TABLE staff (
                       id UUID PRIMARY KEY,
                       branch_id UUID REFERENCES branch(id),
                       tarif_id UUID REFERENCES staff_tarif(id),
                       type_stuff type_stuf_enum,
                       name VARCHAR(255),
                       balance INT,
                       age INT,
                       birthdate DATE ,
                       login VARCHAR(255),
                       password VARCHAR(255)

);

-- Category tablosu
CREATE TABLE category (
                          id UUID PRIMARY KEY,
                          name VARCHAR(255),
                          parent_id UUID
 
);

-- Product tablosu
CREATE TABLE product (
                         id UUID PRIMARY KEY,
                         name VARCHAR(255),
                         price INT,
                         barcode SERIAL,
                         category_id UUID REFERENCES category(id)

);

-- Storage Transaction tablosu
CREATE TABLE storage_transaction (
                                     id UUID PRIMARY KEY,
                                     branch_id UUID REFERENCES branch(id),
                                     staff_id UUID REFERENCES staff(id),
                                     product_id UUID REFERENCES product(id),
                                     transaction_type storage_transaction_type_enum,
                                     price INT,
                                     quantity INT

);

-- Storage tablosu
CREATE TABLE storage (
                         id UUID PRIMARY KEY,
                         product_id UUID REFERENCES product(id),
                         branch_id UUID REFERENCES branch(id),
                         count INT

);

-- Sale tablosu
CREATE TABLE sale (
    id UUID PRIMARY KEY,
    branch_id UUID REFERENCES branch(id),
    shopassistant_id UUID REFERENCES staff(id),
    cashier_id UUID REFERENCES staff(id),
    payment_type payment_type_enum,
    price INT,
    status_type status_enum DEFAULT 'Success',
    clientname VARCHAR(255)

);


-- Transaction tablosu
CREATE TABLE transactions (
                             id UUID PRIMARY KEY,
                             sale_id UUID REFERENCES sale(id),
                             staff_id UUID REFERENCES staff(id),
                             transaction_type transaction_type_enum,
                             sourcetype source_type_enum,
                             amount INT,
                             description VARCHAR(255)
                      
);

-- Basket tablosu
CREATE TABLE basket (
                        id UUID PRIMARY KEY,
                        sale_id UUID REFERENCES sale(id),
                        product_id UUID REFERENCES product(id),
                        quantity INT,
                        price INT

);

-- Trigger fonksiyonu
CREATE OR REPLACE FUNCTION before_product_insert()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.barcode := floor(random() * 899999 + 100000);
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger
CREATE TRIGGER before_product_insert_trigger
    BEFORE INSERT ON product
    FOR EACH ROW
EXECUTE FUNCTION before_product_insert();




alter table branch
    add column if not exists created_at timestamp default now(),
    add column if not exists updated_at timestamp,
    add column if not exists deleted_at integer default 0;

alter table staff_tarif
    add column if not exists created_at timestamp default now(),
    add column if not exists updated_at timestamp,
    add column if not exists deleted_at integer default 0;

alter table staff
    add column if not exists created_at timestamp default now(),
    add column if not exists updated_at timestamp,
    add column if not exists deleted_at integer default 0;

alter table category
    add column if not exists created_at timestamp default now(),
    add column if not exists updated_at timestamp,
    add column if not exists deleted_at integer default 0;

alter table product
    add column if not exists created_at timestamp default now(),
    add column if not exists updated_at timestamp,
    add column if not exists deleted_at integer default 0;

alter table storage_transaction
    add column if not exists created_at timestamp default now(),
    add column if not exists updated_at timestamp,
    add column if not exists deleted_at integer default 0;
alter table storage
    add column if not exists created_at timestamp default now(),
    add column if not exists updated_at timestamp,
    add column if not exists deleted_at integer default 0;

alter table sale
    add column if not exists created_at timestamp default now(),
    add column if not exists updated_at timestamp,
    add column if not exists deleted_at integer default 0;

alter table basket
    add column if not exists created_at timestamp default now(),
    add column if not exists updated_at timestamp,
    add column if not exists deleted_at integer default 0;

alter table transactions
    add column if not exists created_at timestamp default now(),
    add column if not exists updated_at timestamp,
    add column if not exists deleted_at integer default 0;
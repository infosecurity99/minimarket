
CREATE TYPE payment_type_enum AS ENUM ('card', 'cash');
CREATE TYPE status_enum AS ENUM ('In process', 'Success', 'Cancel');
CREATE TYPE transaction_type_enum AS ENUM ('withdraw', 'topup');
CREATE TYPE source_type_enum AS ENUM ('bonus', 'sales');
CREATE TYPE tarif_type_enum AS ENUM ('percent', 'fixed');
CREATE TYPE type_stuf_enum AS ENUM ('Shop Assistant', 'Cashier');
CREATE TYPE storage_transaction_type_enum AS ENUM ('minus', 'plus');

CREATE TABLE branch (
                        id UUID PRIMARY KEY,
                        name VARCHAR(30),
                        address VARCHAR(30),
                        create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
                        
);

CREATE TABLE staff_tarif (
                             id UUID PRIMARY KEY,
                             name VARCHAR(30),
                             tarif_type tarif_type_enum,
                             amount_for_cash INT,
                             amount_for_card INT,
                             create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
                             
);

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

CREATE TABLE category (
                          id UUID PRIMARY KEY,
                          name VARCHAR(30),
                          parent_id UUID,
                         create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE product (
                         id UUID PRIMARY KEY,
                         name VARCHAR(30),
                         price INT,
                         barcode SERIAL,
                         category_id UUID REFERENCES category(id),
                         create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE storage_transaction (
                                     id UUID PRIMARY KEY,
                                     branch_id UUID REFERENCES branch(id),
                                     staff_id UUID REFERENCES staff(id),
                                     product_id UUID REFERENCES product(id),
                                     transaction_type storage_transaction_type_enums,
                                     price INT,
                                     quantity INT,
                                     create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE storage (
                         id UUID PRIMARY KEY,
                         product_id UUID REFERENCES product(id),
                         branch_id UUID REFERENCES branch(id),
                         count INT,
                             create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE sale (
                      id UUID PRIMARY KEY,
                      branch_id UUID REFERENCES branch(id),
                      shopassistant_id UUID REFERENCES staff(id),
                      cashier_id UUID REFERENCES staff(id),
                      payment_type payment_type_enum,
                      price INT,
                      status_type status_enum,
                      clientname VARCHAR(30),
                      create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
                      
);

CREATE TABLE transaction (
                             id UUID PRIMARY KEY,
                             sale_id UUID REFERENCES sale(id),
                             staff_id UUID REFERENCES staff(id),
                             transaction_type transaction_type_enum,
                             sourcetype source_type_enum,
                             amount INT,
                             description VARCHAR(30),
                             create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
                             
);

CREATE TABLE basket (
                       id uuid PRIMARY KEY,
                        sale_id UUID REFERENCES sale(id),
                        product_id UUID REFERENCES product(id),
                        quantity INT,
                        price INT,
                        create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
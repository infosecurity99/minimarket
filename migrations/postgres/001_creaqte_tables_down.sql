drop table if exists branch;

drop table if exists staff_tarif;

drop table if exists staff;

drop table if  exists category;

drop table if exists product;

drop  table if exists storage_transaction;

drop table if exists storage;

drop  table if exists sale;

drop table if exists basket;

drop table if exists   transactions;


drop table if  exists payment_type_enum;

drop table if exists status_enum;

drop table if exists transaction_type_enum;

drop table if exists source_type_enum;

drop table if exists tarif_type_enum;

drop table if exists type_stuf_enum;

drop table if exists storage_transaction_type_enum;




alter table branch
    drop column if ebranchxists created_at,
    drop column if exists updated_at,
    drop column if exists deleted_at;

alter table staff_tarif
    drop column if exists created_at,
    drop column if exists updated_at,
    drop column if exists deleted_at;

alter table staff
    drop column if exists created_at,
    drop column if exists updated_at,
    drop column if exists deleted_at;

alter table category
    drop column if exists created_at,
    drop column if exists updated_at,
    drop column if exists deleted_at;

alter table product
    drop column if exists created_at,
    drop column if exists updated_at,
    drop column if exists deleted_at;

alter table storage_transaction
    drop column if exists created_at,
    drop column if exists updated_at,
    drop column if exists deleted_at;

alter table storage
    drop column if exists created_at,
    drop column if exists updated_at,
    drop column if exists deleted_at;

alter table sale
    drop column if exists created_at,
    drop column if exists updated_at,
    drop column if exists deleted_at;

alter table basket
    drop column if exists created_at,
    drop column if exists updated_at,
    drop column if exists deleted_at;

alter table transactions
    drop column if exists created_at,
    drop column if exists updated_at,
    drop column if exists deleted_at;
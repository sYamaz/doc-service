-- DB作成

CREATE DATABASE doc_db;

-- DB切り替え
\c doc_db

-- schema作成
CREATE SCHEMA user_schema; -- 一般ユーザー用

-- role作成
CREATE ROLE user; -- 一般ユーザー

-- schemaへのアクセス権限をロールに割り当てる
GRANT ALL PRIVILEGES ON SCHEMA user_schema TO user;
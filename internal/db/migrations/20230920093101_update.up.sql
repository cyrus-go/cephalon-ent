-- create index "wallet_user_id_symbol_id_deleted_at" to table: "wallets"
CREATE UNIQUE INDEX "wallet_user_id_symbol_id_deleted_at" ON "wallets" ("user_id", "symbol_id", "deleted_at");

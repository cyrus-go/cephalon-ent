-- modify "invites" table
ALTER TABLE "invites" ADD COLUMN "first_recharge_cep" bigint NOT NULL DEFAULT 0;
-- set comment to column: "first_recharge_cep" on table: "invites"
COMMENT ON COLUMN "invites" ."first_recharge_cep" IS '通过此邀请码邀请用户注册并首次充值能获得的收益';

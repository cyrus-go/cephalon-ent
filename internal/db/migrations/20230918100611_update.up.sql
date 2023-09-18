-- modify "invites" table
ALTER TABLE "invites" ADD COLUMN "first_reg_cep" bigint NOT NULL DEFAULT 0;
-- set comment to column: "first_reg_cep" on table: "invites"
COMMENT ON COLUMN "invites" ."first_reg_cep" IS '通过此邀请码邀请用户注册并首次充值能获得的收益';

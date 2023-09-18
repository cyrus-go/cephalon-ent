-- reverse: set comment to column: "first_recharge_cep" on table: "invites"
COMMENT ON COLUMN "invites" ."first_recharge_cep" IS '';
-- reverse: modify "invites" table
ALTER TABLE "invites" DROP COLUMN "first_recharge_cep";

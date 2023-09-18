-- reverse: set comment to column: "first_reg_cep" on table: "invites"
COMMENT ON COLUMN "invites" ."first_reg_cep" IS '';
-- reverse: modify "invites" table
ALTER TABLE "invites" DROP COLUMN "first_reg_cep";

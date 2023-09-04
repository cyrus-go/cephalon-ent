-- reverse: set comment to column: "temp_hmac_secret" on table: "missions"
COMMENT ON COLUMN "missions" ."temp_hmac_secret" IS '';
-- reverse: set comment to column: "temp_hmac_key" on table: "missions"
COMMENT ON COLUMN "missions" ."temp_hmac_key" IS '';
-- reverse: set comment to column: "inner_method" on table: "missions"
COMMENT ON COLUMN "missions" ."inner_method" IS '';
-- reverse: set comment to column: "inner_api" on table: "missions"
COMMENT ON COLUMN "missions" ."inner_api" IS '';
-- reverse: set comment to column: "resp_body" on table: "missions"
COMMENT ON COLUMN "missions" ."resp_body" IS '返回内容体 json 序列化为 string';
-- reverse: set comment to column: "unit_cep" on table: "missions"
COMMENT ON COLUMN "missions" ."unit_cep" IS '任务单价，按次就是 unit_cep/次，按时就是 unit_cep/分钟';
-- reverse: modify "missions" table
ALTER TABLE "missions" DROP COLUMN "temp_hmac_secret", DROP COLUMN "temp_hmac_key", DROP COLUMN "inner_method", DROP COLUMN "inner_api";

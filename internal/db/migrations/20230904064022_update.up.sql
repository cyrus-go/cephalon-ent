-- modify "missions" table
ALTER TABLE "missions" ADD COLUMN "inner_api" character varying NOT NULL DEFAULT '', ADD COLUMN "inner_method" character varying NOT NULL DEFAULT 'POST', ADD COLUMN "temp_hmac_key" character varying NOT NULL DEFAULT '', ADD COLUMN "temp_hmac_secret" character varying NOT NULL DEFAULT '';
-- set comment to column: "unit_cep" on table: "missions"
COMMENT ON COLUMN "missions" ."unit_cep" IS '任务单价，按次(count)就是 unit_cep/次，按时(time)就是 unit_cep/分钟';
-- set comment to column: "resp_body" on table: "missions"
COMMENT ON COLUMN "missions" ."resp_body" IS '返回内容体 json 转 string';
-- set comment to column: "inner_api" on table: "missions"
COMMENT ON COLUMN "missions" ."inner_api" IS '当 type 为 sd_api 时使用，为转发的 sd 内部接口路径';
-- set comment to column: "inner_method" on table: "missions"
COMMENT ON COLUMN "missions" ."inner_method" IS '内部转发接口的请求方式，POST 或者 GET 等';
-- set comment to column: "temp_hmac_key" on table: "missions"
COMMENT ON COLUMN "missions" ."temp_hmac_key" IS '当 type 为 key_pair 时，使用的临时密钥对的键';
-- set comment to column: "temp_hmac_secret" on table: "missions"
COMMENT ON COLUMN "missions" ."temp_hmac_secret" IS '当 type 为 key_pair 时，使用的临时密钥对的值';

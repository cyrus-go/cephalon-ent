-- create "artworks" table
CREATE TABLE "artworks" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "name" character varying NOT NULL DEFAULT '', "url" character varying NOT NULL DEFAULT '', "author_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- set comment to table: "artworks"
COMMENT ON TABLE "artworks" IS '艺术作品，参与投票等逻辑；Artwork';
-- set comment to column: "id" on table: "artworks"
COMMENT ON COLUMN "artworks" ."id" IS '19 位雪花 ID';
-- set comment to column: "created_by" on table: "artworks"
COMMENT ON COLUMN "artworks" ."created_by" IS '创建者 ID';
-- set comment to column: "updated_by" on table: "artworks"
COMMENT ON COLUMN "artworks" ."updated_by" IS '更新者 ID';
-- set comment to column: "created_at" on table: "artworks"
COMMENT ON COLUMN "artworks" ."created_at" IS '创建时刻，带时区';
-- set comment to column: "updated_at" on table: "artworks"
COMMENT ON COLUMN "artworks" ."updated_at" IS '更新时刻，带时区';
-- set comment to column: "deleted_at" on table: "artworks"
COMMENT ON COLUMN "artworks" ."deleted_at" IS '软删除时刻，带时区';
-- set comment to column: "name" on table: "artworks"
COMMENT ON COLUMN "artworks" ."name" IS '作品名称';
-- set comment to column: "url" on table: "artworks"
COMMENT ON COLUMN "artworks" ."url" IS '作品链接';
-- set comment to column: "author_id" on table: "artworks"
COMMENT ON COLUMN "artworks" ."author_id" IS '作者的用户 id';
-- create "artwork_likes" table
CREATE TABLE "artwork_likes" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "date" integer NOT NULL DEFAULT 0, "artwork_id" bigint NOT NULL DEFAULT 0, "user_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- set comment to table: "artwork_likes"
COMMENT ON TABLE "artwork_likes" IS '用户投赞成票给艺术作品；ArtworkLike';
-- set comment to column: "id" on table: "artwork_likes"
COMMENT ON COLUMN "artwork_likes" ."id" IS '19 位雪花 ID';
-- set comment to column: "created_by" on table: "artwork_likes"
COMMENT ON COLUMN "artwork_likes" ."created_by" IS '创建者 ID';
-- set comment to column: "updated_by" on table: "artwork_likes"
COMMENT ON COLUMN "artwork_likes" ."updated_by" IS '更新者 ID';
-- set comment to column: "created_at" on table: "artwork_likes"
COMMENT ON COLUMN "artwork_likes" ."created_at" IS '创建时刻，带时区';
-- set comment to column: "updated_at" on table: "artwork_likes"
COMMENT ON COLUMN "artwork_likes" ."updated_at" IS '更新时刻，带时区';
-- set comment to column: "deleted_at" on table: "artwork_likes"
COMMENT ON COLUMN "artwork_likes" ."deleted_at" IS '软删除时刻，带时区';
-- set comment to column: "date" on table: "artwork_likes"
COMMENT ON COLUMN "artwork_likes" ."date" IS '投票日期';
-- set comment to column: "artwork_id" on table: "artwork_likes"
COMMENT ON COLUMN "artwork_likes" ."artwork_id" IS '外键，作品 id';
-- set comment to column: "user_id" on table: "artwork_likes"
COMMENT ON COLUMN "artwork_likes" ."user_id" IS '外键，用户 id';

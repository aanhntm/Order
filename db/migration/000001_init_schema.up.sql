CREATE TABLE "orders" (
	"id" int PRIMARY KEY,
	"user_name" varchar,
	"product_name" varchar,
	"amount" int NOT NULL
);

CREATE INDEX ON "orders" ("id");
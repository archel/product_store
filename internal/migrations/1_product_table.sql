-- up
CREATE TABLE "products" (
 "id" uuid PRIMARY KEY,
 "price" DECIMAL(10, 2) NOT NULL,
 "description" VARCHAR(200) NOT NULL,
 "creation" timestamp with time zone NOT NULL
)

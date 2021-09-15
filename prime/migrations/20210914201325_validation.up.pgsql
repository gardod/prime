CREATE TABLE "validation" (
    "id" BIGINT GENERATED ALWAYS AS IDENTITY,
    "number" BIGINT NOT NULL,
    "is_prime" BOOLEAN NOT NULL,
    "started_at" TIMESTAMP WITH TIME ZONE NOT NULL,
    "duration_in_μs" INT NOT NULL
);

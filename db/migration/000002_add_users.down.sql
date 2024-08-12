-- Remove the foreign key constraint from the account table
ALTER TABLE "account" DROP CONSTRAINT "account_owner_fkey";

-- Drop the unique constraint on the owner and currency columns
ALTER TABLE "account" DROP CONSTRAINT "owner_currency_key";

-- Drop the users table
DROP TABLE "users";

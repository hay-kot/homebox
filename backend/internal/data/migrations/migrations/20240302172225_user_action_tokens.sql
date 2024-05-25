-- Create "action_tokens" table
CREATE TABLE `action_tokens` (`id` uuid NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `action` text NOT NULL DEFAULT ('reset_password'), `token` blob NOT NULL, `user_id` uuid NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `action_tokens_users_action_tokens` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE);
-- Create index "action_tokens_token_key" to table: "action_tokens"
CREATE UNIQUE INDEX `action_tokens_token_key` ON `action_tokens` (`token`);
-- Create index "actiontoken_token" to table: "action_tokens"
CREATE INDEX `actiontoken_token` ON `action_tokens` (`token`);
-- Create index "actiontoken_action" to table: "action_tokens"
CREATE INDEX `actiontoken_action` ON `action_tokens` (`action`);
-- Create index "actiontoken_user_id" to table: "action_tokens"
CREATE INDEX `actiontoken_user_id` ON `action_tokens` (`user_id`);

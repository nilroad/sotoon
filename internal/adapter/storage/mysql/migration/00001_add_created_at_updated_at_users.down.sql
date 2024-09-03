begin;

alter table `users` drop column `created_at`;
alter table `users` drop column `updated_at`;

commit;
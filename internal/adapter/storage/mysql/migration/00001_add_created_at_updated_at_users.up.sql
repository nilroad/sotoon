begin;

alter table `users` add column `created_at` timestamp not null;
alter table `users` add column `updated_at` timestamp not null;

commit;
alter table billing.tariff alter column created_at set default now();
alter table billing.tariff alter column updated_at set default now();
alter table billing.tariff alter column id set default uuid_generate_v4();

alter table billing.snapshots alter column created_at set default now();
alter table billing.snapshots alter column updated_at set default now();

alter table billing.events
	add created_at timestamp default now();

alter table billing.account alter column id set default uuid_generate_v4();

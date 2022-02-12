### Commands


#### Migrations
- Create migration: `NODE_ENV=environment knex migrate:make migration_name -x ts`
- Run migrations: `NODE_ENV=environment knex migrate:latest`
- Rollback migration: `NODE_ENV=environment knex migrate:rollback`

import { Knex } from "knex";


export async function up(knex: Knex): Promise<void> {
    if (await knex.schema.hasTable('users')) {
        return;
    }
    return knex.schema.createTable('users', table => {
        table.increments('id').primary();
        table.string('email', 255).unique();
        table.string('name', 25);
        table.timestamp('createdAt').notNullable().defaultTo(knex.fn.now());
    });
}


export async function down(knex: Knex): Promise<void> {
    return knex.schema
        .dropTableIfExists('users');
}


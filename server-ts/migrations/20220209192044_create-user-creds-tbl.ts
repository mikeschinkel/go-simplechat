import { Knex } from "knex";


export async function up(knex: Knex): Promise<void> {
    if (await knex.schema.hasTable('userCreds')) {
        return;
    }
    return knex.schema.createTable('userCreds', table => {
        table.increments('id').primary();
        table.string('pwdHash', 255);
        table.bigInteger('userId')
                .notNullable()
                .unsigned()
                .index()
                .references('id')
                .inTable('users')
                .onDelete('CASCADE');
    });
}


export async function down(knex: Knex): Promise<void> {
    return knex.schema
    .dropTableIfExists('userCreds');
}


package com.phuckhoa.order.model.abstracts.postgres;

import java.util.UUID;

import com.phuckhoa.order.model.abstracts.base.DatabaseBaseModelAbstract;

/**
 * Abstract class {@link DatabasePostgresModelAbstract} with a generic type parameter.
 * This class is used as a base for entities that need to be stored in a Postgres database.
 * 
 * @property id The id of the model.
 * */

public abstract class DatabasePostgresModelAbstract extends DatabaseBaseModelAbstract<String> {

    protected DatabasePostgresModelAbstract() {
        super(UUID.randomUUID().toString());
    }
}
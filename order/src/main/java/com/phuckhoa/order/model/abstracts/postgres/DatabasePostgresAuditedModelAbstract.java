package com.phuckhoa.order.model.abstracts.postgres;

import java.util.UUID;

import com.phuckhoa.order.model.abstracts.base.DatabaseBaseAuditedModelAbstract;

/**
 * Abstract class {@link DatabasePostgresAuditedModelAbstract} with a generic type parameter.
 * This class is used as a base for entities that need to be audited and have a soft delete and need to be stored in a Postgres database.
 * 
 * @property id The id of the model.
 * 
 * @property createdAt The date when the entity was created.
 * 
 * @property updatedAt The date when the entity was last updated.
 * 
 * @property createdByUserId The id of the user who created the entity.
 * 
 * @property updatedByUserId The id of the user who last updated the entity.
 * */
public abstract class DatabasePostgresAuditedModelAbstract extends DatabaseBaseAuditedModelAbstract<String> {
    
    protected DatabasePostgresAuditedModelAbstract() {
        super(UUID.randomUUID().toString());
    }
}
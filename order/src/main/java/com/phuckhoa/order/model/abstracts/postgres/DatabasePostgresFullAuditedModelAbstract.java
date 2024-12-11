package com.phuckhoa.order.model.abstracts.postgres;


import java.util.UUID;

import com.phuckhoa.order.model.abstracts.base.DatabaseBaseFullAuditedModelAbstract;

/**
 * Abstract class {@link DatabasePostgresFullAuditedModelAbstract} with a generic type parameter.
 * This class is used as a base for entities that need to be audited and have a soft delete and need to be stored in a Postgres database.
 * 
 * @property id The id of the model.
 * 
 * @property isDeleted A flag to indicate if the entity is deleted.
 * 
 * @property createdAt The date when the entity was created.
 * 
 * @property updatedAt The date when the entity was last updated.
 * 
 * @property createdByUserId The id of the user who created the entity.
 * 
 * @property updatedByUserId The id of the user who last updated the entity.
 * */
public abstract class DatabasePostgresFullAuditedModelAbstract extends DatabaseBaseFullAuditedModelAbstract<String> {
    protected DatabasePostgresFullAuditedModelAbstract() {
        super(UUID.randomUUID().toString());
    }
}
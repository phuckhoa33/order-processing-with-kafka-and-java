package com.phuckhoa.order.model.abstracts.base;

import java.util.Date;

/**
 * Abstract class {@link DatabaseBaseAuditedModelAbstract} with a generic type parameter.
 * This class is used as a base for entities that need to be audited.
 * 
 * @property createdAt The date when the entity was created.
 * 
 * @property updatedAt The date when the entity was last updated.
 * 
 * @property createdByUserId The id of the user who created the entity.
 * 
 * @property updatedByUserId The id of the user who last updated the entity.
 * */
public abstract class DatabaseBaseAuditedModelAbstract<T> extends DatabaseBaseModelAbstract<T>{
    protected Date createdAt;
    protected T createdByUserId;
    protected Date updatedAt;
    protected T updatedByUserId;
    
    protected DatabaseBaseAuditedModelAbstract(T id) {
        super(id);
        this.createdByUserId = id;
        this.updatedByUserId = id;
        this.createdAt = new Date();
        this.updatedAt = new Date();
    }
}
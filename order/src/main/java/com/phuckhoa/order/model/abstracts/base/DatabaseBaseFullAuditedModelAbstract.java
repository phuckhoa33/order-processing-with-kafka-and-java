package com.phuckhoa.order.model.abstracts.base;

/**
 * Abstract class {@link DatabaseBaseFullAuditedModelAbstract} with a generic type parameter.
 * This class is used as a base for entities that need to be audited and have a soft delete.
 * 
 * @property isDeleted A flag to indicate if the entity is deleted.
 * */
public abstract class DatabaseBaseFullAuditedModelAbstract<T> extends DatabaseBaseAuditedModelAbstract<T>{
    protected boolean isDeleted;
    
    protected DatabaseBaseFullAuditedModelAbstract(T id) {
        super(id);
        this.isDeleted = false;
    }
}
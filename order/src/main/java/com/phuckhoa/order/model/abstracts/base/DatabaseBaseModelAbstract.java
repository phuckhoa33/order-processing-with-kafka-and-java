package com.phuckhoa.order.model.abstracts.base;

import lombok.Data;

/**
 * Abstract class {@link DatabaseBaseModelAbstract} with a generic type parameter.
 * This class is used as a base for entities.
 * 
 * @property id The id of the model.
 *
 * */
@Data
public abstract class DatabaseBaseModelAbstract<T> {
    protected T id;
    
    protected DatabaseBaseModelAbstract(T id) {
        this.id = id;
    }
}
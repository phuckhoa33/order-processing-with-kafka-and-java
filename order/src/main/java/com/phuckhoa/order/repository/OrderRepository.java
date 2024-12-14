package com.phuckhoa.order.repository;


import com.phuckhoa.order.model.OrderModel;
import com.phuckhoa.order.repository.abstracts.base.CrudRepository;

public interface OrderRepository extends CrudRepository<OrderModel, String> {}

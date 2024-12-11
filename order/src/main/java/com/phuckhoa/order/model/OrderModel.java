package com.phuckhoa.order.model;


import java.util.Date;
import java.util.Set;

import com.phuckhoa.order.enums.OrderStatus;
import com.phuckhoa.order.enums.OrderType;
import com.phuckhoa.order.model.abstracts.postgres.DatabasePostgresFullAuditedModelAbstract;

import jakarta.persistence.CascadeType;
import jakarta.persistence.Entity;
import jakarta.persistence.EnumType;
import jakarta.persistence.Enumerated;
import jakarta.persistence.OneToMany;
import jakarta.persistence.Table;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@Getter
@Setter
@Builder
@AllArgsConstructor
@NoArgsConstructor
@Entity()
@Table(name = "orders")
public class OrderModel extends DatabasePostgresFullAuditedModelAbstract {
    // this id will be used as the actual order id for user to track the order
    private String actualOrderId;
    // this date will be used as the order date for user to track the order
    private Date orderDate;
    // this date will be used as the completed date for user to track the order
    private Date completedDate;
    // this user id will be used as the customer id for user to track the order
    private String customerId;
    // product id
    private String productId;
    // this user status display user status
    @Enumerated(EnumType.STRING)
    private OrderStatus orderStatus;
    // this order type display order type
    @Enumerated(EnumType.STRING)
    private OrderType orderType;
    // cancelable date
    private Date cancelableDate;
    // total price
    private double totalPrice;
    // all product price
    private double productPrice;
    // discount price
    private double discountPrice;

    @OneToMany(mappedBy = "order", cascade = CascadeType.ALL, orphanRemoval = true)
    private Set<OrderDiscount> discounts;
}

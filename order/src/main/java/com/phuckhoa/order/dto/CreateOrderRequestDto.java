package com.phuckhoa.order.dto;

import java.util.Date;

import com.phuckhoa.order.enums.OrderType;

import jakarta.persistence.EnumType;
import jakarta.persistence.Enumerated;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@Getter
@Setter
@AllArgsConstructor
@NoArgsConstructor
public class CreateOrderRequestDto {
    private Date orderDate;
    // this date will be used as the completed date for user to track the order
    private Date completedDate;
    // this user id will be used as the customer id for user to track the order
    private String customerId;
    // product id
    private String productId;
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
}

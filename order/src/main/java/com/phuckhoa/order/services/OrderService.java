package com.phuckhoa.order.services;

import com.phuckhoa.order.dto.CreateOrderRequestDto;

public interface OrderService {
    public String createOrder(CreateOrderRequestDto dto);
}

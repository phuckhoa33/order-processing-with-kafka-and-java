package com.phuckhoa.order.services;

import org.springframework.stereotype.Service;

import com.phuckhoa.order.dto.CreateOrderRequestDto;

@Service
public interface OrderService {
    public String createOrder(CreateOrderRequestDto dto);
}

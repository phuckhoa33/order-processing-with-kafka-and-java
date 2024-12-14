package com.phuckhoa.order.services.impl;

import org.springframework.beans.factory.annotation.Autowired;

import com.phuckhoa.order.client.DiscountClient;
import com.phuckhoa.order.client.UserClient;
import com.phuckhoa.order.dto.CreateOrderRequestDto;
import com.phuckhoa.order.exception.InvalidDiscountException;
import com.phuckhoa.order.repository.OrderRepository;
import com.phuckhoa.order.services.OrderService;

public class OrderServiceImpl implements OrderService {

    @Autowired
    private OrderRepository orderRepository;

    @Autowired
    private DiscountClient discountClient;

    @Autowired
    private UserClient userClient;

    /**
     * Creates a new order with the provided details.
     * 
     * This method accepts an order request, validates any discount applied by calling the
     * Discount Service, and creates an order with the associated discount information.
     * 
     * @param dto the request containing the details of the order.
     *                      Includes fields like customerId, totalAmount, and optional discountId.
     * @return Order the created order with discount details (if valid).
     * @throws InvalidDiscountException if the provided discount ID is not valid or expired.
     */

    @Override
    public String createOrder(CreateOrderRequestDto dto) throws InvalidDiscountException {
        orderRepository.save(null);
        userClient.getUserById(null);
        discountClient.getDiscountById(null);
        return "HELLO GUY";
    }
    
}

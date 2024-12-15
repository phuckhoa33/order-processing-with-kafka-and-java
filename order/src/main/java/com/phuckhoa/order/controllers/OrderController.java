package com.phuckhoa.order.controllers;

import java.util.HashMap;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.phuckhoa.order.dto.CreateOrderRequestDto;
import com.phuckhoa.order.services.OrderService;

import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;


@RestController
@CrossOrigin("*")
@RequestMapping("/api/order")
public class OrderController {
    @Autowired
    private OrderService orderService;


    @PostMapping("")
    public ResponseEntity<?> postMethodName(@RequestBody CreateOrderRequestDto dto) {
        System.out.println(dto);
        
        HashMap<String, Object> result = new HashMap<>();

        try {
            String a = orderService.createOrder(dto);
            result.put("success", true);
            result.put("message", "Success to call api doInsertBill");
            result.put("data", a);
        } catch (Exception e) {
            result.put("success", false);
            result.put("message", "Fail to call api doInsertBill");
            result.put("data", null);
            e.printStackTrace();
        }

        return ResponseEntity.ok(result);
    }
    
}

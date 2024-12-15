package com.phuckhoa.order.client;

import org.springframework.cloud.openfeign.FeignClient;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;

/**
 * Feign client for interacting with the Discount Service.
 * This client will be used to query the Discount Service for discount
 * information.
 */
@FeignClient(name = "user-service", url = "http://localhost:8888")
public interface UserClient {

    /**
     * Retrieves user information by its ID.
     * 
     * @param userId the ID of the user.
     * @return ResponseEntity containing the user details.
     */
    @GetMapping("/api/v1/public/users/{userId}")
    ResponseEntity<Object> getUserById(@PathVariable("userId") String userId);
}
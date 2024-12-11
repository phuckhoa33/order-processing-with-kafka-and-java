package com.phuckhoa.order.client;

import org.springframework.cloud.openfeign.FeignClient;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;

/**
 * Feign client for interacting with the Discount Service.
 * This client will be used to query the Discount Service for discount information.
 */
@FeignClient(name = "discount-service", url = "http://discount-service-url")
public interface DiscountClient {

    /**
     * Retrieves discount information by its ID.
     * 
     * @param discountId the ID of the discount.
     * @return ResponseEntity containing the discount details.
     */
    @GetMapping("/discounts/{discountId}")
    ResponseEntity<Object> getDiscountById(@PathVariable("discountId") Long discountId);
}
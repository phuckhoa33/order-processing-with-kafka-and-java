package com.phuckhoa.order.model;

import java.util.Date;

import com.phuckhoa.order.model.abstracts.postgres.DatabasePostgresFullAuditedModelAbstract;

import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import jakarta.persistence.JoinColumn;
import jakarta.persistence.ManyToOne;
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
@Table(name = "order-discounts")
public class OrderDiscount extends DatabasePostgresFullAuditedModelAbstract {
    @Id
    @GeneratedValue(strategy = GenerationType.UUID)
    private String id;
    @ManyToOne
    @JoinColumn(name = "orderId", nullable = false)
    private OrderModel order;
    private String discountId;
    private int discountAmount;
    private Date validDate;
}

package com.phuckhoa.order.enums;

public enum OrderType {
    REGULAR("REGULAR"),
    PRIORITY("PRIORITY"),
    BULK("BULK");

    private final String value;

    OrderType(String value) {
        this.value = value;
    }

    public String getValue() {
        return value;
    }

    public static OrderType fromValue(String value) {
        for (OrderType type : OrderType.values()) {
            if (type.value.equalsIgnoreCase(value)) {
                return type;
            }
        }
        throw new IllegalArgumentException("Unknown enum type " + value);
    }
}
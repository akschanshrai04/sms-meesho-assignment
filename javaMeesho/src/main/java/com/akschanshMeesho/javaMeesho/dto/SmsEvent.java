package com.akschanshMeesho.javaMeesho.dto;

import lombok.Builder;
import lombok.Data;

@Data
@Builder
public class SmsEvent {
    private String phoneNumber;
    private String message;
    private String status;
}

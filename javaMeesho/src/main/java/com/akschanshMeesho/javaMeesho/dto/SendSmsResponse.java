package com.akschanshMeesho.javaMeesho.dto;

import lombok.AllArgsConstructor;
import lombok.Data;

@Data
@AllArgsConstructor
public class SendSmsResponse {
    private String status;
    private String message;
}

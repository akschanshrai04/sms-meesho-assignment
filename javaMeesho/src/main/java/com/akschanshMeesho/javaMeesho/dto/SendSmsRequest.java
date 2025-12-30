package com.akschanshMeesho.javaMeesho.dto;

import jakarta.validation.constraints.NotBlank;
import jakarta.validation.constraints.Pattern;
import lombok.Data;

@Data
public class SendSmsRequest {
    @NotBlank(message = "PhoneNumber is mandatory")
    @Pattern(regexp = "^[6-9][0-9]{9}$", message = "Invalid phone number")
    private String phoneNumber;

    @NotBlank(message = "message is mandatory")
    private String message;
}

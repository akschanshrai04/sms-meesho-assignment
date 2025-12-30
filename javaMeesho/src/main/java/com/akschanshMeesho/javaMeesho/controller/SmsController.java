package com.akschanshMeesho.javaMeesho.controller;

import com.akschanshMeesho.javaMeesho.dto.SendSmsRequest;
import com.akschanshMeesho.javaMeesho.dto.SendSmsResponse;
import com.akschanshMeesho.javaMeesho.service.SmsService;
import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/v0/sms")
@RequiredArgsConstructor
public class SmsController {
    private final SmsService smsService;

    @PostMapping("/send")
    public ResponseEntity<SendSmsResponse> sendSms(@Valid @RequestBody SendSmsRequest smsRequest){
        return ResponseEntity.ok(smsService.sendSms(smsRequest));
    }
}

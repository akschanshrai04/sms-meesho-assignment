package com.akschanshMeesho.javaMeesho.service;

import com.akschanshMeesho.javaMeesho.dto.SendSmsRequest;
import com.akschanshMeesho.javaMeesho.dto.SendSmsResponse;
import com.akschanshMeesho.javaMeesho.dto.SmsEvent;
import com.akschanshMeesho.javaMeesho.kafka.KafkaMessageProducer;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;

@Service
@RequiredArgsConstructor
public class SmsService {

    private final BlockedNumberService blockedNumberService;
    private final VendorSmsService vendorSmsService;
//    private final SmsKafkaProducer smsKafkaProducer;
    private final KafkaMessageProducer kafkaMessageProducer;

    public SendSmsResponse sendSms(SendSmsRequest request){
        //ill check if the number is blocked
        if(blockedNumberService.isBlocked(request.getPhoneNumber())){
            return new SendSmsResponse("BLOCKED" , "the number is blocked");
        }

        boolean sent = vendorSmsService.sendSms(request.getPhoneNumber() , request.getMessage());

        SmsEvent event = SmsEvent.builder()
                .phoneNumber(request.getPhoneNumber())
                .message(request.getMessage())
                .status(sent ? "SUCCESS" : "FAILED")
                .build();

        kafkaMessageProducer.sendMessage(event);

        return new SendSmsResponse(
                sent ? "SUCCESS" : "FAILED",
                sent ? "SMS processed" : "SMS not sent"
        );
    }
}

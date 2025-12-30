package com.akschanshMeesho.javaMeesho.service;

import org.springframework.stereotype.Service;

@Service
public class VendorSmsService {
    public boolean sendSms(String phoneNumber, String message) {
        return Math.random() > .1;
    }
}

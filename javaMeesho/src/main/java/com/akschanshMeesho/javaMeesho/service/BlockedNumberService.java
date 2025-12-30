package com.akschanshMeesho.javaMeesho.service;

import lombok.RequiredArgsConstructor;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.stereotype.Service;

@RequiredArgsConstructor
@Service
public class BlockedNumberService {
    private final RedisTemplate<String, Object> redisTemplate;
    public boolean isBlocked(String phoneNumber){
        Object value = redisTemplate.opsForValue().get(phoneNumber);
        return value != null;
    }
}

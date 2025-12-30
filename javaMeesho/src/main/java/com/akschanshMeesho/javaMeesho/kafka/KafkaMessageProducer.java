package com.akschanshMeesho.javaMeesho.kafka;

import com.akschanshMeesho.javaMeesho.dto.SmsEvent;
import lombok.AllArgsConstructor;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.core.env.Environment;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.stereotype.Component;
import tools.jackson.databind.ObjectMapper;

@Component
@AllArgsConstructor
public class KafkaMessageProducer {
    @Autowired
    private Environment environment;
    private final KafkaTemplate<String, String> kafkaTemplate;
    private final ObjectMapper objectMapper;

    public void sendMessage(SmsEvent event){
        try {
            String json = objectMapper.writeValueAsString(event);
            kafkaTemplate.send(environment.getProperty("kafka.topicname"), event.getPhoneNumber(), json);
        } catch (Exception e) {
            throw new RuntimeException("Failed to serialize SmsEvent", e);
        }
    }

}

package com.example.rolldice;

import org.springframework.boot.Banner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class RolldiceApplication {

	public static void main(String[] args) {
		SpringApplication app = new SpringApplication(RolldiceApplication.class);
		app.setBannerMode(Banner.Mode.OFF);
		app.run(args);

	}

}

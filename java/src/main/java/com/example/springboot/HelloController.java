package com.example.springboot;

// import java.io.File;
// import java.io.FileNotFoundException;
// import java.util.Scanner;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class HelloController {

    @GetMapping("/")
    public String index() {
        return "Greetings from Spring Boot";
    }

    @GetMapping("/hello")
    public String helloPage() {
        // String info = getInfoFromFile();
        return "Hit on the Hello Page";
    }

    // private String getInfoFromFile() {
    //     try {
    //         File myFile = new File("myfile.json");
    //         Scanner myReader = new Scanner(myFile);
    //     } catch (FileNotFoundException e) {
    //         System.out.println("Couldn't find file");
    //         e.printStackTrace();
    //     }

    //     return "ok";
    // }
    
}

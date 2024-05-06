package com.example.springboot;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.stream.IntStream;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class HelloController {

    @GetMapping("/")
    public String index() {
        return "Greetings from Spring Boot!";
    }

    @GetMapping("/test")
    public String testFunction() {
        int[] myIntArray = new int[3];
        int[] myIntArray2 = {1, 2, 3};
        int[] myIntArray3 = new int[]{1, 2, 3};

        int[] myIntArray4 = IntStream.range(0, 100).toArray();
        int[] myIntArray5 = IntStream.rangeClosed(0, 100).toArray();
        int[] myIntArray6 = IntStream.of(12, 23, 25, 77, 1, 50).toArray();
        int[] myIntArray7 = IntStream.of(12, 23, 25, 77, 1, 50).sorted().map(number -> number*number).toArray();

        Arrays.stream(myIntArray4).forEach(number -> {System.out.println(number);});

        List<Integer> myList = new ArrayList<>();
        myList.add(5);
        List<Integer> myList2 = new ArrayList<>();
        myList2.add(60);
        for (int i = 0; i < 10; i ++){
            myList.add(i);
        }

        myList.addAll(myList2);

        myList.remove(5);
        myList.remove(Integer.valueOf(2));

        myList.get(0);

        myList.set(1, Integer.valueOf(230239));

        myList.contains(Integer.valueOf("8"));

        myList.size();

        myList.isEmpty();

        myList2.clear();
        myList.toArray();
        myList.subList(0, myList.size()/2);
        myList.sort((a, b) -> a < b ? a : b);
        myList.replaceAll(number -> number * 2);




        

        return "";
    }
}

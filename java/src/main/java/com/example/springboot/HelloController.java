package com.example.springboot;

import java.lang.reflect.Array;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.Comparator;
import java.util.List;
import java.util.stream.IntStream;

import org.springframework.util.comparator.Comparators;
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
        System.out.println("my int array operations");
        System.out.println(Arrays.toString(myIntArray));

        int[] myIntArray4 = IntStream.range(0, 100).toArray();
        System.out.println(Arrays.toString(myIntArray4));
        int[] myIntArray5 = IntStream.rangeClosed(0, 100).toArray();
        System.out.println(Arrays.toString(myIntArray5));
        int[] myIntArray6 = IntStream.of(12, 23, 25, 77, 1, 50).toArray();
        System.out.println(Arrays.toString(myIntArray6));
        int[] myIntArray7 = IntStream.of(12, 23, 25, 77, 1, 50).sorted().map(number -> number*number).toArray();
        System.out.println(Arrays.toString(myIntArray7));

        List<Integer> myList = new ArrayList<>();
        myList.add(5);
        List<Integer> myList2 = new ArrayList<>();
        myList2.add(60);
        for (int i = 0; i < 10; i ++){
            myList2.add(i);
        }
        System.out.println(myList2);


        myList.addAll(myList2);
        System.out.println("my list with mylist2");

        System.out.println(myList);

        myList.remove(5);
        System.out.println("removed index 5");
        System.out.println(myList);

        System.out.println("contains 2 ?");
        System.out.println(myList.contains(Integer.valueOf("2")));
        System.out.println("removed value of 2");
        myList.remove(Integer.valueOf(2));
        System.out.println(myList);

        System.out.println("contains 2 ?");
        System.out.println(myList.contains(Integer.valueOf("2")));

        System.out.println("get first rlement"); 
        System.out.println(myList.get(0));

        System.out.println("set value of 1 index");
        myList.set(1, Integer.valueOf(230239));
        System.out.println(myList);


        System.out.println("contains 8 ?");
        System.out.println(myList.contains(Integer.valueOf("8")));
        System.out.println("contains 230239 ?");
        System.out.println(myList.contains(Integer.valueOf("230239")));

        System.out.println("size");

        System.out.println(myList.size());

        System.out.println(myList.isEmpty());

        System.out.println("before clear");
        System.out.println(myList2);
        myList2.clear();
        System.out.println("after clear");
        System.out.println(myList2);

        myList.toArray();
        System.out.println(myList.subList(0, myList.size()/2));
        // myList.sort(Comparator.comparing(Integer::valueOf));
        Collections.sort(myList);
        System.out.println(myList);
        // Override default sort function
        Collections.sort(myList, Collections.reverseOrder());
        System.out.println(myList);
        myList.replaceAll(number -> number * 2);
        System.out.println(myList);





        

        return "";
    }
}

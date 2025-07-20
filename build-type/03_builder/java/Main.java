package main;

public class Main {

    public static void main(String[] args) {
        
        var person1 = Person.builder()
                .name("John Doe")
                .age(30)
                .address("123 Main St")
                .build();

        var person2 = Person.builder()
                .name("Jane Smith")
                .age(25)
                .build();

        // 使用建造者设计模式创建不同属性的 Person 对象
        System.out.println("Person1: " + person1);
        System.out.println("Person2: " + person2);

    }

}

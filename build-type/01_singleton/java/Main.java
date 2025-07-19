package main;

public class Main {

    public static void main(String[] args) {

        var singleton = SingletonEager.getInstance();
        System.out.println(singleton);

        var lazySingleton1 = SingletonLazy.getLazyInstance1();
        var lazySingleton2 = SingletonLazy.getLazyInstance1();

        System.out.println(lazySingleton1);
        System.out.println(lazySingleton2);

    }

}

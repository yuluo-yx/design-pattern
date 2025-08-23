package main;

public class Main {
 
    public static void main(String[] args) {
        System.out.println("============= 原始的正方形绘制类 ==============");
        Draw square = new Square();
        System.out.println(square.draw());

        System.out.println("============= 装饰后的正方形绘制类 ==============");
        Draw squareWrapper = new ColorSquareWrapper(square);
        System.out.println(squareWrapper.draw());
    }

}

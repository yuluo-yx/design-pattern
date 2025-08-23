package main;

// 这里可以使用抽象类在封装一层，使其更有结构性
public class ColorSquareWrapper implements Draw {

    private final Draw draw;

    public ColorSquareWrapper(Draw draw) {
        this.draw = draw;
        // 如果使用抽象类时，这里直接调用父类的构造方法
        // super(draw);
    }

    @Override
    public String draw() {
        return "Drawing a decorated square, the color is red";
    }

}

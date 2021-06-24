package flyweight.pattern;

// step 4 使用该工厂，通过传递颜色信息来获取实体类的对象。
public class FlyweightPatternDemo {
    private static final String colors[] = {"Red", "Green", "Blue", "White", "Black"};

    public static void main(String[] args) {
        for (int i = 0; i < 20; i++) {
            Circle circle = (Circle) ShapeFactory.getCircle(getRandomColor());
            circle.setX(getRandomX());
            circle.setY(getRandomY());
            circle.setRadius(100);
            circle.draw();
        }
    }

    private static String getRandomColor() {
        return colors[(int) (Math.random() * colors.length)];
    }

    private static int getRandomX() {
        return (int) (Math.random() * 100);
    }

    private static int getRandomY() {
        return (int) (Math.random() * 100);
    }
}

/*
* output
creating circle of color: White
Circle: Draw() {color='White', x=48, y=33, radius=100}
creating circle of color: Black
Circle: Draw() {color='Black', x=51, y=67, radius=100}
Circle: Draw() {color='Black', x=44, y=61, radius=100}
Circle: Draw() {color='White', x=63, y=88, radius=100}
creating circle of color: Green
Circle: Draw() {color='Green', x=38, y=84, radius=100}
creating circle of color: Red
Circle: Draw() {color='Red', x=86, y=58, radius=100}
Circle: Draw() {color='Green', x=28, y=22, radius=100}
Circle: Draw() {color='Red', x=6, y=43, radius=100}
Circle: Draw() {color='White', x=1, y=61, radius=100}
Circle: Draw() {color='White', x=43, y=61, radius=100}
Circle: Draw() {color='Black', x=76, y=96, radius=100}
creating circle of color: Blue
Circle: Draw() {color='Blue', x=12, y=88, radius=100}
Circle: Draw() {color='White', x=17, y=79, radius=100}
Circle: Draw() {color='Green', x=75, y=43, radius=100}
Circle: Draw() {color='Red', x=84, y=84, radius=100}
Circle: Draw() {color='Green', x=63, y=58, radius=100}
Circle: Draw() {color='Green', x=77, y=53, radius=100}
Circle: Draw() {color='Green', x=32, y=56, radius=100}
Circle: Draw() {color='Black', x=22, y=99, radius=100}
Circle: Draw() {color='White', x=43, y=59, radius=100}

*
* */

/**
 * Class对象的生成方式如下：
 *
 * 1、类名.class
 *
 * 说明： JVM将使用类装载器, 将类装入内存(前提是:类还没有装入内存),不做类的初始化工作.返回Class的对象
 *
 * 2、Class.forName("类名字符串")  （注：类名字符串是包名+类名）
 *
 * 说明：装入类,并做类的静态初始化，返回Class的对象
 *
 * 3、实例对象.getClass()
 *
 * 说明：对类进行静态初始化、非静态初始化；返回引用o运行时真正所指的对象(因为:子对象的引用可能会赋给父对象的引用变量中)所属的类的Class的对象
 */
public class TestClass {
    public static void main(String[] args) {
        try {
            //测试.class
            Class testTypeClass = TestClassType.class;
            System.out.println("testTypeClass---" + testTypeClass);

            //测试Class.forName()
            Class testTypeForName = Class.forName("TestClassType");
            System.out.println("testTypeForName---" + testTypeForName);

            //测试Object.getClass()
            TestClassType testTypeGetClass = new TestClassType();
            System.out.println("testTypeGetClass---" + testTypeGetClass.getClass());
        } catch (Exception e) {

// TODO Auto-generated catch block
            e.printStackTrace();
        }
    }

}

class TestClassType {

    //构造函数
    public TestClassType() {
        System.out.println("----构造函数---");
    }

    //静态的参数初始化
    static {
        System.out.println("---静态的参数初始化---");
    }

    //非静态的参数初始化
    {
        System.out.println("----非静态的参数初始化---");
    }

}
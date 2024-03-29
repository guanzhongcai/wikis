package single.pattern.n4doublecheck;

/**
 * 双检锁/双重校验锁（DCL，即 double-checked locking）
 * 采用双锁机制，安全且在多线程情况下能保持高性能。
 * getInstance() 的性能对应用程序很关键。
 */
public class Singleton {
    private volatile static Singleton singleton;
    private Singleton(){}
    public static Singleton getSingleton(){
        if (singleton == null) {
            synchronized(Singleton.class) {
                if (singleton == null) {
                    singleton = new Singleton();
                }
            }
        }
        return singleton;
    }
}

package learn.reflect;

import java.io.File;
import java.lang.reflect.Field;
import java.lang.reflect.Method;
import java.net.URL;
import java.net.URLClassLoader;
import java.util.Enumeration;
import java.util.jar.JarEntry;
import java.util.jar.JarFile;

/**
 * @ClassName: ReflexDemo
 * 通过反射获取类、属性及方法
 */
public class ReflectDemo {

    private static StringBuffer sBuffer;

    public static void getJar(String jar) throws Exception {
        try {
            File file = new File(jar);
            URL url = file.toURI().toURL();
            URLClassLoader classLoader = new URLClassLoader(new URL[]{url},
                    Thread.currentThread().getContextClassLoader());

            JarFile jarFile = new JarFile(jar);
            Enumeration<JarEntry> enumeration = jarFile.entries();
            JarEntry jarEntry;

            sBuffer = new StringBuffer();    //存数据

            while (enumeration.hasMoreElements()) {
                jarEntry = enumeration.nextElement();

                if (!jarEntry.getName().contains("META-INF")) {
                    String classFullName = jarEntry.getName();
                    if (!classFullName.contains(".class")) {
                        classFullName = classFullName.substring(0, classFullName.length() - 1);
                    } else {
                        // 去除后缀.class，获得类名
                        String className = classFullName.substring(0, classFullName.length() - 6).replace("/", ".");
                        Class<?> myClass = classLoader.loadClass(className);
                        sBuffer.append("类名\t：").append(className);
                        System.out.println("类名\t：" + className);

                        // 获得属性名
                        Class<?> clazz = Class.forName(className);
                        Field[] fields = clazz.getDeclaredFields();
                        for (Field field : fields) {
                            sBuffer.append("属性名\t：").append(field.getName()).append("\n");
                            System.out.println("属性名\t：" + field.getName());
                            sBuffer.append("-属性类型\t：").append(field.getType()).append("\n");
                            System.out.println("-属性类型\t：" + field.getType());
                        }

                        // 获得方法名
                        Method[] methods = myClass.getMethods();
                        for (Method method : methods) {
                            if (method.toString().indexOf(className) > 0) {
                                sBuffer.append("方法名\t：").append(method.toString().substring(method.toString().indexOf(className))).append("\n");
                                System.out.println("方法名\t：" + method.toString().substring(method.toString().indexOf(className)));
                            }
                        }
                        sBuffer.append("--------------------------------------------------------------------------------" + "\n");
                        System.out.println("--------------------------------------------------------------------------------");
                    }
                }
            }
        } catch (Exception e) {
            e.printStackTrace();
        } finally {
            sBuffer.append("End");
            System.out.println("End");

            WriteFile.write(sBuffer);    //写文件
        }
    }

}
package learn.reflect;

import java.io.BufferedWriter;
import java.io.File;
import java.io.FileWriter;

public class WriteFile {
    private static String pathname = "src/com/adamjwh/jnp/ex14/out.txt";

    public static void write(StringBuffer sBuffer) throws Exception {
        File file = new File(pathname);
        BufferedWriter bw = new BufferedWriter(new FileWriter(file));

        bw.write(sBuffer.toString());
        bw.close();
    }
}

import com.alibaba.fastjson.JSONObject;

public class FastJsonTest {
    public static void main(String[] args) {
        JSONObject object = new JSONObject();
        object.put("string", "hello");
        object.put("int", 2);
        object.replace("int", 3);
        System.out.println(object);
    }
}

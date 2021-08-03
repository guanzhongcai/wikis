package p1.transfer.object.pattern;

public class TransferObjectPatternDemo {
    public static void main(String[] args) {
        StudentBO studentBusinessObject = new StudentBO();
        for (StudentVO student : studentBusinessObject.getAllStudents()) {
            System.out.println("Student: [RollNo : "
                    + student.getRollNo() + ", Name : " + student.getName() + " ]");
        }

        StudentVO student = studentBusinessObject.getAllStudents().get(0);
        student.setName("Jack");
        studentBusinessObject.updateStudent(student);

        studentBusinessObject.getStudent(0);
        System.out.println("Student: [RollNo : "
                + student.getRollNo() + ", Name : " + student.getName() + " ]");
    }
}

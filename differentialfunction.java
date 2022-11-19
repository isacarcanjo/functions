
import java.io.File;

public class file {
    private File file;
    private double[] values;

    public file() {
        file = new File("C:\\Users\\a1285\\Desktop\\Differential\\Differential.txt");
    }

    public file(String filename) {
        file = new File(filename);
    }

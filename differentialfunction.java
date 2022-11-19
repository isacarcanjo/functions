
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

    public void start() throws IOException {

        try (BufferedReader bufferedReader = new BufferedReader(new FileReader(file))) {
            String line;
            int size = 0;
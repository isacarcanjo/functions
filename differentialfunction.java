
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
            while((line = bufferedReader.readLine()) != null) {
                size++;
            }
            values = new double[size];
            for(int i = 0; i < size; i++) {
                values[i] = Double.parseDouble(bufferedReader.readLine());
            }
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    public double[] getValues() {
        return values;
    }

    public double[] calc(int n) throws Exception {
        if(values == null) throw new Exception("Values not set");


public class Main {

    public static double func(double x) {
        return Math.sin(x);
    }

    public static double diff(double x) {

        double h = 0.0001;
        double der = (func(x+h) - func(x))/h;
        return der;
    }

    public static void main(String[] args) {

        System.out.println(diff(0.1));
    }
}
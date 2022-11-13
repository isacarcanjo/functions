



double integrate(double (*f)(double), double a, double b, int n);
double f(double x);

int main()
{
    double (*fp) (double);
    fp = &f;
    printf("%lf\n", integrate(f, 1, 2, N));
    printf("%lf\n", integrate(fp, 1, 2, N));

    return 0;
}

double f(double x)
{
    return x*x;
}

double integrate(double (*f)(double), double a, double b, int n)
{
    double delta = (b-a)/n;
    double sum = 0;
    int i;
    for (i =0; i < n; i++)
    {
        double x = a + i*delta;
        sum += delta*(f(x) + f(x+delta))/2;
    }

    return sum;
}




"""@author: github.com/isacarcanjo"""




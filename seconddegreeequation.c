


void second_degree_equation(float a, float b, float c);

int main()
{
    float a, b, c;
    printf("Enter a,b,c of the second degree equation ax^2+bx+c=0:\n");
    scanf("%f %f %f", &a, &b, &c);

    second_degree_equation(a, b, c);

    return 0;
}

void second_degree_equation(float a, float b, float c)
{
    float D, x, x1, x2;

    D = b*b - 4*a*c;

    if(D > 0)
    {
        x1 = (-b + sqrt(D)) / (2*a);
        x2 = (-b - sqrt(D)) / (2*a);

        printf("Equation roots are:\nx1 = %f\nx2 = %f\n", x1, x2);
    }
    else if(D == 0)
    {
        x = -b / (2*a);
        printf("Equation root is:\nx = %f\n", x);
    }
    else
    {
        printf("No real roots\n");
    }
}




"""@author: github.com/isacarcanjo"""




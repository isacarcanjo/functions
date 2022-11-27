

int main()
{
	// use the for loop
	int i,j;
	int a,b;
	int result;
	printf("enter the first number: ");
	scanf("%d",&a);
	printf("enter the second number: ");
	scanf("%d",&b);
	result = 0;
	for(j=a;j<=b;j++)
	{
		 i = j*j;
		 result = result + i;
	}
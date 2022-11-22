import java.util.Scanner;
import java.io.File;
public class file
{
   public static void main(String[] args)
   {
       Scanner myScanner = null;
       try
       {
           myScanner = new Scanner(new File("test.txt"));
           while (myScanner.hasNextLine())
           {
               System.out.println(myScanner.nextLine());
           }
       }
       catch (Exception ex)
       {
           ex.printStackTrace();
       }
       finally
       {
           if (myScanner != null)
           {
               myScanner.close();
           }
       }
   }
}




"""@author: github.com/isacarcanjo"""




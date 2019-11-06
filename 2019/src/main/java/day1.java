import java.io.File;
import java.io.FileNotFoundException;
import java.util.Scanner;

public class day1 {
    public static void main(String[] args) throws FileNotFoundException {

        File input = new File("inputs/day1_input.txt");
        Scanner s = new Scanner(input);
        int fuelSum = 0;
        int fuelRequiredSum = 0;
        while (s.hasNext()) {
            int mass = s.nextInt();
            fuelSum += fuelRequired(mass);
            fuelRequiredSum += fuelRequiredRecursive(mass);
        }
        System.out.println(fuelSum);

        System.out.println(fuelRequiredSum);
    }

    public static int fuelRequired(int mass) {
        return (mass / 3) - 2;
    }

    public static int fuelRequiredRecursive(int mass) {
        int fuel = fuelRequired(mass);
        if (fuel > 0) {
            return fuel + fuelRequiredRecursive(fuel);
        }
        return 0;
    }

//    // non-recursive "alternative"
//    public static int fuelRequiredForFuel(int mass) {
//        int fuel = 0;
//        int fuelIterator = mass;
//        while (fuelIterator > 0) {
//            fuelIterator = fuelRequired(fuelIterator);
//            if (fuelIterator <= 0) {
//                break;
//            }
//            fuel = fuel + fuelIterator;
//        }
//        return fuel;
//    }
}

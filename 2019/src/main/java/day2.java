import java.io.File;
import java.io.FileNotFoundException;
import java.util.*;

public class day2 {
    int offset = 0;
    int one = 0;
    int two = 0;
    int answerSlot = 0;
    int[] numbers;

    public day2(int[] numbers){
        this.numbers = numbers;
    }

    public static void main(String[] args) throws FileNotFoundException {

        File input = new File("inputs/day2_input.txt");
        Scanner s = new Scanner(input);

        // https://stackoverflow.com/questions/18838781/converting-string-array-to-an-integer-array/37093052#37093052
        int[] numbers = Arrays.stream(s.next().split(",")).mapToInt(Integer::parseInt).toArray();

        // day 1
        int[] day1numbers = numbers.clone();
        System.out.println("day2, part 1 answer: " + runIntCode(day1numbers, 12, 2));

        // day 2
        int result = 0;
        for (int i = 0; i < 100; i++) {
            for (int j = 0; j < 100; j++) {
                int[] day2numbers = numbers.clone();
                if (runIntCode(day2numbers, i, j) == 19690720) {
                    System.out.print("day2, part 2 answer: ");
                    System.out.println(i * 100 + j);
                }
            }
        }
    }

    public static int runIntCode(int[] numbers, int noun, int verb) {
        // day 2 part 1 tweak:
        numbers[1] = noun;
        numbers[2] = verb;

        boolean complete = false;
        int offset = 0;
        int one = 0;
        int two = 0;
        int answerSlot = 0;
        while (!complete) {
            int command = numbers[offset];

            if (command == 99) {
                complete = true;
            } else {
                one = numbers[1 + offset];
                two = numbers[2 + offset];
                answerSlot = numbers[3 + offset];
                offset += 4;
            }

            if (command == 1) {
                numbers[answerSlot] = numbers[one] + numbers[two];
            } else if (command == 2) {
                numbers[answerSlot] = numbers[one] * numbers[two];
            }
        }
        return numbers[0];
    }
}


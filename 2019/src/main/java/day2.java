import java.io.File;
import java.io.FileNotFoundException;
import java.util.*;
import java.io.PrintStream;

public class day2 {

    public static void main(String[] args) throws FileNotFoundException {
        final PrintStream log = System.out;
        final File input = new File("inputs/day2_input.txt");
        try (final Scanner s = new Scanner(input)) {
            // https://stackoverflow.com/questions/18838781/converting-string-array-to-an-integer-array/37093052#37093052
            final int[] numbers = Arrays.stream(s.next().split(",")).mapToInt(Integer::parseInt).toArray();
            final OpcodeComputer computer = new OpcodeComputer(numbers, log);

            // day 1
            log.println(computer.runIntCode(12, 2));

            // day 2
            log.println(computer.searchNounAndVerb(19690720));
        }
    }
}


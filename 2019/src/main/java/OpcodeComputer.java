import java.io.PrintStream;

public class OpcodeComputer {
    private int[] memory;
    private PrintStream log;

    public OpcodeComputer(final int[] numbers, final PrintStream out)
    {
        this.memory = numbers.clone();
        this.log = out;
    }

    public int runIntCode(int noun, int verb) {
        log.println("Running program with noun " + noun + " and verb " + verb);
        int numbers[] = memory.clone();

        numbers[1] = noun;
        numbers[2] = verb;

        int offset = 0;
        int one = 0;
        int two = 0;
        int answerSlot = 0;
        int command = numbers[offset];
        while (command != 99) {
                one = numbers[1 + offset];
                two = numbers[2 + offset];
                answerSlot = numbers[3 + offset];

            if (command == 1) {
                numbers[answerSlot] = numbers[one] + numbers[two];
            } else if (command == 2) {
                numbers[answerSlot] = numbers[one] * numbers[two];
            }
            offset += 4;
            command = numbers[offset];

        }
        return numbers[0];
    }

    // Returns the noun and verb, concatenated.
    public int searchNounAndVerb(final int goal)
    {
        final int limit = memory.length;
        for (int i = 0; i < limit; i++) {
            for (int j = 0; j < limit; j++) {
                if (runIntCode(i, j) == goal) {
                    return i * 100 + j;
                }
            }
        }
        return -1;
    }

}
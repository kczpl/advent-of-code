# opcodes
# 0 - adv - performs division
# 1 - mul - bitwise XOR of B and literal
# 2 - bst value of combo modulo 8
# 3 - jnz does nothing if A is 0
# 4 - bxc - withwise XOR of B and C
# 5 - out - claculates the value of combo modulo 8
# 6 - bdv - adv but with B
# 7 - cdv - adv but with C

# The value of a combo operand can be found as follows:
# Combo operands 0 through 3 represent literal values 0 through 3.
# Combo operand 4 represents the value of register A.
# Combo operand 5 represents the value of register B.
# Combo operand 6 represents the value of register C.
# Combo operand 7 is reserved and will not appear in valid programs.


with open("data/input17.txt", "r") as f:
    data = f.read()


def solve(registers, program):
    def get_combo_val(op):
        if op < 4:
            return op
        elif op == 4:
            return registers["A"]
        elif op == 5:
            return registers["B"]
        elif op == 6:
            return registers["C"]

    outputs = []
    init = 0

    while init < len(program):
        op = program[init]
        operand = program[init + 1] if init + 1 < len(program) else 0

        if op == 0:  # adv - divide A by 2^combo_operand
            registers["A"] //= 2 ** get_combo_val(operand)
        elif op == 1:  # bxl - XOR of B and literal
            registers["B"] ^= operand
        elif op == 2:  # bst - store combo_operand mod 8 in B
            registers["B"] = get_combo_val(operand) % 8
        elif op == 3:  # jnz - jump if A is not zero
            if registers["A"] != 0:
                init = operand
                continue
        elif op == 4:  # bxc - XOR B with C
            registers["B"] ^= registers["C"]
        elif op == 5:  # out - output combo_operand mod 8
            outputs.append(str(get_combo_val(operand) % 8))
        elif op == 6:  # bdv - like adv but store in B
            registers["B"] = registers["A"] // 2 ** get_combo_val(operand)
        elif op == 7:  # cdv - like adv but store in C
            registers["C"] = registers["A"] // 2 ** get_combo_val(operand)

        init += 2

    return outputs

def part1(data):
    lines = data.split("\n")
    registers = {
        "A": int(lines[0].split(": ")[1]),
        "B": int(lines[1].split(": ")[1]),
        "C": int(lines[2].split(": ")[1]),
    }

    program = [int(x) for x in lines[-1].split(": ")[1].split(",")]

    outputs = solve(registers, program)
    return ",".join(outputs)

print(part1(data))

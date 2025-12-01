with open("data/input22.txt", "r") as f:
    data = f.read()

def get_next(secret):
    # multiply by 64, mix (XOR), and prune
    secret = (secret ^ (secret * 64)) % 16777216

    # divide by 32, mix (XOR), and prune
    secret = (secret ^ (secret // 32)) % 16777216

    # multiply by 2048, mix (XOR), and prune
    secret = (secret ^ (secret * 2048)) % 16777216

    return secret

def find_num(initial):
    secret = initial
    for _ in range(2000):
        secret = get_next(secret)
    return secret

def part1(data):
    init = [int(x) for x in data.strip().splitlines()]
    final = [find_num(num) for num in init]
    print(sum(final))

part1(data)

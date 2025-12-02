#imports 

#fucntional code
def validate_product_ids(product_ids):
    invalid_ids = []
    for pid in product_ids:
        #each id is first and last id 
        (first_id, last_id) = pid.split("-")
        for i in range(int(first_id), int(last_id) + 1):
            #for invalid ids where the id repeats itself
            str_i = str(i)
            if str_i[0] == "0":
                str_i = str_i[1:]
            if str_i[0:round(len(str_i)/2)] == str_i[round(len(str_i)/2):]:
                print(f"Invalid id found: {i}")
                invalid_ids.append(i)
    return sum(invalid_ids)

#part 2
def invalid_repeats(product_ids):
    invalid_ids = []
    for pid in product_ids:
        (first_id, last_id) = pid.split("-")
        for i in range(int(first_id), int(last_id) + 1):
            #for invalid ids where the id repeats itself
            str_i = str(i)
            if str_i[0] == "0":
                str_i = str_i[1:]
            for j in range(0,round(len(str_i)/2)):
                sequence = str_i[0:j+1]
                if len(sequence) == 0:
                    print(f"Skipping empty sequence for id: {i}")
                    continue
                if str_i.count(sequence) == len(str_i)/len(sequence):
                    print(f"Invalid id found with repeating sequence: {i}")
                    invalid_ids.append(i)
                    break
    return sum(invalid_ids)


#call main function
def main():
    #get input files 
    with open("input.txt", "r") as file:
        line = file.readline().strip()
    product_ids = line.split(",")
    invalid_id_sum = validate_product_ids(product_ids)
    #print(f"Sum of invalid ids: {invalid_id_sum}")
    invalid_ids_with_repeats = []
    invalid_ids_with_repeats = invalid_repeats(product_ids)
    print(f"Invalid ids with repeating sequences: {invalid_ids_with_repeats}")

if __name__ == "__main__":
    main()
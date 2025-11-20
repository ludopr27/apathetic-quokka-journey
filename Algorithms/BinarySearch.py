'''binary_search'''
def binary_search(list , item):
    low = 0
    high = len(list)-1
    while low <= high:
        mid = int((low + high) / 2)
        guess = list[mid]
        if guess == item:
            return mid
        if guess > item:
            high = mid - 1
        else:
            low = mid + 1
    return None

'''sum'''
def sum(list):
    if list == []:
        return 0
    return list[0] + sum (list[1:])

'''smallest value in array'''
def findSmallest(arr):
    smallest = arr[0] 
    smallest_index = 0 
    for i in range(1, len(arr)):
        if arr[i] < smallest:
            smallest = arr[i]
            smallest_index = i
    return smallest_index

'''sorting array'''
def selectionSort(arr):
    newArr = []
    for i in range(len(arr)):
        smallest = findSmallest(arr)
        newArr.append(arr.pop(smallest))
    return newArr

        
'''chiamata codice'''
my_list = [1, 3, 5, 7, 9, 0]
my_list2 = [2, 6, 5, 3, 9, 0, 1, 8]
print(binary_search(my_list, 3))   # Output: 1
print(binary_search(my_list, -1))  # Output: None

print(sum(my_list))
print(findSmallest(my_list2))
print(selectionSort(my_list2))

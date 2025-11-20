'''smallest value in array'''
def findSmallest(arr):
    smallest = arr[0] 
    smallest_index = 0 
    for i in range(1, len(arr)):
        if arr[i] < smallest:
            smallest = arr[i]
            smallest_index = i
    return smallest_index

'''biggest value in array'''
def findBiggest(arr):
    biggest = arr[0]
    biggest_index = 0
    iNum = 0
    for i in range(1, len(arr)):
        if arr[i] > biggest:
            biggest = arr[i]
            biggest_index = i
    return str(biggest) + " e l'indice è: " + str(biggest_index) + " e i numeri totali sono " + str(i+1)

'''sorting array'''
def selectionSort(arr):
    newArr = []
    for i in range(len(arr)):
        smallest = findSmallest(arr)
        newArr.append(arr.pop(smallest))
    return newArr

def quicksort(array):
    if len(array) < 2:
        return array 
    else:
        pivot = array[0]
        less = [i for i in array[1:] if i <= pivot] 
        greater = [i for i in array[1:] if i > pivot] 
        return quicksort(less) + [pivot] + quicksort(greater)



my_arr = [2, 6, 5, 3, 9, 0, 10, 8]
print(findBiggest(my_arr))
print(quicksort(my_arr))
public class BinarySearch {

    // Restituisce l'indice della chiave se trovata, altrimenti -1
    public static int search(int[] a, int key) {
        int lo = 0;
        int hi = a.length - 1;

        while (lo <= hi) {
            int mid = lo + (hi - lo) / 2;

            if (key < a[mid]) {
                hi = mid - 1;
            } else if (key > a[mid]) {
                lo = mid + 1;
            } else {
                return mid;  // trovato
            }
        }
        return -1;  // non trovato
    }

    public static void main(String[] args) {
        int[] arr = {1, 3, 4, 7, 9, 12, 15};
        int target = 7;

        int pos = search(arr, target);
        System.out.println("Risultato: " + pos);
    }
}

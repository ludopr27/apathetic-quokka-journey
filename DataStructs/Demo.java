public class Demo {

    public static void main(String[] args) {
        System.out.println(true && false || true && true);

        int[] arr = {5, 5, 4};

        System.out.println(areEquals(arr));

    }

    public static Boolean areEquals(int[] integers){

        if (integers[0] == integers[1] && integers[1] == integers[2]) return true;
        return false;

    }

}

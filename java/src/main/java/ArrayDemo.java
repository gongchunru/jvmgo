/**
 * @author gongchunru
 * @Package PACKAGE_NAME
 * @date 2017/10/17 00:58
 */
public class ArrayDemo {
    public static void main(String[] args) {
        int[] a1 = new int[10];       // newarray
        String[] a2 = new String[10]; // anewarray
        int[][] a3 = new int[10][10]; // multianewarray
        int x = a1.length;            // arraylength
        a1[0] = 100;                  // iastore
        int y = a1[0];                // iaload
        a2[0] = "abc";                // aastore
        String s = a2[0];             // aaload
    }

    public void test() {
        int [][][] x = new int [3][4][5];
    }
}

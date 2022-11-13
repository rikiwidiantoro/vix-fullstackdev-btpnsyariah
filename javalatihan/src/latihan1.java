//public class latihan1 {
//
//}
class Laptop {
    int harga = 7500000;
    String cekHarga () {
        return "Harga Laptop adalah " + harga;
    }
}
class BelajarJava {
    public static void main(String args[]) {

        Laptop laptop = new Laptop();
        String hargaLaptop = laptop.cekHarga();
        System.out.println(hargaLaptop);
    }
}